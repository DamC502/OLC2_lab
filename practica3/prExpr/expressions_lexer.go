// Code generated from Expressions.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prExpr

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 55, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 6, 6, 29, 10, 6, 13, 6, 14, 6, 30, 3, 6, 3, 6, 6, 6, 35, 10, 6,
	13, 6, 14, 6, 36, 3, 7, 6, 7, 40, 10, 7, 13, 7, 14, 7, 41, 3, 8, 6, 8,
	45, 10, 8, 13, 8, 14, 8, 46, 3, 8, 3, 8, 3, 9, 5, 9, 52, 10, 9, 3, 9, 3,
	9, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2,
	5, 3, 2, 50, 59, 4, 2, 46, 46, 48, 48, 4, 2, 11, 11, 34, 34, 2, 59, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	3, 19, 3, 2, 2, 2, 5, 21, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 25, 3, 2, 2,
	2, 11, 28, 3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15, 44, 3, 2, 2, 2, 17, 51,
	3, 2, 2, 2, 19, 20, 7, 44, 2, 2, 20, 4, 3, 2, 2, 2, 21, 22, 7, 45, 2, 2,
	22, 6, 3, 2, 2, 2, 23, 24, 7, 42, 2, 2, 24, 8, 3, 2, 2, 2, 25, 26, 7, 43,
	2, 2, 26, 10, 3, 2, 2, 2, 27, 29, 9, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30,
	3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2,
	32, 34, 9, 3, 2, 2, 33, 35, 9, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 36, 3,
	2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 12, 3, 2, 2, 2, 38,
	40, 9, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 39, 3, 2, 2,
	2, 41, 42, 3, 2, 2, 2, 42, 14, 3, 2, 2, 2, 43, 45, 9, 4, 2, 2, 44, 43,
	3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2,
	47, 48, 3, 2, 2, 2, 48, 49, 8, 8, 2, 2, 49, 16, 3, 2, 2, 2, 50, 52, 7,
	15, 2, 2, 51, 50, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53,
	54, 7, 12, 2, 2, 54, 18, 3, 2, 2, 2, 8, 2, 30, 36, 41, 46, 51, 3, 8, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'+'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "DECIMAL", "INT", "WS", "NL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "DECIMAL", "INT", "WS", "NL",
}

type ExpressionsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewExpressionsLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ExpressionsLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewExpressionsLexer(input antlr.CharStream) *ExpressionsLexer {
	l := new(ExpressionsLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expressions.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionsLexer tokens.
const (
	ExpressionsLexerT__0    = 1
	ExpressionsLexerT__1    = 2
	ExpressionsLexerT__2    = 3
	ExpressionsLexerT__3    = 4
	ExpressionsLexerDECIMAL = 5
	ExpressionsLexerINT     = 6
	ExpressionsLexerWS      = 7
	ExpressionsLexerNL      = 8
)
