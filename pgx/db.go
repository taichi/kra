// Copyright 2021 taichi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pgx

import (
	"context"
	"time"

	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/taichi/kra"
)

func Open(ctx context.Context, connString string, hooks ...interface{}) (*DB, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	return OpenConfig(ctx, config, hooks...)
}

type Identifier pgx.Identifier

func OpenConfig(ctx context.Context, config *pgxpool.Config, hooks ...interface{}) (*DB, error) {
	if pool, err := pgxpool.ConnectConfig(ctx, config); err != nil {
		return nil, err
	} else {
		return NewDB(pool, NewCore(kra.NewCore(kra.PostgreSQL), hooks...)), nil
	}
}

type Conn struct {
	conn  *pgx.Conn
	core  *Core
	count int64
}

func Connect(ctx context.Context, connString string, hooks ...interface{}) (*Conn, error) {
	if config, err := pgx.ParseConfig(connString); err != nil {
		return nil, err
	} else {
		return ConnectConfig(ctx, config, hooks...)
	}
}

func ConnectConfig(ctx context.Context, connConfig *pgx.ConnConfig, hooks ...interface{}) (*Conn, error) {
	if conn, err := pgx.ConnectConfig(ctx, connConfig); err != nil {
		return nil, err
	} else {
		return NewConn(conn, NewCore(kra.NewCore(kra.PostgreSQL), hooks...)), nil
	}
}

func NewConn(conn *pgx.Conn, core *Core) *Conn {
	return &Conn{conn, core, 0}
}

func (conn *Conn) Conn() *pgx.Conn {
	return conn.conn
}

func (conn *Conn) Close(ctx context.Context) error {
	return NewConnClose(conn, conn.conn.Close).Proceed(ctx)
}

func (conn *Conn) Begin(ctx context.Context) (*Tx, error) {
	return conn.BeginTx(ctx, pgx.TxOptions{})
}

func (conn *Conn) BeginFunc(ctx context.Context, f func(*Tx) error) error {
	return conn.BeginTxFunc(ctx, pgx.TxOptions{}, f)
}

func (conn *Conn) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*Tx, error) {
	return NewConnBeginTx(conn, func(c context.Context, o pgx.TxOptions) (*Tx, error) {
		if tx, err := conn.conn.BeginTx(c, o); err != nil {
			return nil, err
		} else {
			return &Tx{tx, conn.conn, conn.core, &conn.count}, nil
		}
	}).Proceed(ctx, txOptions)
}

func (conn *Conn) BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, fn func(*Tx) error) error {
	return NewConnBeginTxFunc(conn, func(c context.Context, o pgx.TxOptions, f func(*Tx) error) error {
		return conn.conn.BeginTxFunc(c, o, func(tx pgx.Tx) error {
			return f(&Tx{tx, conn.conn, conn.core, &conn.count})
		})
	}).Proceed(ctx, txOptions, fn)
}

func (conn *Conn) CopyFrom(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	return NewConnCopyFrom(conn, func(c context.Context, i Identifier, rs interface{}) (int64, error) {
		return doCopyFrom(conn.core, conn.conn.CopyFrom, c, i, rs)
	}).Proceed(ctx, tableName, rowSrc)
}

