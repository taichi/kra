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
	"errors"
	"fmt"

	"github.com/taichi/kra"
)

type ExecFn func(context.Context, string, ...interface{}) (sql.Result, error)

func doExec(core *kra.Core, exec ExecFn, ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if rawQuery, bindArgs, err := core.Analyze(query, args...); err != nil {
		return nil, err
	} else {
		return exec(ctx, rawQuery, bindArgs...)
	}
}

type PrepareFn func(context.Context, string) (*sql.Stmt, error)

var ErrLackOfQueryParameters = errors.New("require example parameters for prepare query with IN operator")

func doPrepare(core *kra.Core, prepare PrepareFn, ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	if query, err := core.Parse(query); err != nil {
		return nil, err
	} else {
		resolver, err := core.NewResolver(examples...)
		if err != nil {
			return nil, err
		}
		var requires []string
		for _, name := range query.DynamicParameters() {
			if val, err := resolver.ByName(name); err != nil || val == nil {
				requires = append(requires, name)
			}
		}

		if 0 < len(requires) {
			return nil, fmt.Errorf("require parameters: %v %w", requires, ErrLackOfQueryParameters)
		}

		if rawQuery, _, err := query.Analyze(resolver); err != nil {
			return nil, err
		} else if stmt, err := prepare(ctx, rawQuery); err != nil {
			return nil, err
		} else {
			return &Stmt{stmt, core, query}, nil
		}
	}
}

type QueryFn func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

func doQuery(core *kra.Core, query QueryFn, ctx context.Context, queryString string, args ...interface{}) (*Rows, error) {
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

func doFind(core *kra.Core, query QueryFn, ctx context.Context, dst interface{}, queryString string, args ...interface{}) error {
	if rows, err := doQuery(core, query, ctx, queryString, args...); err != nil {
		return err
	} else if err := rows.Scan(dst); err != nil {
		return err
	}
	return nil
}

func doFindAll(core *kra.Core, query QueryFn, ctx context.Context, dst interface{}, queryString string, args ...interface{}) error {
	if rows, err := doQuery(core, query, ctx, queryString, args...); err != nil {
		return err
	} else if err := rows.ScanAll(dst); err != nil {
		return err
	}
	return nil
}
