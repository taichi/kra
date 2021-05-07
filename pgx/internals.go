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
	"crypto/sha256"
	"encoding/hex"

	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"

	"github.com/taichi/kra"
)

type ExecFn func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)

func doExec(core *kra.Core, exec ExecFn, ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	if rawQuery, bindArgs, err := core.Analyze(query, args...); err != nil {
		return nil, err
	} else {
		return exec(ctx, rawQuery, bindArgs...)
	}
}

type PrepareFn func(ctx context.Context, name, query string) (sd *pgconn.StatementDescription, err error)

func doPrepare(core *kra.Core, conn *pgx.Conn, prepare PrepareFn, ctx context.Context, query string, examples ...interface{}) (*Stmt, error) {
	if query, err := core.Parse(query); err != nil {
		return nil, err
	} else if resolver, err := core.NewResolver(examples...); err != nil {
		return nil, err
	} else if err := query.Verify(resolver); err != nil {
		return nil, err
	} else if rawQuery, _, err := query.Analyze(kra.KeepSilent(resolver)); err != nil {
		return nil, err
	} else if name, err := toName(rawQuery); err != nil {
		return nil, err
	} else if stmt, err := prepare(ctx, name, rawQuery); err != nil {
		return nil, err
	} else {
		return &Stmt{stmt, conn, core, query}, nil
	}
}

func toName(query string) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(query)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

type QueryFn func(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)

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
