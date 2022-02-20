// Code generated from HexExp.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prHex // HexExp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by HexExpParser.
type HexExpVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HexExpParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by HexExpParser#L1.
	VisitL1(ctx *L1Context) interface{}

	// Visit a parse tree produced by HexExpParser#L2.
	VisitL2(ctx *L2Context) interface{}

	// Visit a parse tree produced by HexExpParser#a.
	VisitA(ctx *AContext) interface{}
}
