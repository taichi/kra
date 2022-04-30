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
