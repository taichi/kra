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

func (hook *RowsHook) Fill() {
	if hook.Next == nil {
		hook.Next = func(invocation *RowsNext) bool {
			return invocation.original()
		}
	}
	if hook.Err == nil {
		hook.Err = func(invocation *RowsErr) error {
			return invocation.original()
		}
	}
	if hook.Scan == nil {
		hook.Scan = func(invocation *RowsScan, dest interface{}) error {
			return invocation.original(dest)
		}
	}
	if hook.Close == nil {
		hook.Close = func(invocation *RowsClose) error {
			return invocation.original()
		}
	}
}

type rowsInvocation struct {
	Receiver *Rows
	hooks    []*RowsHook
	length   int
	index    int
}

type RowsNext struct {
	rowsInvocation
	original func() bool
}

func NewRowsNext(recv *Rows, original func() bool) *RowsNext {
	hooks := recv.core.hooks.Rows
	return &RowsNext{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsNext) Proceed() bool {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Next(invocation)
	}
	return invocation.original()
}

type RowsErr struct {
	rowsInvocation
	original func() error
}

func NewRowsErr(recv *Rows, original func() error) *RowsErr {
	hooks := recv.core.hooks.Rows
	return &RowsErr{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsErr) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Err(invocation)
	}
	return invocation.original()
}

type RowsScan struct {
	rowsInvocation
	original func(dest interface{}) error
}

func NewRowsScan(recv *Rows, original func(dest interface{}) error) *RowsScan {
	hooks := recv.core.hooks.Rows
	return &RowsScan{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsScan) Proceed(dest interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Scan(invocation, dest)
	}
	return invocation.original(dest)
}

type RowsClose struct {
	rowsInvocation
	original func() error
}

func NewRowsClose(recv *Rows, original func() error) *RowsClose {
	hooks := recv.core.hooks.Rows
	return &RowsClose{rowsInvocation{recv, hooks, len(hooks), 0}, original}
}

func (invocation *RowsClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}
