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
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type TypeRepository struct {
	mapping sync.Map // map[reflect.Type]*StructDef
	core    *Core
}

func NewTypeRepository(core *Core) *TypeRepository {
	return &TypeRepository{core: core}
}

func (repo *TypeRepository) Lookup(root reflect.Type) (*StructDef, error) {
	root = Indirect(root)
	if def, err := repo.LookupOrTraverse(root); err != nil {
		return nil, err
	} else {
		return def, nil
	}
}

func (repo *TypeRepository) LookupOrTraverse(target reflect.Type, history ...*StructDef) (*StructDef, error) {
	for _, def := range history {
		if target == def.target { // skip recursive type reference
			return def, nil
		}
	}

	if def, ok := repo.mapping.Load(target); ok {
		return def.(*StructDef), nil
	}

	newDef := &StructDef{target: target}
	history = append(history, newDef)
	if def, err := repo.Traverse(target, history...); err != nil {
		return nil, err
	} else {
		*newDef = *def
		repo.mapping.Store(target, newDef)
		return def, nil
	}
}

var ErrUnsupportedValueType = errors.New("unsupported value type")

func (repo *TypeRepository) Traverse(target reflect.Type, history ...*StructDef) (*StructDef, error) {
	targetType := Indirect(target)
	if targetType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type: %s %w", targetType.Name(), ErrUnsupportedValueType)
	}

	length := targetType.NumField()
	members := map[string]FieldDef{}
	for index := 0; index < length; index++ {
		field := targetType.Field(index)
		name, options := parseTag(repo.core, &field)
		if 0 < len(field.PkgPath) { // skip unexported field
			members[name] = FieldDef{nil, nil, true, options}
			continue
		}

		var child *StructDef
		if Indirect(field.Type).Kind() == reflect.Struct {
			if found, err := repo.LookupOrTraverse(field.Type, history...); err != nil {
				return nil, err
			} else {
				child = found
			}
		}
		members[name] = FieldDef{[]int{index}, child, false, options}

		if field.Anonymous && child != nil && child.members != nil {
			for key, val := range child.members {
				if _, ok := members[key]; ok == false { // don't override by embedded members
					members[key] = FieldDef{append([]int{index}, val.Indices...), val.Self, false, options}
				}
			}
		}
	}

	return &StructDef{target, members}, nil
}

func parseTag(core *Core, field *reflect.StructField) (name string, options map[string]string) {
	name = field.Name
	options = map[string]string{}

	if tag, ok := field.Tag.Lookup(core.TagName); ok {
		var tagged string
		if index := strings.Index(tag, ","); index < 0 {
			tagged = tag
		} else {
			tagged = tag[:index]
			for _, elem := range strings.Split(tag[index+1:], ",") {
				kv := strings.Split(elem, "=")
				if 1 < len(kv) {
					options[kv[0]] = kv[1]
				} else {
					options[kv[0]] = ""
				}
			}
		}
		if 0 < len(tagged) {
			options["name"] = tagged
			if tagged != "-" {
				name = tagged
			}
		}
	}

	return strings.ToLower(name), options
}

type StructDef struct {
	target  reflect.Type
	members map[string]FieldDef
}

type FieldDef struct {
	Indices    []int
	Self       *StructDef
	Unexported bool
	Options    map[string]string
}

var ErrFieldNotFound = errors.New("field not found")
var ErrFieldUnexported = errors.New("field not exported")

func (def *StructDef) ByName(root reflect.Value, name string) (*FieldDef, *reflect.Value, error) {
	names := strings.Split(strings.ToLower(name), ".")
	return visitByName(def, &root, names)
}

func visitByName(def *StructDef, value *reflect.Value, names []string) (*FieldDef, *reflect.Value, error) {
	cur := names[0]
	if fdef, ok := def.members[cur]; ok {
		if fdef.Unexported {
			return nil, nil, fmt.Errorf("name: %s %w", cur, ErrFieldUnexported)
		}

		var val reflect.Value = *value
		for _, index := range fdef.Indices {
			if val.Kind() == reflect.Ptr && val.IsNil() {
				break
			}
			val = reflect.Indirect(val).Field(index)
		}
		if fdef.Self != nil && 1 < len(names) {
			return visitByName(fdef.Self, &val, names[1:])
		}
		return &fdef, &val, nil
	} else {
		return nil, nil, fmt.Errorf("name: %s %w", cur, ErrFieldNotFound)
	}
}

func Indirect(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
