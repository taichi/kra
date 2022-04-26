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
}

func NewHook() *Hook {
	return &Hook{
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
	}
}
