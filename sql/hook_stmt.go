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

type StmtExec invocation[Stmt, StmtHook, func(ctx context.Context, args ...interface{}) (sql.Result, error)]

func NewStmtExec(recv *Stmt, original func(ctx context.Context, args ...interface{}) (sql.Result, error)) *StmtExec {
	hooks := recv.core.hooks.Stmt
	return &StmtExec{recv, hooks, len(hooks), 0, original}
}

func (invocation *StmtExec) Proceed(ctx context.Context, args ...interface{}) (sql.Result, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Exec != nil {
		return invocation.hooks[current].Exec(invocation, ctx, args...)
	}
	return invocation.original(ctx, args...)
}

type StmtQuery invocation[Stmt, StmtHook, func(ctx context.Context, args ...interface{}) (*Rows, error)]

func NewStmtQuery(recv *Stmt, original func(ctx context.Context, args ...interface{}) (*Rows, error)) *StmtQuery {
	hooks := recv.core.hooks.Stmt
	return &StmtQuery{recv, hooks, len(hooks), 0, original}
}

func (invocation *StmtQuery) Proceed(ctx context.Context, args ...interface{}) (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Query != nil {
		return invocation.hooks[current].Query(invocation, ctx, args...)
	}
	return invocation.original(ctx, args...)
}

type StmtClose invocation[Stmt, StmtHook, func() error]

func NewStmtClose(recv *Stmt, original func() error) *StmtClose {
	hooks := recv.core.hooks.Stmt
	return &StmtClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *StmtClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
