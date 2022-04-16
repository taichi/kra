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
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("pgx", "user=test password=test host=localhost port=5432 database=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	type Product struct {
		Id   int
		Name string
	}
	rows, err := db.QueryContext(ctx, "select id, name from products where unit_price < $1", 100)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		prod := Product{}
		if se := rows.Scan(&prod.Id, &prod.Name); se != nil {
			fmt.Println(se)
		}
		fmt.Printf("%+v\n", &prod)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
