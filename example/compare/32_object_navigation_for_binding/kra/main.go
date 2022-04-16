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

	"github.com/taichi/kra/pgx"
)

func main() {
	ctx := context.Background()

	db, err := pgx.Open(ctx, "user=test password=test host=localhost port=5432 database=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	type Category struct {
		Name string
	}
	type Product struct {
		Category
	}
	cond := Product{
		Category{Name: "生鮮食品"},
	}
	sql := "SELECT COUNT(c.id) FROM products AS p JOIN categories AS c" +
		" ON c.id = p.category_id WHERE c.name = :category.name"
	rows, err := db.Query(ctx, sql, &cond)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var count int
		if se := rows.Scan(&count); se != nil {
			panic(se)
		}
		fmt.Println(count)
	}
}
