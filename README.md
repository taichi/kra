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

```go
package main

import (
	"context"
	"fmt"
	"time"

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
	Len      string
}

func main() {
	ctx := context.Background()

	db, err := sql.Open(kra.NewCore(kra.PostgreSQL), "pgx", "user=test password=test host=localhost port=5432 database=test sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	var films []Film
	cond := map[string]interface{}{"kind": []string{"CDR", "ZDE"}}
	if err := db.FindAll(ctx, &films, "SELECT * FROM films WHERE kind IN (:kind)", cond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", films)
}
```
