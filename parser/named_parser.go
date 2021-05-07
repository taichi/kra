// Code generated from c:\dev\repo\go\kra\parser\Named.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Named

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 92, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 6, 2, 23, 10, 2, 13, 2,
	14, 2, 24, 3, 2, 7, 2, 28, 10, 2, 12, 2, 14, 2, 31, 11, 2, 3, 2, 7, 2,
	34, 10, 2, 12, 2, 14, 2, 37, 11, 2, 3, 3, 3, 3, 3, 3, 6, 3, 42, 10, 3,
	13, 3, 14, 3, 43, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4,
	14, 4, 54, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 62, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 68, 10, 6, 12, 6, 14, 6, 71, 11, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 82, 10, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 90, 10, 10, 3, 10, 2, 2,
	11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 4, 3, 2, 15, 16, 4, 2, 14, 14, 19,
	19, 2, 100, 2, 20, 3, 2, 2, 2, 4, 41, 3, 2, 2, 2, 6, 45, 3, 2, 2, 2, 8,
	61, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 14, 74, 3, 2, 2,
	2, 16, 76, 3, 2, 2, 2, 18, 89, 3, 2, 2, 2, 20, 29, 5, 4, 3, 2, 21, 23,
	7, 17, 2, 2, 22, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2,
	24, 25, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 28, 5, 4, 3, 2, 27, 22, 3,
	2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30,
	35, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 34, 7, 17, 2, 2, 33, 32, 3, 2,
	2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 3,
	3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 42, 5, 6, 4, 2, 39, 42, 5, 18, 10,
	2, 40, 42, 5, 8, 5, 2, 41, 38, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 40,
	3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2,
	44, 5, 3, 2, 2, 2, 45, 46, 7, 6, 2, 2, 46, 47, 7, 7, 2, 2, 47, 52, 5, 8,
	5, 2, 48, 49, 7, 11, 2, 2, 49, 51, 5, 8, 5, 2, 50, 48, 3, 2, 2, 2, 51,
	54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2,
	2, 54, 52, 3, 2, 2, 2, 55, 56, 7, 8, 2, 2, 56, 7, 3, 2, 2, 2, 57, 62, 5,
	10, 6, 2, 58, 62, 5, 12, 7, 2, 59, 62, 5, 14, 8, 2, 60, 62, 5, 16, 9, 2,
	61, 57, 3, 2, 2, 2, 61, 58, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3,
	2, 2, 2, 62, 9, 3, 2, 2, 2, 63, 64, 9, 2, 2, 2, 64, 69, 7, 14, 2, 2, 65,
	66, 7, 18, 2, 2, 66, 68, 7, 14, 2, 2, 67, 65, 3, 2, 2, 2, 68, 71, 3, 2,
	2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 11, 3, 2, 2, 2, 71, 69,
	3, 2, 2, 2, 72, 73, 7, 9, 2, 2, 73, 13, 3, 2, 2, 2, 74, 75, 7, 10, 2, 2,
	75, 15, 3, 2, 2, 2, 76, 77, 7, 12, 2, 2, 77, 17, 3, 2, 2, 2, 78, 81, 7,
	14, 2, 2, 79, 80, 7, 18, 2, 2, 80, 82, 9, 3, 2, 2, 81, 79, 3, 2, 2, 2,
	81, 82, 3, 2, 2, 2, 82, 90, 3, 2, 2, 2, 83, 90, 7, 7, 2, 2, 84, 90, 7,
	8, 2, 2, 85, 90, 7, 11, 2, 2, 86, 90, 7, 19, 2, 2, 87, 90, 7, 20, 2, 2,
	88, 90, 7, 13, 2, 2, 89, 78, 3, 2, 2, 2, 89, 83, 3, 2, 2, 2, 89, 84, 3,
	2, 2, 2, 89, 85, 3, 2, 2, 2, 89, 86, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89,
	88, 3, 2, 2, 2, 90, 19, 3, 2, 2, 2, 12, 24, 29, 35, 41, 43, 52, 61, 69,
	81, 89,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'('", "')'", "'?'", "", "','", "", "", "", "'@'",
	"':'", "';'", "'.'", "'*'",
}
var symbolicNames = []string{
	"", "SPACES", "BLOCK_COMMENT", "LINE_COMMENT", "IN", "OPEN_PAREN", "CLOSE_PAREN",
	"QMARK", "DDEC", "COMMA", "STRING", "NUMBER", "IDENTIFIER", "AT", "COLON",
	"SEMI", "DOT", "STAR", "ANY_SYMBOL",
}

