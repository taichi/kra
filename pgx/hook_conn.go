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

package pgx

import (
	"context"

	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

type ConnHook struct {
	BeginTx     func(invocation *ConnBeginTx, ctx context.Context, txOptions pgx.TxOptions) (*Tx, error)
	BeginTxFunc func(invocation *ConnBeginTxFunc, ctx context.Context, txOptions pgx.TxOptions, fn func(*Tx) error) error
	CopyFrom    func(invocation *ConnCopyFrom, ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error)
	Exec        func(invocation *ConnExec, ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Ping        func(invocation *ConnPing, ctx context.Context) error
	Prepare     func(invocation *ConnPrepare, ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
	Query       func(invocation *ConnQuery, ctx context.Context, query string, args ...interface{}) (*Rows, error)
	SendBatch   func(invocation *ConnSendBatch, ctx context.Context, batch *Batch) *BatchResults
	Find        func(invocation *ConnFind, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	FindAll     func(invocation *ConnFindAll, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Close       func(invocation *ConnClose, ctx context.Context) error
}

type ConnBeginTx invocation[Conn, ConnHook, func(ctx context.Context, txOptions pgx.TxOptions) (*Tx, error)]

func NewConnBeginTx(recv *Conn, original func(ctx context.Context, txOptions pgx.TxOptions) (*Tx, error)) *ConnBeginTx {
	hooks := recv.core.hooks.Conn
	return &ConnBeginTx{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnBeginTx) Proceed(ctx context.Context, txOptions pgx.TxOptions) (*Tx, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].BeginTx != nil {
		return invocation.hooks[current].BeginTx(invocation, ctx, txOptions)
	}
	return invocation.original(ctx, txOptions)
}

type ConnBeginTxFunc invocation[Conn, ConnHook, func(ctx context.Context, txOptions pgx.TxOptions, fn func(*Tx) error) error]

func NewConnBeginTxFunc(recv *Conn, original func(ctx context.Context, txOptions pgx.TxOptions, fn func(*Tx) error) error) *ConnBeginTxFunc {
	hooks := recv.core.hooks.Conn
	return &ConnBeginTxFunc{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnBeginTxFunc) Proceed(ctx context.Context, txOptions pgx.TxOptions, fn func(*Tx) error) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].BeginTx != nil {
		return invocation.hooks[current].BeginTxFunc(invocation, ctx, txOptions, fn)
	}
	return invocation.original(ctx, txOptions, fn)
}

type ConnCopyFrom invocation[Conn, ConnHook, func(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error)]

func NewConnCopyFrom(recv *Conn, original func(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error)) *ConnCopyFrom {
	hooks := recv.core.hooks.Conn
	return &ConnCopyFrom{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnCopyFrom) Proceed(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].CopyFrom != nil {
		return invocation.hooks[current].CopyFrom(invocation, ctx, tableName, rowSrc)
	}
	return invocation.original(ctx, tableName, rowSrc)
}

type ConnExec invocation[Conn, ConnHook, func(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)]

func NewConnExec(recv *Conn, original func(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)) *ConnExec {
	hooks := recv.core.hooks.Conn
	return &ConnExec{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnExec) Proceed(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
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

type ConnSendBatch invocation[Conn, ConnHook, func(ctx context.Context, batch *Batch) *BatchResults]

func NewConnSendBatch(recv *Conn, original func(ctx context.Context, batch *Batch) *BatchResults) *ConnSendBatch {
	hooks := recv.core.hooks.Conn
	return &ConnSendBatch{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnSendBatch) Proceed(ctx context.Context, batch *Batch) *BatchResults {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].SendBatch != nil {
		return invocation.hooks[current].SendBatch(invocation, ctx, batch)
	}
	return invocation.original(ctx, batch)
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

type ConnClose invocation[Conn, ConnHook, func(ctx context.Context) error]

func NewConnClose(recv *Conn, original func(ctx context.Context) error) *ConnClose {
	hooks := recv.core.hooks.Conn
	return &ConnClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *ConnClose) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation, ctx)
	}
	return invocation.original(ctx)
}
