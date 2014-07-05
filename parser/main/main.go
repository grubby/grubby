package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/grubby/grubby/parser"
)

func main() {
	stdin := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		fmt.Printf("Give me your best line of ruby: ")

		if line, err := stdin.ReadString('\n'); err == nil {
			lexer := parser.NewLexer(line)

			if parser.RubyParse(lexer) != 0 {
				break
			}

			statements := parser.Statements
			fmt.Printf("There are now %d statements\n", len(statements))
			fmt.Printf("The latest statement is %#v\n", statements[len(statements)-1])
		} else {
			println("FAIL:", err.Error())
			break
		}
	}
}