var ruleNames = []string{
	"parse", "stmt", "inExpr", "parameter", "namedParamter", "qmarkParameter",
	"dDecParameter", "staticParameter", "anyStmtParts",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type NamedParser struct {
	*antlr.BaseParser
}

func NewNamedParser(input antlr.TokenStream) *NamedParser {
	this := new(NamedParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Named.g4"

	return this
}

// NamedParser tokens.
const (
	NamedParserEOF           = antlr.TokenEOF
	NamedParserSPACES        = 1
	NamedParserBLOCK_COMMENT = 2
	NamedParserLINE_COMMENT  = 3
	NamedParserIN            = 4
	NamedParserOPEN_PAREN    = 5
	NamedParserCLOSE_PAREN   = 6
	NamedParserQMARK         = 7
	NamedParserDDEC          = 8
	NamedParserCOMMA         = 9
	NamedParserSTRING        = 10
	NamedParserNUMBER        = 11
	NamedParserIDENTIFIER    = 12
	NamedParserAT            = 13
	NamedParserCOLON         = 14
	NamedParserSEMI          = 15
	NamedParserDOT           = 16
	NamedParserSTAR          = 17
	NamedParserANY_SYMBOL    = 18
)

// NamedParser rules.
const (
	NamedParserRULE_parse           = 0
	NamedParserRULE_stmt            = 1
	NamedParserRULE_inExpr          = 2
	NamedParserRULE_parameter       = 3
	NamedParserRULE_namedParamter   = 4
	NamedParserRULE_qmarkParameter  = 5
	NamedParserRULE_dDecParameter   = 6
	NamedParserRULE_staticParameter = 7
	NamedParserRULE_anyStmtParts    = 8
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *ParseContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ParseContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(NamedParserSEMI)
}

func (s *ParseContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(NamedParserSEMI, i)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NamedParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Stmt()
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NamedParserSEMI {
				{
					p.SetState(19)
					p.Match(NamedParserSEMI)
				}

				p.SetState(22)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(24)
				p.Stmt()
			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NamedParserSEMI {
		{
			p.SetState(30)
			p.Match(NamedParserSEMI)
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) AllInExpr() []IInExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInExprContext)(nil)).Elem())
	var tst = make([]IInExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInExprContext)
		}
	}

	return tst
}

func (s *StmtContext) InExpr(i int) IInExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInExprContext)
}

func (s *StmtContext) AllAnyStmtParts() []IAnyStmtPartsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnyStmtPartsContext)(nil)).Elem())
	var tst = make([]IAnyStmtPartsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnyStmtPartsContext)
		}
	}

	return tst
}

func (s *StmtContext) AnyStmtParts(i int) IAnyStmtPartsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyStmtPartsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnyStmtPartsContext)
}

func (s *StmtContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *StmtContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NamedParserRULE_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NamedParserIN)|(1<<NamedParserOPEN_PAREN)|(1<<NamedParserCLOSE_PAREN)|(1<<NamedParserQMARK)|(1<<NamedParserDDEC)|(1<<NamedParserCOMMA)|(1<<NamedParserSTRING)|(1<<NamedParserNUMBER)|(1<<NamedParserIDENTIFIER)|(1<<NamedParserAT)|(1<<NamedParserCOLON)|(1<<NamedParserSTAR)|(1<<NamedParserANY_SYMBOL))) != 0) {
		p.SetState(39)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NamedParserIN:
			{
				p.SetState(36)
				p.InExpr()
			}

		case NamedParserOPEN_PAREN, NamedParserCLOSE_PAREN, NamedParserCOMMA, NamedParserNUMBER, NamedParserIDENTIFIER, NamedParserSTAR, NamedParserANY_SYMBOL:
			{
				p.SetState(37)
				p.AnyStmtParts()
			}

		case NamedParserQMARK, NamedParserDDEC, NamedParserSTRING, NamedParserAT, NamedParserCOLON:
			{
				p.SetState(38)
				p.Parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInExprContext is an interface to support dynamic dispatch.
type IInExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInExprContext differentiates from other interfaces.
	IsInExprContext()
}

type InExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInExprContext() *InExprContext {
	var p = new(InExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_inExpr
	return p
}

func (*InExprContext) IsInExprContext() {}

func NewInExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InExprContext {
	var p = new(InExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_inExpr

	return p
}

func (s *InExprContext) GetParser() antlr.Parser { return s.parser }

func (s *InExprContext) IN() antlr.TerminalNode {
	return s.GetToken(NamedParserIN, 0)
}

func (s *InExprContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(NamedParserOPEN_PAREN, 0)
}

func (s *InExprContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *InExprContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *InExprContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(NamedParserCLOSE_PAREN, 0)
}

func (s *InExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NamedParserCOMMA)
}

func (s *InExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NamedParserCOMMA, i)
}

func (s *InExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitInExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) InExpr() (localctx IInExprContext) {
	localctx = NewInExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NamedParserRULE_inExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(NamedParserIN)
	}
	{
		p.SetState(44)
		p.Match(NamedParserOPEN_PAREN)
	}
	{
		p.SetState(45)
		p.Parameter()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NamedParserCOMMA {
		{
			p.SetState(46)
			p.Match(NamedParserCOMMA)
		}
		{
			p.SetState(47)
			p.Parameter()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(NamedParserCLOSE_PAREN)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) NamedParamter() INamedParamterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedParamterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedParamterContext)
}

func (s *ParameterContext) QmarkParameter() IQmarkParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQmarkParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQmarkParameterContext)
}

