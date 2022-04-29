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
	return conn.core.hook.Conn.Close(conn.conn.Close)
}

func (conn *Conn) Begin(ctx context.Context) (*Tx, error) {
	return conn.BeginTx(ctx, nil)
}

func (conn *Conn) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	return conn.core.hook.BeginTx(func(c context.Context, o *sql.TxOptions) (*Tx, error) {
		if tx, err := conn.conn.BeginTx(c, o); err != nil {
			return nil, err
		} else {
			return &Tx{tx, conn.core}, nil
		}
	}, ctx, opts)
}

func (conn *Conn) Ping(ctx context.Context) error {
	return conn.core.hook.Ping(conn.conn.PingContext, ctx)
}

func (conn *Conn) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return doExec(conn.core, conn.conn.ExecContext, ctx, query, args...)
}

func (conn *Conn) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(conn.core, conn.conn.PrepareContext, ctx, query, examples...)
}

func (conn *Conn) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return conn.core.hook.Query(func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(conn.core, conn.conn.QueryContext, c, q, a...)
	}, ctx, query, args...)
}

func (conn *Conn) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return conn.core.hook.Find(func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(conn.core, conn.conn.QueryContext, c, d, q, a...)
	}, ctx, dest, query, args...)
}

func (conn *Conn) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return conn.core.hook.FindAll(func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(conn.core, conn.conn.QueryContext, c, d, q, a...)
	}, ctx, dest, query, args...)
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
	return db.core.hook.DB.Close(db.db.Close)
}

func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	return db.BeginTx(ctx, &sql.TxOptions{})
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	return db.core.hook.BeginTx(func(c context.Context, o *sql.TxOptions) (*Tx, error) {
		if tx, err := db.db.BeginTx(c, o); err != nil {
			return nil, err
		} else {
			return &Tx{tx, db.core}, nil
		}
	}, ctx, opts)
}

func (db *DB) Ping(ctx context.Context) error {
	return db.core.hook.Ping(db.db.PingContext, ctx)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return doExec(db.core, db.db.ExecContext, ctx, query, args...)
}

func (db *DB) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(db.core, db.db.PrepareContext, ctx, query, examples...)
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return db.core.hook.Query(func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(db.core, db.db.QueryContext, c, q, a...)
	}, ctx, query, args...)
}

func (db *DB) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.core.hook.Find(func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(db.core, db.db.QueryContext, c, d, q, a...)
	}, ctx, dest, query, args...)
}

func (db *DB) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.core.hook.FindAll(func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(db.core, db.db.QueryContext, c, d, q, a...)
	}, ctx, dest, query, args...)
}

type Tx struct {
	tx   *sql.Tx
	core *Core
}

func (tx *Tx) Tx() *sql.Tx {
	return tx.tx
}

func (tx *Tx) Commit() error {
	return tx.core.hook.Tx.Commit(tx.tx.Commit)
}

func (tx *Tx) Rollback() error {
	return tx.core.hook.Tx.Rollback(tx.tx.Rollback)
}

func (tx *Tx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return doExec(tx.core, tx.tx.ExecContext, ctx, query, args...)
}

func (tx *Tx) Prepare(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	return doPrepare(tx.core, tx.tx.PrepareContext, ctx, query, examples...)
}

func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return tx.core.hook.Query(func(c context.Context, q string, a ...interface{}) (*Rows, error) {
		return doQuery(tx.core, tx.tx.QueryContext, c, q, a...)
	}, ctx, query, args...)
}

func (tx *Tx) Find(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return tx.core.hook.Find(func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFind(tx.core, tx.tx.QueryContext, c, d, q, a...)
	}, ctx, dest, query, args...)
}

func (tx *Tx) FindAll(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return tx.core.hook.FindAll(func(c context.Context, d interface{}, q string, a ...interface{}) error {
		return doFindAll(tx.core, tx.tx.QueryContext, c, d, q, a...)
	}, ctx, dest, query, args...)
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
	return stmt.core.hook.Stmt.Close(stmt.stmt.Close)
}

func (stmt *Stmt) Exec(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return stmt.core.hook.Stmt.Exec(func(c context.Context, a ...interface{}) (sql.Result, error) {
		if resolver, err := stmt.core.NewResolver(a...); err != nil {
			return nil, err
		} else if _, bindArgs, err := stmt.query.Analyze(resolver); err != nil {
			return nil, err
		} else {
			return stmt.stmt.Exec(bindArgs...)
		}
	}, ctx, args...)
}

func (stmt *Stmt) Query(ctx context.Context, args ...interface{}) (*Rows, error) {
	return stmt.core.hook.Stmt.Query(func(c context.Context, a ...interface{}) (*Rows, error) {
		if resolver, err := stmt.core.NewResolver(a...); err != nil {
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
	}, ctx, args...)
}

type Rows struct {
	rows        *sql.Rows
	core        *Core
	transformer kra.Transformer
}

func NewRows(core *Core, rows *sql.Rows) *Rows {
	return &Rows{rows, core, core.hook.NewTransformer(core.NewTransformer)}
}

func (rows *Rows) Rows() *sql.Rows {
	return rows.rows
}

func (rows *Rows) Close() error {
	return rows.core.hook.Rows.Close(rows.rows.Close)
}

func (rows *Rows) Scan(dest interface{}) error {
	return rows.core.hook.Rows.Scan(func(d interface{}) error {
		return rows.transformer.Transform(rows.rows, d)
	}, dest)
}
