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
)

type QueryAnalyzer interface {
	Verify(ValueResolver) error
	Analyze(ValueResolver) (query string, vars []interface{}, err error)
}

type ValueResolver interface {
	BindVar(index int) string
	ByIndex(index int) (interface{}, error)
	ByName(name string) (interface{}, error)
}

type Rows interface {
	Close() error
	Err() error
	Next() bool
	Columns() ([]string, error)
	Scan(dest ...interface{}) error
}

type Transformer interface {
	Transform(src Rows, dest interface{}) error
	TransformAll(src Rows, dest interface{}) error
}

type Core struct {
	BindVar        func(index int) string
	Parse          func(query string) (QueryAnalyzer, error)
	NewResolver    func(args ...interface{}) (ValueResolver, error)
	NewTransformer func() Transformer
	TagName        string
	Repository     *TypeRepository
}

func (core *Core) Analyze(hooks []*CoreHook, query string, args ...interface{}) (rawQuery string, vars []interface{}, err error) {
	if analyzer, err := NewCoreParse(core, hooks).Proceed(query); err != nil {
		return "", nil, err
	} else if resolver, err := NewCoreNewResolver(core, hooks).Proceed(args...); err != nil {
		return "", nil, err
	} else {
		return analyzer.Analyze(resolver)
	}
}

var ErrLackOfQueryParameters = errors.New("kra: require example parameters for prepare query with IN operator")
var ErrNoRecord = errors.New("kra: no record")

type NamedArg sql.NamedArg
