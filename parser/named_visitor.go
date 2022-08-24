// Code generated from C:/dev/repo/go/kra/parser\Named.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Named

import "github.com/antlr/antlr4/runtime/Go/antlr"

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

// A complete Visitor for a parse tree produced by NamedParser.
type NamedVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by NamedParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by NamedParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by NamedParser#inExpr.
	VisitInExpr(ctx *InExprContext) interface{}

	// Visit a parse tree produced by NamedParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by NamedParser#namedParameter.
	VisitNamedParameter(ctx *NamedParameterContext) interface{}

	// Visit a parse tree produced by NamedParser#qmarkParameter.
	VisitQmarkParameter(ctx *QmarkParameterContext) interface{}

	// Visit a parse tree produced by NamedParser#decParameter.
	VisitDecParameter(ctx *DecParameterContext) interface{}

	// Visit a parse tree produced by NamedParser#staticParameter.
	VisitStaticParameter(ctx *StaticParameterContext) interface{}

	// Visit a parse tree produced by NamedParser#anyStmtParts.
	VisitAnyStmtParts(ctx *AnyStmtPartsContext) interface{}
}
