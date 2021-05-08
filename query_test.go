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
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestVR struct {
	values map[string]interface{}
}

func (vr *TestVR) BindVar(index int) string {
	return fmt.Sprintf("$%d", index)
}

func (vr *TestVR) ByIndex(index int) (interface{}, error) {
	return vr.values[fmt.Sprintf("%d", index)], nil
}
func (vr *TestVR) ByName(name string) (interface{}, error) {
	return vr.values[name], nil
}

func TestParse(t *testing.T) {
	if fixtureFile, err := os.ReadFile(filepath.Join("test", "select.sql")); err != nil {
		t.Error(err)
	} else if query, err := NewQuery(string(fixtureFile)); err != nil {
		t.Error(err)
	} else if sql, _, err := query.Analyze(&TestVR{
		map[string]interface{}{
			"予算": []string{"foo", "bar", "baz"},
		},
	}); err != nil {
		t.Error(err)
	} else {
		t.Log(sql)
	}
}

func Test_2orMoreStyles(t *testing.T) {
	visitor := PartsCollector{}
	visitor.Use(NAMED)

	assert.False(t, visitor.Use2orMoreStyles())

	visitor.Use(QMARK)

	assert.True(t, visitor.Use2orMoreStyles())
}

func TestAsSlice(t *testing.T) {

	assert.Nil(t, AsSlice(nil))

	strSlice := []string{"foo", "bar", "baz"}
	assert.ElementsMatch(t, AsSlice(strSlice), strSlice)

	var strArray [2]string
	strArray[0] = "foo"
	strArray[1] = "bar"
	assert.ElementsMatch(t, AsSlice(strArray), strArray)

	assert.ElementsMatch(t, AsSlice("foo"), []string{"foo"})

	byteSlice := []byte{'f', 'b', 'z'}
	assert.ElementsMatch(t, AsSlice(byteSlice), []interface{}{byteSlice})

	var byteArray [2]byte
	byteArray[0] = '@'
	byteArray[1] = 'c'
	assert.ElementsMatch(t, AsSlice(byteArray), []interface{}{byteArray})
}

func TestQMark(t *testing.T) {
	if query, err := NewQuery("INSERT INTO foo (bar, baz) VALUES (?, ?)"); err != nil {
		t.Error(err)
		return
	} else if raw, _, err := query.Analyze(&TestVR{}); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "INSERT INTO foo ( bar , baz ) VALUES ( $1 , $2 )", raw)
	}
}

func TestDMark(t *testing.T) {
	if query, err := NewQuery("INSERT INTO foo (bar, baz) VALUES ($1, $2)"); err != nil {
		t.Error(err)
		return
	} else if raw, _, err := query.Analyze(&TestVR{}); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "INSERT INTO foo ( bar , baz ) VALUES ( $1 , $2 )", raw)
	}
}

func TestATMark(t *testing.T) {
	if query, err := NewQuery("INSERT INTO foo (bar, baz) VALUES (@p1, @p2)"); err != nil {
		t.Error(err)
		return
	} else if raw, _, err := query.Analyze(&TestVR{}); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "INSERT INTO foo ( bar , baz ) VALUES ( $1 , $2 )", raw)
	}
}

func TestNamedParameter(t *testing.T) {
	if query, err := NewQuery("SELECT foo, bar FROM baz WHERE foo = :饅頭 AND baz = :予算"); err != nil {
		t.Error(err)
		return
	} else if raw, vars, err := query.Analyze(&TestVR{
		map[string]interface{}{
			"饅頭": "11111",
			"予算": []string{"foo", "bar", "baz"},
		},
	}); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "SELECT foo , bar FROM baz WHERE foo = $1 AND baz = $2", raw)
		assert.Equal(t, []interface{}{"11111", []string{"foo", "bar", "baz"}}, vars)
	}
}

func TestNamedATParameter(t *testing.T) {
	if query, err := NewQuery("SELECT foo, bar FROM baz WHERE foo = @饅頭.こしあん AND baz = @予算"); err != nil {
		t.Error(err)
		return
	} else if raw, vars, err := query.Analyze(&TestVR{
		map[string]interface{}{
			"饅頭.こしあん": "11111",
			"予算":      "foo",
		},
	}); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "SELECT foo , bar FROM baz WHERE foo = $1 AND baz = $2", raw)
		assert.Equal(t, []interface{}{"11111", "foo"}, vars)
	}
}

func TestINOperator_AutoExpansion(t *testing.T) {
	if query, err := NewQuery("SELECT foo, bar FROM baz WHERE foo IN (@予算, ?)"); err != nil {
		t.Error(err)
		return
	} else if raw, vars, err := query.Analyze(&TestVR{
		map[string]interface{}{
			"予算": []string{"foo", "bar", "baz"},
		},
	}); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "SELECT foo , bar FROM baz WHERE foo IN ($1 , $2 , $3)", raw)
		assert.Equal(t, []interface{}{"foo", "bar", "baz"}, vars)
	}
}

func TestINOperator_WithoutExpansion_Q(t *testing.T) {
	if query, err := NewQuery("SELECT foo, bar FROM baz WHERE foo IN (?, ?, ?)"); err != nil {
		t.Error(err)
		return
	} else if raw, vars, err := query.Analyze(&TestVR{
		map[string]interface{}{
			"1": "foo",
			"2": "bar",
			"3": "baz",
		},
	}); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "SELECT foo , bar FROM baz WHERE foo IN ( $1 , $2 , $3 )", raw)
		assert.Equal(t, []interface{}{"foo", "bar", "baz"}, vars)
	}
}
