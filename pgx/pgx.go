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
	baseHook := NewHook()
	if hook != nil {
		if hook.Parse != nil {
			baseHook.Parse = hook.Parse
		}
		if hook.NewResolver != nil {
			baseHook.NewResolver = hook.NewResolver
		}
		if hook.NewTransformer != nil {
			baseHook.NewTransformer = hook.NewTransformer
		}
		if hook.BeginTx != nil {
			baseHook.BeginTx = hook.BeginTx
		}
		if hook.BeginTxFunc != nil {
			baseHook.BeginTxFunc = hook.BeginTxFunc
		}
		if hook.CopyFrom != nil {
			baseHook.CopyFrom = hook.CopyFrom
		}
		if hook.Exec != nil {
			baseHook.Exec = hook.Exec
		}
		if hook.Ping != nil {
			baseHook.Ping = hook.Ping
		}
		if hook.Prepare != nil {
			baseHook.Prepare = hook.Prepare
		}
		if hook.Query != nil {
			baseHook.Query = hook.Query
		}
		if hook.SendBatch != nil {
			baseHook.SendBatch = hook.SendBatch
		}
		if hook.Find != nil {
			baseHook.Find = hook.Find
		}
		if hook.FindAll != nil {
			baseHook.FindAll = hook.FindAll
		}
	}
	return &Core{kra.NewCore(db), baseHook}
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
