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

package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/taichi/kra"
	"github.com/taichi/kra/pgx"
)

type TraceQA struct {
	delegate kra.QueryAnalyzer
}

func (qa *TraceQA) Verify(vr kra.ValueResolver) error {
	return qa.delegate.Verify(vr)
}

var BUFFER_SIZE = 4
var MAIN = 7

func (qa *TraceQA) Analyze(vr kra.ValueResolver) (query string, vars []interface{}, err error) {
	q, v, e := qa.delegate.Analyze(vr)
	frame := kra.FindCaller(MAIN, BUFFER_SIZE, func(frame *runtime.Frame) (found bool) {
		return kra.ToPackageName(frame.Function) == "main"
	})
	return fmt.Sprintf("/* %s:%d */%s", frame.File, frame.Line, q), v, e
}

func main() {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig("user=test password=test database=test")
	if err != nil {
		panic(err)
	}
	db, err := pgx.OpenConfig(ctx, config, &kra.CoreHook{
		Parse: func(invocation *kra.CoreParse, query string) (kra.QueryAnalyzer, error) {
			if qa, er1 := invocation.Proceed(query); er1 != nil {
				return nil, er1
			} else {
				return &TraceQA{qa}, nil
			}
		},
	}, &pgx.DBHook{
		Prepare: func(invocation *pgx.DBPrepare, ctx context.Context, query string, examples ...interface{}) (*pgx.PooledStmt, error) {
			fmt.Println("prepare:", query)
			stmt, er2 := invocation.Proceed(ctx, query, examples...)
			fmt.Printf("postPrepare: %+v, %v \n", stmt.Stmt().SQL, er2)
			return stmt, er2
		},
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cond := map[string]interface{}{
		"name": "いぶりがっこ",
	}
	stmt, err := db.Prepare(ctx, "INSERT INTO products(name) VALUES (:name)")
	defer stmt.Close(ctx)
	if err != nil {
		panic(err)
	} else if res, err := stmt.Exec(ctx, cond); err != nil {
		panic(err)
	} else {
		rowCnt := res.RowsAffected()
		fmt.Printf("affected = %d\n", rowCnt)
	}
}
