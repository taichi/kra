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

func (hook *DBHook) Fill() {
	if hook.BeginTx == nil {
		hook.BeginTx = func(invocation *DBBeginTx, ctx context.Context, txOptions *sql.TxOptions) (*Tx, error) {
			return invocation.Proceed(ctx, txOptions)
		}
	}
	if hook.Exec == nil {
		hook.Exec = func(invocation *DBExec, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
			return invocation.Proceed(ctx, query, args...)
		}
	}
	if hook.Ping == nil {
		hook.Ping = func(invocation *DBPing, ctx context.Context) error {
			return invocation.Proceed(ctx)
		}
	}
	if hook.Prepare == nil {
		hook.Prepare = func(invocation *DBPrepare, ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
			return invocation.Proceed(ctx, query, examples...)
		}
	}
	if hook.Query == nil {
		hook.Query = func(invocation *DBQuery, ctx context.Context, query string, args ...interface{}) (*Rows, error) {
			return invocation.Proceed(ctx, query, args...)
		}
	}
	if hook.Find == nil {
		hook.Find = func(invocation *DBFind, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
			return invocation.Proceed(ctx, dest, query, args...)
		}
	}
	if hook.FindAll == nil {
		hook.FindAll = func(invocation *DBFindAll, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
			return invocation.Proceed(ctx, dest, query, args...)
		}
	}
	if hook.Close == nil {
		hook.Close = func(invocation *DBClose) error {
			return invocation.Proceed()
		}
	}
}

type dbInvocation struct {
	Receiver *DB
	hooks    []*DBHook
	length   int
	index    int
}

type DBBeginTx struct {
	dbInvocation
	original func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)
}

func NewDBBeginTx(recv *DB, original func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)) *DBBeginTx {
	hooks := recv.core.hooks.DB
	return &DBBeginTx{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBBeginTx) Proceed(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].BeginTx(invocation, ctx, txOptions)
	}
	return invocation.original(ctx, txOptions)
}

type DBExec struct {
	dbInvocation
	original func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func NewDBExec(recv *DB, original func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)) *DBExec {
	hooks := recv.core.hooks.DB
	return &DBExec{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBExec) Proceed(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Exec(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type DBPing struct {
	dbInvocation
	original func(ctx context.Context) error
}

func NewDBPing(recv *DB, original func(ctx context.Context) error) *DBPing {
	hooks := recv.core.hooks.DB
	return &DBPing{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBPing) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Ping(invocation, ctx)
	}
	return invocation.original(ctx)
}

type DBPrepare struct {
	dbInvocation
	original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
}

func NewDBPrepare(recv *DB, original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)) *DBPrepare {
	hooks := recv.core.hooks.DB
	return &DBPrepare{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBPrepare) Proceed(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Prepare(invocation, ctx, query, examples...)
	}
	return invocation.original(ctx, query, examples...)
}

type DBQuery struct {
	dbInvocation
	original func(ctx context.Context, query string, args ...interface{}) (*Rows, error)
}

func NewDBQuery(recv *DB, original func(ctx context.Context, query string, args ...interface{}) (*Rows, error)) *DBQuery {
	hooks := recv.core.hooks.DB
	return &DBQuery{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBQuery) Proceed(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Query(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type DBFind struct {
	dbInvocation
	original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func NewDBFind(recv *DB, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *DBFind {
	hooks := recv.core.hooks.DB
	return &DBFind{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBFind) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Find(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type DBFindAll struct {
	dbInvocation
	original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func NewDBFindAll(recv *DB, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *DBFindAll {
	hooks := recv.core.hooks.DB
	return &DBFindAll{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBFindAll) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].FindAll(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type DBClose struct {
	dbInvocation
	original func() error
}

func NewDBClose(recv *DB, original func() error) *DBClose {
	hooks := recv.core.hooks.DB
	return &DBClose{dbInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *DBClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