func (conn *Conn) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return NewConnExec(conn, func(c context.Context, q string, a ...interface{}) (pgconn.CommandTag, error) {
		return doExec(conn.core, conn.conn.Exec, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (conn *Conn) Ping(ctx context.Context) error {
	return NewConnPing(conn, conn.conn.Ping).Proceed(ctx)
}

func (conn *Conn) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return NewConnPrepare(conn, func(c context.Context, q string, e ...interface{}) (*Stmt, error) {
		return doPrepare(conn.core, conn.conn, &conn.count, conn.conn.Prepare, c, q, e...)
	}).Proceed(ctx, query, examples...)
}

func (conn *Conn) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return NewConnQuery(conn, func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(conn.core, conn.conn.Query, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (conn *Conn) SendBatch(ctx context.Context, batch *Batch) *BatchResults {
	return NewConnSendBatch(conn, func(c context.Context, b *Batch) *BatchResults {
		results := conn.conn.SendBatch(c, b.batch)
		return &BatchResults{results, conn.core}
	}).Proceed(ctx, batch)
}

func (conn *Conn) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewConnFind(conn, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(conn.core, conn.conn.Query, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

func (conn *Conn) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewConnFindAll(conn, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(conn.core, conn.conn.Query, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

type DB struct {
	pool  *pgxpool.Pool
	core  *Core
	count int64
}

func NewDB(db *pgxpool.Pool, core *Core) *DB {
	return &DB{db, core, 0}
}

func (db *DB) Pool() *pgxpool.Pool {
	return db.pool
}

func (db *DB) Close() error {
	return NewDBClose(db, func() error {
		db.pool.Close()
		return nil
	}).Proceed()
}

func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	return db.BeginTx(ctx, pgx.TxOptions{})
}

func (db *DB) BeginFunc(ctx context.Context, f func(*Tx) error) error {
	return db.BeginTxFunc(ctx, pgx.TxOptions{}, f)
}

func (db *DB) BeginTx(ctx context.Context, opts pgx.TxOptions) (*Tx, error) {
	return NewDBBeginTx(db, func(c context.Context, o pgx.TxOptions) (*Tx, error) {
		if tx, err := db.pool.BeginTx(c, o); err != nil {
			return nil, err
		} else {
			return &Tx{tx, tx.Conn(), db.core, &db.count}, nil
		}
	}).Proceed(ctx, opts)
}

func (db *DB) BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, fn func(*Tx) error) error {
	return NewDBBeginTxFunc(db, func(c context.Context, o pgx.TxOptions, f func(*Tx) error) error {
		return db.pool.BeginTxFunc(c, o, func(tx pgx.Tx) error {
			return f(&Tx{tx, tx.Conn(), db.core, &db.count})
		})
	}).Proceed(ctx, txOptions, fn)
}

func (db *DB) Ping(ctx context.Context) error {
	return NewDBPing(db, db.pool.Ping).Proceed(ctx)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return NewDBExec(db, func(c context.Context, q string, a ...interface{}) (pgconn.CommandTag, error) {
		return doExec(db.core, db.pool.Exec, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (db *DB) CopyFrom(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	return NewDBCopyFrom(db, func(c context.Context, i Identifier, rs interface{}) (int64, error) {
		return doCopyFrom(db.core, db.pool.CopyFrom, c, i, rs)
	}).Proceed(ctx, tableName, rowSrc)
}

func (db *DB) SendBatch(ctx context.Context, batch *Batch) *BatchResults {
	return NewDBSendBatch(db, func(c context.Context, b *Batch) *BatchResults {
		results := db.pool.SendBatch(ctx, batch.batch)
		return &BatchResults{results, db.core}
	}).Proceed(ctx, batch)
}

func (db *DB) Prepare(ctx context.Context, query string, examples ...interface{}) (*PooledStmt, error) {
	return NewDBPrepare(db, func(c context.Context, q string, e ...interface{}) (*PooledStmt, error) {
		conn, err := db.pool.Acquire(ctx)
		if err != nil {
			return nil, err
		}

		pooled := conn.Conn()

		if stmt, err := doPrepare(db.core, pooled, &db.count, pooled.Prepare, ctx, query, examples...); err != nil {
			return nil, err
		} else {
			return &PooledStmt{stmt, conn}, nil
		}
	}).Proceed(ctx, query, examples...)
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return NewDBQuery(db, func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(db.core, db.pool.Query, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (db *DB) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewDBFind(db, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(db.core, db.pool.Query, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

func (db *DB) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewDBFindAll(db, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(db.core, db.pool.Query, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

type Tx struct {
	tx    pgx.Tx
	conn  *pgx.Conn
	core  *Core
	count *int64
}

func (tx *Tx) Tx() pgx.Tx {
	return tx.tx
}

func (tx *Tx) Begin(ctx context.Context) (*Tx, error) {
	return NewTxBegin(tx, func(c context.Context) (*Tx, error) {
		if newone, err := tx.tx.Begin(c); err != nil {
			return nil, err
		} else {
			return &Tx{newone, tx.conn, tx.core, tx.count}, nil
		}
	}).Proceed(ctx)
}

func (tx *Tx) BeginFunc(ctx context.Context, fn func(*Tx) error) error {
	return NewTxBeginFunc(tx, func(c context.Context, f func(*Tx) error) error {
		return tx.tx.BeginFunc(c, func(newone pgx.Tx) error {
			return f(&Tx{newone, tx.conn, tx.core, tx.count})
		})
	}).Proceed(ctx, fn)
}

func (tx *Tx) Commit(ctx context.Context) error {
	return NewTxCommit(tx, tx.tx.Commit).Proceed(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	return NewTxRollback(tx, tx.tx.Rollback).Proceed(ctx)
}

func (tx *Tx) CopyFrom(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	return NewTxCopyFrom(tx, func(c context.Context, i Identifier, rs interface{}) (int64, error) {
		return doCopyFrom(tx.core, tx.tx.CopyFrom, c, i, rs)
	}).Proceed(ctx, tableName, rowSrc)
}

func (tx *Tx) SendBatch(ctx context.Context, batch *Batch) *BatchResults {
	return NewTxSendBatch(tx, func(c context.Context, b *Batch) *BatchResults {
		results := tx.tx.SendBatch(ctx, b.batch)
		return &BatchResults{results, tx.core}
	}).Proceed(ctx, batch)
}

func (tx *Tx) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return NewTxPrepare(tx, func(c context.Context, q string, e ...interface{}) (*Stmt, error) {
		return doPrepare(tx.core, tx.conn, tx.count, tx.tx.Prepare, c, q, e...)
	}).Proceed(ctx, query, examples...)
}

func (tx *Tx) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return NewTxExec(tx, func(c context.Context, q string, a ...interface{}) (pgconn.CommandTag, error) {
		return doExec(tx.core, tx.tx.Exec, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return NewTxQuery(tx, func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(tx.core, tx.tx.Query, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (tx *Tx) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewTxFind(tx, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(tx.core, tx.tx.Query, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

func (tx *Tx) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewTxFindAll(tx, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(tx.core, tx.tx.Query, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

type Stmt struct {
	stmt  *pgconn.StatementDescription
	conn  *pgx.Conn
	core  *Core
	query kra.QueryAnalyzer
}

func (stmt *Stmt) Stmt() *pgconn.StatementDescription {
	return stmt.stmt
}

const smalltime = 5

func (stmt *Stmt) Close(ctx context.Context) error {
	return NewStmtClose(stmt, func(c context.Context) error {
		x, cancel := context.WithTimeout(c, time.Second*smalltime)
		defer cancel()
		return stmt.conn.Deallocate(x, stmt.stmt.Name)
	}).Proceed(ctx)
}

func (stmt *Stmt) Exec(ctx context.Context, args ...interface{}) (pgconn.CommandTag, error) {
	return NewStmtExec(stmt, func(c context.Context, a ...interface{}) (pgconn.CommandTag, error) {
		if resolver, err := kra.NewCoreNewResolver(stmt.core.Core, stmt.core.hooks.Core).Proceed(a...); err != nil {
			return nil, err
		} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
			return nil, err
		} else {
			return stmt.conn.Exec(c, stmt.stmt.Name, bindArgs...)
		}
	}).Proceed(ctx, args...)
}

func (stmt *Stmt) Query(ctx context.Context, args ...interface{}) (*Rows, error) {
	return NewStmtQuery(stmt, func(c context.Context, a ...interface{}) (*Rows, error) {
		if resolver, err := kra.NewCoreNewResolver(stmt.core.Core, stmt.core.hooks.Core).Proceed(a...); err != nil {
			return nil, err
		} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
			return nil, err
		} else if rows, err := stmt.conn.Query(ctx, stmt.stmt.Name, bindArgs...); err != nil {
			return nil, err
		} else if rows.Err() != nil {
			return nil, rows.Err()
		} else {
			return NewRows(stmt.core, rows), nil
		}
	}).Proceed(ctx, args...)
}

type PooledStmt struct {
	delegate *Stmt
	conn     *pgxpool.Conn
}

func (stmt *PooledStmt) Stmt() *pgconn.StatementDescription {
	return stmt.delegate.stmt
}

func (stmt *PooledStmt) Close(ctx context.Context) error {
	defer stmt.conn.Release()
	return stmt.delegate.Close(ctx)
}

func (stmt *PooledStmt) Exec(ctx context.Context, args ...interface{}) (pgconn.CommandTag, error) {
	return stmt.delegate.Exec(ctx, args...)
}

func (stmt *PooledStmt) Query(ctx context.Context, args ...interface{}) (*Rows, error) {
	return stmt.delegate.Query(ctx, args...)
}

type Rows struct {
	rows        *rowsAdapter
	core        *Core
	transformer kra.Transformer
}

func NewRows(core *Core, rows pgx.Rows) *Rows {
	return &Rows{&rowsAdapter{rows}, core, kra.NewCoreNewTransformer(core.Core, core.hooks.Core).Proceed()}
}

func (rows *Rows) Next() bool {
	return NewRowsNext(rows, rows.rows.Next).Proceed()
}

func (rows *Rows) Err() error {
	return NewRowsErr(rows, rows.rows.Err).Proceed()
}

func (rows *Rows) Rows() pgx.Rows {
	return rows.rows.Rows
}

func (rows *Rows) Close() error {
	return NewRowsClose(rows, func() error {
		rows.rows.Close()
		return rows.rows.Err()
	}).Proceed()
}

func (rows *Rows) Scan(dest interface{}) error {
	return NewRowsScan(rows, func(d interface{}) error {
		return rows.transformer.Transform(rows.rows, d)
	}).Proceed(dest)
}

type rowsAdapter struct {
	pgx.Rows
}

func (adapeter *rowsAdapter) Columns() ([]string, error) {
	columns := make([]string, len(adapeter.Rows.FieldDescriptions()))
	for i, fd := range adapeter.Rows.FieldDescriptions() {
		columns[i] = string(fd.Name)
	}
	return columns, nil
}

func (adapeter *rowsAdapter) Close() error {
	adapeter.Rows.Close()
	return adapeter.Rows.Err()
}

type Batch struct {
	batch *pgx.Batch
	core  *Core
}

func (batch *Batch) Batch() *pgx.Batch {
	return batch.batch
}

func (batch *Batch) Queue(query string, args ...interface{}) error {
	return NewBatchQueue(batch, func(q string, a ...interface{}) error {
		if rawQuery, bindArgs, err := batch.core.Analyze(batch.core.hooks.Core, q, a...); err != nil {
			return err
		} else {
			batch.batch.Queue(rawQuery, bindArgs...)
			return nil
		}
	}).Proceed(query, args...)
}

type BatchResults struct {
	batchResults pgx.BatchResults
	core         *Core
}

func (batchResults *BatchResults) BatchResults() pgx.BatchResults {
	return batchResults.batchResults
}

func (batchResults *BatchResults) Close() error {
	return NewBatchResultsClose(batchResults, batchResults.batchResults.Close).Proceed()
}

func (batchResults *BatchResults) Exec() (pgconn.CommandTag, error) {
	return NewBatchResultsExec(batchResults, batchResults.batchResults.Exec).Proceed()
}

func (batchResults *BatchResults) Query() (*Rows, error) {
	return NewBatchResultsQuery(batchResults, func() (*Rows, error) {
		if rows, err := batchResults.batchResults.Query(); err != nil {
			return nil, err
		} else {
			return NewRows(batchResults.core, rows), nil
		}
	}).Proceed()
}
