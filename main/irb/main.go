package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/grubby/grubby/interpreter/vm"
)

func main() {
	vm := vm.NewVM("(grubby irb")

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
