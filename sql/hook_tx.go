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

type TxHook struct {
	BeginTx  func(invocation *TxBeginTx, ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)
	Exec     func(invocation *TxExec, ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Ping     func(invocation *TxPing, ctx context.Context) error
	Prepare  func(invocation *TxPrepare, ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
	Query    func(invocation *TxQuery, ctx context.Context, query string, args ...interface{}) (*Rows, error)
	Find     func(invocation *TxFind, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	FindAll  func(invocation *TxFindAll, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Commit   func(invocation *TxCommit) error
	Rollback func(invocation *TxRollback) error
	Close    func(invocation *TxClose) error
}

type TxBeginTx invocation[Tx, TxHook, func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)]

func NewTxBeginTx(recv *Tx, original func(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error)) *TxBeginTx {
	hooks := recv.core.hooks.Tx
	return &TxBeginTx{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxBeginTx) Proceed(ctx context.Context, txOptions *sql.TxOptions) (*Tx, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].BeginTx != nil {
		return invocation.hooks[current].BeginTx(invocation, ctx, txOptions)
	}
	return invocation.original(ctx, txOptions)
}

type TxExec invocation[Tx, TxHook, func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)]

func NewTxExec(recv *Tx, original func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)) *TxExec {
	hooks := recv.core.hooks.Tx
	return &TxExec{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxExec) Proceed(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Exec != nil {
		return invocation.hooks[current].Exec(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type TxPing invocation[Tx, TxHook, func(ctx context.Context) error]

func NewTxPing(recv *Tx, original func(ctx context.Context) error) *TxPing {
	hooks := recv.core.hooks.Tx
	return &TxPing{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxPing) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Ping != nil {
		return invocation.hooks[current].Ping(invocation, ctx)
	}
	return invocation.original(ctx)
}

type TxPrepare invocation[Tx, TxHook, func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)]

func NewTxPrepare(recv *Tx, original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error)) *TxPrepare {
	hooks := recv.core.hooks.Tx
	return &TxPrepare{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxPrepare) Proceed(ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Prepare != nil {
		return invocation.hooks[current].Prepare(invocation, ctx, query, examples...)
	}
	return invocation.original(ctx, query, examples...)
}

type TxQuery invocation[Tx, TxHook, func(ctx context.Context, query string, args ...interface{}) (*Rows, error)]

func NewTxQuery(recv *Tx, original func(ctx context.Context, query string, args ...interface{}) (*Rows, error)) *TxQuery {
	hooks := recv.core.hooks.Tx
	return &TxQuery{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxQuery) Proceed(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Query != nil {
		return invocation.hooks[current].Query(invocation, ctx, query, args...)
	}
	return invocation.original(ctx, query, args...)
}

type TxFind invocation[Tx, TxHook, func(ctx context.Context, dest interface{}, query string, args ...interface{}) error]

func NewTxFind(recv *Tx, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *TxFind {
	hooks := recv.core.hooks.Tx
	return &TxFind{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxFind) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Find != nil {
		return invocation.hooks[current].Find(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type TxFindAll invocation[Tx, TxHook, func(ctx context.Context, dest interface{}, query string, args ...interface{}) error]

func NewTxFindAll(recv *Tx, original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error) *TxFindAll {
	hooks := recv.core.hooks.Tx
	return &TxFindAll{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxFindAll) Proceed(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].FindAll != nil {
		return invocation.hooks[current].FindAll(invocation, ctx, dest, query, args...)
	}
	return invocation.original(ctx, dest, query, args...)
}

type TxClose invocation[Tx, TxHook, func() error]

func NewTxClose(recv *Tx, original func() error) *TxClose {
	hooks := recv.core.hooks.Tx
	return &TxClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}

type TxCommit invocation[Tx, TxHook, func() error]

func NewTxCommit(recv *Tx, original func() error) *TxCommit {
	hooks := recv.core.hooks.Tx
	return &TxCommit{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxCommit) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Commit != nil {
		return invocation.hooks[current].Commit(invocation)
	}
	return invocation.original()
}

type TxRollback invocation[Tx, TxHook, func() error]

func NewTxRollback(recv *Tx, original func() error) *TxRollback {
	hooks := recv.core.hooks.Tx
	return &TxRollback{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxRollback) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Rollback != nil {
		return invocation.hooks[current].Rollback(invocation)
	}
	return invocation.original()
}
