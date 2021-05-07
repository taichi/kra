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

package kra

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name    string
	create  string
	drop    string
	insert  string
	find    string
	findAll string
	count   string
	core    *Core
	db      *sql.DB
}

type fixture struct {
	TestKey   string `db:"test_key"`
	TestValue string `db:"test_value"`
	Len       pgtype.Interval
}

func setup(t *testing.T) (*TestTable, error) {

	table := newTestTable()
	table.core = NewCore(PostgreSQL)

	rawDb, err := sql.Open("pgx", "user=test password=test host=localhost port=5432 database=test sslmode=disable")
	if err != nil {
		return nil, err
	}
	table.db = rawDb

	if _, eErr := table.db.ExecContext(context.Background(), table.create); eErr != nil {
		return nil, eErr
	}

	t.Cleanup(cleanup(t, table))

	if err := insertData(t, table); err != nil {
		return nil, err
	}

	return table, nil
}

func newTestTable() *TestTable {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	result := TestTable{
		name: fmt.Sprintf("test_%s_%x", time.Now().Format("20060102_150405"), b),
	}

	result.create = fmt.Sprintf(`CREATE TABLE %s (
		test_key VARCHAR,
		test_value VARCHAR,
		len interval hour to minute
);`, result.name)
	result.drop = fmt.Sprintf("DROP TABLE IF EXISTS %s", result.name)
	result.insert = fmt.Sprintf("INSERT INTO %s (test_key, test_value, len) VALUES ($1, $2, $3)", result.name)
	result.find = fmt.Sprintf("SELECT test_key, test_value, len FROM %s WHERE test_key= $1", result.name)
	result.findAll = fmt.Sprintf("SELECT test_key, test_value, len FROM %s ORDER BY test_key", result.name)
	result.count = fmt.Sprintf("SELECT COUNT(*) FROM %s", result.name)

	return &result
}

func cleanup(t *testing.T, table *TestTable) func() {
	return func() {
		if _, err := table.db.ExecContext(context.Background(), table.drop); err != nil {
			t.Error(err)
			return
		}
		if err := table.db.Close(); err != nil {
			t.Error(err)
		}
	}
}

func insertData(t *testing.T, table *TestTable) error {
	data := [][]interface{}{
		{"1111", "aaaa", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
		{"2222", "bbbb", pgtype.Interval{Microseconds: 9000000000, Status: pgtype.Present}},
		{"3333", "cccc", pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}},
	}

	for _, fix := range data {
		if res, err := table.db.ExecContext(context.Background(), table.insert, fix[0], fix[1], fix[2]); err != nil {
			return err
		} else if count, err := res.RowsAffected(); err != nil {
			return err
		} else {
			assert.Equal(t, int64(1), count)
		}
	}
	return nil
}

func TestDefaultTransformer_Transform_Scanner(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	query := fmt.Sprintf("SELECT len FROM %s WHERE test_key= $1", table.name) // nolint:gosec
	rows, err := table.db.Query(query, "1111")
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()
	assert.True(t, rows.Next())

	target := table.core.NewTransformer()

	var result pgtype.Interval
	if err := target.Transform(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}, result)
}

func TestDefaultTransformer_Transform_Struct(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	rows, err := table.db.Query(table.find, "1111")
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()
	assert.True(t, rows.Next())

	target := table.core.NewTransformer()

	var result fixture
	if err := target.Transform(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "1111", result.TestKey)
	assert.Equal(t, pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}, result.Len)
}

func TestDefaultTransformer_Transform_Map(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	rows, err := table.db.Query(table.find, "1111")
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()
	assert.True(t, rows.Next())

	target := table.core.NewTransformer()

	var result map[string]interface{}
	if err := target.Transform(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "1111", result["test_key"])
	assert.Equal(t, "01:30:00", result["len"])
}

func TestDefaultTransformer_Transform_int(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	rows, err := table.db.Query(table.count)
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()
	assert.True(t, rows.Next())

	target := table.core.NewTransformer()

	var result int
	if err := target.Transform(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, 3, result)
}

func TestDefaultTransformer_TransformAll_Scanner(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	query := fmt.Sprintf("SELECT len FROM %s ORDER BY test_key", table.name) // nolint:gosec
	rows, err := table.db.Query(query)
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()

	target := table.core.NewTransformer()

	var result []pgtype.Interval
	if err := target.TransformAll(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, pgtype.Interval{Microseconds: 5400000000, Status: pgtype.Present}, result[0])
	assert.Equal(t, 3, len(result))
}

func TestDefaultTransformer_TransformAll_String(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	query := fmt.Sprintf("SELECT test_key FROM %s ORDER BY test_key", table.name) // nolint:gosec
	rows, err := table.db.Query(query)
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()

	target := table.core.NewTransformer()

	var result []string
	if err := target.TransformAll(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "1111", result[0])
	assert.Equal(t, 3, len(result))
}

func TestDefaultTransformer_TransformAll_Struct(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	rows, err := table.db.Query(table.findAll)
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()

	target := table.core.NewTransformer()

	var result []fixture
	if err := target.TransformAll(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "1111", result[0].TestKey)
}

func TestDefaultTransformer_TransformAll_Map(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	rows, err := table.db.Query(table.findAll)
	if err != nil {
		t.Error(err)
		return
	}
	if rows.Err() != nil {
		t.Error(rows.Err())
		return
	}
	defer rows.Close()

	target := table.core.NewTransformer()

	var result []map[string]interface{}
	if err := target.TransformAll(rows, &result); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "1111", result[0]["test_key"])
	assert.Equal(t, "01:30:00", result[0]["len"])
}
