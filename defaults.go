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

type DB uint8

const (
	PostgreSQL DB = iota
	MySQL
	SQLite
	SQLServer
)

const DefaultTagName = "db"

func NewCore(db DB) *Core {
	core := &Core{}
	core.BindVar = db.toBindVar()
	core.Parse = func(s string) (QueryAnalyzer, error) {
		return NewQuery(s)
	}
	core.NewResolver = func(args ...interface{}) (ValueResolver, error) {
		return NewDefaultResolver(core, args...)
	}

	trans := NewDefaultTransformer(core)
	core.NewTransformer = func() Transformer {
		return trans
	}
	core.TagName = DefaultTagName
	core.Repository = NewTypeRepository(core)

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
