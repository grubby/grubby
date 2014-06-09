package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/grubby/grubby/ast"
)

var (
	whitespace       = regexp.MustCompile("^\\s*$")
	integerRegexp    = regexp.MustCompile("^\\s*[0-9]+$")
	floatRegexp      = regexp.MustCompile("^\\s*[0-9]+\\.[0-9]+$")
	stringRegexp     = regexp.MustCompile("^\\s*'[^']+'$")
	symbolRegexp     = regexp.MustCompile("^\\s*:[a-zA-Z-][a-zA-Z_0-9]*$")
	openingString    = regexp.MustCompile("^\\s*'[^']*$")
	bareRefRegexp    = regexp.MustCompile("^\\s*[a-zA-Z_][a-zA-Z_0-9]*$")
	callExprRegexp   = regexp.MustCompile("^\\s*[a-zA-Z_][a-zA-Z_0-9]*\\((.*)\\)$")
	methodDefnRegexp = regexp.MustCompile("^\\s*def [a-zA-Z_][a-zA-Z_0-9]*(\\(.*\\))?")
)

func Parse(input string) ast.Block {
	block := ast.Block{}

	var line string
	lines := strings.Split(input, "\n")
	for index := 0; index < len(lines); index++ {
		line = lines[index]

		if whitespace.MatchString(line) {
			continue
		}

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
		} else if symbolRegexp.MatchString(line) {
			index := strings.Index(line, ":")
			block.Statement = ast.Symbol{Name: strings.TrimSpace(line[index+1:])}
		} else if callExprRegexp.MatchString(line) {
			index := strings.Index(line, "(")
			remainingArgs := line[index+1 : len(line)-1]

			args := make([]ast.Node, 0)
			if remainingArgs != "" {
				for _, piece := range strings.Split(remainingArgs, ",") {

					args = append(args, Parse(piece).Statement)
				}
			}

			block.Statement = ast.CallExpression{
				Func: strings.TrimSpace(line[0:index]),
				Args: args,
			}
		} else if openingString.MatchString(line) {
			linesRead, stringBlock, err := ParseMultilineString(lines[index:])
			if err != nil {
				panic(err)
			}

			index += linesRead
			block.Statement = stringBlock
		} else if methodDefnRegexp.MatchString(line) {
			linesRead, method, err := ParseFunctionDefinition(lines[index:])
			if err != nil {
				panic(err)
			}

			index += linesRead
			block.Statement = method
		} else {
			panic(fmt.Sprintf("unknown type %T, (%#v)", line, line))
		}
	}

	return block
}
