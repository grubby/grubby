package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/grubby/grubby/interpreter/vm"
)

func main() {
	pathToExecutable, err := filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0])))

	if err != nil {
		panic(err)
	}

	vm := vm.NewVM(pathToExecutable, "(grubby irb")

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
