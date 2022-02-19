// Code generated from Expressions.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prExpr // Expressions
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExpressionsListener is a complete listener for a parse tree produced by ExpressionsParser.
type ExpressionsListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterAdd is called when entering the Add production.
	EnterAdd(c *AddContext)

	// EnterDecimal is called when entering the Decimal production.
	EnterDecimal(c *DecimalContext)

	// EnterMul is called when entering the Mul production.
	EnterMul(c *MulContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitAdd is called when exiting the Add production.
	ExitAdd(c *AddContext)

	// ExitDecimal is called when exiting the Decimal production.
	ExitDecimal(c *DecimalContext)

	// ExitMul is called when exiting the Mul production.
	ExitMul(c *MulContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)
}
