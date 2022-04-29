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

	"github.com/taichi/kra"
)

type HookHolster struct {
	Core []*kra.CoreHook
	Conn []*ConnHook
	DB   []*DBHook
	Tx   []*TxHook
	Stmt []*StmtHook
	Rows []*RowsHook
}

func NewHookHolster(hooks ...kra.Hook) *HookHolster {
	holster := new(HookHolster)
	for _, hook := range hooks {
		if hook == nil {
			continue
		}
		hook.Fill()
		switch h := hook.(type) {
		case *kra.CoreHook:
			holster.Core = append(holster.Core, h)
		case *ConnHook:
			holster.Conn = append(holster.Conn, h)
		case *DBHook:
			holster.DB = append(holster.DB, h)
		case *TxHook:
			holster.Tx = append(holster.Tx, h)
		case *StmtHook:
			holster.Stmt = append(holster.Stmt, h)
		case *RowsHook:
			holster.Rows = append(holster.Rows, h)
		default:
			// do nothing
		}
	}
	return holster
}

type StmtHook struct {
	Exec  func(invocation *StmtExec, ctx context.Context, args ...interface{}) (sql.Result, error)
	Query func(invocation *StmtQuery, ctx context.Context, args ...interface{}) (*Rows, error)
	Close func(invocation *StmtClose) error
}

func (hook *StmtHook) Fill() {
	if hook.Exec == nil {
		hook.Exec = func(invocation *StmtExec, ctx context.Context, args ...interface{}) (sql.Result, error) {
			return invocation.original(ctx, args...)
		}
	}
	if hook.Query == nil {
		hook.Query = func(invocation *StmtQuery, ctx context.Context, args ...interface{}) (*Rows, error) {
			return invocation.original(ctx, args...)
		}
	}
	if hook.Close == nil {
		hook.Close = func(invocation *StmtClose) error {
			return invocation.original()
		}
	}
}

type stmtInvocation struct {
	Receiver *Stmt
	hooks    []*StmtHook
	length   int
	index    int
}

type StmtExec struct {
	stmtInvocation
	original func(ctx context.Context, args ...interface{}) (sql.Result, error)
}

func NewStmtExec(recv *Stmt, original func(ctx context.Context, args ...interface{}) (sql.Result, error)) *StmtExec {
	hooks := recv.core.hooks.Stmt
	return &StmtExec{stmtInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *StmtExec) Proceed(ctx context.Context, args ...interface{}) (sql.Result, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Exec(invocation, ctx, args...)
	}
	return invocation.original(ctx, args...)
}

type StmtQuery struct {
	stmtInvocation
	original func(ctx context.Context, args ...interface{}) (*Rows, error)
}

func NewStmtQuery(recv *Stmt, original func(ctx context.Context, args ...interface{}) (*Rows, error)) *StmtQuery {
	hooks := recv.core.hooks.Stmt
	return &StmtQuery{stmtInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *StmtQuery) Proceed(ctx context.Context, args ...interface{}) (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Query(invocation, ctx, args...)
	}
	return invocation.original(ctx, args...)
}

type StmtClose struct {
	stmtInvocation
	original func() error
}

func NewStmtClose(recv *Stmt, original func() error) *StmtClose {
	hooks := recv.core.hooks.Stmt
	return &StmtClose{stmtInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *StmtClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}

type RowsHook struct {
	Next  func(invocation *RowsNext) bool
	Err   func(invocation *RowsErr) error
	Scan  func(invocation *RowsScan, dest interface{}) error
	Close func(invocation *RowsClose) error
}

func (hook *RowsHook) Fill() {
	if hook.Next == nil {
		hook.Next = func(invocation *RowsNext) bool {
			return invocation.original()
		}
	}
	if hook.Err == nil {
		hook.Err = func(invocation *RowsErr) error {
			return invocation.original()
		}
	}
	if hook.Scan == nil {
		hook.Scan = func(invocation *RowsScan, dest interface{}) error {
			return invocation.original(dest)
		}
	}
	if hook.Close == nil {
		hook.Close = func(invocation *RowsClose) error {
			return invocation.original()
		}
	}
}

type rowsInvocation struct {
	Receiver *Rows
	hooks    []*RowsHook
	length   int
	index    int
}

type RowsNext struct {
	rowsInvocation
	original func() bool
}

func NewRowsNext(recv *Rows, original func() bool) *RowsNext {
	hooks := recv.core.hooks.Rows
	return &RowsNext{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsNext) Proceed() bool {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Next(invocation)
	}
	return invocation.original()
}

type RowsErr struct {
	rowsInvocation
	original func() error
}

func NewRowsErr(recv *Rows, original func() error) *RowsErr {
	hooks := recv.core.hooks.Rows
	return &RowsErr{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsErr) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Err(invocation)
	}
	return invocation.original()
}

type RowsScan struct {
	rowsInvocation
	original func(dest interface{}) error
}

func NewRowsScan(recv *Rows, original func(dest interface{}) error) *RowsScan {
	hooks := recv.core.hooks.Rows
	return &RowsScan{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsScan) Proceed(dest interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Scan(invocation, dest)
	}
	return invocation.original(dest)
}

type RowsClose struct {
	rowsInvocation
	original func() error
}

func NewRowsClose(recv *Rows, original func() error) *RowsClose {
	hooks := recv.core.hooks.Rows
	return &RowsClose{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
