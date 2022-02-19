package main

import (
	"fmt"
	"log"
	"practica3/Interpreter"
	"practica3/prExpr"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ExpressionsVisitorImpl struct {
	antlr.ParseTreeVisitor
}

func NewExpressionsVisitorImp() prExpr.ExpressionsVisitor {
	return &ExpressionsVisitorImpl{
		ParseTreeVisitor: &prExpr.BaseExpressionsVisitor{},
	}
}

func (e *ExpressionsVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(t.GetText())
		return Interpreter.ExprNode{}
	default:
		if node, ok := tree.Accept(e).(Interpreter.ExprInterpreter); ok {
			return node
		}

	}
	return Interpreter.ExprNode{}
}

func (e *ExpressionsVisitorImpl) VisitS(ctx *prExpr.SContext) interface{} {
	nodes := ctx.AllExpr()
	expr := Interpreter.ExprNode{}
	rootResult := new(Interpreter.Result)
	for _, node := range nodes {
		expr = node.Accept(e).(Interpreter.ExprNode)
		result := expr.Interpret(rootResult)
		fmt.Print(node.GetText(), " = ")
		fmt.Println(result.String())
	}
	return expr
}

func (e *ExpressionsVisitorImpl) VisitAdd(ctx *prExpr.AddContext) interface{} {
	left := ctx.Expr(0).Accept(e).(Interpreter.ExprInterpreter)
	right := ctx.Expr(1).Accept(e).(Interpreter.ExprInterpreter)

	return Interpreter.NewExprNode(left, right, "+")

}

func (e *ExpressionsVisitorImpl) VisitDecimal(ctx *prExpr.DecimalContext) interface{} {
	return Interpreter.NewDoubleNode(ctx.DECIMAL().GetText())
}

func (e *ExpressionsVisitorImpl) VisitMul(ctx *prExpr.MulContext) interface{} {
	left := ctx.Expr(0).Accept(e).(Interpreter.ExprInterpreter)
	right := ctx.Expr(1).Accept(e).(Interpreter.ExprInterpreter)
	return Interpreter.NewExprNode(left, right, "*")

}

func (e *ExpressionsVisitorImpl) VisitParens(ctx *prExpr.ParensContext) interface{} {
	return ctx.Expr().Accept(e)
}

func (e *ExpressionsVisitorImpl) VisitInt(ctx *prExpr.IntContext) interface{} {
	return Interpreter.NewIntNode(ctx.INT().GetText())
}
