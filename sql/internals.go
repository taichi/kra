// Copyright 2021 taichi
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

type execFn func(context.Context, string, ...interface{}) (sql.Result, error)

func doExec(core *kra.Core, exec execFn, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if rawQuery, bindArgs, err := core.Analyze(query, args...); err != nil {
		return nil, err
	} else {
		return exec(ctx, rawQuery, bindArgs...)
	}
}

type prepareFn func(context.Context, string) (*sql.Stmt, error)

func doPrepare(core *kra.Core, prepare prepareFn, ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	if query, err := core.Parse(query); err != nil {
		return nil, err
	} else if resolver, err := core.NewResolver(examples...); err != nil {
		return nil, err
	} else if err := query.Verify(resolver); err != nil {
		return nil, err
	} else if rawQuery, _, err := query.Analyze(kra.KeepSilent(resolver)); err != nil {
		return nil, err
	} else if stmt, err := prepare(ctx, rawQuery); err != nil {
		return nil, err
	} else {
		return &Stmt{stmt, core, query}, nil
	}
}

type queryFn func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

func doQuery(core *kra.Core, query queryFn, ctx context.Context, queryString string, args ...interface{}) (*Rows, error) {
	if rawQuery, bindArgs, err := core.Analyze(queryString, args...); err != nil {
		return nil, err
	} else if rows, err := query(ctx, rawQuery, bindArgs...); err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	} else {
		return &Rows{rows, core.NewTransformer()}, nil
	}
}

func doFind(core *kra.Core, query queryFn, ctx context.Context, dest interface{}, queryString string, args ...interface{}) error {
	if rows, err := doQuery(core, query, ctx, queryString, args...); err != nil {
		return err
	} else {
		defer rows.Close()
		if rows.rows.Next() == false {
			return kra.ErrNoRecord
		} else if err := rows.Scan(dest); err != nil {
			return err
		}
	}
	return nil
}

func doFindAll(core *kra.Core, query queryFn, ctx context.Context, dest interface{}, queryString string, args ...interface{}) error {
	if rows, err := doQuery(core, query, ctx, queryString, args...); err != nil {
		return err
	} else {
		defer rows.Close()
		if err := rows.transformer.TransformAll(rows.rows, dest); err != nil {
			return err
		}
	}
	return nil
}
