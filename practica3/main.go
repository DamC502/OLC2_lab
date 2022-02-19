// practica3 project main.go
package main

import (
	"log"
	"practica3/prExpr"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	file, err := antlr.NewFileStream("entrada.txt")
	if err != nil {
		log.Fatal(err)
	}
	lexer := prExpr.NewExpressionsLexer(file)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := prExpr.NewExpressionsParser(tokens)
	visitor := NewExpressionsVisitorImp()
	tree := parser.S()
	visitor.Visit(tree)
}
