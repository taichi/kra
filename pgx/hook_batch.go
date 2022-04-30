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

import "github.com/jackc/pgconn"

type BatchHook struct {
	Queue func(invocation *BatchQueue, query string, args ...interface{}) error
}

type BatchQueue invocation[Batch, BatchHook, func(query string, args ...interface{}) error]

func NewBatchQueue(recv *Batch, original func(query string, args ...interface{}) error) *BatchQueue {
	hooks := recv.core.hooks.Batch
	return &BatchQueue{recv, hooks, len(hooks), 0, original}
}

func (invocation *BatchQueue) Proceed(query string, args ...interface{}) error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Queue != nil {
		return invocation.hooks[current].Queue(invocation, query, args...)
	}
	return invocation.original(query, args...)
}

type BatchResultsHook struct {
	Close func(invocation *BatchResultsClose) error
	Exec  func(invocation *BatchResultsExec) (pgconn.CommandTag, error)
	Query func(invocation *BatchResultsQuery) (*Rows, error)
}

type BatchResultsClose invocation[BatchResults, BatchResultsHook, func() error]

func NewBatchResultsClose(recv *BatchResults, original func() error) *BatchResultsClose {
	hooks := recv.core.hooks.BatchResults
	return &BatchResultsClose{recv, hooks, len(hooks), 0, original}
}

func (invocation *BatchResultsClose) Proceed() error {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Close != nil {
		return invocation.hooks[current].Close(invocation)
	}
	return invocation.original()
}

type BatchResultsExec invocation[BatchResults, BatchResultsHook, func() (pgconn.CommandTag, error)]

func NewBatchResultsExec(recv *BatchResults, original func() (pgconn.CommandTag, error)) *BatchResultsExec {
	hooks := recv.core.hooks.BatchResults
	return &BatchResultsExec{recv, hooks, len(hooks), 0, original}
}

func (invocation *BatchResultsExec) Proceed() (pgconn.CommandTag, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Exec != nil {
		return invocation.hooks[current].Exec(invocation)
	}
	return invocation.original()
}

type BatchResultsQuery invocation[BatchResults, BatchResultsHook, func() (*Rows, error)]

func NewBatchResultsQuery(recv *BatchResults, original func() (*Rows, error)) *BatchResultsQuery {
	hooks := recv.core.hooks.BatchResults
	return &BatchResultsQuery{recv, hooks, len(hooks), 0, original}
}

func (invocation *BatchResultsQuery) Proceed() (*Rows, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Query != nil {
		return invocation.hooks[current].Query(invocation)
	}
	return invocation.original()
}
