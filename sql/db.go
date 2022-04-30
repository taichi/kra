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

package sql

import (
	"context"
	"database/sql"

	"github.com/taichi/kra"
)

func Open(core *Core, driverName, dataSourceName string) (*DB, error) {
	if db, err := sql.Open(driverName, dataSourceName); err != nil {
		return nil, err
	} else {
		return NewDB(db, core), nil
	}
}

type Conn struct {
	conn *sql.Conn
	core *Core
}

func NewConn(conn *sql.Conn, core *Core) *Conn {
	return &Conn{conn, core}
}

func (conn *Conn) Conn() *sql.Conn {
	return conn.conn
}

func (conn *Conn) Close() error {
	return NewConnClose(conn, conn.conn.Close).Proceed()
}

func (conn *Conn) Begin(ctx context.Context) (*Tx, error) {
	return conn.BeginTx(ctx, nil)
}

func (conn *Conn) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	return NewConnBeginTx(conn, func(c context.Context, o *sql.TxOptions) (*Tx, error) {
		if tx, err := conn.conn.BeginTx(c, o); err != nil {
			return nil, err
		} else {
			return &Tx{tx, conn.core}, nil
		}
	}).Proceed(ctx, opts)
}

func (conn *Conn) Ping(ctx context.Context) error {
	return NewConnPing(conn, conn.conn.PingContext).Proceed(ctx)
}

func (conn *Conn) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return NewConnExec(conn, func(c context.Context, q string, a ...interface{}) (sql.Result, error) {
		return doExec(conn.core, conn.conn.ExecContext, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (conn *Conn) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return NewConnPrepare(conn, func(c context.Context, q string, e ...interface{}) (*Stmt, error) {
		return doPrepare(conn.core, conn.conn.PrepareContext, c, q, e...)
	}).Proceed(ctx, query, examples...)
}

func (conn *Conn) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return NewConnQuery(conn, func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(conn.core, conn.conn.QueryContext, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (conn *Conn) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewConnFind(conn, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(conn.core, conn.conn.QueryContext, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

func (conn *Conn) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewConnFindAll(conn, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(conn.core, conn.conn.QueryContext, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

type DB struct {
	db   *sql.DB
	core *Core
}

func NewDB(db *sql.DB, core *Core) *DB {
	return &DB{db, core}
}

func (db *DB) DB() *sql.DB {
	return db.db
}

func (db *DB) Conn(ctx context.Context) (*Conn, error) {
	raw, err := db.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	return NewConn(raw, db.core), nil
}

func (db *DB) Close() error {
	return NewDBClose(db, db.db.Close).Proceed()
}

func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	return db.BeginTx(ctx, &sql.TxOptions{})
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	return NewDBBeginTx(db, func(c context.Context, o *sql.TxOptions) (*Tx, error) {
		if tx, err := db.db.BeginTx(c, o); err != nil {
			return nil, err
		} else {
			return &Tx{tx, db.core}, nil
		}
	}).Proceed(ctx, opts)
}

func (db *DB) Ping(ctx context.Context) error {
	return NewDBPing(db, db.db.PingContext).Proceed(ctx)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return NewDBExec(db, func(c context.Context, q string, a ...interface{}) (sql.Result, error) {
		return doExec(db.core, db.db.ExecContext, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (db *DB) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return NewDBPrepare(db, func(c context.Context, q string, e ...interface{}) (*Stmt, error) {
		return doPrepare(db.core, db.db.PrepareContext, c, q, e...)
	}).Proceed(ctx, query, examples...)
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return NewDBQuery(db, func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(db.core, db.db.QueryContext, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (db *DB) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewDBFind(db, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(db.core, db.db.QueryContext, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

func (db *DB) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewDBFindAll(db, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(db.core, db.db.QueryContext, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

type Tx struct {
	tx   *sql.Tx
	core *Core
}

func (tx *Tx) Tx() *sql.Tx {
	return tx.tx
}

func (tx *Tx) Commit() error {
	return NewTxCommit(tx, tx.tx.Commit).Proceed()
}

func (tx *Tx) Rollback() error {
	return NewTxRollback(tx, tx.tx.Rollback).Proceed()
}

func (tx *Tx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return NewTxExec(tx, func(c context.Context, q string, a ...interface{}) (sql.Result, error) {
		return doExec(tx.core, tx.tx.ExecContext, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (tx *Tx) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return NewTxPrepare(tx, func(c context.Context, q string, e ...interface{}) (*Stmt, error) {
		return doPrepare(tx.core, tx.tx.PrepareContext, c, q, e...)
	}).Proceed(ctx, query, examples...)
}

func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return NewTxQuery(tx, func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(tx.core, tx.tx.QueryContext, c, q, a...)
	}).Proceed(ctx, query, args...)
}

func (tx *Tx) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewTxFind(tx, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(tx.core, tx.tx.QueryContext, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

func (tx *Tx) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return NewTxFindAll(tx, func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(tx.core, tx.tx.QueryContext, c, d, q, a...)
	}).Proceed(ctx, dest, query, args...)
}

type Stmt struct {
	stmt  *sql.Stmt
	core  *Core
	query kra.QueryAnalyzer
}

func (stmt *Stmt) Stmt() *sql.Stmt {
	return stmt.stmt
}

func (stmt *Stmt) Close() error {
	return NewStmtClose(stmt, stmt.stmt.Close).Proceed()
}

func (stmt *Stmt) Exec(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return NewStmtExec(stmt, func(c context.Context, a ...interface{}) (sql.Result, error) {
		if resolver, err := kra.NewCoreNewResolver(stmt.core.Core, stmt.core.hooks.Core).Proceed(a...); err != nil {
			return nil, err
		} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
			return nil, err
		} else {
			return stmt.stmt.Exec(bindArgs...)
		}
	}).Proceed(ctx, args...)
}

func (stmt *Stmt) Query(ctx context.Context, args ...interface{}) (*Rows, error) {
	return NewStmtQuery(stmt, func(c context.Context, a ...interface{}) (*Rows, error) {
		if resolver, err := kra.NewCoreNewResolver(stmt.core.Core, stmt.core.hooks.Core).Proceed(a...); err != nil {
			return nil, err
		} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
			return nil, err
		} else if rows, err := stmt.stmt.Query(bindArgs...); err != nil {
			return nil, err
		} else if rows.Err() != nil {
			return nil, rows.Err()
		} else {
			return NewRows(stmt.core, rows), nil
		}
	}).Proceed(ctx, args...)
}

type Rows struct {
	rows        *sql.Rows
	core        *Core
	transformer kra.Transformer
}

func NewRows(core *Core, rows *sql.Rows) *Rows {
	return &Rows{rows, core, kra.NewCoreNewTransformer(core.Core, core.hooks.Core).Proceed()}
}

func (rows *Rows) Rows() *sql.Rows {
	return rows.rows
}

func (rows *Rows) Close() error {
	return NewRowsClose(rows, rows.rows.Close).Proceed()
}

func (rows *Rows) Scan(dest interface{}) error {
	return NewRowsScan(rows, func(d interface{}) error {
		return rows.transformer.Transform(rows.rows, d)
	}).Proceed(dest)
}
