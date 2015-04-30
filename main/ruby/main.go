package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grubby/grubby/interpreter/vm"
	"github.com/grubby/grubby/parser"
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

	home := os.Getenv("HOME")
	grubbyHome := filepath.Join(home, ".grubby")

	rubyVM := vm.NewVM(grubbyHome, flag.Args()[0])
	defer rubyVM.Exit()

	_, err = rubyVM.Run(string(bytes))

	switch err.(type) {
	case *vm.ParseError:
		offendingFilename := err.(*vm.ParseError).Filename
		println(fmt.Sprintf("Error parsing ruby script %s", offendingFilename))
		println("last 20 statements from the parser:")
		println("")

		debugStatements := []string{}
		for _, d := range parser.DebugStatements {
			debugStatements = append(debugStatements, d)
		}

		threshold := 61
		debugCount := len(debugStatements)
		if debugCount <= threshold {
			for _, stmt := range debugStatements {
				fmt.Printf("\t%s\n", stmt)
			}
		} else {
			for _, stmt := range debugStatements[debugCount-threshold:] {
				fmt.Printf("\t%s\n", stmt)
			}
		}

		os.Exit(1)
	case nil:
	case error:
		panic(err.Error())
	default:
		panic(fmt.Sprintf("%#v", err))
	}
}
