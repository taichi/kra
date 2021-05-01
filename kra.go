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
)

type QueryAnalyzer interface {
	Analyze(ValueResolver) (query string, vars []interface{}, err error)
	DynamicParameters() []string
}

type ValueResolver interface {
	BindVar(index int) string
	ByIndex(index int) (interface{}, error)
	ByName(name string) (interface{}, error)
}

type Transformer interface {
	Transform(src *sql.Rows, dst interface{}) error
	TransformAll(src *sql.Rows, dst interface{}) error
}

type Core struct {
	BindVar        func(index int) string
	Parse          func(query string) (QueryAnalyzer, error)
	NewResolver    func(args ...interface{}) (ValueResolver, error)
	NewTransformer func() Transformer
}
