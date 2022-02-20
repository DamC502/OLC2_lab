package main

import (
	"corto1/prHex"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

func main() {
	file, err := antlr.NewFileStream("entrada.txt")
	if err != nil {
		log.Fatalln(err)
	}
	lexer := prHex.NewHexExpLexer(file)
	tokens := antlr.NewCommonTokenStream( lexer, antlr.TokenDefaultChannel)
	parser := prHex.NewHexExpParser(tokens)
	visitor := NewHexVisitorImp()
	tree := parser.S()
	visitor.Visit(tree)

}
