// Code generated from Gramatica.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 44, 242,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 40, 6, 40, 211, 10, 40, 13, 40, 14, 40, 212,
	3, 41, 3, 41, 7, 41, 217, 10, 41, 12, 41, 14, 41, 220, 11, 41, 3, 42, 3,
	42, 3, 42, 3, 42, 7, 42, 226, 10, 42, 12, 42, 14, 42, 229, 11, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 6, 43, 237, 10, 43, 13, 43, 14, 43,
	238, 3, 43, 3, 43, 3, 227, 2, 44, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	3, 2, 6, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67,
	92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 245, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73,
	3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2,
	81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 3, 87, 3, 2, 2, 2,
	5, 89, 3, 2, 2, 2, 7, 91, 3, 2, 2, 2, 9, 93, 3, 2, 2, 2, 11, 95, 3, 2,
	2, 2, 13, 97, 3, 2, 2, 2, 15, 99, 3, 2, 2, 2, 17, 102, 3, 2, 2, 2, 19,
	104, 3, 2, 2, 2, 21, 107, 3, 2, 2, 2, 23, 109, 3, 2, 2, 2, 25, 111, 3,
	2, 2, 2, 27, 116, 3, 2, 2, 2, 29, 122, 3, 2, 2, 2, 31, 124, 3, 2, 2, 2,
	33, 127, 3, 2, 2, 2, 35, 130, 3, 2, 2, 2, 37, 132, 3, 2, 2, 2, 39, 134,
	3, 2, 2, 2, 41, 137, 3, 2, 2, 2, 43, 140, 3, 2, 2, 2, 45, 142, 3, 2, 2,
	2, 47, 144, 3, 2, 2, 2, 49, 146, 3, 2, 2, 2, 51, 148, 3, 2, 2, 2, 53, 154,
	3, 2, 2, 2, 55, 157, 3, 2, 2, 2, 57, 161, 3, 2, 2, 2, 59, 164, 3, 2, 2,
	2, 61, 169, 3, 2, 2, 2, 63, 176, 3, 2, 2, 2, 65, 178, 3, 2, 2, 2, 67, 183,
	3, 2, 2, 2, 69, 185, 3, 2, 2, 2, 71, 193, 3, 2, 2, 2, 73, 195, 3, 2, 2,
	2, 75, 201, 3, 2, 2, 2, 77, 204, 3, 2, 2, 2, 79, 210, 3, 2, 2, 2, 81, 214,
	3, 2, 2, 2, 83, 221, 3, 2, 2, 2, 85, 236, 3, 2, 2, 2, 87, 88, 7, 45, 2,
	2, 88, 4, 3, 2, 2, 2, 89, 90, 7, 47, 2, 2, 90, 6, 3, 2, 2, 2, 91, 92, 7,
	35, 2, 2, 92, 8, 3, 2, 2, 2, 93, 94, 7, 44, 2, 2, 94, 10, 3, 2, 2, 2, 95,
	96, 7, 49, 2, 2, 96, 12, 3, 2, 2, 2, 97, 98, 7, 39, 2, 2, 98, 14, 3, 2,
	2, 2, 99, 100, 7, 40, 2, 2, 100, 101, 7, 40, 2, 2, 101, 16, 3, 2, 2, 2,
	102, 103, 7, 96, 2, 2, 103, 18, 3, 2, 2, 2, 104, 105, 7, 126, 2, 2, 105,
	106, 7, 126, 2, 2, 106, 20, 3, 2, 2, 2, 107, 108, 7, 42, 2, 2, 108, 22,
	3, 2, 2, 2, 109, 110, 7, 43, 2, 2, 110, 24, 3, 2, 2, 2, 111, 112, 7, 118,
	2, 2, 112, 113, 7, 116, 2, 2, 113, 114, 7, 119, 2, 2, 114, 115, 7, 103,
	2, 2, 115, 26, 3, 2, 2, 2, 116, 117, 7, 104, 2, 2, 117, 118, 7, 99, 2,
	2, 118, 119, 7, 110, 2, 2, 119, 120, 7, 117, 2, 2, 120, 121, 7, 103, 2,
	2, 121, 28, 3, 2, 2, 2, 122, 123, 7, 95, 2, 2, 123, 30, 3, 2, 2, 2, 124,
	125, 7, 63, 2, 2, 125, 126, 7, 63, 2, 2, 126, 32, 3, 2, 2, 2, 127, 128,
	7, 35, 2, 2, 128, 129, 7, 63, 2, 2, 129, 34, 3, 2, 2, 2, 130, 131, 7, 64,
	2, 2, 131, 36, 3, 2, 2, 2, 132, 133, 7, 62, 2, 2, 133, 38, 3, 2, 2, 2,
	134, 135, 7, 64, 2, 2, 135, 136, 7, 63, 2, 2, 136, 40, 3, 2, 2, 2, 137,
	138, 7, 62, 2, 2, 138, 139, 7, 63, 2, 2, 139, 42, 3, 2, 2, 2, 140, 141,
	7, 63, 2, 2, 141, 44, 3, 2, 2, 2, 142, 143, 7, 61, 2, 2, 143, 46, 3, 2,
	2, 2, 144, 145, 7, 46, 2, 2, 145, 48, 3, 2, 2, 2, 146, 147, 7, 93, 2, 2,
	147, 50, 3, 2, 2, 2, 148, 149, 7, 99, 2, 2, 149, 150, 7, 116, 2, 2, 150,
	151, 7, 116, 2, 2, 151, 152, 7, 99, 2, 2, 152, 153, 7, 123, 2, 2, 153,
	52, 3, 2, 2, 2, 154, 155, 7, 48, 2, 2, 155, 156, 7, 48, 2, 2, 156, 54,
	3, 2, 2, 2, 157, 158, 7, 107, 2, 2, 158, 159, 7, 112, 2, 2, 159, 160, 7,
	118, 2, 2, 160, 56, 3, 2, 2, 2, 161, 162, 7, 107, 2, 2, 162, 163, 7, 104,
	2, 2, 163, 58, 3, 2, 2, 2, 164, 165, 7, 103, 2, 2, 165, 166, 7, 110, 2,
	2, 166, 167, 7, 117, 2, 2, 167, 168, 7, 103, 2, 2, 168, 60, 3, 2, 2, 2,
	169, 170, 7, 117, 2, 2, 170, 171, 7, 121, 2, 2, 171, 172, 7, 107, 2, 2,
	172, 173, 7, 118, 2, 2, 173, 174, 7, 101, 2, 2, 174, 175, 7, 106, 2, 2,
	175, 62, 3, 2, 2, 2, 176, 177, 7, 125, 2, 2, 177, 64, 3, 2, 2, 2, 178,
	179, 7, 101, 2, 2, 179, 180, 7, 99, 2, 2, 180, 181, 7, 117, 2, 2, 181,
	182, 7, 103, 2, 2, 182, 66, 3, 2, 2, 2, 183, 184, 7, 60, 2, 2, 184, 68,
	3, 2, 2, 2, 185, 186, 7, 102, 2, 2, 186, 187, 7, 103, 2, 2, 187, 188, 7,
	104, 2, 2, 188, 189, 7, 99, 2, 2, 189, 190, 7, 119, 2, 2, 190, 191, 7,
	110, 2, 2, 191, 192, 7, 118, 2, 2, 192, 70, 3, 2, 2, 2, 193, 194, 7, 127,
	2, 2, 194, 72, 3, 2, 2, 2, 195, 196, 7, 121, 2, 2, 196, 197, 7, 106, 2,
	2, 197, 198, 7, 107, 2, 2, 198, 199, 7, 110, 2, 2, 199, 200, 7, 103, 2,
	2, 200, 74, 3, 2, 2, 2, 201, 202, 7, 102, 2, 2, 202, 203, 7, 113, 2, 2,
	203, 76, 3, 2, 2, 2, 204, 205, 7, 110, 2, 2, 205, 206, 7, 113, 2, 2, 206,
	207, 7, 113, 2, 2, 207, 208, 7, 114, 2, 2, 208, 78, 3, 2, 2, 2, 209, 211,
	9, 2, 2, 2, 210, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 210, 3, 2,
	2, 2, 212, 213, 3, 2, 2, 2, 213, 80, 3, 2, 2, 2, 214, 218, 9, 3, 2, 2,
	215, 217, 9, 4, 2, 2, 216, 215, 3, 2, 2, 2, 217, 220, 3, 2, 2, 2, 218,
	216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 82, 3, 2, 2, 2, 220, 218, 3,
	2, 2, 2, 221, 222, 7, 49, 2, 2, 222, 223, 7, 44, 2, 2, 223, 227, 3, 2,
	2, 2, 224, 226, 11, 2, 2, 2, 225, 224, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2,
	227, 228, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229,
	227, 3, 2, 2, 2, 230, 231, 7, 44, 2, 2, 231, 232, 7, 49, 2, 2, 232, 233,
	3, 2, 2, 2, 233, 234, 8, 42, 2, 2, 234, 84, 3, 2, 2, 2, 235, 237, 9, 5,
	2, 2, 236, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2,
	238, 239, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 241, 8, 43, 2, 2, 241,
	86, 3, 2, 2, 2, 7, 2, 212, 218, 227, 238, 3, 8, 2, 2,
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
	"", "'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "'&&'", "'^'", "'||'", "'('",
	"')'", "'true'", "'false'", "']'", "'=='", "'!='", "'>'", "'<'", "'>='",
	"'<='", "'='", "';'", "','", "'['", "'array'", "'..'", "'int'", "'if'",
	"'else'", "'switch'", "'{'", "'case'", "':'", "'default'", "'}'", "'while'",
	"'do'", "'loop'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "NUM", "ID", "COMMENT", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "NUM", "ID", "COMMENT", "WHITESPACE",
}

