package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/grubby/grubby/parser"
)

var lexer parser.RubyLexer

func main() {
	for {
		lexer = parser.NewLexer(readInput())
		parser.RubyParse(lexer)
		println("parsed", len(parser.Statements), "statements")
		for _, stmt := range parser.Statements {
			fmt.Printf("%#v\n", stmt)
		}
	}
}

func readInput() string {
	print(" > ")
	bio := bufio.NewReader(os.Stdin)
	userInput, err := bio.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return userInput
}
