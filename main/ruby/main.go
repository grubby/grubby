package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/grubby/grubby/interpreter/vm"
)

var verboseFlag = flag.Bool("verbose", false, "enables verbose mode")

func init() {
	flag.BoolVar(verboseFlag, "v", false, "enables verbose mode")
}

func main() {
	flag.Parse()

	// for now, assumes this is only being invoked with a filename to interpret
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	vm := vm.NewVM(flag.Args()[1])
	vm.Run(string(bytes))
}
