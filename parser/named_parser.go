// Code generated from C:/dev/repo/go/kra/parser\Named.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Named

import (
	"fmt"
	"strconv"
	"sync"

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
var _ = strconv.Itoa
var _ = sync.Once{}

type NamedParser struct {
	*antlr.BaseParser
}

var namedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func namedParserInit() {
	staticData := &namedParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "", "", "", "", "'('", "')'", "'?'", "','", "'@'",
		"':'", "';'", "'.'", "'*'",
	}
	staticData.symbolicNames = []string{
		"", "SPACES", "BLOCK_COMMENT", "LINE_COMMENT", "IN", "STRING", "NUMBER",
		"IDENTIFIER", "DECPARAM", "OPEN_PAREN", "CLOSE_PAREN", "QMARK", "COMMA",
		"AT", "COLON", "SEMI", "DOT", "STAR", "ANY_SYMBOL",
	}
	staticData.ruleNames = []string{
		"parse", "stmt", "inExpr", "parameter", "namedParameter", "qmarkParameter",
		"decParameter", "staticParameter", "anyStmtParts",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 93, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 4, 0, 21,
		8, 0, 11, 0, 12, 0, 22, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1,
		0, 5, 0, 32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 1, 1, 1, 1, 1, 4, 1, 40,
		8, 1, 11, 1, 12, 1, 41, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 49, 8, 2, 10,
		2, 12, 2, 52, 9, 2, 1, 2, 3, 2, 55, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 63, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 69, 8, 4, 10, 4, 12,
		4, 72, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8,
		83, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 91, 8, 8, 1, 8, 0,
		0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 2, 1, 0, 13, 14, 2, 0, 7, 7, 17,
		17, 102, 0, 18, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 43, 1, 0, 0, 0, 6, 62,
		1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 73, 1, 0, 0, 0, 12, 75, 1, 0, 0, 0,
		14, 77, 1, 0, 0, 0, 16, 90, 1, 0, 0, 0, 18, 27, 3, 2, 1, 0, 19, 21, 5,
		15, 0, 0, 20, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22,
		23, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 20, 1, 0, 0,
		0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 33,
		1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 32, 5, 15, 0, 0, 31, 30, 1, 0, 0, 0,
		32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 1, 1, 0,
		0, 0, 35, 33, 1, 0, 0, 0, 36, 40, 3, 4, 2, 0, 37, 40, 3, 16, 8, 0, 38,
		40, 3, 6, 3, 0, 39, 36, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0,
		0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 3, 1,
		0, 0, 0, 43, 44, 5, 4, 0, 0, 44, 54, 5, 9, 0, 0, 45, 50, 3, 6, 3, 0, 46,
		47, 5, 12, 0, 0, 47, 49, 3, 6, 3, 0, 48, 46, 1, 0, 0, 0, 49, 52, 1, 0,
		0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 55, 1, 0, 0, 0, 52, 50,
		1, 0, 0, 0, 53, 55, 3, 2, 1, 0, 54, 45, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 57, 5, 10, 0, 0, 57, 5, 1, 0, 0, 0, 58, 63, 3,
		8, 4, 0, 59, 63, 3, 10, 5, 0, 60, 63, 3, 12, 6, 0, 61, 63, 3, 14, 7, 0,
		62, 58, 1, 0, 0, 0, 62, 59, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 61, 1,
		0, 0, 0, 63, 7, 1, 0, 0, 0, 64, 65, 7, 0, 0, 0, 65, 70, 5, 7, 0, 0, 66,
		67, 5, 16, 0, 0, 67, 69, 5, 7, 0, 0, 68, 66, 1, 0, 0, 0, 69, 72, 1, 0,
		0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 9, 1, 0, 0, 0, 72, 70,
		1, 0, 0, 0, 73, 74, 5, 11, 0, 0, 74, 11, 1, 0, 0, 0, 75, 76, 5, 8, 0, 0,
		76, 13, 1, 0, 0, 0, 77, 78, 5, 5, 0, 0, 78, 15, 1, 0, 0, 0, 79, 82, 5,
		7, 0, 0, 80, 81, 5, 16, 0, 0, 81, 83, 7, 1, 0, 0, 82, 80, 1, 0, 0, 0, 82,
		83, 1, 0, 0, 0, 83, 91, 1, 0, 0, 0, 84, 91, 5, 9, 0, 0, 85, 91, 5, 10,
		0, 0, 86, 91, 5, 12, 0, 0, 87, 91, 5, 17, 0, 0, 88, 91, 5, 18, 0, 0, 89,
		91, 5, 6, 0, 0, 90, 79, 1, 0, 0, 0, 90, 84, 1, 0, 0, 0, 90, 85, 1, 0, 0,
		0, 90, 86, 1, 0, 0, 0, 90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89,
		1, 0, 0, 0, 91, 17, 1, 0, 0, 0, 11, 22, 27, 33, 39, 41, 50, 54, 62, 70,
		82, 90,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// NamedParserInit initializes any static state used to implement NamedParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNamedParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NamedParserInit() {
	staticData := &namedParserStaticData
	staticData.once.Do(namedParserInit)
}

// NewNamedParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNamedParser(input antlr.TokenStream) *NamedParser {
	NamedParserInit()
	this := new(NamedParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &namedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
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
	NamedParserSTRING        = 5
	NamedParserNUMBER        = 6
	NamedParserIDENTIFIER    = 7
	NamedParserDECPARAM      = 8
	NamedParserOPEN_PAREN    = 9
	NamedParserCLOSE_PAREN   = 10
	NamedParserQMARK         = 11
	NamedParserCOMMA         = 12
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
	NamedParserRULE_namedParameter  = 4
	NamedParserRULE_qmarkParameter  = 5
	NamedParserRULE_decParameter    = 6
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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInExprContext); ok {
			len++
		}
	}

	tst := make([]IInExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInExprContext); ok {
			tst[i] = t.(IInExprContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) InExpr(i int) IInExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInExprContext)
}