func (s *ParameterContext) DDecParameter() IDDecParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDDecParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDDecParameterContext)
}

func (s *ParameterContext) StaticParameter() IStaticParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStaticParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStaticParameterContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NamedParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NamedParserAT, NamedParserCOLON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.NamedParamter()
		}

	case NamedParserQMARK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.QmarkParameter()
		}

	case NamedParserDDEC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.DDecParameter()
		}

	case NamedParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.StaticParameter()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamedParamterContext is an interface to support dynamic dispatch.
type INamedParamterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedParamterContext differentiates from other interfaces.
	IsNamedParamterContext()
}

type NamedParamterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedParamterContext() *NamedParamterContext {
	var p = new(NamedParamterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_namedParamter
	return p
}

func (*NamedParamterContext) IsNamedParamterContext() {}

func NewNamedParamterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedParamterContext {
	var p = new(NamedParamterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_namedParamter

	return p
}

func (s *NamedParamterContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedParamterContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(NamedParserIDENTIFIER)
}

func (s *NamedParamterContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(NamedParserIDENTIFIER, i)
}

func (s *NamedParamterContext) AT() antlr.TerminalNode {
	return s.GetToken(NamedParserAT, 0)
}

func (s *NamedParamterContext) COLON() antlr.TerminalNode {
	return s.GetToken(NamedParserCOLON, 0)
}

func (s *NamedParamterContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(NamedParserDOT)
}

func (s *NamedParamterContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(NamedParserDOT, i)
}

func (s *NamedParamterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedParamterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedParamterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitNamedParamter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) NamedParamter() (localctx INamedParamterContext) {
	localctx = NewNamedParamterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NamedParserRULE_namedParamter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NamedParserAT || _la == NamedParserCOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(62)
		p.Match(NamedParserIDENTIFIER)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NamedParserDOT {
		{
			p.SetState(63)
			p.Match(NamedParserDOT)
		}
		{
			p.SetState(64)
			p.Match(NamedParserIDENTIFIER)
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQmarkParameterContext is an interface to support dynamic dispatch.
type IQmarkParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQmarkParameterContext differentiates from other interfaces.
	IsQmarkParameterContext()
}

type QmarkParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQmarkParameterContext() *QmarkParameterContext {
	var p = new(QmarkParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_qmarkParameter
	return p
}

func (*QmarkParameterContext) IsQmarkParameterContext() {}

func NewQmarkParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QmarkParameterContext {
	var p = new(QmarkParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_qmarkParameter

	return p
}

func (s *QmarkParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *QmarkParameterContext) QMARK() antlr.TerminalNode {
	return s.GetToken(NamedParserQMARK, 0)
}

func (s *QmarkParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QmarkParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QmarkParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitQmarkParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) QmarkParameter() (localctx IQmarkParameterContext) {
	localctx = NewQmarkParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NamedParserRULE_qmarkParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(NamedParserQMARK)
	}

	return localctx
}

// IDDecParameterContext is an interface to support dynamic dispatch.
type IDDecParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDDecParameterContext differentiates from other interfaces.
	IsDDecParameterContext()
}

type DDecParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDDecParameterContext() *DDecParameterContext {
	var p = new(DDecParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_dDecParameter
	return p
}

func (*DDecParameterContext) IsDDecParameterContext() {}

func NewDDecParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DDecParameterContext {
	var p = new(DDecParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_dDecParameter

	return p
}

func (s *DDecParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *DDecParameterContext) DDEC() antlr.TerminalNode {
	return s.GetToken(NamedParserDDEC, 0)
}

func (s *DDecParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DDecParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DDecParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitDDecParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) DDecParameter() (localctx IDDecParameterContext) {
	localctx = NewDDecParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NamedParserRULE_dDecParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(NamedParserDDEC)
	}

	return localctx
}

// IStaticParameterContext is an interface to support dynamic dispatch.
type IStaticParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStaticParameterContext differentiates from other interfaces.
	IsStaticParameterContext()
}

type StaticParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStaticParameterContext() *StaticParameterContext {
	var p = new(StaticParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_staticParameter
	return p
}

func (*StaticParameterContext) IsStaticParameterContext() {}

func NewStaticParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StaticParameterContext {
	var p = new(StaticParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_staticParameter

	return p
}

func (s *StaticParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *StaticParameterContext) STRING() antlr.TerminalNode {
	return s.GetToken(NamedParserSTRING, 0)
}

func (s *StaticParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StaticParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StaticParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitStaticParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) StaticParameter() (localctx IStaticParameterContext) {
	localctx = NewStaticParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NamedParserRULE_staticParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(NamedParserSTRING)
	}

	return localctx
}

// IAnyStmtPartsContext is an interface to support dynamic dispatch.
type IAnyStmtPartsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnyStmtPartsContext differentiates from other interfaces.
	IsAnyStmtPartsContext()
}

type AnyStmtPartsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyStmtPartsContext() *AnyStmtPartsContext {
	var p = new(AnyStmtPartsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_anyStmtParts
	return p
}

func (*AnyStmtPartsContext) IsAnyStmtPartsContext() {}

func NewAnyStmtPartsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyStmtPartsContext {
	var p = new(AnyStmtPartsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_anyStmtParts

	return p
}

func (s *AnyStmtPartsContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyStmtPartsContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(NamedParserIDENTIFIER)
}

func (s *AnyStmtPartsContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(NamedParserIDENTIFIER, i)
}

func (s *AnyStmtPartsContext) DOT() antlr.TerminalNode {
	return s.GetToken(NamedParserDOT, 0)
}

func (s *AnyStmtPartsContext) STAR() antlr.TerminalNode {
	return s.GetToken(NamedParserSTAR, 0)
}

func (s *AnyStmtPartsContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(NamedParserOPEN_PAREN, 0)
}

func (s *AnyStmtPartsContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(NamedParserCLOSE_PAREN, 0)
}

func (s *AnyStmtPartsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(NamedParserCOMMA, 0)
}

func (s *AnyStmtPartsContext) ANY_SYMBOL() antlr.TerminalNode {
	return s.GetToken(NamedParserANY_SYMBOL, 0)
}

func (s *AnyStmtPartsContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NamedParserNUMBER, 0)
}

func (s *AnyStmtPartsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyStmtPartsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyStmtPartsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitAnyStmtParts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) AnyStmtParts() (localctx IAnyStmtPartsContext) {
	localctx = NewAnyStmtPartsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NamedParserRULE_anyStmtParts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NamedParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Match(NamedParserIDENTIFIER)
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NamedParserDOT {
			{
				p.SetState(77)
				p.Match(NamedParserDOT)
			}
			{
				p.SetState(78)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NamedParserIDENTIFIER || _la == NamedParserSTAR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case NamedParserOPEN_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(NamedParserOPEN_PAREN)
		}

	case NamedParserCLOSE_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Match(NamedParserCLOSE_PAREN)
		}

	case NamedParserCOMMA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Match(NamedParserCOMMA)
		}

	case NamedParserSTAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.Match(NamedParserSTAR)
		}

	case NamedParserANY_SYMBOL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(85)
			p.Match(NamedParserANY_SYMBOL)
		}

	case NamedParserNUMBER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(86)
			p.Match(NamedParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
