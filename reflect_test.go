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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructDef_Simple(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		Cccc int
		dddd uint
	}

	aaa := Aaaa{"foo", 32, 11}
	core := NewCore(PostgreSQL)
	core.TagName = ""
	repo := NewTypeRepository(core)

	value := reflect.ValueOf(aaa)
	def, err := repo.Lookup(value.Type())
	if err != nil {
		t.Error(err)
		return
	}

	if _, found, err := def.ByName(value, "bbbb"); err != nil {
		t.Error(err)
	} else if assert.NotNil(t, found) {
		assert.Equal(t, "foo", found.Interface())
	}
}

func TestStructDef_NamedByTag(t *testing.T) {
	type Aaaa struct {
		Bbbb string `db:"zzz"`
		Cccc int
		dddd uint
	}

	aaa := Aaaa{"foo", 32, 11}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(aaa)
	def, err := repo.Lookup(value.Type())
	if err != nil {
		t.Error(err)
		return
	}

	if _, value, err := def.ByName(value, "zzz"); err != nil {
		t.Error(err)
	} else if assert.NotNil(t, value) {
		assert.Equal(t, "foo", value.Interface())
	}
}

func TestStructDef_Unexported(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		Cccc int
		dddd uint
	}

	aaa := Aaaa{"foo", 32, 11}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(aaa)

	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		_, _, err := def.ByName(value, "dddd")
		assert.ErrorIs(t, err, ErrFieldUnexported)
	}
}

func TestStructDef_NotFound(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		Cccc int
		dddd uint
	}

	aaa := Aaaa{"foo", 32, 11}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(aaa)

	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		_, _, err := def.ByName(value, "zzz")
		assert.ErrorIs(t, err, ErrFieldNotFound)
	}
}

func TestStructDef_Nested(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		dddd uint
	}

	type Eeee struct {
		Aa *Aaaa
		Bb string
	}

	type Ffff struct {
		Eee Eeee
	}

	fff := Ffff{Eeee{&Aaaa{"foo", 32}, "bar"}}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)

	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		if _, found, err := def.ByName(value, "eee.aa.bbbb"); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, "foo", found.Interface())
		}
	}
}

func TestStructDef_Nested_VisitBreak(t *testing.T) {
	type Aaaa struct {
		Bbbb string
	}

	type Eeee struct {
		Aa *Aaaa
		Bb string
	}

	type Ffff struct {
		Eee Eeee
		Aaa *Aaaa
	}

	fff := Ffff{Eeee{nil, "bar"}, nil}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)

	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		if _, found, err := def.ByName(value, "eee.aa.bbbb"); err != nil {
			t.Error(err)
		} else {
			assert.Nil(t, found.Interface())
		}
		if _, found, err := def.ByName(value, "aaa.bbbb"); err != nil {
			t.Error(err)
		} else {
			assert.Nil(t, found.Interface())
		}
	}
}

func TestStructDef_Embedded(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		Cccc int
		dddd uint
	}

	type Eeee struct {
		*Aaaa
		Bb string
	}

	type Ffff struct {
		Eee Eeee
	}

	fff := Ffff{Eeee{&Aaaa{"foo", 32, 11}, "bar"}}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)
	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else if _, found, err := def.ByName(value, "eee.bbbb"); err != nil {
		t.Error(err)
	} else if assert.NotNil(t, found) {
		assert.Equal(t, "foo", found.Interface())
	}
}

func TestStructDef_Embedded_Multiple(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		Cccc int
		dddd uint
	}

	type Eeee struct {
		Bb string
	}

	type Ffff struct {
		Aaaa
		Eeee
	}

	fff := Ffff{Aaaa{"foo", 32, 21}, Eeee{"bar"}}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)
	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		if _, found, err := def.ByName(value, "bbbb"); err != nil {
			t.Error(err)
		} else if assert.NotNil(t, found) {
			assert.Equal(t, "foo", found.Interface())
		}
		if _, found, err := def.ByName(value, "bb"); err != nil {
			t.Error(err)
		} else if assert.NotNil(t, found) {
			assert.Equal(t, "bar", found.Interface())
		}
	}
}

func TestStructDef_Embedded_CacheHit(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		Cccc int
		dddd uint
	}

	type Eeee struct {
		Aaaa
		Optional *Aaaa
	}

	fff := Eeee{Aaaa{"foo", 32, 21}, nil}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)
	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		if _, found, err := def.ByName(value, "bbbb"); err != nil {
			t.Error(err)
		} else if assert.NotNil(t, found) {
			assert.Equal(t, "foo", found.Interface())
		}
		if _, found, err := def.ByName(value, "optional.cccc"); err != nil {
			t.Error(err)
		} else {
			assert.Nil(t, found.Interface())
		}
	}
}