func (s *StmtContext) AllAnyStmtParts() []IAnyStmtPartsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnyStmtPartsContext); ok {
			len++
		}
	}

	tst := make([]IAnyStmtPartsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnyStmtPartsContext); ok {
			tst[i] = t.(IAnyStmtPartsContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) AnyStmtParts(i int) IAnyStmtPartsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyStmtPartsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyStmtPartsContext)
}

func (s *StmtContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NamedParserRULE_stmt)

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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(39)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case NamedParserIN:
				{
					p.SetState(36)
					p.InExpr()
				}

			case NamedParserNUMBER, NamedParserIDENTIFIER, NamedParserOPEN_PAREN, NamedParserCLOSE_PAREN, NamedParserCOMMA, NamedParserSTAR, NamedParserANY_SYMBOL:
				{
					p.SetState(37)
					p.AnyStmtParts()
				}

			case NamedParserSTRING, NamedParserDECPARAM, NamedParserQMARK, NamedParserAT, NamedParserCOLON:
				{
					p.SetState(38)
					p.Parameter()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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

type  InExprContext struct {
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

func (s *InExprContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(NamedParserCLOSE_PAREN, 0)
}

func (s *InExprContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *InExprContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *InExprContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
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
	this := p
	_ = this

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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
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

	case 2:
		{
			p.SetState(53)
			p.Stmt()
		}

	}
	{
		p.SetState(56)
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

func (s *ParameterContext) NamedParameter() INamedParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedParameterContext)
}

func (s *ParameterContext) QmarkParameter() IQmarkParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQmarkParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQmarkParameterContext)
}

func (s *ParameterContext) DecParameter() IDecParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecParameterContext)
}

func (s *ParameterContext) StaticParameter() IStaticParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStaticParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NamedParserAT, NamedParserCOLON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.NamedParameter()
		}

	case NamedParserQMARK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.QmarkParameter()
		}

	case NamedParserDECPARAM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.DecParameter()
		}

	case NamedParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.StaticParameter()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamedParameterContext is an interface to support dynamic dispatch.
type INamedParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedParameterContext differentiates from other interfaces.
	IsNamedParameterContext()
}

type NamedParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedParameterContext() *NamedParameterContext {
	var p = new(NamedParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_namedParameter
	return p
}

func (*NamedParameterContext) IsNamedParameterContext() {}

func NewNamedParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedParameterContext {
	var p = new(NamedParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_namedParameter

	return p
}

func (s *NamedParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedParameterContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(NamedParserIDENTIFIER)
}

func (s *NamedParameterContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(NamedParserIDENTIFIER, i)
}

func (s *NamedParameterContext) AT() antlr.TerminalNode {
	return s.GetToken(NamedParserAT, 0)
}

func (s *NamedParameterContext) COLON() antlr.TerminalNode {
	return s.GetToken(NamedParserCOLON, 0)
}

func (s *NamedParameterContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(NamedParserDOT)
}

func (s *NamedParameterContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(NamedParserDOT, i)
}

func (s *NamedParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitNamedParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) NamedParameter() (localctx INamedParameterContext) {
	this := p
	_ = this

	localctx = NewNamedParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NamedParserRULE_namedParameter)
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
		p.SetState(64)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NamedParserAT || _la == NamedParserCOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(65)
		p.Match(NamedParserIDENTIFIER)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NamedParserDOT {
		{
			p.SetState(66)
			p.Match(NamedParserDOT)
		}
		{
			p.SetState(67)
			p.Match(NamedParserIDENTIFIER)
		}

		p.SetState(72)
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
	this := p
	_ = this

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
		p.SetState(73)
		p.Match(NamedParserQMARK)
	}

	return localctx
}

// IDecParameterContext is an interface to support dynamic dispatch.
type IDecParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecParameterContext differentiates from other interfaces.
	IsDecParameterContext()
}

type DecParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecParameterContext() *DecParameterContext {
	var p = new(DecParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NamedParserRULE_decParameter
	return p
}

func (*DecParameterContext) IsDecParameterContext() {}

func NewDecParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecParameterContext {
	var p = new(DecParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NamedParserRULE_decParameter

	return p
}

func (s *DecParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *DecParameterContext) DECPARAM() antlr.TerminalNode {
	return s.GetToken(NamedParserDECPARAM, 0)
}

func (s *DecParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NamedVisitor:
		return t.VisitDecParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NamedParser) DecParameter() (localctx IDecParameterContext) {
	this := p
	_ = this

	localctx = NewDecParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NamedParserRULE_decParameter)

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
		p.SetState(75)
		p.Match(NamedParserDECPARAM)
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
	this := p
	_ = this

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
		p.SetState(77)
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
	this := p
	_ = this

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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NamedParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(NamedParserIDENTIFIER)
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NamedParserDOT {
			{
				p.SetState(80)
				p.Match(NamedParserDOT)
			}
			{
				p.SetState(81)
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
			p.SetState(84)
			p.Match(NamedParserOPEN_PAREN)
		}

	case NamedParserCLOSE_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.Match(NamedParserCLOSE_PAREN)
		}

	case NamedParserCOMMA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)
			p.Match(NamedParserCOMMA)
		}

	case NamedParserSTAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(87)
			p.Match(NamedParserSTAR)
		}

	case NamedParserANY_SYMBOL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)
			p.Match(NamedParserANY_SYMBOL)
		}

	case NamedParserNUMBER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(89)
			p.Match(NamedParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
