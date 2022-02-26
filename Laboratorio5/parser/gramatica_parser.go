// Code generated from Gramatica.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import _ "Laboratorio5/generador"
import "Laboratorio5/entorno"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 65, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 20, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 31, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 44, 10, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7,
	6, 60, 10, 6, 12, 6, 14, 6, 63, 11, 6, 3, 6, 2, 3, 10, 7, 2, 4, 6, 8, 10,
	2, 4, 3, 2, 7, 8, 3, 2, 5, 6, 2, 64, 2, 19, 3, 2, 2, 2, 4, 21, 3, 2, 2,
	2, 6, 30, 3, 2, 2, 2, 8, 32, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12, 13, 5,
	4, 3, 2, 13, 14, 5, 6, 4, 2, 14, 15, 8, 2, 1, 2, 15, 20, 3, 2, 2, 2, 16,
	17, 5, 10, 6, 2, 17, 18, 7, 2, 2, 3, 18, 20, 3, 2, 2, 2, 19, 12, 3, 2,
	2, 2, 19, 16, 3, 2, 2, 2, 20, 3, 3, 2, 2, 2, 21, 22, 8, 3, 1, 2, 22, 5,
	3, 2, 2, 2, 23, 24, 5, 8, 5, 2, 24, 25, 7, 12, 2, 2, 25, 26, 8, 4, 1, 2,
	26, 27, 7, 3, 2, 2, 27, 28, 5, 6, 4, 2, 28, 31, 3, 2, 2, 2, 29, 31, 3,
	2, 2, 2, 30, 23, 3, 2, 2, 2, 30, 29, 3, 2, 2, 2, 31, 7, 3, 2, 2, 2, 32,
	33, 7, 4, 2, 2, 33, 34, 8, 5, 1, 2, 34, 9, 3, 2, 2, 2, 35, 36, 8, 6, 1,
	2, 36, 37, 7, 9, 2, 2, 37, 38, 5, 10, 6, 2, 38, 39, 7, 10, 2, 2, 39, 40,
	8, 6, 1, 2, 40, 44, 3, 2, 2, 2, 41, 42, 7, 11, 2, 2, 42, 44, 8, 6, 1, 2,
	43, 35, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 61, 3, 2, 2, 2, 45, 46, 12,
	6, 2, 2, 46, 47, 8, 6, 1, 2, 47, 48, 9, 2, 2, 2, 48, 49, 8, 6, 1, 2, 49,
	50, 5, 10, 6, 7, 50, 51, 8, 6, 1, 2, 51, 60, 3, 2, 2, 2, 52, 53, 12, 5,
	2, 2, 53, 54, 8, 6, 1, 2, 54, 55, 9, 3, 2, 2, 55, 56, 8, 6, 1, 2, 56, 57,
	5, 10, 6, 6, 57, 58, 8, 6, 1, 2, 58, 60, 3, 2, 2, 2, 59, 45, 3, 2, 2, 2,
	59, 52, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3,
	2, 2, 2, 62, 11, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 7, 19, 30, 43, 59, 61,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'int'", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "MAS", "MEN", "POR", "DIV", "PARI", "PARD", "NUM", "ID", "WHITESPACE",
}

var ruleNames = []string{
	"start", "marcador", "dec", "tipo", "exp",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GramaticaParser struct {
	*antlr.BaseParser
}

func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

type Nodo struct {
	dir  string
	tipo string
	num  int
}

var desp int = 0
var temp int = 1

// var env []*entorno.Entorno
var sup *entorno.Entorno

// func push(e *entorno.Entorno) {
//     env = append(env, e)
// }

// func pop() *entorno.Entorno {
//     if len(env) < 1 {
//         panic("empty env")
//     }
//     result := env[len(env) - 1]
//     env = env[:len(env) - 1]
//     return result
// }

func gen(out string) {
	fmt.Println(out)
}

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserT__0       = 1
	GramaticaParserT__1       = 2
	GramaticaParserMAS        = 3
	GramaticaParserMEN        = 4
	GramaticaParserPOR        = 5
	GramaticaParserDIV        = 6
	GramaticaParserPARI       = 7
	GramaticaParserPARD       = 8
	GramaticaParserNUM        = 9
	GramaticaParserID         = 10
	GramaticaParserWHITESPACE = 11
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start    = 0
	GramaticaParserRULE_marcador = 1
	GramaticaParserRULE_dec      = 2
	GramaticaParserRULE_tipo     = 3
	GramaticaParserRULE_exp      = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Marcador() IMarcadorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarcadorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarcadorContext)
}

