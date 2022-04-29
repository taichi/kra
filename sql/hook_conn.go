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

type ConnHook struct {
	BeginTx func(invocation *ConnBeginTx, ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)
	Exec    func(invocation *ConnExec, ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Ping    func(invocation *ConnPing, ctx context.Context) error
	Prepare func(invocation *ConnPrepare, ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
	Query   func(invocation *ConnQuery, ctx context.Context, query string, args ...interface{}) (*Rows, error)
	Find    func(invocation *ConnFind, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	FindAll func(invocation *ConnFindAll, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Close   func(invocation *ConnClose) error
}

func (hook *ConnHook) Fill() {
	if hook.BeginTx == nil {
		hook.BeginTx = func(invocation *ConnBeginTx, ctx context.Context, txOptions *sql.TxOptions) (*Tx, error) {
			return invocation.Proceed(ctx, txOptions)
		}
	}
	if hook.Exec == nil {
		hook.Exec = func(invocation *ConnExec, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
			return invocation.Proceed(ctx, query, args...)
		}
	}
	if hook.Ping == nil {
		hook.Ping = func(invocation *ConnPing, ctx context.Context) error {
			return invocation.Proceed(ctx)
		}
	}
	if hook.Prepare == nil {
		hook.Prepare = func(invocation *ConnPrepare, ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
			return invocation.Proceed(ctx, query, examples...)
		}
	}
	if hook.Query == nil {
		hook.Query = func(invocation *ConnQuery, ctx context.Context, query string, args ...interface{}) (*Rows, error) {
			return invocation.Proceed(ctx, query, args...)
		}
	}
	if hook.Find == nil {
		hook.Find = func(invocation *ConnFind, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
			return invocation.Proceed(ctx, dest, query, args...)
		}
	}
	if hook.FindAll == nil {
		hook.FindAll = func(invocation *ConnFindAll, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
			return invocation.Proceed(ctx, dest, query, args...)
		}
	}
	if hook.Close == nil {
		hook.Close = func(invocation *ConnClose) error {
			return invocation.Proceed()
		}
	}
}

type connInvocation struct {
	Receiver *Conn
	hooks    []*ConnHook
	length   int
	index    int
}

type ConnBeginTx struct {
	connInvocation
	original func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)
}

func NewConnBeginTx(recv *Conn, original func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)) *ConnBeginTx {
	hooks := recv.core.hooks.Conn
	return &ConnBeginTx{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnBeginTx) Proceed(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].BeginTx(invocation, ctx, txOptions)
	}
	return invocation.original(ctx, txOptions)
}

type ConnExec struct {
	connInvocation
	original func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func NewConnExec(recv *Conn, original func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)) *ConnExec {
	hooks := recv.core.hooks.Conn
	return &ConnExec{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnExec) Proceed(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Exec(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type ConnPing struct {
	connInvocation
	original func(ctx context.Context) error
}

func NewConnPing(recv *Conn, original func(ctx context.Context) error) *ConnPing {
	hooks := recv.core.hooks.Conn
	return &ConnPing{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnPing) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Ping(invocation, ctx)
	}
	return invocation.original(ctx)
}

type ConnPrepare struct {
	connInvocation
	original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
}

func NewConnPrepare(recv *Conn, original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)) *ConnPrepare {
	hooks := recv.core.hooks.Conn
	return &ConnPrepare{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnPrepare) Proceed(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Prepare(invocation, ctx, query, examples...)
	}
	return invocation.original(ctx, query, examples...)
}

type ConnQuery struct {
	connInvocation
	original func(ctx context.Context, query string, args ...interface{}) (*Rows, error)
}

func NewConnQuery(recv *Conn, original func(ctx context.Context, query string, args ...interface{}) (*Rows, error)) *ConnQuery {
	hooks := recv.core.hooks.Conn
	return &ConnQuery{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnQuery) Proceed(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Query(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type ConnFind struct {
	connInvocation
	original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func NewConnFind(recv *Conn, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *ConnFind {
	hooks := recv.core.hooks.Conn
	return &ConnFind{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnFind) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Find(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type ConnFindAll struct {
	connInvocation
	original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func NewConnFindAll(recv *Conn, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *ConnFindAll {
	hooks := recv.core.hooks.Conn
	return &ConnFindAll{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnFindAll) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].FindAll(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type ConnClose struct {
	connInvocation
	original func() error
}

func NewConnClose(recv *Conn, original func() error) *ConnClose {
	hooks := recv.core.hooks.Conn
	return &ConnClose{connInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *ConnClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
