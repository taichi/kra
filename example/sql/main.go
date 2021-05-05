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

package main

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/taichi/kra"
	"github.com/taichi/kra/sql"
)

type Film struct {
	Code     string
	Title    string
	Did      int
	DateProd time.Time
	Kind     string
	Len      pgtype.Interval
}

func main() {
	ctx := context.Background()

	db, err := sql.Open(kra.NewCore(kra.PostgreSQL), "pgx", "user=test password=test host=localhost port=5432 database=test sslmode=disable")
	if err != nil {
		fmt.Println("open", err)
		return
	}
	defer db.Close()

	if _, err := db.Exec(ctx, `CREATE TABLE IF NOT EXISTS films (
	    code        char(5) PRIMARY KEY,
	    title       varchar(40) NOT NULL,
	    did         integer NOT NULL,
	    date_prod   date,
	    kind        varchar(10),
	    len         interval hour to minute
	);`); err != nil {
		fmt.Println("create", err)
		return
	}
	defer func() {
		if _, err := db.Exec(ctx, "DROP TABLE films"); err != nil {
			fmt.Println(err)
		}
	}()

	if stmt, err := db.Prepare(ctx, "INSERT INTO films (code, title, did, date_prod, kind, len) VALUES (:code, :title, :did, :dateprod, :kind, :len)"); err != nil {
		fmt.Println("prepare", err)
		return
	} else {
		testdata := []map[string]interface{}{
			{"code": "1111", "title": "aaaa", "did": 32, "dateprod": time.Now(), "kind": "CDR", "len": pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
			{"code": "2222", "title": "bbbb", "did": 34, "dateprod": time.Now(), "kind": "ZDE", "len": pgtype.Interval{Microseconds: 9000000000, Status: pgtype.Present}},
			{"code": "3333", "title": "cccc", "did": 65, "dateprod": time.Now(), "kind": "IOM", "len": pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
			{"code": "4444", "title": "dddd", "did": 72, "dateprod": time.Now(), "kind": "ERW", "len": pgtype.Interval{Microseconds: 7200000000, Status: pgtype.Present}},
		}
		for _, data := range testdata {
			if _, err := stmt.Exec(ctx, data); err != nil {
				fmt.Println("insert", err)
				return
			}
		}
		if err := stmt.Close(); err != nil {
			fmt.Println("close", err)
			return
		}
	}
	var films []Film
	if err := db.FindAll(ctx, &films, "SELECT * FROM films WHERE kind IN (:kind)", stdsql.NamedArg{Name: "kind", Value: []string{"CDR", "ZDE"}}); err != nil {
		fmt.Println("find", err)
		return
	}

	fmt.Printf("%v\n", films)
}