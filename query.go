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
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/taichi/kra/parser"
)

type Query struct {
	parts             []StmtPart
	dynamicParameters []string
}

func NewQuery(src string) (*Query, error) {
	visitor := new(PartsCollector)
	tree := parser.Parse(src)
	if res := tree.Accept(visitor); res != nil {
		if res, ok := res.(error); ok {
			return nil, res
		}
	}

	if err := visitor.Validate(); err != nil {
		return nil, err
	}

	return &Query{visitor.parts, visitor.dynamicParameters}, nil
}

func (query *Query) Verify(resolver ValueResolver) error {
	var notFound []string
	for _, name := range query.dynamicParameters {
		if val, err := resolver.ByName(name); err != nil || val == nil {
			notFound = append(notFound, name)
		}
	}

	if 0 < len(notFound) {
		return fmt.Errorf("require parameters=%v %w", notFound, ErrLackOfQueryParameters)
	}
	return nil
}

func (query *Query) Analyze(resolver ValueResolver) (rawQuery string, vars []interface{}, err error) {
	state := ResolvingState{0, &strings.Builder{}, nil}

	for _, part := range query.parts {
		if err := part(&state, resolver); err != nil {
			return "", nil, err
		}
	}

	return strings.TrimSpace(state.query.String()), state.values, nil
}

type ResolvingState struct {
	parameterIndex int
	query          *strings.Builder
	values         []interface{}
}

func (state *ResolvingState) NextIndex() int {
	state.parameterIndex += 1
	return state.parameterIndex
}

func (state *ResolvingState) AppendStmt(part string) {
	state.query.WriteRune(' ')
	state.query.WriteString(part)
}

func (state *ResolvingState) AppendVar(val interface{}) {
	state.values = append(state.values, val)
}

func (state *ResolvingState) ConcatVar(val []interface{}) {
	state.values = append(state.values, val...)
}

type BindingStyle uint

const (
	NAMED BindingStyle = 1 << iota
	QMARK
	DDEC
)

type PartsCollector struct {
	parts             []StmtPart
	statements        int
	style             BindingStyle
	dynamicParameters []string
}

func (collector *PartsCollector) Add(fn func() (StmtPart, error)) error {
	if part, err := fn(); err != nil {
		return err
	} else {
		collector.parts = append(collector.parts, part)
		return nil
	}
}

func (collector *PartsCollector) Use(style BindingStyle) {
	collector.style |= style
}

func (collector *PartsCollector) Use2orMoreStyles() bool {
	NQ := NAMED | QMARK
	ND := NAMED | DDEC
	QD := QMARK | DDEC
	return collector.style&NQ == NQ || collector.style&ND == ND || collector.style&QD == QD
}

var Err2orMoreStatements = errors.New("kra: 2 or more statements in 1 query. Use batch queries. ")
var Err2orMoreParameterStyles = errors.New("kra: 2 or more bind variables style contains in 1 statement. Use only 1 bind variables style in 1 query, such as ? or $1,$2,$3... or :foo,:bar,:baz... ")

func (collector *PartsCollector) Validate() error {
	if 1 < collector.statements {
		return Err2orMoreStatements
	}
	if collector.Use2orMoreStyles() {
		return Err2orMoreParameterStyles
	}
	return nil
}

func VisitChildren(visitor parser.NamedVisitor, node antlr.RuleNode) interface{} {
	for _, kid := range node.GetChildren() {
		if tree, ok := kid.(antlr.ParseTree); ok {
			if err := tree.Accept(visitor); err != nil {
				return err
			}
		}
	}
	return nil
}

func (collector *PartsCollector) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(collector)
}

func (collector *PartsCollector) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (collector *PartsCollector) VisitErrorNode(node antlr.ErrorNode) interface{} { return nil }

func (collector *PartsCollector) VisitChildren(node antlr.RuleNode) interface{} {
	return VisitChildren(collector, node)
}

func (collector *PartsCollector) VisitParse(ctx *parser.ParseContext) interface{} {
	return VisitChildren(collector, ctx)
}

func (collector *PartsCollector) VisitStmt(ctx *parser.StmtContext) interface{} {
	collector.statements++
	return VisitChildren(collector, ctx)
}

type ParameterVisitor struct {
	parser.BaseNamedVisitor
	named []string
	other []string
}

func (visitor *ParameterVisitor) VisitInExpr(ctx *parser.InExprContext) interface{} {
	return VisitChildren(visitor, ctx)
}

func (visitor *ParameterVisitor) VisitParameter(ctx *parser.ParameterContext) interface{} {
	return VisitChildren(visitor, ctx)
}

func (visitor *ParameterVisitor) VisitNamedParamter(ctx *parser.NamedParamterContext) interface{} {
	visitor.named = append(visitor.named, ctx.GetText())
	return nil
}

func (visitor *ParameterVisitor) VisitQmarkParameter(ctx *parser.QmarkParameterContext) interface{} {
	visitor.other = append(visitor.other, ctx.GetText())
	return nil
}

func (visitor *ParameterVisitor) VisitDDecParameter(ctx *parser.DDecParameterContext) interface{} {
	visitor.other = append(visitor.other, ctx.GetText())
	return nil
}

