// Code generated from Distance.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prDistance // Distance
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DistanceParser.
type DistanceVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DistanceParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by DistanceParser#L1.
	VisitL1(ctx *L1Context) interface{}

	// Visit a parse tree produced by DistanceParser#L2.
	VisitL2(ctx *L2Context) interface{}
}
