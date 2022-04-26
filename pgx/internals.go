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

package pgx

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync/atomic"

	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"

	"github.com/taichi/kra"
)

type execFn func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)

func doExec(core *Core, exec execFn, ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return core.hook.Exec(func(c context.Context, q string, a ...interface{}) (pgconn.CommandTag, error) {
		if rawQuery, bindArgs, err := core.Analyze(q, a...); err != nil {
			return nil, err
		} else {
			return exec(c, rawQuery, bindArgs...)
		}
	}, ctx, query, args...)
}

type prepareFn func(ctx context.Context, name, query string) (sd *pgconn.StatementDescription, err error)

func doPrepare(core *Core, conn *pgx.Conn, count *int64, prepare prepareFn, ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	atomic.AddInt64(count, 1)
	return core.hook.Prepare(func(c context.Context, q string, e ...interface{}) (*Stmt, error) {
		if query, err := core.hook.Parse(core.Parse, q); err != nil {
			return nil, err
		} else if resolver, err := core.hook.NewResolver(core.NewResolver, e...); err != nil {
			return nil, err
		} else if err := query.Verify(resolver); err != nil {
			return nil, err
		} else if rawQuery, _, err := query.Analyze(kra.KeepSilent(resolver)); err != nil {
			return nil, err
		} else if stmt, err := prepare(c, toName(*count), rawQuery); err != nil {
			return nil, err
		} else {
			return &Stmt{stmt, conn, core, query}, nil
		}
	}, ctx, query, examples...)

}

func toName(count int64) string {
	return fmt.Sprintf("kra-%d", count)
}

type queryFn func(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)

func doQuery(core *Core, query queryFn, ctx context.Context, queryString string, args ...interface{}) (*Rows, error) {
	if rawQuery, bindArgs, err := core.Analyze(queryString, args...); err != nil {
		return nil, err
	} else if rows, err := query(ctx, rawQuery, bindArgs...); err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	} else {
		return NewRows(core, rows), nil
	}
}

func doFind(core *Core, query queryFn, ctx context.Context, dst interface{}, queryString string, args ...interface{}) error {
	if rows, err := doQuery(core, query, ctx, queryString, args...); err != nil {
		return err
	} else {
		defer rows.Close()
		if rows.rows.Next() == false {
			return kra.ErrNoRecord
		} else if err := rows.Scan(dst); err != nil {
			return err
		}
	}
	return nil
}

func doFindAll(core *Core, query queryFn, ctx context.Context, dst interface{}, queryString string, args ...interface{}) error {
	if rows, err := doQuery(core, query, ctx, queryString, args...); err != nil {
		return err
	} else {
		defer rows.Close()
		if err := rows.transformer.TransformAll(rows.rows, dst); err != nil {
			return err
		}
	}
	return nil
}

type copyFromFn func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)

var ErrEmptySlice = errors.New("kra: empty slice")

var ErrCopyFromSource = errors.New("kra: cannot use pgx.CopyFromSource as src, use the underlying object")

func validateSrc(src interface{}) (*reflect.Value, int, error) {
	if _, ok := src.(pgx.CopyFromSource); ok {
		return nil, 0, ErrCopyFromSource
	}

	directValue := reflect.Indirect(reflect.ValueOf(src))

	if directValue.Kind() != reflect.Slice {
		return nil, 0, fmt.Errorf("type=%v %w", directValue.Type(), kra.ErrNoSlice)
	}
	length := directValue.Len()
	if length < 1 {
		return nil, 0, ErrEmptySlice
	}

	return &directValue, length, nil
}

func doCopyFrom(core *Core, copyFrom copyFromFn, ctx context.Context, tableName Identifier, src interface{}) (int64, error) {
	return core.hook.CopyFrom(func(c context.Context, tn Identifier, s interface{}) (int64, error) {
		dv, length, err := validateSrc(s)
		if err != nil {
			return 0, err
		}

		directValue := *dv

		var elementDef *kra.StructDef
		var columnNames []string
		var columnLength int
		rowSrc := make([][]interface{}, length)
		for index := 0; index < length; index++ {
			element := reflect.Indirect(directValue.Index(index))
			if element.Kind() != reflect.Struct {
				return 0, fmt.Errorf("type=%v %w", element.Kind(), kra.ErrUnsupportedValueType)
			} else if columnNames == nil {
				elementType := element.Type()
				if def, err := core.Repository.Lookup(elementType); err != nil {
					return 0, err
				} else {
					elementDef = def
				}
				for col, def := range elementDef.Members {
					if isCopyable(def) {
						columnNames = append(columnNames, col)
						columnLength++
					}
				}
			}

			values := make([]interface{}, columnLength)
			for i, col := range columnNames {
				if def, val, err := elementDef.ByName(element, col); err != nil {
					return 0, err
				} else if isCopyable(def) {
					values[i] = val.Interface()
				}
			}
			rowSrc[index] = values
		}

		return copyFrom(c, pgx.Identifier(tn), columnNames, pgx.CopyFromRows(rowSrc))
	}, ctx, tableName, src)
}

func isCopyable(def *kra.FieldDef) bool {
	return def.Unexported == false && def.Options["name"] != "-"
}
