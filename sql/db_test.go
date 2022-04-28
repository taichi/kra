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

package sql

import (
	"context"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/taichi/kra"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type TestTable struct {
	name    string
	create  string
	drop    string
	insert  string
	find    string
	findAll string
	count   string
	db      *DB
}

type fixture struct {
	TestKey   string `db:"test_key"`
	TestValue string `db:"test_value"`
}

func setup(t *testing.T) (*TestTable, error) {

	core := NewCore(kra.PostgreSQL, nil)

	table := newTestTable()

	rawDb, err := Open(core, "pgx", "user=test password=test host=localhost port=5432 database=test sslmode=disable")
	if err != nil {
		return nil, err
	}
	table.db = rawDb

	if _, eErr := table.db.db.ExecContext(context.Background(), table.create); eErr != nil {
		return nil, eErr
	}

	t.Cleanup(cleanup(t, table))

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
		test_value VARCHAR
);`, result.name)
	result.drop = fmt.Sprintf("DROP TABLE IF EXISTS %s", result.name)
	result.insert = fmt.Sprintf("INSERT INTO %s (test_key, test_value) VALUES (:test_key, :test_value)", result.name)
	result.find = fmt.Sprintf("SELECT test_key, test_value FROM %s WHERE test_key= ?", result.name)
	result.findAll = fmt.Sprintf("SELECT test_key, test_value FROM %s", result.name)
	result.count = fmt.Sprintf("SELECT COUNT(*) FROM %s", result.name)

	return &result
}

func cleanup(t *testing.T, table *TestTable) func() {
	return func() {
		if _, err := table.db.db.ExecContext(context.Background(), table.drop); err != nil {
			t.Error(err)
			return
		}
		if err := table.db.Close(); err != nil {
			t.Error(err)
		}
	}
}

func insertData(t *testing.T, table *TestTable) error {
	data := []fixture{
		{"111", "aa"},
		{"222", "bbbb"},
		{"333", "ccc"},
	}

	for _, fix := range data {
		if res, err := table.db.Exec(context.Background(), table.insert, fix); err != nil {
			return err
		} else if count, err := res.RowsAffected(); err != nil {
			return err
		} else {
			assert.Equal(t, int64(1), count)
		}
	}
	return nil
}

func TestExec(t *testing.T) {

	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	if res, err := table.db.Exec(context.Background(), table.insert, &fixture{"111", "bbbb"}); err != nil {
		t.Error(err)
		return
	} else if count, err := res.RowsAffected(); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, int64(1), count)
	}
}

func TestFind(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}
	if err := insertData(t, table); err != nil {
		t.Error(err)
		return
	}

	var dst fixture
	if err := table.db.Find(context.Background(), &dst, table.find, "222"); err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "bbbb", dst.TestValue)

	var dstMap map[string]interface{}
	if err := table.db.Find(context.Background(), &dstMap, table.find, "222"); err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "bbbb", dstMap["test_value"])

	var count int
	if err := table.db.Find(context.Background(), &count, table.count); err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 3, count)
}

func TestFindAll(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}
	if err := insertData(t, table); err != nil {
		t.Error(err)
		return
	}

	var dstAry []*fixture
	if err := table.db.FindAll(context.Background(), &dstAry, table.findAll); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, 3, len(dstAry))
	assert.Equal(t, "111", dstAry[0].TestKey)
}

func TestFindAllMap(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}
	if err := insertData(t, table); err != nil {
		t.Error(err)
		return
	}

	var dstAry []map[string]interface{}
	if err := table.db.FindAll(context.Background(), &dstAry, table.findAll); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, 3, len(dstAry))
	assert.Equal(t, "111", dstAry[0]["test_key"])
}

func TestPrepare_Exec(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()

	stmt, err := table.db.Prepare(ctx, table.insert)
	if err != nil {
		t.Error(err)
		return
	}

	defer stmt.Close()

	if res, err := stmt.Exec(ctx, &fixture{"4444", "dddd"}); err != nil {
		t.Error(err)
		return
	} else if count, err := res.RowsAffected(); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, int64(1), count)
	}
}

func TestPrepare_Query(t *testing.T) {
	table, err := setup(t)
	if err != nil {
		t.Error(err)
		return
	}

	err = insertData(t, table)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()

	stmt, err := table.db.Prepare(ctx, table.find)
	if err != nil {
		t.Error(err)
		return
	}

	defer stmt.Close()

	if rows, err := stmt.Query(ctx, "111"); err != nil {
		t.Error(err)
		return
	} else if rows.rows.Next() == false {
		t.Fail()
	} else {
		var data fixture
		sErr := rows.Scan(&data)
		if sErr != nil {
			t.Error(sErr)
			return
		}
		assert.Equal(t, "aa", data.TestValue)
	}
}
