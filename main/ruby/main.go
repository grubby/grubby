package main

import (
	"io/ioutil"
	"os"

	"github.com/grubby/grubby/interpreter/vm"
)

func main() {
	// for now, assumes this is only being invoked with a filename to interpret
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	vm := vm.NewVM(os.Args[1])
	vm.Run(string(bytes))
}
