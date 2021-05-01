// Code generated from c:\dev\repo\go\kra\parser\Named.g4 by ANTLR 4.8. DO NOT EDIT.

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

func (v *BaseNamedVisitor) VisitNamedParamter(ctx *NamedParamterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitQmarkParameter(ctx *QmarkParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitDDecParameter(ctx *DDecParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitStaticParameter(ctx *StaticParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNamedVisitor) VisitAnyStmtParts(ctx *AnyStmtPartsContext) interface{} {
	return v.VisitChildren(ctx)
}