func (s *StartContext) Dec() IDecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecContext)
}

func (s *StartContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserEOF, GramaticaParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Marcador()
		}
		{
			p.SetState(11)
			p.Dec()
		}
		entorno.Mostrar(sup)

	case GramaticaParserPARI, GramaticaParserNUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(14)
			p.exp(0)
		}
		{
			p.SetState(15)
			p.Match(GramaticaParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMarcadorContext is an interface to support dynamic dispatch.
type IMarcadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMarcadorContext differentiates from other interfaces.
	IsMarcadorContext()
}

type MarcadorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarcadorContext() *MarcadorContext {
	var p = new(MarcadorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_marcador
	return p
}

func (*MarcadorContext) IsMarcadorContext() {}

func NewMarcadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarcadorContext {
	var p = new(MarcadorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_marcador

	return p
}

func (s *MarcadorContext) GetParser() antlr.Parser { return s.parser }
func (s *MarcadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarcadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarcadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterMarcador(s)
	}
}

func (s *MarcadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitMarcador(s)
	}
}

func (p *GramaticaParser) Marcador() (localctx IMarcadorContext) {
	localctx = NewMarcadorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_marcador)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	sup = entorno.NewEntorno(nil)
	desp = 0

	return localctx
}

// IDecContext is an interface to support dynamic dispatch.
type IDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITipoContext

	// SetT sets the t rule contexts.
	SetT(ITipoContext)

	// IsDecContext differentiates from other interfaces.
	IsDecContext()
}

type DecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      ITipoContext
	id     antlr.Token
}

func NewEmptyDecContext() *DecContext {
	var p = new(DecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_dec
	return p
}

func (*DecContext) IsDecContext() {}

func NewDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecContext {
	var p = new(DecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_dec

	return p
}

func (s *DecContext) GetParser() antlr.Parser { return s.parser }

func (s *DecContext) GetId() antlr.Token { return s.id }

func (s *DecContext) SetId(v antlr.Token) { s.id = v }

func (s *DecContext) GetT() ITipoContext { return s.t }

func (s *DecContext) SetT(v ITipoContext) { s.t = v }

func (s *DecContext) Dec() IDecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecContext)
}

func (s *DecContext) Tipo() ITipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DecContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *DecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDec(s)
	}
}

func (s *DecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDec(s)
	}
}

func (p *GramaticaParser) Dec() (localctx IDecContext) {
	localctx = NewDecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GramaticaParserRULE_dec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)

			var _x = p.Tipo()

			localctx.(*DecContext).t = _x
		}
		{
			p.SetState(22)

			var _m = p.Match(GramaticaParserID)

			localctx.(*DecContext).id = _m
		}

		sim := entorno.NewSimbolo((func() string {
			if localctx.(*DecContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*DecContext).GetId().GetText()
			}
		}()), localctx.(*DecContext).GetT().GetTipoDato(), desp)
		sup.Put((func() string {
			if localctx.(*DecContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*DecContext).GetId().GetText()
			}
		}()), sim)
		desp = desp + 1

		{
			p.SetState(24)
			p.Match(GramaticaParserT__0)
		}
		{
			p.SetState(25)
			p.Dec()
		}

	case GramaticaParserEOF:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetTipoDato returns the tipoDato attribute.
	GetTipoDato() string

	// SetTipoDato sets the tipoDato attribute.
	SetTipoDato(string)

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	tipoDato string
	t        antlr.Token
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_tipo
	return p
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) GetT() antlr.Token { return s.t }

func (s *TipoContext) SetT(v antlr.Token) { s.t = v }

func (s *TipoContext) GetTipoDato() string { return s.tipoDato }

func (s *TipoContext) SetTipoDato(v string) { s.tipoDato = v }

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *GramaticaParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_tipo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)

		var _m = p.Match(GramaticaParserT__1)

		localctx.(*TipoContext).t = _m
	}

	localctx.(*TipoContext).tipoDato = (func() string {
		if localctx.(*TipoContext).GetT() == nil {
			return ""
		} else {
			return localctx.(*TipoContext).GetT().GetText()
		}
	}())

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpContext

	// GetE returns the e rule contexts.
	GetE() IExpContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpContext)

	// SetE sets the e rule contexts.
	SetE(IExpContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpContext)

	// GetNodo returns the nodo attribute.
	GetNodo() *Nodo

	// SetNodo sets the nodo attribute.
	SetNodo(*Nodo)

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	nodo   *Nodo
	e1     IExpContext
	e      IExpContext
	n      antlr.Token
	op     antlr.Token
	e2     IExpContext
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetN() antlr.Token { return s.n }

