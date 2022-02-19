// Code generated from Expressions.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prExpr // Expressions
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExpressionsParser.
type ExpressionsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExpressionsParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#Add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#Decimal.
	VisitDecimal(ctx *DecimalContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#Mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#Int.
	VisitInt(ctx *IntContext) interface{}
}
