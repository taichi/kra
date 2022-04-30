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

type ConnBeginTx invocation[Conn, ConnHook, func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)]

func NewConnBeginTx(recv *Conn, original func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)) *ConnBeginTx {
	hooks := recv.core.hooks.Conn
	return &ConnBeginTx{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnBeginTx) Proceed(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].BeginTx != nil {
		return invocation.hooks[current].BeginTx(invocation, ctx, txOptions)
	}
	return invocation.original(ctx, txOptions)
}

type ConnExec invocation[Conn, ConnHook, func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)]

func NewConnExec(recv *Conn, original func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)) *ConnExec {
	hooks := recv.core.hooks.Conn
	return &ConnExec{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnExec) Proceed(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Exec != nil {
		return invocation.hooks[current].Exec(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type ConnPing invocation[Conn, ConnHook, func(ctx context.Context) error]

func NewConnPing(recv *Conn, original func(ctx context.Context) error) *ConnPing {
	hooks := recv.core.hooks.Conn
	return &ConnPing{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnPing) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Ping != nil {
		return invocation.hooks[current].Ping(invocation, ctx)
	}
	return invocation.original(ctx)
}

type ConnPrepare invocation[Conn, ConnHook, func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)]

func NewConnPrepare(recv *Conn, original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)) *ConnPrepare {
	hooks := recv.core.hooks.Conn
	return &ConnPrepare{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnPrepare) Proceed(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Prepare != nil {
		return invocation.hooks[current].Prepare(invocation, ctx, query, examples...)
	}
	return invocation.original(ctx, query, examples...)
}

type ConnQuery invocation[Conn, ConnHook, func(ctx context.Context, query string, args ...interface{}) (*Rows, error)]

func NewConnQuery(recv *Conn, original func(ctx context.Context, query string, args ...interface{}) (*Rows, error)) *ConnQuery {
	hooks := recv.core.hooks.Conn
	return &ConnQuery{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnQuery) Proceed(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Query != nil {
		return invocation.hooks[current].Query(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type ConnFind invocation[Conn, ConnHook, func(ctx context.Context, dest interface{}, query string, args ...interface{}) error]

func NewConnFind(recv *Conn, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *ConnFind {
	hooks := recv.core.hooks.Conn
	return &ConnFind{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnFind) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Find != nil {
		return invocation.hooks[current].Find(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type ConnFindAll invocation[Conn, ConnHook, func(ctx context.Context, dest interface{}, query string, args ...interface{}) error]

func NewConnFindAll(recv *Conn, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *ConnFindAll {
	hooks := recv.core.hooks.Conn
	return &ConnFindAll{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnFindAll) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].FindAll != nil {
		return invocation.hooks[current].FindAll(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type ConnClose invocation[Conn, ConnHook, func() error]

func NewConnClose(recv *Conn, original func() error) *ConnClose {
	hooks := recv.core.hooks.Conn
	return &ConnClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
