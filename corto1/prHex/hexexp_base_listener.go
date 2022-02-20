// Code generated from HexExp.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prHex // HexExp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHexExpListener is a complete listener for a parse tree produced by HexExpParser.
type BaseHexExpListener struct{}

var _ HexExpListener = &BaseHexExpListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHexExpListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHexExpListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHexExpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHexExpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseHexExpListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseHexExpListener) ExitS(ctx *SContext) {}

// EnterL1 is called when production L1 is entered.
func (s *BaseHexExpListener) EnterL1(ctx *L1Context) {}

// ExitL1 is called when production L1 is exited.
func (s *BaseHexExpListener) ExitL1(ctx *L1Context) {}

// EnterL2 is called when production L2 is entered.
func (s *BaseHexExpListener) EnterL2(ctx *L2Context) {}

// ExitL2 is called when production L2 is exited.
func (s *BaseHexExpListener) ExitL2(ctx *L2Context) {}

// EnterA is called when production a is entered.
func (s *BaseHexExpListener) EnterA(ctx *AContext) {}

// ExitA is called when production a is exited.
func (s *BaseHexExpListener) ExitA(ctx *AContext) {}
