# kra

simple RDB access library

# Main Features

- named parameter
- IN statement variable expansion
- structure mapping

# Getting Started

## Install

```
go get github.com/taichi/kra
```

## Usage

### native pgx API

```go
package main

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	stdpgx "github.com/jackc/pgx/v4"

	"github.com/taichi/kra/pgx"
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

	db, err := pgx.Open(ctx, "user=test password=test host=localhost port=5432 database=test sslmode=disable")
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

	values := [][]interface{}{
		{"1111", "aaaa", 32, time.Now(), "CDR", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
		{"2222", "bbbb", 34, time.Now(), "ZDE", pgtype.Interval{Microseconds: 9000000000, Status: pgtype.Present}},
		{"3333", "cccc", 65, time.Now(), "IOM", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
		{"4444", "dddd", 72, time.Now(), "ERW", pgtype.Interval{Microseconds: 7200000000, Status: pgtype.Present}},
	}

	if _, err := db.CopyFrom(ctx, stdpgx.Identifier{"films"}, []string{"code", "title", "did", "date_prod", "kind", "len"}, stdpgx.CopyFromRows(values)); err != nil {
		fmt.Println("CopyFrom", err)
		return
	}

	var films []Film
	if err := db.FindAll(ctx, &films, "SELECT * FROM films WHERE kind IN (:kind)", stdsql.NamedArg{Name: "kind", Value: []string{"CDR", "ZDE"}}); err != nil {
		fmt.Println("find", err)
		return
	}

	fmt.Printf("%v\n", films)
}
```

### standard database/sql based API

```go
package main

import (
	"context"
	"fmt"
	"time"

	stdsql "database/sql"

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
```
