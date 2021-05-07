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

func Open(ctx context.Context, connString string) (*DB, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	return OpenConfig(ctx, config)
}

type Identifier pgx.Identifier

func OpenConfig(ctx context.Context, config *pgxpool.Config) (*DB, error) {
	if pool, err := pgxpool.ConnectConfig(ctx, config); err != nil {
		return nil, err
	} else {
		return NewDB(pool, kra.NewCore(kra.PostgreSQL)), nil
	}
}

type Conn struct {
	conn *pgx.Conn
	core *kra.Core
}

func Connect(ctx context.Context, connString string) (*Conn, error) {
	if conn, err := pgx.Connect(ctx, connString); err != nil {
		return nil, err
	} else {
		return NewConn(conn, kra.NewCore(kra.PostgreSQL)), nil
	}
}

func NewConn(conn *pgx.Conn, core *kra.Core) *Conn {
	return &Conn{conn, core}
}

func (conn *Conn) Conn() *pgx.Conn {
	return conn.conn
}

func (conn *Conn) Close(ctx context.Context) error {
	return conn.conn.Close(ctx)
}

func (conn *Conn) Begin(ctx context.Context) (*Tx, error) {
	if tx, err := conn.conn.Begin(ctx); err != nil {
		return nil, err
	} else {
		return &Tx{tx, conn.conn, conn.core}, nil
	}
}

func (conn *Conn) BeginFunc(ctx context.Context, f func(*Tx) error) error {
	return conn.conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		return f(&Tx{tx, conn.conn, conn.core})
	})
}

func (conn *Conn) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (*Tx, error) {
	if tx, err := conn.conn.BeginTx(ctx, txOptions); err != nil {
		return nil, err
	} else {
		return &Tx{tx, conn.conn, conn.core}, nil
	}
}

func (conn *Conn) BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(*Tx) error) error {
	return conn.conn.BeginTxFunc(ctx, txOptions, func(tx pgx.Tx) error {
		return f(&Tx{tx, conn.conn, conn.core})
	})
}

func (conn *Conn) CopyFrom(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	return doCopyFrom(conn.core, conn.conn.CopyFrom, ctx, tableName, rowSrc)
}

func (conn *Conn) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return doExec(conn.core, conn.conn.Exec, ctx, query, args...)
}

func (conn *Conn) Ping(ctx context.Context) error {
	return conn.conn.Ping(ctx)
}

func (conn *Conn) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(conn.core, conn.conn, conn.conn.Prepare, ctx, query, examples...)
}

func (conn *Conn) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return doQuery(conn.core, conn.conn.Query, ctx, query, args...)
}

func (conn *Conn) SendBatch(ctx context.Context, batch *Batch) *BatchResults {
	results := conn.conn.SendBatch(ctx, batch.batch)
	return &BatchResults{results, conn.core}
}

func (conn *Conn) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFind(conn.core, conn.conn.Query, ctx, dest, query, args...)
}

func (conn *Conn) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFindAll(conn.core, conn.conn.Query, ctx, dest, query, args...)
}

type DB struct {
	pool *pgxpool.Pool
	core *kra.Core
}

func NewDB(db *pgxpool.Pool, core *kra.Core) *DB {
	return &DB{db, core}
}

func (db *DB) Pool() *pgxpool.Pool {
	return db.pool
}

func (db *DB) Close() error {
	db.pool.Close()
	return nil
}

func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	return db.BeginTx(ctx, pgx.TxOptions{})
}

func (db *DB) BeginTx(ctx context.Context, opts pgx.TxOptions) (*Tx, error) {
	if tx, err := db.pool.BeginTx(ctx, opts); err != nil {
		return nil, err
	} else {
		return &Tx{tx, tx.Conn(), db.core}, nil
	}
}

func (db *DB) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return doExec(db.core, db.pool.Exec, ctx, query, args...)
}

func (db *DB) CopyFrom(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	return doCopyFrom(db.core, db.pool.CopyFrom, ctx, tableName, rowSrc)
}

func (db *DB) SendBatch(ctx context.Context, batch *Batch) *BatchResults {
	results := db.pool.SendBatch(ctx, batch.batch)
	return &BatchResults{results, db.core}
}

func (db *DB) Prepare(ctx context.Context, query string, examples ...interface{}) (*PooledStmt, error) {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	pooled := conn.Conn()

	if stmt, err := doPrepare(db.core, pooled, pooled.Prepare, ctx, query, examples...); err != nil {
		return nil, err
	} else {
		return &PooledStmt{stmt, conn}, nil
	}
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return doQuery(db.core, db.pool.Query, ctx, query, args...)
}

func (db *DB) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFind(db.core, db.pool.Query, ctx, dest, query, args...)
}

