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
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultResolver_Map(t *testing.T) {
	params := map[string]interface{}{
		"foo": "111",
		"bar": "222",
		"baz": "333",
	}
	assertDefaultResolver(t, params)
	assertDefaultResolver(t, &params)
}

func assertDefaultResolver(t *testing.T, params interface{}) {
	target, err := NewDefaultResolver(NewCore(MySQL), params)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "?", target.BindVar(1))

	if val, err := target.ByName("bar"); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "222", val)
	}

	if _, err := target.ByName("boz"); err != nil {
		assert.ErrorIs(t, err, ErrParameterNotFound)
	} else {
		t.Fail()
	}
}

func TestNewDefaultResolver_2Maps(t *testing.T) {
	core := NewCore(MySQL)
	target, err := NewDefaultResolver(core, map[string]interface{}{
		"foo": "111",
		"bar": "222",
		"baz": "333",
	}, map[string]interface{}{
		"fooFoo": "111111",
		"barBar": "222222",
		"bazBaz": "333333",
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "?", target.BindVar(1))

	if val, err := target.ByName("bazbaz"); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "333333", val)
	}

	if _, err := target.ByName("boz"); err != nil {
		assert.ErrorIs(t, err, ErrParameterNotFound)
	} else {
		t.Fail()
	}
}

func TestNewDefaultResolver_NamedArg(t *testing.T) {
	core := NewCore(MySQL)
	target, err := NewDefaultResolver(core,
		sql.NamedArg{Name: "foo", Value: "111"},
		NamedArg{Name: "bar", Value: "222"},
		NamedArg{Name: "baz", Value: "333"},
	)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "?", target.BindVar(1))

	if val, err := target.ByName("bar"); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "222", val)
	}

	if _, err := target.ByName("boz"); err != nil {
		assert.ErrorIs(t, err, ErrParameterNotFound)
	} else {
		t.Fail()
	}
}

func TestNewDefaultResolver_2NamedArgs(t *testing.T) {
	core := NewCore(MySQL)
	target, err := NewDefaultResolver(core,
		sql.NamedArg{Name: "foo", Value: "111"},
		NamedArg{Name: "bar", Value: "222"},
		NamedArg{Name: "baz", Value: "333"},
		sql.NamedArg{Name: "fooFoo", Value: "111111"},
		NamedArg{Name: "barBar", Value: "222222"},
		NamedArg{Name: "bazBaz", Value: "333333"},
	)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "?", target.BindVar(1))

	if val, err := target.ByName("foofoo"); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "111111", val)
	}

	if _, err := target.ByName("boz"); err != nil {
		assert.ErrorIs(t, err, ErrParameterNotFound)
	} else {
		t.Fail()
	}
}

type Fixture struct {
	Foo string
	Bar string
	Baz string
}

func TestNewDefaultResolver_Struct(t *testing.T) {
	core := NewCore(MySQL)
	target, err := NewDefaultResolver(core, &Fixture{"111", "222", "333"})

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "?", target.BindVar(1))

	if val, err := target.ByName("bar"); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "222", val)
	}

	if _, err := target.ByName("boz"); err != nil {
		assert.ErrorIs(t, err, ErrParameterNotFound)
	} else {
		t.Fail()
	}
}

type Fixture2 struct {
	Fix    Fixture
	FooFoo string
	BarFoo string
	BazFoo string
}

func TestNewDefaultResolver_2Structs(t *testing.T) {
	core := NewCore(SQLServer)
	target, err := NewDefaultResolver(core,
		&Fixture{"111", "222", "333"},
		&Fixture2{Fixture{"444", "555", "666"}, "777", "888", "999"},
	)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "@p1", target.BindVar(1))

	if val, err := target.ByName("fix.foo"); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "444", val)
	}

	if val, err := target.ByName("barfoo"); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "888", val)
	}

	if _, err := target.ByName("boz"); err != nil {
		assert.ErrorIs(t, err, ErrParameterNotFound)
	} else {
		t.Fail()
	}
}

func TestNewDefaultResolver_ByIndex(t *testing.T) {
	core := NewCore(MySQL)
	target, err := NewDefaultResolver(core, "111", "222", "333")

	if err != nil {
		t.Error(err)
		return
	}

	if val, err := target.ByIndex(1); err != nil {
		t.Error(err)
		return
	} else {
		assert.Equal(t, "111", val)
	}

	if _, err := target.ByIndex(4); err != nil {
		assert.ErrorIs(t, err, ErrParameterNotFound)
	} else {
		t.Fail()
	}
}
