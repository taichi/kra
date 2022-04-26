// Copyright 2022 taichi
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

package pgx

import (
	"github.com/taichi/kra"
)

type Core struct {
	*kra.Core
	hook *Hook
}

func NewCore(db kra.DB, hook *Hook) *Core {
	return &Core{kra.NewCore(db), NewHook(hook)}
}

func (core *Core) Analyze(query string, args ...interface{}) (rawQuery string, vars []interface{}, err error) {
	if analyzer, err := core.hook.Parse(core.Parse, query); err != nil {
		return "", nil, err
	} else if resolver, err := core.hook.NewResolver(core.NewResolver, args...); err != nil {
		return "", nil, err
	} else {
		return analyzer.Analyze(resolver)
	}
}
