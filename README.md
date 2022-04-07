# kra

![](https://github.com/taichi/kra/actions/workflows/push.yml/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/https://pkg.go.dev/github.comtaichi/kra)](https://pkg.go.dev/github.com/taichi/kra)

**SQL is the best way to access database**.

kra is a relational database access helper library on top of go.

kra works with `database/sql`, so all of database with `database/sql` based driver is supported.
and kra also works with `pgx` native API. kra focuses on the convenient use of `CopyFrom`.

# Features

- Named parameter support with dot notation
- IN statement variable expansion
- Rows to structure/map mapping
- Selectable base API. pgx or database/sql
- Highly configurable behavior
- Context is required for network access APIs
- All wrapper object has escape hatches

# Getting Started

## Install

```
go get github.com/taichi/kra
```

## Usage

### native pgx based API

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgtype"

	"github.com/taichi/kra/pgx"
)

type Film struct {
	Code     string
	Title    string
	Did      int
	DateProd time.Time `db:"date_prod"`
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

	testdata := []Film{
		{"1111", "aaaa", 32, time.Now(), "CDR", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
		{"2222", "bbbb", 34, time.Now(), "ZDE", pgtype.Interval{Microseconds: 9000000000, Status: pgtype.Present}},
		{"3333", "cccc", 65, time.Now(), "IOM", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
		{"4444", "dddd", 72, time.Now(), "ERW", pgtype.Interval{Microseconds: 7200000000, Status: pgtype.Present}},
	}

	if _, err := db.CopyFrom(ctx, pgx.Identifier{"films"}, testdata); err != nil {
		fmt.Println("CopyFrom", err)
		return
	}

	var films []Film
	if err := db.FindAll(ctx, &films, "SELECT * FROM films WHERE kind IN (:kind)", sql.NamedArg{Name: "kind", Value: []string{"CDR", "ZDE"}}); err != nil {
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

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/taichi/kra"
	"github.com/taichi/kra/sql"
)

type Film struct {
	Code     string
	Title    string
	Did      int
	DateProd time.Time `db:"date_prod"`
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

	if stmt, err := db.Prepare(ctx, "INSERT INTO films (code, title, did, date_prod, kind, len) VALUES (:code, :title, :did, :date_prod, :kind, :len)"); err != nil {
		fmt.Println("prepare", err)
		return
	} else {
		testdata := []Film{
			{"1111", "aaaa", 32, time.Now(), "CDR", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
			{"2222", "bbbb", 34, time.Now(), "ZDE", pgtype.Interval{Microseconds: 9000000000, Status: pgtype.Present}},
			{"3333", "cccc", 65, time.Now(), "IOM", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
			{"4444", "dddd", 72, time.Now(), "ERW", pgtype.Interval{Microseconds: 7200000000, Status: pgtype.Present}},
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
	if err := db.FindAll(ctx, &films, "SELECT * FROM films WHERE kind IN (:kind)", kra.NamedArg{Name: "kind", Value: []string{"CDR", "ZDE"}}); err != nil {
		fmt.Println("find", err)
		return
	}

	fmt.Printf("%v\n", films)
}
```

# Related OSS

- [pgx](https://github.com/jackc/pgx)
- [scany](https://github.com/georgysavva/scany)
- [sqlx](https://github.com/jmoiron/sqlx)