func (db *DB) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFindAll(db.core, db.pool.Query, ctx, dest, query, args...)
}

type Tx struct {
	tx   pgx.Tx
	conn *pgx.Conn
	core *kra.Core
}

func (tx *Tx) Tx() pgx.Tx {
	return tx.tx
}

func (tx *Tx) Begin(ctx context.Context) (*Tx, error) {
	if newone, err := tx.tx.Begin(ctx); err != nil {
		return nil, err
	} else {
		return &Tx{newone, tx.conn, tx.core}, nil
	}
}

func (tx *Tx) BeginFunc(ctx context.Context, f func(*Tx) error) error {
	return tx.tx.BeginFunc(ctx, func(newone pgx.Tx) error {
		return f(&Tx{newone, tx.conn, tx.core})
	})
}

func (tx *Tx) Commit(ctx context.Context) error {
	return tx.tx.Commit(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	return tx.tx.Rollback(ctx)
}

func (tx *Tx) CopyFrom(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	return doCopyFrom(tx.core, tx.tx.CopyFrom, ctx, tableName, rowSrc)
}

func (tx *Tx) SendBatch(ctx context.Context, batch *Batch) *BatchResults {
	results := tx.tx.SendBatch(ctx, batch.batch)
	return &BatchResults{results, tx.core}
}

func (tx *Tx) LargeObjects() pgx.LargeObjects {
	return tx.tx.LargeObjects()
}

func (tx *Tx) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(tx.core, tx.conn, tx.tx.Prepare, ctx, query, examples...)
}

func (tx *Tx) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return doExec(tx.core, tx.tx.Exec, ctx, query, args...)
}

func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return doQuery(tx.core, tx.tx.Query, ctx, query, args...)
}

func (tx *Tx) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFind(tx.core, tx.tx.Query, ctx, dest, query, args...)
}

func (tx *Tx) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFindAll(tx.core, tx.tx.Query, ctx, dest, query, args...)
}

type Stmt struct {
	stmt  *pgconn.StatementDescription
	conn  *pgx.Conn
	core  *kra.Core
	query kra.QueryAnalyzer
}

func (stmt *Stmt) Stmt() *pgconn.StatementDescription {
	return stmt.stmt
}

const smalltime = 5

func (stmt *Stmt) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*smalltime)
	defer cancel()
	return stmt.conn.Deallocate(ctx, stmt.stmt.Name)
}

func (stmt *Stmt) Exec(ctx context.Context, args ...interface{}) (pgconn.CommandTag, error) {
	if resolver, err := stmt.core.NewResolver(args...); err != nil {
		return nil, err
	} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
		return nil, err
	} else {
		return stmt.conn.Exec(ctx, stmt.stmt.Name, bindArgs...)
	}
}

func (stmt *Stmt) Query(ctx context.Context, args ...interface{}) (*Rows, error) {
	if resolver, err := stmt.core.NewResolver(args...); err != nil {
		return nil, err
	} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
		return nil, err
	} else if rows, err := stmt.conn.Query(ctx, stmt.stmt.Name, bindArgs...); err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	} else {
		return &Rows{rows, stmt.core.NewTransformer()}, nil
	}
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
	rows        pgx.Rows
	transformer kra.Transformer
}

func (rows *Rows) Next() bool {
	return rows.rows.Next()
}

func (rows *Rows) Err() error {
	return rows.rows.Err()
}

func (rows *Rows) Rows() pgx.Rows {
	return rows.rows
}

func (rows *Rows) Close() error {
	rows.rows.Close()
	return rows.rows.Err()
}

func (rows *Rows) Scan(dest interface{}) error {
	return rows.transformer.Transform(&rowsAdapter{rows.rows}, dest)
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
	core  *kra.Core
}

func (batch *Batch) Batch() *pgx.Batch {
	return batch.batch
}

func (batch *Batch) Queue(query string, args ...interface{}) error {
	if rawQuery, bindArgs, err := batch.core.Analyze(query, args...); err != nil {
		return err
	} else {
		batch.batch.Queue(rawQuery, bindArgs...)
		return nil
	}
}

type BatchResults struct {
	batchResults pgx.BatchResults
	core         *kra.Core
}

func (batchResults *BatchResults) BatchResults() pgx.BatchResults {
	return batchResults.batchResults
}

func (batchResults *BatchResults) Close() error {
	return batchResults.batchResults.Close()
}

func (batchResults *BatchResults) Exec() (pgconn.CommandTag, error) {
	return batchResults.batchResults.Exec()
}

func (batchResults *BatchResults) Query() (*Rows, error) {
	if rows, err := batchResults.batchResults.Query(); err != nil {
		return nil, err
	} else {
		return &Rows{rows, batchResults.core.NewTransformer()}, nil
	}
}
