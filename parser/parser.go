package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/grubby/grubby/ast"
)

var (
	integerRegexp = regexp.MustCompile("^[0-9]+$")
	floatRegexp   = regexp.MustCompile("^[0-9]+\\.[0-9]+$")
	stringRegexp  = regexp.MustCompile("^'[^']+'$")
)

func Parse(input string) ast.Block {
	block := ast.Block{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if integerRegexp.MatchString(line) {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			block.Statement = ast.ConstantInt{Value: val}
		} else if stringRegexp.MatchString(line) {
			block.Statement = ast.SimpleString{Value: line[1 : len(line)-1]}
		} else if floatRegexp.MatchString(line) {
			val, err := strconv.ParseFloat(line, 10)
			if err != nil {
				panic(err)
			}

			block.Statement = ast.ConstantFloat{Value: val}
		} else {
			println("unknown type:", line)
		}
	}

	return block
}
