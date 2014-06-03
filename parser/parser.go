package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/grubby/grubby/ast"
)

var (
	integerRegexp  = regexp.MustCompile("^[0-9]+$")
	floatRegexp    = regexp.MustCompile("^[0-9]+\\.[0-9]+$")
	stringRegexp   = regexp.MustCompile("^'[^']+'$")
	bareRefRegexp  = regexp.MustCompile("^[a-zA-Z_][a-zA-Z_0-9]*$")
	callExprRegexp = regexp.MustCompile("^[a-zA-Z_][a-zA-Z_0-9]*\\((.*)\\)$")
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
		} else if bareRefRegexp.MatchString(line) {
			block.Statement = ast.BareReference{Name: line}
		} else if callExprRegexp.MatchString(line) {
			index := strings.Index(line, "(")
			remainingArgs := line[index+1 : len(line)-1]

			args := make([]ast.Node, 0)
			if remainingArgs != "" {
				args = append(args, Parse(remainingArgs).Statement)
			}

			block.Statement = ast.CallExpression{
				Func: line[0:index],
				Args: args,
			}
		} else {
			panic(fmt.Sprintf("unknown type %T, (%#v)", line, line))
		}
	}

	return block
}
