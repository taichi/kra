// Code generated from C:/dev/repo/go/kra/parser\Named.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Named

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseNamedVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseNamedVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitInExpr(ctx *InExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitNamedParameter(ctx *NamedParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitQmarkParameter(ctx *QmarkParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitDecParameter(ctx *DecParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitStaticParameter(ctx *StaticParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitAnyStmtParts(ctx *AnyStmtPartsContext) interface{} {
	return v.VisitChildren(ctx)
}
