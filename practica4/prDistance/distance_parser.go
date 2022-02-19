// Code generated from Distance.g4 by ANTLR 4.9.2. DO NOT EDIT.

package prDistance // Distance
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7,
	3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 2, 2,
	30, 2, 6, 3, 2, 2, 2, 4, 8, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 3, 3, 2, 2,
	2, 8, 9, 8, 3, 1, 2, 9, 10, 7, 3, 2, 2, 10, 11, 7, 6, 2, 2, 11, 12, 7,
	4, 2, 2, 12, 13, 7, 6, 2, 2, 13, 14, 7, 5, 2, 2, 14, 15, 7, 3, 2, 2, 15,
	16, 7, 6, 2, 2, 16, 17, 7, 4, 2, 2, 17, 18, 7, 6, 2, 2, 18, 19, 7, 5, 2,
	2, 19, 28, 3, 2, 2, 2, 20, 21, 12, 4, 2, 2, 21, 22, 7, 3, 2, 2, 22, 23,
	7, 6, 2, 2, 23, 24, 7, 4, 2, 2, 24, 25, 7, 6, 2, 2, 25, 27, 7, 5, 2, 2,
	26, 20, 3, 2, 2, 2, 27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3,
	2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 3, 28,
}
var literalNames = []string{
	"", "'('", "','", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "INT", "WS",
}

var ruleNames = []string{
	"s", "list",
}

type DistanceParser struct {
	*antlr.BaseParser
}

// NewDistanceParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DistanceParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDistanceParser(input antlr.TokenStream) *DistanceParser {
	this := new(DistanceParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Distance.g4"

	return this
}

// DistanceParser tokens.
const (
	DistanceParserEOF  = antlr.TokenEOF
	DistanceParserT__0 = 1
	DistanceParserT__1 = 2
	DistanceParserT__2 = 3
	DistanceParserINT  = 4
	DistanceParserWS   = 5
)

// DistanceParser rules.
const (
	DistanceParserRULE_s    = 0
	DistanceParserRULE_list = 1
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
	p.RuleIndex = DistanceParserRULE_s
	return p
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DistanceParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DistanceVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DistanceParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DistanceParserRULE_s)

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
		p.SetState(4)
		p.list(0)
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DistanceParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DistanceParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) CopyFrom(ctx *ListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L1Context struct {
	*ListContext
}

func NewL1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *L1Context {
	var p = new(L1Context)

	p.ListContext = NewEmptyListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ListContext))

	return p
}

func (s *L1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L1Context) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *L1Context) AllINT() []antlr.TerminalNode {
	return s.GetTokens(DistanceParserINT)
}

func (s *L1Context) INT(i int) antlr.TerminalNode {
	return s.GetToken(DistanceParserINT, i)
}

func (s *L1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.EnterL1(s)
	}
}

func (s *L1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.ExitL1(s)
	}
}

func (s *L1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DistanceVisitor:
		return t.VisitL1(s)

	default:
		return t.VisitChildren(s)
	}
}

type L2Context struct {
	*ListContext
}

func NewL2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *L2Context {
	var p = new(L2Context)

	p.ListContext = NewEmptyListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ListContext))

	return p
}

func (s *L2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L2Context) AllINT() []antlr.TerminalNode {
	return s.GetTokens(DistanceParserINT)
}

func (s *L2Context) INT(i int) antlr.TerminalNode {
	return s.GetToken(DistanceParserINT, i)
}

func (s *L2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.EnterL2(s)
	}
}

func (s *L2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.ExitL2(s)
	}
}

func (s *L2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DistanceVisitor:
		return t.VisitL2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DistanceParser) List() (localctx IListContext) {
	return p.list(0)
}

func (p *DistanceParser) list(_p int) (localctx IListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, DistanceParserRULE_list, _p)

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
	localctx = NewL2Context(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(7)
		p.Match(DistanceParserT__0)
	}
	{
		p.SetState(8)
		p.Match(DistanceParserINT)
	}
	{
		p.SetState(9)
		p.Match(DistanceParserT__1)
	}
	{
		p.SetState(10)
		p.Match(DistanceParserINT)
	}
	{
		p.SetState(11)
		p.Match(DistanceParserT__2)
	}
	{
		p.SetState(12)
		p.Match(DistanceParserT__0)
	}
	{
		p.SetState(13)
		p.Match(DistanceParserINT)
	}
	{
		p.SetState(14)
		p.Match(DistanceParserT__1)
	}
	{
		p.SetState(15)
		p.Match(DistanceParserINT)
	}
	{
		p.SetState(16)
		p.Match(DistanceParserT__2)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewL1Context(p, NewListContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, DistanceParserRULE_list)
			p.SetState(18)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(19)
				p.Match(DistanceParserT__0)
			}
			{
				p.SetState(20)
				p.Match(DistanceParserINT)
			}
			{
				p.SetState(21)
				p.Match(DistanceParserT__1)
			}
			{
				p.SetState(22)
				p.Match(DistanceParserINT)
			}
			{
				p.SetState(23)
				p.Match(DistanceParserT__2)
			}

		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

func (p *DistanceParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ListContext = nil
		if localctx != nil {
			t = localctx.(*ListContext)
		}
		return p.List_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DistanceParser) List_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
