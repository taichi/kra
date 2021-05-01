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

type TestVR struct{}

func (vr *TestVR) BindVar(index int) string {
	return fmt.Sprintf("$%d", index)
}

func (vr *TestVR) ByIndex(index int) (interface{}, error) {
	fmt.Println("ByIndex", index)
	return nil, nil
}
func (vr *TestVR) ByName(name string) (interface{}, error) {
	if name == "予算" {
		return []string{"foo", "bar", "baz"}, nil
	}
	return nil, nil
}

func TestParse(t *testing.T) {
	if fixtureFile, err := os.ReadFile(filepath.Join("test", "select.sql")); err != nil {
		t.Error(err)
	} else if query, err := NewQuery(string(fixtureFile)); err != nil {
		t.Error(err)
	} else if sql, _, err := query.Analyze(&TestVR{}); err != nil {
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
