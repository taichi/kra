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
	"fmt"
	"reflect"
	"strings"

	"github.com/georgysavva/scany/dbscan"
	"github.com/mitchellh/mapstructure"
)

type DB uint8

const (
	PostgreSQL DB = iota
	MySQL
	SQLite
	SQLServer
)

func NewCore(db DB) *Core {
	core := &Core{}
	core.BindVar = db.toBindVar()
	core.Parse = func(s string) (QueryAnalyzer, error) {
		return NewQuery(s)
	}
	core.NewResolver = func(args ...interface{}) (ValueResolver, error) {
		return NewDefaultResolver(core, args...)
	}

	trans := &DefaultTransformer{}
	core.NewTransformer = func() Transformer {
		return trans
	}
	core.TagName = "db"

	return core
}

func (core *Core) Analyze(query string, args ...interface{}) (rawQuery string, vars []interface{}, err error) {
	if analyzer, err := core.Parse(query); err != nil {
		return "", nil, err
	} else if resolver, err := core.NewResolver(args...); err != nil {
		return "", nil, err
	} else {
		return analyzer.Analyze(resolver)
	}
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

type DefaultValueResolver struct {
	bindVar        BindVar
	originalLength int
	original       []interface{}
	values         map[string]interface{}
}

func (resolver *DefaultValueResolver) BindVar(index int) string {
	return resolver.bindVar(index)
}

func (resolver *DefaultValueResolver) ByIndex(index int) (interface{}, error) {
	aryIndex := index - 1
	if aryIndex < resolver.originalLength {
		return resolver.original[aryIndex], nil
	}
	return nil, nil
}

func (resolver *DefaultValueResolver) ByName(name string) (interface{}, error) {
	val := resolver.values[strings.ToLower(name)]
	return val, nil
}

func NewDefaultResolver(core *Core, args ...interface{}) (ValueResolver, error) {
	result := map[string]interface{}{}
	maps := []map[string]interface{}{}

	for _, arg := range args {
		switch val := arg.(type) {
		case map[string]interface{}:
			maps = append(maps, val)
		case sql.NamedArg:
			result[strings.ToLower(val.Name)] = val.Value
		default:
			if isStruct(arg) {
				if newmap, err := toMap(arg); err != nil {
					return nil, err
				} else {
					maps = append(maps, newmap)
				}
			}
		}
	}

	for _, m := range maps {
		for key, value := range m {
			result[strings.ToLower(key)] = value
		}
	}

	return &DefaultValueResolver{core.BindVar, len(args), args, result}, nil
}

func toMap(arg interface{}) (map[string]interface{}, error) {
	var output map[string]interface{}
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &output,
		TagName:  "db",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(arg); err != nil {
		return nil, err
	}

	return output, nil
}

func isStruct(arg interface{}) bool {
	value := reflect.ValueOf(arg)
	kind := value.Kind()
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
	}
	return kind == reflect.Struct
}

type DefaultTransformer struct{}

func (*DefaultTransformer) Transform(src Rows, dst interface{}) error {
	return dbscan.ScanOne(dst, src)
}
func (*DefaultTransformer) TransformAll(src Rows, dst interface{}) error {
	return dbscan.ScanAll(dst, src)
}