func TestStructDef_Embedded_Nested(t *testing.T) {
	type Aaaa struct {
		Bbbb string
		Cccc int
		dddd uint
	}

	type Eeee struct {
		*Aaaa
		Bb string
	}

	type Ffff struct {
		Eeee
	}

	fff := Ffff{Eeee{&Aaaa{"foo", 32, 11}, "bar"}}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)
	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else if _, found, err := def.ByName(value, "bbbb"); err != nil {
		t.Error(err)
	} else if assert.NotNil(t, found) {
		assert.Equal(t, "foo", found.Interface())
	}
}

func TestStructDef_Embedded_Tag(t *testing.T) {
	type Aaaa struct {
		Bbbb string `db:"bb"`
		Cccc int    `db:"zzz"`
		dddd uint
	}

	type Eeee struct {
		*Aaaa
		Bb string
	}

	type Ffff struct {
		Eee Eeee
	}

	fff := Ffff{Eeee{&Aaaa{"foo", 32, 11}, "bar"}}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)
	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		if _, found, err := def.ByName(value, "eee.bb"); err != nil {
			t.Error(err)
		} else if assert.NotNil(t, found) {
			assert.Equal(t, "bar", found.Interface())
		}
		if _, found, err := def.ByName(value, "eee.zzz"); err != nil {
			t.Error(err)
		} else if assert.NotNil(t, found) {
			assert.Equal(t, 32, found.Interface())
		}
	}
}

func TestStructDef_Embedded_WithSameName(t *testing.T) {
	type Aaaa struct {
		Bbbb string `db:"aaaa"`
		Cccc int
	}

	type Eeee struct {
		*Aaaa
		Bb string
	}

	type Ffff struct {
		Eee Eeee
	}

	fff := Ffff{Eeee{&Aaaa{"foo", 32}, "bar"}}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(fff)
	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else {
		if _, found, err := def.ByName(value, "eee.aaaa"); err != nil {
			t.Error(err)
		} else if assert.NotNil(t, found) {
			assert.Equal(t, fff.Eee.Aaaa, found.Interface())
		}
	}
}

type Recur_Aaaa struct {
	Bbbb string
	Cccc int
	dddd uint
}

type Recur_Eeee struct {
	Aa *Recur_Aaaa
	Gg *Recur_Gggg
}

type Recur_Ffff struct {
	Eee *Recur_Eeee
}

type Recur_Gggg struct {
	Fff *Recur_Ffff
}

func TestStructDef_Recursive(t *testing.T) {

	ggg := &Recur_Gggg{&Recur_Ffff{&Recur_Eeee{&Recur_Aaaa{"foo", 32, 11},
		&Recur_Gggg{&Recur_Ffff{&Recur_Eeee{&Recur_Aaaa{"bar", 79, 43},
			nil}}}}}}

	repo := NewTypeRepository(NewCore(PostgreSQL))

	value := reflect.ValueOf(ggg)
	if def, err := repo.Lookup(value.Type()); err != nil {
		t.Error(err)
	} else if _, found, err := def.ByName(value, "fff.eee.gg.fff.eee.aa.bbbb"); err != nil {
		t.Error(err)
	} else if assert.NotNil(t, found) {
		assert.Equal(t, "bar", found.Interface())
	}
}

func Test_Traverse(t *testing.T) {
	repo := NewTypeRepository(NewCore(PostgreSQL))

	_, err := repo.Lookup(reflect.TypeOf("33"))
	assert.ErrorIs(t, err, ErrUnsupportedValueType)
}

func TestStructDef_FieldOptions(t *testing.T) {
	type Aaaa struct {
		Bbbb string `db:",foo=bar"`
		Cccc int    `db:"-"`
		dddd uint
	}

	aaa := Aaaa{"foo", 32, 11}
	core := NewCore(PostgreSQL)
	core.TagName = "db"
	repo := NewTypeRepository(core)

	value := reflect.ValueOf(aaa)
	def, err := repo.Lookup(value.Type())
	if err != nil {
		t.Error(err)
		return
	}

	if def, found, err := def.ByName(value, "bbbb"); err != nil {
		t.Error(err)
	} else if assert.NotNil(t, found) {
		assert.Equal(t, "foo", found.Interface())
		assert.Equal(t, "bar", def.Options["foo"])
		assert.Equal(t, "", def.Options["name"])
	}

	if def, found, err := def.ByName(value, "cccc"); err != nil {
		t.Error(err)
	} else if assert.NotNil(t, found) {
		assert.Equal(t, 32, found.Interface())
		assert.Equal(t, "-", def.Options["name"])
	}
}
