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

package kra

type CoreHook struct {
	Parse          func(invocation *CoreParse, query string) (QueryAnalyzer, error)
	NewResolver    func(invocation *CoreNewResolver, args ...interface{}) (ValueResolver, error)
	NewTransformer func(invocation *CoreNewTransformer) Transformer
}

type coreInvocation[FN any] struct {
	Receiver *Core
	hooks    []*CoreHook
	length   int
	index    int
	original FN
}

type CoreParse coreInvocation[func(query string) (QueryAnalyzer, error)]

func NewCoreParse(recv *Core, hooks []*CoreHook) *CoreParse {
	return &CoreParse{recv, hooks, len(hooks), 0, recv.Parse}
}

func (invocation *CoreParse) Proceed(query string) (QueryAnalyzer, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].Parse != nil {
		return invocation.hooks[current].Parse(invocation, query)
	}
	return invocation.original(query)
}

type CoreNewResolver coreInvocation[func(args ...interface{}) (ValueResolver, error)]

func NewCoreNewResolver(recv *Core, hooks []*CoreHook) *CoreNewResolver {
	return &CoreNewResolver{recv, hooks, len(hooks), 0, recv.NewResolver}
}

func (invocation *CoreNewResolver) Proceed(args ...interface{}) (ValueResolver, error) {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].NewResolver != nil {
		return invocation.hooks[current].NewResolver(invocation, args...)
	}
	return invocation.original(args...)
}

type CoreNewTransformer coreInvocation[func() Transformer]

func NewCoreNewTransformer(recv *Core, hooks []*CoreHook) *CoreNewTransformer {
	return &CoreNewTransformer{recv, hooks, len(hooks), 0, recv.NewTransformer}
}

func (invocation *CoreNewTransformer) Proceed() Transformer {
	current := invocation.index
	invocation.index++
	if current < invocation.length && invocation.hooks[current].NewTransformer != nil {
		return invocation.hooks[current].NewTransformer(invocation)
	}
	return invocation.original()
}
