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

	"github.com/taichi/kra"
)

type Hook struct {
	// kra core hooks
	Parse          func(original func(query string) (kra.QueryAnalyzer, error), query string) (kra.QueryAnalyzer, error)
	NewResolver    func(original func(args ...interface{}) (kra.ValueResolver, error), args ...interface{}) (kra.ValueResolver, error)
	NewTransformer func(original func() kra.Transformer) kra.Transformer

	// pgx hooks
	BeginTx     func(original func(ctx context.Context, txOptions pgx.TxOptions) (*Tx, error), ctx context.Context, txOptions pgx.TxOptions) (*Tx, error)
	BeginTxFunc func(original func(ctx context.Context, txOptions pgx.TxOptions, f func(*Tx) error) error, ctx context.Context, txOptions pgx.TxOptions, f func(*Tx) error) error
	CopyFrom    func(original func(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error), ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error)
	Exec        func(original func(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error), ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Ping        func(original func(ctx context.Context) error, ctx context.Context) error
	Prepare     func(original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error), ctx context.Context, query string, examples ...interface{}) (*Stmt, error)
	Query       func(original func(ctx context.Context, query string, args ...interface{}) (*Rows, error), ctx context.Context, query string, args ...interface{}) (*Rows, error)
	SendBatch   func(original func(ctx context.Context, batch *Batch) *BatchResults, ctx context.Context, batch *Batch) *BatchResults
	Find        func(original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	FindAll     func(original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error, ctx context.Context, dest interface{}, query string, args ...interface{}) error

	Conn         *ConnHook
	DB           *DBHook
	Tx           *TxHook
	Stmt         *StmtHook
	Rows         *RowsHook
	Batch        *BatchHook
	BatchResults *BatchResultsHook
}

type ConnHook struct {
	Close func(original func(ctx context.Context) error, ctx context.Context) error
}

type DBHook struct {
	Close func(original func() error) error
}

type TxHook struct {
	Commit   func(original func(ctx context.Context) error, ctx context.Context) error
	Rollback func(original func(ctx context.Context) error, ctx context.Context) error

	Begin     func(original func(ctx context.Context) (*Tx, error), ctx context.Context) (*Tx, error)
	BeginFunc func(original func(ctx context.Context, f func(*Tx) error) error, ctx context.Context, f func(*Tx) error) error
}

type StmtHook struct {
	Exec  func(original func(ctx context.Context, args ...interface{}) (pgconn.CommandTag, error), ctx context.Context, args ...interface{}) (pgconn.CommandTag, error)
	Query func(original func(ctx context.Context, args ...interface{}) (*Rows, error), ctx context.Context, args ...interface{}) (*Rows, error)
	Close func(original func(ctx context.Context) error, ctx context.Context) error
}

type RowsHook struct {
	Next  func(original func() bool) bool
	Err   func(original func() error) error
	Scan  func(original func(dest interface{}) error, dest interface{}) error
	Close func(original func() error) error
}

type BatchHook struct {
	Queue func(original func(query string, args ...interface{}) error, query string, args ...interface{}) error
}

type BatchResultsHook struct {
	Exec  func(original func() (pgconn.CommandTag, error)) (pgconn.CommandTag, error)
	Query func(original func() (*Rows, error)) (*Rows, error)
}

func NewHook(hook *Hook) *Hook {
	baseHook := &Hook{
		Parse: func(original func(query string) (kra.QueryAnalyzer, error), query string) (kra.QueryAnalyzer, error) {
			return original(query)
		},
		NewResolver: func(original func(args ...interface{}) (kra.ValueResolver, error), args ...interface{}) (kra.ValueResolver, error) {
			return original(args...)
		},
		NewTransformer: func(original func() kra.Transformer) kra.Transformer {
			return original()
		},
		BeginTx: func(original func(ctx context.Context, txOptions pgx.TxOptions) (*Tx, error), ctx context.Context, txOptions pgx.TxOptions) (*Tx, error) {
			return original(ctx, txOptions)
		},
		BeginTxFunc: func(original func(ctx context.Context, txOptions pgx.TxOptions, f func(*Tx) error) error, ctx context.Context, txOptions pgx.TxOptions, f func(*Tx) error) error {
			return original(ctx, txOptions, f)
		},
		CopyFrom: func(original func(ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error), ctx context.Context, tableName Identifier, rowSrc interface{}) (int64, error) {
			return original(ctx, tableName, rowSrc)
		},
		Exec: func(original func(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error), ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
			return original(ctx, query, args...)
		},
		Ping: func(original func(ctx context.Context) error, ctx context.Context) error {
			return original(ctx)
		},
		Prepare: func(original func(ctx context.Context, query string, examples ...interface{}) (*Stmt, error), ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
			return original(ctx, query, examples...)
		},
		Query: func(original func(ctx context.Context, query string, args ...interface{}) (*Rows, error), ctx context.Context, query string, args ...interface{}) (*Rows, error) {
			return original(ctx, query, args...)
		},
		SendBatch: func(original func(ctx context.Context, batch *Batch) *BatchResults, ctx context.Context, batch *Batch) *BatchResults {
			return original(ctx, batch)
		},
		Find: func(original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
			return original(ctx, dest, query, args...)
		},
		FindAll: func(original func(ctx context.Context, dest interface{}, query string, args ...interface{}) error, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
			return original(ctx, dest, query, args...)
		},
		Conn:         NewConnHook(),
		DB:           NewDBHook(),
		Tx:           NewTxHook(),
		Stmt:         NewStmtHook(),
		Rows:         NewRowsHook(),
		Batch:        NewBatchHook(),
		BatchResults: NewBatchResultsHook(),
	}

	if hook != nil {
		baseHook.Merge(hook)
		if hook.Conn != nil {
			baseHook.Conn.Merge(hook.Conn)
		}
		if hook.DB != nil {
			baseHook.DB.Merge(hook.DB)
		}
		if hook.Tx != nil {
			baseHook.Tx.Merge(hook.Tx)
		}
		if hook.Stmt != nil {
			baseHook.Stmt.Merge(hook.Stmt)
		}
		if hook.Rows != nil {
			baseHook.Rows.Merge(hook.Rows)
		}
		if hook.Batch != nil {
			baseHook.Batch.Merge(hook.Batch)
		}
		if hook.BatchResults != nil {
			baseHook.BatchResults.Merge(hook.BatchResults)
		}
	}

	return baseHook
}

func (baseHook *Hook) Merge(hook *Hook) {
	if hook.Parse != nil {
		baseHook.Parse = hook.Parse
	}
	if hook.NewResolver != nil {
		baseHook.NewResolver = hook.NewResolver
	}
	if hook.NewTransformer != nil {
		baseHook.NewTransformer = hook.NewTransformer
	}
	if hook.BeginTx != nil {
		baseHook.BeginTx = hook.BeginTx
	}
	if hook.BeginTxFunc != nil {
		baseHook.BeginTxFunc = hook.BeginTxFunc
	}
	if hook.CopyFrom != nil {
		baseHook.CopyFrom = hook.CopyFrom
	}
	if hook.Exec != nil {
		baseHook.Exec = hook.Exec
	}
	if hook.Ping != nil {
		baseHook.Ping = hook.Ping
	}
	if hook.Prepare != nil {
		baseHook.Prepare = hook.Prepare
	}
	if hook.Query != nil {
		baseHook.Query = hook.Query
	}
	if hook.SendBatch != nil {
		baseHook.SendBatch = hook.SendBatch
	}
	if hook.Find != nil {
		baseHook.Find = hook.Find
	}
	if hook.FindAll != nil {
		baseHook.FindAll = hook.FindAll
	}
}

func NewConnHook() *ConnHook {
	return &ConnHook{
		Close: func(original func(ctx context.Context) error, ctx context.Context) error {
			return original(ctx)
		},
	}
}

func (baseHook *ConnHook) Merge(hook *ConnHook) {
	if hook.Close != nil {
		baseHook.Close = hook.Close
	}
}

func NewDBHook() *DBHook {
	return &DBHook{
		Close: func(original func() error) error {
			return original()
		},
	}
}

func (baseHook *DBHook) Merge(hook *DBHook) {
	if hook.Close != nil {
		baseHook.Close = hook.Close
	}
}

func NewTxHook() *TxHook {
	return &TxHook{
		Commit: func(original func(ctx context.Context) error, ctx context.Context) error {
			return original(ctx)
		},
		Rollback: func(original func(ctx context.Context) error, ctx context.Context) error {
			return original(ctx)
		},
		Begin: func(original func(ctx context.Context) (*Tx, error), ctx context.Context) (*Tx, error) {
			return original(ctx)
		},
		BeginFunc: func(original func(ctx context.Context, f func(*Tx) error) error, ctx context.Context, f func(*Tx) error) error {
			return original(ctx, f)
		},
	}
}

func (baseHook *TxHook) Merge(hook *TxHook) {
	if hook.Commit != nil {
		baseHook.Commit = hook.Commit
	}
	if hook.Rollback != nil {
		baseHook.Rollback = hook.Rollback
	}
	if hook.Begin != nil {
		baseHook.Begin = hook.Begin
	}
	if hook.BeginFunc != nil {
		baseHook.BeginFunc = hook.BeginFunc
	}
}

func NewStmtHook() *StmtHook {
	return &StmtHook{
		Exec: func(original func(ctx context.Context, args ...interface{}) (pgconn.CommandTag, error), ctx context.Context, args ...interface{}) (pgconn.CommandTag, error) {
			return original(ctx, args...)
		},
		Query: func(original func(ctx context.Context, args ...interface{}) (*Rows, error), ctx context.Context, args ...interface{}) (*Rows, error) {
			return original(ctx, args...)
		},
		Close: func(original func(ctx context.Context) error, ctx context.Context) error {
			return original(ctx)
		},
	}
}

func (basehook *StmtHook) Merge(hook *StmtHook) {
	if hook.Exec != nil {
		basehook.Exec = hook.Exec
	}
	if hook.Query != nil {
		basehook.Query = hook.Query
	}
	if hook.Close != nil {
		basehook.Close = hook.Close
	}
}

func NewRowsHook() *RowsHook {
	return &RowsHook{
		Next: func(original func() bool) bool {
			return original()
		},
		Err: func(original func() error) error {
			return original()
		},
		Scan: func(original func(dest interface{}) error, dest interface{}) error {
			return original(dest)
		},
		Close: func(original func() error) error {
			return original()
		},
	}
}

func (basehook *RowsHook) Merge(hook *RowsHook) {
	if hook.Next != nil {
		basehook.Next = hook.Next
	}
	if hook.Err != nil {
		basehook.Err = hook.Err
	}
	if hook.Scan != nil {
		basehook.Scan = hook.Scan
	}
	if hook.Close != nil {
		basehook.Close = hook.Close
	}
}

func NewBatchHook() *BatchHook {
	return &BatchHook{
		Queue: func(original func(query string, args ...interface{}) error, query string, args ...interface{}) error {
			return original(query, args...)
		},
	}
}

func (basehook *BatchHook) Merge(hook *BatchHook) {
	if hook.Queue != nil {
		basehook.Queue = hook.Queue
	}
}

func NewBatchResultsHook() *BatchResultsHook {
	return &BatchResultsHook{
		Exec: func(original func() (pgconn.CommandTag, error)) (pgconn.CommandTag, error) {
			return original()
		},
		Query: func(original func() (*Rows, error)) (*Rows, error) {
			return original()
		},
	}
}

func (basehook *BatchResultsHook) Merge(hook *BatchResultsHook) {
	if hook.Exec != nil {
		basehook.Exec = hook.Exec
	}
	if hook.Query != nil {
		basehook.Query = hook.Query
	}
}
