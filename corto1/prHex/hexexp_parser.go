// Code generated from HexExp.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prHex // HexExp
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 29, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 5, 2, 11, 10, 2, 6, 2, 13,
	10, 2, 13, 2, 14, 2, 14, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 22, 10, 3,
	3, 4, 6, 4, 25, 10, 4, 13, 4, 14, 4, 26, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2,
	2, 29, 2, 12, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 24, 3, 2, 2, 2, 8, 10,
	5, 4, 3, 2, 9, 11, 7, 3, 2, 2, 10, 9, 3, 2, 2, 2, 10, 11, 3, 2, 2, 2, 11,
	13, 3, 2, 2, 2, 12, 8, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 12, 3, 2, 2,
	2, 14, 15, 3, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 17, 5, 6, 4, 2, 17, 18, 7,
	4, 2, 2, 18, 19, 5, 6, 4, 2, 19, 22, 3, 2, 2, 2, 20, 22, 5, 6, 4, 2, 21,
	16, 3, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 25, 7, 5, 2,
	2, 24, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27,
	3, 2, 2, 2, 27, 7, 3, 2, 2, 2, 6, 10, 14, 21, 26,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'.'",
}
var symbolicNames = []string{
	"", "", "", "CHAR",
}

var ruleNames = []string{
	"s", "l", "a",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type HexExpParser struct {
	*antlr.BaseParser
}

func NewHexExpParser(input antlr.TokenStream) *HexExpParser {
	this := new(HexExpParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "HexExp.g4"

	return this
}

// HexExpParser tokens.
const (
	HexExpParserEOF  = antlr.TokenEOF
	HexExpParserT__0 = 1
	HexExpParserT__1 = 2
	HexExpParserCHAR = 3
)

// HexExpParser rules.
const (
	HexExpParserRULE_s = 0
	HexExpParserRULE_l = 1
	HexExpParserRULE_a = 2
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HexExpParserRULE_s
	return p
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HexExpParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) AllL() []ILContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILContext)(nil)).Elem())
	var tst = make([]ILContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILContext)
		}
	}

	return tst
}

func (s *SContext) L(i int) ILContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HexExpVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HexExpParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HexExpParserRULE_s)
	var _la int

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
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == HexExpParserCHAR {
		{
			p.SetState(6)
			p.L()
		}
		p.SetState(8)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == HexExpParserT__0 {
			{
				p.SetState(7)
				p.Match(HexExpParserT__0)
			}

		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILContext is an interface to support dynamic dispatch.
type ILContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLContext differentiates from other interfaces.
	IsLContext()
}

type LContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLContext() *LContext {
	var p = new(LContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HexExpParserRULE_l
	return p
}

func (*LContext) IsLContext() {}

func NewLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LContext {
	var p = new(LContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HexExpParserRULE_l

	return p
}

func (s *LContext) GetParser() antlr.Parser { return s.parser }

func (s *LContext) CopyFrom(ctx *LContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L1Context struct {
	*LContext
}

func NewL1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *L1Context {
	var p = new(L1Context)

	p.LContext = NewEmptyLContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LContext))

	return p
}

func (s *L1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L1Context) AllA() []IAContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAContext)(nil)).Elem())
	var tst = make([]IAContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAContext)
		}
	}

	return tst
}

func (s *L1Context) A(i int) IAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAContext)
}

func (s *L1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.EnterL1(s)
	}
}

func (s *L1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.ExitL1(s)
	}
}

func (s *L1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HexExpVisitor:
		return t.VisitL1(s)

	default:
		return t.VisitChildren(s)
	}
}

type L2Context struct {
	*LContext
}

func NewL2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *L2Context {
	var p = new(L2Context)

	p.LContext = NewEmptyLContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LContext))

	return p
}

func (s *L2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L2Context) A() IAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAContext)
}

func (s *L2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.EnterL2(s)
	}
}

func (s *L2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.ExitL2(s)
	}
}

func (s *L2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HexExpVisitor:
		return t.VisitL2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HexExpParser) L() (localctx ILContext) {
	localctx = NewLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HexExpParserRULE_l)

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

	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewL1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.A()
		}
		{
			p.SetState(15)
			p.Match(HexExpParserT__1)
		}
		{
			p.SetState(16)
			p.A()
		}

	case 2:
		localctx = NewL2Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.A()
		}

	}

	return localctx
}

// IAContext is an interface to support dynamic dispatch.
type IAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAContext differentiates from other interfaces.
	IsAContext()
}

type AContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAContext() *AContext {
	var p = new(AContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HexExpParserRULE_a
	return p
}

func (*AContext) IsAContext() {}

func NewAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AContext {
	var p = new(AContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HexExpParserRULE_a

	return p
}

func (s *AContext) GetParser() antlr.Parser { return s.parser }

func (s *AContext) AllCHAR() []antlr.TerminalNode {
	return s.GetTokens(HexExpParserCHAR)
}

func (s *AContext) CHAR(i int) antlr.TerminalNode {
	return s.GetToken(HexExpParserCHAR, i)
}

func (s *AContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.EnterA(s)
	}
}

func (s *AContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HexExpListener); ok {
		listenerT.ExitA(s)
	}
}

func (s *AContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HexExpVisitor:
		return t.VisitA(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HexExpParser) A() (localctx IAContext) {
	localctx = NewAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HexExpParserRULE_a)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(21)
				p.Match(HexExpParserCHAR)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}