func (visitor *ParameterVisitor) VisitStaticParameter(ctx *parser.StaticParameterContext) interface{} {
	visitor.other = append(visitor.other, ctx.GetText())
	return nil
}

func (collector *PartsCollector) VisitInExpr(ctx *parser.InExprContext) interface{} {
	pVisitor := ParameterVisitor{}
	ctx.Accept(&pVisitor)
	if 0 < len(pVisitor.named) {
		// IN句の一部にnamed parameterが含まれている場合のみ、BindVarを自動的に決める
		collector.dynamicParameters = append(collector.dynamicParameters, pVisitor.named...)
		return collector.Add(func() (StmtPart, error) {
			return NewInPart(ctx.IN().GetText(), pVisitor.named[0])
		})
	} else {
		return collector.Add(func() (StmtPart, error) {
			return NewStringPart(ctx.GetText())
		})
	}
}

func (collector *PartsCollector) VisitParameter(ctx *parser.ParameterContext) interface{} {
	return VisitChildren(collector, ctx)
}

func (collector *PartsCollector) VisitAnyStmtParts(ctx *parser.AnyStmtPartsContext) interface{} {
	return collector.Add(func() (StmtPart, error) {
		return NewStringPart(ctx.GetText())
	})
}

func (collector *PartsCollector) VisitNamedParamter(ctx *parser.NamedParamterContext) interface{} {
	collector.Use(NAMED)
	return collector.Add(func() (StmtPart, error) {
		return NewNamedParameterPart(ctx.GetText())
	})
}

func (collector *PartsCollector) VisitQmarkParameter(ctx *parser.QmarkParameterContext) interface{} {
	collector.Use(QMARK)
	return collector.Add(func() (StmtPart, error) {
		return NewQMarkParameterPart(ctx.GetText())
	})
}

func (collector *PartsCollector) VisitDDecParameter(ctx *parser.DDecParameterContext) interface{} {
	collector.Use(DDEC)
	return collector.Add(func() (StmtPart, error) {
		return NewDMarkParameterPart(ctx.GetText())
	})
}

func (collector *PartsCollector) VisitStaticParameter(ctx *parser.StaticParameterContext) interface{} {
	return collector.Add(func() (StmtPart, error) {
		return NewStringPart(ctx.GetText())
	})
}

type StmtPart func(*ResolvingState, ValueResolver) error

func NewStringPart(src string) (StmtPart, error) {
	return func(state *ResolvingState, resolver ValueResolver) error {
		state.AppendStmt(src)
		return nil
	}, nil
}

var ErrEmptySlice = errors.New("kra: empty slice set to in query parameter")

func NewInPart(in, src string) (StmtPart, error) {
	name := src[1:]
	return func(state *ResolvingState, resolver ValueResolver) error {
		if val, err := resolver.ByName(name); err != nil {
			return err
		} else {
			values := AsSlice(val)
			length := len(values)
			if length < 1 {
				return fmt.Errorf("name=%s %w", name, ErrEmptySlice)
			}
			var vars []string
			for i := 0; i < length; i++ {
				vars = append(vars, resolver.BindVar(state.NextIndex()))
			}
			stmt := fmt.Sprintf("%s (%s)", in, strings.Join(vars, " ,"))
			state.AppendStmt(stmt)
			state.ConcatVar(values)
		}
		return nil
	}, nil
}

func NewNamedParameterPart(src string) (StmtPart, error) {
	name := src[1:]
	return func(state *ResolvingState, resolver ValueResolver) error {
		if val, err := resolver.ByName(name); err != nil {
			return err
		} else {
			state.AppendStmt(resolver.BindVar(state.NextIndex()))
			state.AppendVar(val)
		}
		return nil
	}, nil
}

func NewQMarkParameterPart(src string) (StmtPart, error) {
	return func(state *ResolvingState, resolver ValueResolver) error {
		index := state.NextIndex()
		if val, err := resolver.ByIndex(index); err != nil {
			return err
		} else {
			state.AppendStmt(resolver.BindVar(index))
			state.AppendVar(val)
		}
		return nil
	}, nil
}

func NewDMarkParameterPart(src string) (StmtPart, error) {
	srcIndex, err := strconv.Atoi(src[1:])
	if err != nil {
		return nil, err
	}
	return func(state *ResolvingState, resolver ValueResolver) error {
		if val, err := resolver.ByIndex(srcIndex); err != nil {
			return err
		} else {
			state.AppendStmt(resolver.BindVar(state.NextIndex()))
			state.AppendVar(val)
		}
		return nil
	}, nil
}

func NewStaticParameterPart(src string) (StmtPart, error) {
	return NewStringPart(src)
}

func AsSlice(object interface{}) []interface{} {
	if object == nil {
		return nil
	}

	val := reflect.ValueOf(object)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	kind := val.Kind()
	if (kind != reflect.Slice && kind != reflect.Array) || val.Type().Elem().Kind() == reflect.Uint8 {
		return []interface{}{object}
	}

	length := val.Len()
	result := make([]interface{}, length)

	for i := 0; i < length; i++ {
		result[i] = val.Index(i).Interface()
	}

	return result
}
