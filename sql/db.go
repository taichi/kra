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

func Open(core *kra.Core, driverName, dataSourceName string) (*DB, error) {
	if db, err := sql.Open(driverName, dataSourceName); err != nil {
		return nil, err
	} else {
		return NewDB(db, core), nil
	}
}

type Conn struct {
	conn *sql.Conn
	core *kra.Core
}

func NewConn(conn *sql.Conn, core *kra.Core) *Conn {
	return &Conn{conn, core}
}

func (conn *Conn) Conn() *sql.Conn {
	return conn.conn
}

func (conn *Conn) Close() error {
	return conn.conn.Close()
}

func (conn *Conn) Begin(ctx context.Context) (*Tx, error) {
	return conn.BeginTx(ctx, nil)
}

func (conn *Conn) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if tx, err := conn.conn.BeginTx(ctx, opts); err != nil {
		return nil, err
	} else {
		return &Tx{tx, conn.core}, nil
	}
}

func (conn *Conn) Ping(ctx context.Context) error {
	return conn.conn.PingContext(ctx)
}

func (conn *Conn) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return doExec(conn.core, conn.conn.ExecContext, ctx, query, args...)
}

func (conn *Conn) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(conn.core, conn.conn.PrepareContext, ctx, query, examples...)
}

func (conn *Conn) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return doQuery(conn.core, conn.conn.QueryContext, ctx, query, args...)
}

func (conn *Conn) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFind(conn.core, conn.conn.QueryContext, ctx, dest, query, args...)
}

func (conn *Conn) FindAll(ctx context.Context, dest []interface{}, query string, args ...interface{}) error {
	return doFindAll(conn.core, conn.conn.QueryContext, ctx, dest, query, args...)
}

type DB struct {
	db   *sql.DB
	core *kra.Core
}

func NewDB(db *sql.DB, core *kra.Core) *DB {
	return &DB{db, core}
}

func (db *DB) DB() *sql.DB {
	return db.db
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	return db.BeginTx(ctx, nil)
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if tx, err := db.db.BeginTx(ctx, opts); err != nil {
		return nil, err
	} else {
		return &Tx{tx, db.core}, nil
	}
}

func (db *DB) Ping(ctx context.Context) error {
	return db.db.PingContext(ctx)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return doExec(db.core, db.db.ExecContext, ctx, query, args...)
}

func (db *DB) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(db.core, db.db.PrepareContext, ctx, query, examples...)
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return doQuery(db.core, db.db.QueryContext, ctx, query, args...)
}

func (db *DB) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFind(db.core, db.db.QueryContext, ctx, dest, query, args...)
}

func (db *DB) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFindAll(db.core, db.db.QueryContext, ctx, dest, query, args...)
}

type Tx struct {
	tx   *sql.Tx
	core *kra.Core
}

func (tx *Tx) Tx() *sql.Tx {
	return tx.tx
}

func (tx *Tx) Commit() error {
	return tx.tx.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.tx.Rollback()
}

func (tx *Tx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return doExec(tx.core, tx.tx.ExecContext, ctx, query, args...)
}

func (tx *Tx) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(tx.core, tx.tx.PrepareContext, ctx, query, examples...)
}

func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return doQuery(tx.core, tx.tx.QueryContext, ctx, query, args...)
}

func (tx *Tx) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return doFind(tx.core, tx.tx.QueryContext, ctx, dest, query, args...)
}

func (tx *Tx) FindAll(ctx context.Context, dest []interface{}, query string, args ...interface{}) error {
	return doFindAll(tx.core, tx.tx.QueryContext, ctx, dest, query, args...)
}

type Stmt struct {
	stmt  *sql.Stmt
	core  *kra.Core
	query kra.QueryAnalyzer
}

func (stmt *Stmt) Stmt() *sql.Stmt {
	return stmt.stmt
}

func (stmt *Stmt) Close() error {
	return stmt.stmt.Close()
}

func (stmt *Stmt) Exec(ctx context.Context, args ...interface{}) (sql.Result, error) {
	if resolver, err := stmt.core.NewResolver(args...); err != nil {
		return nil, err
	} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
		return nil, err
	} else {
		return stmt.stmt.Exec(bindArgs...)
	}
}

func (stmt *Stmt) Query(ctx context.Context, args ...interface{}) (*Rows, error) {
	if resolver, err := stmt.core.NewResolver(args...); err != nil {
		return nil, err
	} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
		return nil, err
	} else if rows, err := stmt.stmt.Query(bindArgs...); err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	} else {
		return &Rows{rows, stmt.core.NewTransformer()}, nil
	}
}

type Rows struct {
	rows        *sql.Rows
	transformer kra.Transformer
}

func (rows *Rows) Rows() *sql.Rows {
	return rows.rows
}

func (rows *Rows) Close() error {
	return rows.rows.Close()
}

func (rows *Rows) Scan(dest interface{}) error {
	return rows.transformer.Transform(rows.rows, dest)
}
