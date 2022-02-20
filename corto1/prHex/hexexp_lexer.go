// Code generated from HexExp.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prHex

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 5, 16, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 5,
	4, 15, 10, 4, 2, 2, 5, 3, 3, 5, 4, 7, 5, 3, 2, 3, 5, 2, 50, 59, 67, 92,
	99, 124, 2, 15, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 3,
	9, 3, 2, 2, 2, 5, 11, 3, 2, 2, 2, 7, 14, 3, 2, 2, 2, 9, 10, 7, 46, 2, 2,
	10, 4, 3, 2, 2, 2, 11, 12, 7, 48, 2, 2, 12, 6, 3, 2, 2, 2, 13, 15, 9, 2,
	2, 2, 14, 13, 3, 2, 2, 2, 15, 8, 3, 2, 2, 2, 4, 2, 14, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "CHAR",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "CHAR",
}

type HexExpLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewHexExpLexer(input antlr.CharStream) *HexExpLexer {

	l := new(HexExpLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "HexExp.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HexExpLexer tokens.
const (
	HexExpLexerT__0 = 1
	HexExpLexerT__1 = 2
	HexExpLexerCHAR = 3
)
