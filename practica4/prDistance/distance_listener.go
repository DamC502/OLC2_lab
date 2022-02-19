// Code generated from Distance.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prDistance // Distance
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DistanceListener is a complete listener for a parse tree produced by DistanceParser.
type DistanceListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterL1 is called when entering the L1 production.
	EnterL1(c *L1Context)

	// EnterL2 is called when entering the L2 production.
	EnterL2(c *L2Context)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitL1 is called when exiting the L1 production.
	ExitL1(c *L1Context)

	// ExitL2 is called when exiting the L2 production.
	ExitL2(c *L2Context)
}
