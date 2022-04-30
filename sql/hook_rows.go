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

type RowsHook struct {
	Next  func(invocation *RowsNext) bool
	Err   func(invocation *RowsErr) error
	Scan  func(invocation *RowsScan, dest interface{}) error
	Close func(invocation *RowsClose) error
}

type RowsNext invocation[Rows, RowsHook, func() bool]

func NewRowsNext(recv *Rows, original func() bool) *RowsNext {
	hooks := recv.core.hooks.Rows
	return &RowsNext{recv, hooks, len(hooks), 0, original}
}

func (invocation *RowsNext) Proceed() bool {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Next != nil {
		return invocation.hooks[current].Next(invocation)
	}
	return invocation.original()
}

type RowsErr invocation[Rows, RowsHook, func() error]

func NewRowsErr(recv *Rows, original func() error) *RowsErr {
	hooks := recv.core.hooks.Rows
	return &RowsErr{recv, hooks, len(hooks), 0, original}
}

func (invocation *RowsErr) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Err != nil {
		return invocation.hooks[current].Err(invocation)
	}
	return invocation.original()
}

type RowsScan invocation[Rows, RowsHook, func(dest interface{}) error]

func NewRowsScan(recv *Rows, original func(dest interface{}) error) *RowsScan {
	hooks := recv.core.hooks.Rows
	return &RowsScan{recv, hooks, len(hooks), 0, original}
}

func (invocation *RowsScan) Proceed(dest interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Scan != nil {
		return invocation.hooks[current].Scan(invocation, dest)
	}
	return invocation.original(dest)
}

type RowsClose invocation[Rows, RowsHook, func() error]

func NewRowsClose(recv *Rows, original func() error) *RowsClose {
	hooks := recv.core.hooks.Rows
	return &RowsClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *RowsClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
