// Code generated from Expressions.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prExpr // Expressions
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExpressionsListener is a complete listener for a parse tree produced by ExpressionsParser.
type BaseExpressionsListener struct{}

var _ ExpressionsListener = &BaseExpressionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpressionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpressionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpressionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpressionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseExpressionsListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseExpressionsListener) ExitS(ctx *SContext) {}

// EnterAdd is called when production Add is entered.
func (s *BaseExpressionsListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production Add is exited.
func (s *BaseExpressionsListener) ExitAdd(ctx *AddContext) {}

// EnterDecimal is called when production Decimal is entered.
func (s *BaseExpressionsListener) EnterDecimal(ctx *DecimalContext) {}

// ExitDecimal is called when production Decimal is exited.
func (s *BaseExpressionsListener) ExitDecimal(ctx *DecimalContext) {}

// EnterMul is called when production Mul is entered.
func (s *BaseExpressionsListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production Mul is exited.
func (s *BaseExpressionsListener) ExitMul(ctx *MulContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseExpressionsListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseExpressionsListener) ExitParens(ctx *ParensContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseExpressionsListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseExpressionsListener) ExitInt(ctx *IntContext) {}
