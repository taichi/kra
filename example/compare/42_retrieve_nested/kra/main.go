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
		Id   int
		Name string
	}
	type Product struct {
		Id   int
		Name string
		Category
	}

	rows, err := db.Query(ctx, "SELECT p.id, p.name, c.id AS \"Category.Id\", c.name AS \"Category.Name\" FROM products AS p JOIN categories AS c ON c.id = p.category_id")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		prod := Product{}
		if se := rows.Scan(&prod); se != nil {
			fmt.Println(se)
		}
		fmt.Println(prod.Id, prod.Name, prod.Category.Id, prod.Category.Name)
	}
}
