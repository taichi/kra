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
)

type TxHook struct {
	Begin     func(invocation *TxBegin, ctx context.Context) (*Tx, error)
	BeginFunc func(invocation *TxBeginFunc, ctx context.Context, f func(*Tx) error) error
	Commit    func(invocation *TxCommit, ctx context.Context) error
	Rollback  func(invocation *TxRollback, ctx context.Context) error
	CopyFrom  func(invocation *TxCopyFrom, ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error)
	Exec      func(invocation *TxExec, ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Ping      func(invocation *TxPing, ctx context.Context) error
	Prepare   func(invocation *TxPrepare, ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
	Query     func(invocation *TxQuery, ctx context.Context, query string, args ...interface{}) (*Rows, error)
	SendBatch func(invocation *TxSendBatch, ctx context.Context, batch *Batch) *BatchResults
	Find      func(invocation *TxFind, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	FindAll   func(invocation *TxFindAll, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Close     func(invocation *TxClose, ctx context.Context) error
}

type TxBegin invocation[Tx, TxHook, func(ctx context.Context) (*Tx, error)]

func NewTxBegin(recv *Tx, original func(ctx context.Context) (*Tx, error)) *TxBegin {
	hooks := recv.core.hooks.Tx
	return &TxBegin{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxBegin) Proceed(ctx context.Context) (*Tx, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Begin != nil {
		return invocation.hooks[current].Begin(invocation, ctx)
	}
	return invocation.original(ctx)
}

type TxBeginFunc invocation[Tx, TxHook, func(ctx context.Context, fn func(*Tx) error) error]

func NewTxBeginFunc(recv *Tx, original func(ctx context.Context, fn func(*Tx) error) error) *TxBeginFunc {
	hooks := recv.core.hooks.Tx
	return &TxBeginFunc{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxBeginFunc) Proceed(ctx context.Context, fn func(*Tx) error) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].BeginFunc != nil {
		return invocation.hooks[current].BeginFunc(invocation, ctx, fn)
	}
	return invocation.original(ctx, fn)
}

type TxCommit invocation[Tx, TxHook, func(ctx context.Context) error]

func NewTxCommit(recv *Tx, original func(ctx context.Context) error) *TxCommit {
	hooks := recv.core.hooks.Tx
	return &TxCommit{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxCommit) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Commit != nil {
		return invocation.hooks[current].Commit(invocation, ctx)
	}
	return invocation.original(ctx)
}

type TxRollback invocation[Tx, TxHook, func(ctx context.Context) error]

func NewTxRollback(recv *Tx, original func(ctx context.Context) error) *TxRollback {
	hooks := recv.core.hooks.Tx
	return &TxRollback{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxRollback) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Rollback != nil {
		return invocation.hooks[current].Rollback(invocation, ctx)
	}
	return invocation.original(ctx)
}

type TxCopyFrom invocation[Tx, TxHook, func(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error)]

func NewTxCopyFrom(recv *Tx, original func(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error)) *TxCopyFrom {
	hooks := recv.core.hooks.Tx
	return &TxCopyFrom{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxCopyFrom) Proceed(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].CopyFrom != nil {
		return invocation.hooks[current].CopyFrom(invocation, ctx, tableName, rowSrc)
	}
	return invocation.original(ctx, tableName, rowSrc)
}

type TxExec invocation[Tx, TxHook, func(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)]

func NewTxExec(recv *Tx, original func(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)) *TxExec {
	hooks := recv.core.hooks.Tx
	return &TxExec{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxExec) Proceed(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
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

type TxSendBatch invocation[Tx, TxHook, func(ctx context.Context, batch *Batch) *BatchResults]

func NewTxSendBatch(recv *Tx, original func(ctx context.Context, batch *Batch) *BatchResults) *TxSendBatch {
	hooks := recv.core.hooks.Tx
	return &TxSendBatch{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxSendBatch) Proceed(ctx context.Context, batch *Batch) *BatchResults {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].SendBatch != nil {
		return invocation.hooks[current].SendBatch(invocation, ctx, batch)
	}
	return invocation.original(ctx, batch)
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

type TxClose invocation[Tx, TxHook, func(ctx context.Context) error]

func NewTxClose(recv *Tx, original func(ctx context.Context) error) *TxClose {
	hooks := recv.core.hooks.Tx
	return &TxClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *TxClose) Proceed(ctx context.Context) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation, ctx)
	}
	return invocation.original(ctx)
}
