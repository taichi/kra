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
	defer db.Close()

	var films []Film
	cond := map[string]interface{}{"kind": []string{"CDR", "ZDE"}}
	if err := db.FindAll(ctx, &films, "SELECT * FROM films WHERE kind IN (:kind)", cond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", films)
}
