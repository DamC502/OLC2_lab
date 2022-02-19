package main

import (
	"fmt"
	"log"
	"math"
	"practica4/prDistance"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type DistanceVisitorImpl struct {
	antlr.ParseTreeVisitor	
}

type PairXY struct {
	x        float64
	y        float64
	distance float64
}

func (p *PairXY) AddPair(x float64, y float64) {
	fmt.Print("Distancia actual: ", p.distance, " desde  x:", p.x, " , y:",
		p.y, " -> x0:", x, " , y0:", y, " nueva distancia: ")
	p.distance = p.distance + math.Sqrt(math.Pow(p.x-x, 2)+
		math.Pow(p.y-y, 2))
	fmt.Println(p.distance)

	p.x = x
	p.y = y
}

func NewDistanceVisitorImp() prDistance.DistanceVisitor {
	return &DistanceVisitorImpl{
		ParseTreeVisitor: &prDistance.BaseDistanceVisitor{},
	}
}

func (d *DistanceVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	// switch val := tree.(type) {
	// 	case *prDistance.L1Context:
	// 		return val.Accept(c)
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(t.GetText())
		return PairXY{0, 0, 0}
	default:
		if node, ok := tree.Accept(d).(PairXY); ok {
			return node
		}

	}
	return PairXY{0, 0, 0}
}

func (d *DistanceVisitorImpl) VisitS(ctx *prDistance.SContext) interface{} {
	return ctx.List().Accept(d).(PairXY)
}

func (d *DistanceVisitorImpl) VisitL1(ctx *prDistance.L1Context) interface{} {

	list := ctx.List().Accept(d).(PairXY)
	x, err := strconv.ParseFloat(ctx.INT(0).GetText(), 10)
	if err != nil {
		panic(err)
	}
	y, err2 := strconv.ParseFloat(ctx.INT(1).GetText(), 10)
	if err2 != nil {
		panic(err)
	}
	list.AddPair(x, y)
	return list
}

func (d *DistanceVisitorImpl) VisitL2(ctx *prDistance.L2Context) interface{} {
	x, err := strconv.ParseFloat(ctx.INT(0).GetText(), 10)
	if err != nil {
		panic(err)
	}
	y, err2 := strconv.ParseFloat(ctx.INT(1).GetText(), 10)
	if err2 != nil {
		panic(err)
	}
	pair := PairXY{x, y, 0}

	x, err = strconv.ParseFloat(ctx.INT(2).GetText(), 10)
	if err != nil {
		panic(err)
	}
	y, err2 = strconv.ParseFloat(ctx.INT(3).GetText(), 10)
	if err2 != nil {
		panic(err)
	}
	pair.AddPair(x, y)

	return pair
}
