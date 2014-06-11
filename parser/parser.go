package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/token"
)

var (
	whitespace       = regexp.MustCompilea("^\\s*$")
	integerRegexp    = regexp.MustCompile("^[0-9]+$")
	floatRegexp      = regexp.MustCompile("^[0-9]+\\.[0-9]+$")
	stringRegexp     = regexp.MustCompile("^'[^']+'$")
	symbolRegexp     = regexp.MustCompile("^:[a-zA-Z-][a-zA-Z_0-9]*$")
	openingString    = regexp.MustCompile("^\\s*'[^']*$")
	bareRefRegexp    = regexp.MustCompile("^[a-zA-Z_][a-zA-Z_0-9]*$")
	callExprRegexp   = regexp.MustCompile("^[a-zA-Z_][a-zA-Z_0-9]*\\((.*)\\)$")
	methodDefnRegexp = regexp.MustCompile("^\\s*def [a-zA-Z_][a-zA-Z_0-9]*(\\(.*\\))?")
	classRegexp      = regexp.MustCompile("^\\s*class [A-Z][a-zA-Z_0-9]*")
)

func Parse(input string) ast.Block {
	block := ast.Block{Statements: []ast.Node{}}

	for _, t := range token.Tokenize(input) {
		switch {
		case bareRefRegexp.MatchString(t):
			block.Statements = append(block.Statements, ast.BareReference{Name: t})
		case integerRegexp.MatchString(t):
			val, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}

			block.Statements = append(block.Statements, ast.ConstantInt{Value: val})
		case floatRegexp.MatchString(t):
			val, err := strconv.ParseFloat(t, 10)
			if err != nil {
				panic(err)
			}

			block.Statements = append(block.Statements, ast.ConstantFloat{val})
		case stringRegexp.MatchString(t):
			block.Statements = append(block.Statements, ast.SimpleString{Value: t[1 : len(t)-1]})
		case symbolRegexp.MatchString(t):
			index := strings.Index(t, ":")
			block.Statements = append(block.Statements, ast.Symbol{Name: t[index+1:]})
		case callExprRegexp.MatchString(t):
			index := strings.Index(t, "(")
			remainingArgs := t[index+1 : len(t)-1]

			args := make([]ast.Node, 0)
			if remainingArgs != "" {
				for _, piece := range strings.Split(remainingArgs, ",") {

					args = append(args, Parse(piece).Statements...)
				}
			}

			block.Statements = append(block.Statements, ast.CallExpression{
				Func: strings.TrimSpace(t[0:index]),
				Args: args,
			})

		default:
			panic(fmt.Sprintf("unknown token (%#v)", t))
		}
	}

	return block
}

// func Parse(input string) ast.Block {
// 	block := ast.Block{}

// 	var line string
// 	lines := strings.Split(input, "\n")
// 	for index := 0; index < len(lines); index++ {
// 		line = lines[index]

// 		if whitespace.MatchString(line) {
// 			continue
// 		}

// 		if integerRegexp.MatchString(line) {
// 			val, err := strconv.Atoi(line)
// 			if err != nil {
// 				panic(err)
// 			}

// 			block.Statement = ast.ConstantInt{Value: val}
// 		} else if stringRegexp.MatchString(line) {
// 			block.Statement = ast.SimpleString{Value: line[1 : len(line)-1]}
// 		} else if floatRegexp.MatchString(line) {
// 			val, err := strconv.ParseFloat(line, 10)
// 			if err != nil {
// 				panic(err)
// 			}

// 			block.Statement = ast.ConstantFloat{Value: val}
// 		} else if bareRefRegexp.MatchString(line) {
// 			block.Statement = ast.BareReference{Name: line}
// 		} else if symbolRegexp.MatchString(line) {
// 			index := strings.Index(line, ":")
// 			block.Statement = ast.Symbol{Name: strings.TrimSpace(line[index+1:])}
// 		} else if callExprRegexp.MatchString(line) {
// 			index := strings.Index(line, "(")
// 			remainingArgs := line[index+1 : len(line)-1]

// 			args := make([]ast.Node, 0)
// 			if remainingArgs != "" {
// 				for _, piece := range strings.Split(remainingArgs, ",") {

// 					args = append(args, Parse(piece).Statement)
// 				}
// 			}

// 			block.Statement = ast.CallExpression{
// 				Func: strings.TrimSpace(line[0:index]),
// 				Args: args,
// 			}
// 		} else if openingString.MatchString(line) {
// 			linesRead, stringBlock, err := ParseMultilineString(lines[index:])
// 			if err != nil {
// 				panic(err)
// 			}

// 			index += linesRead
// 			block.Statement = stringBlock
// 		} else if methodDefnRegexp.MatchString(line) {
// 			linesRead, method, err := ParseFunctionDefinition(lines[index:])
// 			if err != nil {
// 				panic(err)
// 			}

// 			index += linesRead
// 			block.Statement = method
// 		} else if classRegexp.MatchString(line) {
// 			linesRead, class, err := ParseClassDefinition(lines[index:])
// 			if err != nil {
// 				panic(err)
// 			}

// 			index += linesRead
// 			block.Statement = class
// 		} else {
// 			panic(fmt.Sprintf("unknown type %T, (%#v)", line, line))
// 		}
// 	}

// 	return block
// }
