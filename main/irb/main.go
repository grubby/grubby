package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/grubby/grubby/interpreter/vm"
)

func main() {
	home := os.Getenv("HOME")
	grubbyHome := filepath.Join(home, ".grubby")

	vm := vm.NewVM(grubbyHome, "(grubby irb")

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