func (s *ExpContext) GetOp() antlr.Token { return s.op }

func (s *ExpContext) SetN(v antlr.Token) { s.n = v }

func (s *ExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpContext) GetE1() IExpContext { return s.e1 }

func (s *ExpContext) GetE() IExpContext { return s.e }

func (s *ExpContext) GetE2() IExpContext { return s.e2 }

func (s *ExpContext) SetE1(v IExpContext) { s.e1 = v }

func (s *ExpContext) SetE(v IExpContext) { s.e = v }

func (s *ExpContext) SetE2(v IExpContext) { s.e2 = v }

func (s *ExpContext) GetNodo() *Nodo { return s.nodo }

func (s *ExpContext) SetNodo(v *Nodo) { s.nodo = v }

func (s *ExpContext) PARI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARI, 0)
}

func (s *ExpContext) PARD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARD, 0)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *ExpContext) POR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPOR, 0)
}

func (s *ExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIV, 0)
}

func (s *ExpContext) MAS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAS, 0)
}

func (s *ExpContext) MEN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMEN, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *GramaticaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *GramaticaParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, GramaticaParserRULE_exp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserPARI:
		{
			p.SetState(34)
			p.Match(GramaticaParserPARI)
		}
		{
			p.SetState(35)

			var _x = p.exp(0)

			localctx.(*ExpContext).e = _x
		}
		{
			p.SetState(36)
			p.Match(GramaticaParserPARD)
		}

		localctx.(*ExpContext).nodo = &Nodo{}
		localctx.(*ExpContext).nodo.dir = localctx.(*ExpContext).GetE().GetNodo().dir
		localctx.(*ExpContext).nodo.num = localctx.(*ExpContext).GetE().GetNodo().num

	case GramaticaParserNUM:
		{
			p.SetState(39)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*ExpContext).n = _m
		}

		localctx.(*ExpContext).nodo = &Nodo{}
		localctx.(*ExpContext).nodo.num = 0
		localctx.(*ExpContext).nodo.dir = (func() string {
			if localctx.(*ExpContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*ExpContext).GetN().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				fmt.Println("wooooo_mul")
				{
					p.SetState(45)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserPOR || _la == GramaticaParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				fmt.Println("wooooo_mul1")
				{
					p.SetState(47)

					var _x = p.exp(5)

					localctx.(*ExpContext).e2 = _x
				}

				localctx.(*ExpContext).nodo = &Nodo{}
				temp = temp - localctx.(*ExpContext).GetE1().GetNodo().num - localctx.(*ExpContext).GetE2().GetNodo().num
				//  localctx.(*ExpContext).nodo.dir = generador.NewTemp()
				localctx.(*ExpContext).nodo.dir = "t" + strconv.Itoa(temp)
				cad := localctx.(*ExpContext).nodo.dir + " = " + localctx.(*ExpContext).GetE1().GetNodo().dir + " " + (func() string {
					if localctx.(*ExpContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpContext).GetOp().GetText()
					}
				}()) + " "
				cad += localctx.(*ExpContext).GetE2().GetNodo().dir
				gen(cad)
				localctx.(*ExpContext).nodo.num = 1
				temp = temp + 1

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				fmt.Println("wooooo-SUM1")
				{
					p.SetState(52)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserMAS || _la == GramaticaParserMEN) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				fmt.Println("woooooSUM2")
				{
					p.SetState(54)

					var _x = p.exp(4)

					localctx.(*ExpContext).e2 = _x
				}

				localctx.(*ExpContext).nodo = &Nodo{}
				temp = temp - localctx.(*ExpContext).GetE1().GetNodo().num - localctx.(*ExpContext).GetE2().GetNodo().num
				localctx.(*ExpContext).nodo.dir = "t" + strconv.Itoa(temp)
				//  localctx.(*ExpContext).nodo.dir = generador.NewTemp()
				cad := localctx.(*ExpContext).nodo.dir + " = " + localctx.(*ExpContext).GetE1().GetNodo().dir + " " + (func() string {
					if localctx.(*ExpContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpContext).GetOp().GetText()
					}
				}()) + " "
				cad += localctx.(*ExpContext).GetE2().GetNodo().dir
				gen(cad)
				localctx.(*ExpContext).nodo.num = 1
				temp = temp + 1

			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