type GramaticaLexer struct {
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

func NewGramaticaLexer(input antlr.CharStream) *GramaticaLexer {

	l := new(GramaticaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Gramatica.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GramaticaLexer tokens.
const (
	GramaticaLexerT__0       = 1
	GramaticaLexerT__1       = 2
	GramaticaLexerT__2       = 3
	GramaticaLexerT__3       = 4
	GramaticaLexerT__4       = 5
	GramaticaLexerT__5       = 6
	GramaticaLexerT__6       = 7
	GramaticaLexerT__7       = 8
	GramaticaLexerT__8       = 9
	GramaticaLexerT__9       = 10
	GramaticaLexerT__10      = 11
	GramaticaLexerT__11      = 12
	GramaticaLexerT__12      = 13
	GramaticaLexerT__13      = 14
	GramaticaLexerT__14      = 15
	GramaticaLexerT__15      = 16
	GramaticaLexerT__16      = 17
	GramaticaLexerT__17      = 18
	GramaticaLexerT__18      = 19
	GramaticaLexerT__19      = 20
	GramaticaLexerT__20      = 21
	GramaticaLexerT__21      = 22
	GramaticaLexerT__22      = 23
	GramaticaLexerT__23      = 24
	GramaticaLexerT__24      = 25
	GramaticaLexerT__25      = 26
	GramaticaLexerT__26      = 27
	GramaticaLexerT__27      = 28
	GramaticaLexerT__28      = 29
	GramaticaLexerT__29      = 30
	GramaticaLexerT__30      = 31
	GramaticaLexerT__31      = 32
	GramaticaLexerT__32      = 33
	GramaticaLexerT__33      = 34
	GramaticaLexerT__34      = 35
	GramaticaLexerT__35      = 36
	GramaticaLexerT__36      = 37
	GramaticaLexerT__37      = 38
	GramaticaLexerNUM        = 39
	GramaticaLexerID         = 40
	GramaticaLexerCOMMENT    = 41
	GramaticaLexerWHITESPACE = 42
)
