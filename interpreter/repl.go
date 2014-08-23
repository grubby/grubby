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
	vm := vm.NewVM("(irb)")

	for {
		result, err := vm.Run(readInput())
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("=> %s", result.String())
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
