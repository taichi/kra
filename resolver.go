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
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ResolveFn func(string) (interface{}, bool, error)

type DefaultValueResolver struct {
	bindVar        BindVar
	originalLength int
	original       []interface{}
	values         []ResolveFn
}

func NewDefaultResolver(core *Core, args ...interface{}) (ValueResolver, error) {
	var values []ResolveFn

	for _, arg := range args {
		switch val := arg.(type) {
		case map[string]interface{}:
			values = append(values, toMapFn(val))
		case NamedArg:
			values = append(values, toNamedArgFn(sql.NamedArg(val)))
		case sql.NamedArg:
			values = append(values, toNamedArgFn(val))
		default:
			if isStruct(arg) {
				if fn, err := toStructFn(core, arg); err != nil {
					return nil, err
				} else {
					values = append(values, fn)
				}
			}
		}
	}

	return &DefaultValueResolver{core.BindVar, len(args), args, values}, nil
}

type BindVar func(int) string

func (db DB) toBindVar() BindVar {
	switch db {
	case PostgreSQL:
		return func(index int) string {
			return fmt.Sprintf("$%d", index)
		}
	case MySQL, SQLite:
		return func(index int) string {
			return "?"
		}
	case SQLServer:
		return func(index int) string {
			return fmt.Sprintf("@p%d", index)
		}
	}
	panic("unknown DB")
}

func (resolver *DefaultValueResolver) BindVar(index int) string {
	return resolver.bindVar(index)
}

func (resolver *DefaultValueResolver) ByIndex(index int) (interface{}, error) {
	aryIndex := index - 1
	if aryIndex < resolver.originalLength {
		return resolver.original[aryIndex], nil
	}
	return nil, fmt.Errorf("index=%d %w", index, ErrParameterNotFound)
}

var ErrParameterNotFound = errors.New("kra: parameter not found")

func (resolver *DefaultValueResolver) ByName(name string) (interface{}, error) {
	for _, fn := range resolver.values {
		if val, ok, _ := fn(name); ok {
			return val, nil
		}
	}
	return nil, fmt.Errorf("name=%s %w", name, ErrParameterNotFound)
}

func toMapFn(arg map[string]interface{}) ResolveFn {
	lowerMap := make(map[string]interface{}, len(arg))
	for key, val := range arg {
		lowerMap[strings.ToLower(key)] = val
	}
	return func(name string) (interface{}, bool, error) {
		res, ok := lowerMap[strings.ToLower(name)]
		return res, ok, nil
	}
}

func toNamedArgFn(arg sql.NamedArg) ResolveFn {
	return func(name string) (interface{}, bool, error) {
		if strings.EqualFold(arg.Name, name) {
			return arg.Value, true, nil
		}
		return nil, false, nil
	}
}

func toStructFn(core *Core, arg interface{}) (ResolveFn, error) {
	root := reflect.ValueOf(arg)
	if def, err := core.Repository.Lookup(root.Type()); err != nil {
		return nil, err
	} else {
		return func(name string) (interface{}, bool, error) {
			if _, value, err := def.ByName(root, name); err != nil {
				return nil, false, err
			} else {
				return value.Interface(), true, nil
			}
		}, nil
	}
}

func isStruct(arg interface{}) bool {
	value := reflect.TypeOf(arg)
	kind := value.Kind()
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
	}
	return kind == reflect.Struct
}

type SilentValueResolver struct {
	resolver ValueResolver
}

func KeepSilent(resolver ValueResolver) ValueResolver {
	return &SilentValueResolver{resolver}
}

func (resolver *SilentValueResolver) BindVar(index int) string {
	return resolver.resolver.BindVar(index)
}

func (resolver *SilentValueResolver) ByIndex(index int) (interface{}, error) {
	if value, err := resolver.resolver.ByIndex(index); err != nil {
		return nil, nil
	} else {
		return value, nil
	}
}

func (resolver *SilentValueResolver) ByName(name string) (interface{}, error) {
	if value, err := resolver.resolver.ByName(name); err != nil {
		return nil, nil
	} else {
		return value, nil
	}
}
