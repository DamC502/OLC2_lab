// practica4 project main.go
package main

import (
	"fmt"
	"log"

	"practica4/prDistance"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	file, err := antlr.NewFileStream("entrada.txt")
	if err != nil {
		log.Fatal(err)
	}
	lexer := prDistance.NewDistanceLexer(file)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	pr := prDistance.NewDistanceParser(tokens)
	visitor := NewDistanceVisitorImp()
	tree := pr.S()
	root := visitor.Visit(tree).(PairXY)
	fmt.Println("La distancia final es de: ", root.distance)

}
