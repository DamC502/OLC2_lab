package main

import (
	"corto1/prHex"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)
type HexExprVisitorImp struct {
	antlr.ParseTreeVisitor
}

type result struct {
	decimal float64 // decimal es el aceptado
	entero int64
	error bool
}
func NewHexVisitorImp() prHex.HexExpVisitor {
	return &HexExprVisitorImp{
		ParseTreeVisitor: &prHex.BaseHexExpVisitor{},
	}
}

func (h* HexExprVisitorImp)Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatalln(t.GetText())
		return result{0.0,0,true}
	default:
		if node, ok := tree.Accept(h).(result); ok {
			return node
		}
	}
	return result{0.0,0,false}
}

func (h *HexExprVisitorImp) VisitS(ctx *prHex.SContext) interface{} {
	nodo_aux := result{}
	for _,nodo_l := range ctx.AllL(){
		nodo_aux = nodo_l.Accept(h).(result)
		fmt.Println(nodo_aux.decimal, ",")
	}
	return nodo_aux
}

func (h *HexExprVisitorImp) VisitL1(ctx *prHex.L1Context) interface{} {
	nodo_entero := ctx.A(0).Accept(h).(result)
	nodo_decimal := ctx.A(1).Accept(h).(result)
	nodo_decimal.decimal = float64(nodo_entero.entero) + nodo_decimal.decimal
	return nodo_decimal
}

func (h *HexExprVisitorImp) VisitL2(ctx *prHex.L2Context) interface{} {
	nodo_a := ctx.A().Accept(h).(result)
	nodo_a.decimal =  float64(nodo_a.entero)
	return nodo_a;
}

func (h *HexExprVisitorImp) VisitA(ctx *prHex.AContext) interface{} {
	resultado := result{0,0.0, false}
	nivel := 16.0;
	for _,char :=  range ctx.AllCHAR() {
		switch char.GetText()[0] {
		case '0':
			resultado.entero = (resultado.entero*16) + 0
			resultado.decimal = resultado.decimal + (0 / nivel)

			break;
		case '1':
			resultado.entero = (resultado.entero*16) + 1
			resultado.decimal = resultado.decimal + (1 / nivel)
			break;
		case '2':
			resultado.entero = (resultado.entero*16) + 2
			resultado.decimal = resultado.decimal + (2 / nivel)
			break;
		case '3':
			resultado.entero = (resultado.entero*16) + 3
			resultado.decimal = resultado.decimal + (3 / nivel)
			break;
		case '4':
			resultado.entero = (resultado.entero*16) + 4
			resultado.decimal = resultado.decimal + (4 / nivel)
			break;
		case '5':
			resultado.entero = (resultado.entero*16) + 5
			resultado.decimal = resultado.decimal + (5 / nivel)
			break;
		case '6':
			resultado.entero = (resultado.entero*16) + 6
			resultado.decimal = resultado.decimal + (6 / nivel)
			break;
		case '7':
			resultado.entero = (resultado.entero*16) + 7
			resultado.decimal = resultado.decimal + (7 / nivel)
			break;
		case '8':
			resultado.entero = (resultado.entero*16) + 8
			resultado.decimal = resultado.decimal + (8 / nivel)
			break;
		case '9':
			resultado.entero = (resultado.entero*16) + 9
			resultado.decimal = resultado.decimal + (9 / nivel)
			break;
		case 'A':
			resultado.entero = (resultado.entero*16) + 10
			resultado.decimal = resultado.decimal + (10 / nivel)
			break;
		case 'B':
			resultado.entero = (resultado.entero*16) + 11
			resultado.decimal = resultado.decimal + (11 / nivel)
			break;
		case 'C':
			resultado.entero = (resultado.entero*16) + 12
			resultado.decimal = resultado.decimal + (12 / nivel)
			break;
		case 'D':
			resultado.entero = (resultado.entero*16) + 13
			resultado.decimal = resultado.decimal + (13 / nivel)
			break;
		case 'E':
			resultado.entero = (resultado.entero*16) + 14
			resultado.decimal = resultado.decimal + (14 / nivel)
			break;
		case 'F':
			resultado.entero = (resultado.entero*16) + 15
			resultado.decimal = resultado.decimal + (15 / nivel)
		default:
			fmt.Println("Error caracter no reconocido: ", char.GetText() )
			resultado.error = true
		}
		nivel = nivel*16
	}
	return resultado
}

