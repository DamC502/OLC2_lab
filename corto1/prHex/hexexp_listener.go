// Code generated from HexExp.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prHex // HexExp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// HexExpListener is a complete listener for a parse tree produced by HexExpParser.
type HexExpListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterL1 is called when entering the L1 production.
	EnterL1(c *L1Context)

	// EnterL2 is called when entering the L2 production.
	EnterL2(c *L2Context)

	// EnterA is called when entering the a production.
	EnterA(c *AContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitL1 is called when exiting the L1 production.
	ExitL1(c *L1Context)

	// ExitL2 is called when exiting the L2 production.
	ExitL2(c *L2Context)

	// ExitA is called when exiting the a production.
	ExitA(c *AContext)
}
