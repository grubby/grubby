package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/grubby/grubby/interpreter/vm"
	"github.com/grubby/grubby/parser"
)

var lexer parser.RubyLexer

func main() {
	vm := vm.NewVM()

	for {
		vm.Run(readInput())

		println("parsed", len(parser.Statements), "statements")
		for _, stmt := range parser.Statements {
			fmt.Printf("%#v\n", stmt)
		}
	}
}

func readInput() string {
	print("> ")
	bio := bufio.NewReader(os.Stdin)
	userInput, err := bio.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return userInput
}
