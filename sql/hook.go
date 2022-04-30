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

package sql

import (
	"github.com/taichi/kra"
)

type HookHolster struct {
	Core []*kra.CoreHook
	Conn []*ConnHook
	DB   []*DBHook
	Tx   []*TxHook
	Stmt []*StmtHook
	Rows []*RowsHook
}

func NewHookHolster(hooks ...kra.Hook) *HookHolster {
	holster := new(HookHolster)
	for _, hook := range hooks {
		if hook == nil {
			continue
		}
		hook.Fill()
		switch h := hook.(type) {
		case *kra.CoreHook:
			holster.Core = append(holster.Core, h)
		case *ConnHook:
			holster.Conn = append(holster.Conn, h)
		case *DBHook:
			holster.DB = append(holster.DB, h)
		case *TxHook:
			holster.Tx = append(holster.Tx, h)
		case *StmtHook:
			holster.Stmt = append(holster.Stmt, h)
		case *RowsHook:
			holster.Rows = append(holster.Rows, h)
		default:
			// do nothing
		}
	}
	return holster
}
