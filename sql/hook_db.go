// Copyright 2022 taichi
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
)

type DBHook struct {
	BeginTx func(invocation *DBBeginTx, ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)
	Exec    func(invocation *DBExec, ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Ping    func(invocation *DBPing, ctx context.Context) error
	Prepare func(invocation *DBPrepare, ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
	Query   func(invocation *DBQuery, ctx context.Context, query string, args ...interface{}) (*Rows, error)
	Find    func(invocation *DBFind, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	FindAll func(invocation *DBFindAll, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Close   func(invocation *DBClose) error
}

type DBBeginTx invocation[DB, DBHook, func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)]

func NewDBBeginTx(recv *DB, original func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)) *DBBeginTx {
	hooks := recv.core.hooks.DB
	return &DBBeginTx{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBBeginTx) Proceed(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].BeginTx != nil {
		return invocation.hooks[current].BeginTx(invocation, ctx, txOptions)
	}
	return invocation.original(ctx, txOptions)
}

type DBExec invocation[DB, DBHook, func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)]

func NewDBExec(recv *DB, original func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)) *DBExec {
	hooks := recv.core.hooks.DB
	return &DBExec{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBExec) Proceed(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Exec != nil {
		return invocation.hooks[current].Exec(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type DBPing invocation[DB, DBHook, func(ctx context.Context) error]

func NewDBPing(recv *DB, original func(ctx context.Context) error) *DBPing {
	hooks := recv.core.hooks.DB
	return &DBPing{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBPing) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Ping != nil {
		return invocation.hooks[current].Ping(invocation, ctx)
	}
	return invocation.original(ctx)
}

type DBPrepare invocation[DB, DBHook, func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)]

func NewDBPrepare(recv *DB, original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)) *DBPrepare {
	hooks := recv.core.hooks.DB
	return &DBPrepare{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBPrepare) Proceed(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Prepare != nil {
		return invocation.hooks[current].Prepare(invocation, ctx, query, examples...)
	}
	return invocation.original(ctx, query, examples...)
}

type DBQuery invocation[DB, DBHook, func(ctx context.Context, query string, args ...interface{}) (*Rows, error)]

func NewDBQuery(recv *DB, original func(ctx context.Context, query string, args ...interface{}) (*Rows, error)) *DBQuery {
	hooks := recv.core.hooks.DB
	return &DBQuery{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBQuery) Proceed(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Query != nil {
		return invocation.hooks[current].Query(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type DBFind invocation[DB, DBHook, func(ctx context.Context, dest interface{}, query string, args ...interface{}) error]

func NewDBFind(recv *DB, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *DBFind {
	hooks := recv.core.hooks.DB
	return &DBFind{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBFind) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Find != nil {
		return invocation.hooks[current].Find(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type DBFindAll invocation[DB, DBHook, func(ctx context.Context, dest interface{}, query string, args ...interface{}) error]

func NewDBFindAll(recv *DB, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *DBFindAll {
	hooks := recv.core.hooks.DB
	return &DBFindAll{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBFindAll) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].FindAll != nil {
		return invocation.hooks[current].FindAll(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type DBClose invocation[DB, DBHook, func() error]

func NewDBClose(recv *DB, original func() error) *DBClose {
	hooks := recv.core.hooks.DB
	return &DBClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *DBClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
