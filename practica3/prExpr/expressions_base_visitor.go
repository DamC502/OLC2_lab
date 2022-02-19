// Code generated from Expressions.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prExpr // Expressions
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExpressionsVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExpressionsVisitor) VisitS(ctx *SContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitDecimal(ctx *DecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitMul(ctx *MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}
