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

			something := parser.RubyParse(lexer)
			fmt.Printf("Got a %T: %#v\n", something, something)
		} else {
			println("FAIL:", err.Error())
			break
		}
	}
}
