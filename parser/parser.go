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
	whitespace     = regexp.MustCompile("^\\s*$")
	integerRegexp  = regexp.MustCompile("^[0-9]+$")
	floatRegexp    = regexp.MustCompile("^[0-9]+\\.[0-9]+$")
	stringRegexp   = regexp.MustCompile("^'[^']+'$")
	symbolRegexp   = regexp.MustCompile("^:[a-zA-Z-][a-zA-Z_0-9]*$")
	openingString  = regexp.MustCompile("^\\s*'[^']*$")
	bareRefRegexp  = regexp.MustCompile("^[a-zA-Z_][a-zA-Z_0-9]*$")
	callExprRegexp = regexp.MustCompile("^[a-zA-Z_][a-zA-Z_0-9]*\\(")
	classRegexp    = regexp.MustCompile("^\\s*class [A-Z][a-zA-Z_0-9]*")
)

func parseNextTokens(block *ast.Block, index *int, tokens *[]string) error {
	t := (*tokens)[*index]
	defer func() { *index += 1 }()

	switch {
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
		i := strings.Index(t, ":")
		block.Statements = append(block.Statements, ast.Symbol{Name: t[i+1:]})
	case t == "def":
		defn := ParseFunctionDefinition(t, block, index, tokens)
		block.Statements = append(block.Statements, defn)

		// because of defer above (this function handles all of its index incrementing)
		*index--
	case t == "class":
		defn := ParseClassDefinition(t, block, index, tokens)
		block.Statements = append(block.Statements, defn)
		// because of defer above (this function handles all of its index incrementing)
		*index--
	case callExprRegexp.MatchString(t) || peek(tokens, index) == "(":
		// find first and last parens
		firstParen := -1
		lastParen := -1

		for i := *index; i < len(*tokens); i++ {
			if (*tokens)[i] == "(" {
				firstParen = i
				break
			}
		}

		if firstParen == -1 {
			panic("could not find an opening paren")
		}

		for i := firstParen + 1; i < len(*tokens); i++ {
			if (*tokens)[i] == ")" {
				lastParen = i
				break
			}
		}

		if lastParen == -1 {
			panic("could not find a closing paren")
		}

		args := make([]ast.Node, 0)
		for i := firstParen + 1; i < lastParen; i++ {
			piece := (*tokens)[i]
			if strings.HasSuffix(piece, ",") {
				piece = piece[:len(piece)-1]
			}

			args = append(args, Parse(piece).Statements...)
		}

		(*index) = lastParen // move index past last paren of call expr

		block.Statements = append(block.Statements, ast.CallExpression{
			Func: t,
			Args: args,
		})
	case bareRefRegexp.MatchString(t) && !isKeyword(t):
		block.Statements = append(block.Statements, ast.BareReference{Name: t})
	case t == "end":
		return NewUnexpectedEnd()
	default:
		panic(fmt.Sprintf("unknown token (%#v)", t))
	}

	return nil
}

func Parse(input string) ast.Block {
	block := ast.Block{Statements: []ast.Node{}}
	lexer := token.NewLexer()

	index := 0
	tokens := lexer.Tokenize(input)
	for index < len(tokens) {
		err := parseNextTokens(&block, &index, &tokens)
		if err != nil {
			panic(err.Error())
		}
	}

	return block
}

func isKeyword(t string) bool {
	switch t {
	case "def":
		return true
	case "end":
		return true
	case "class":
		return true
	default:
		return false
	}
}

func peek(tokens *[]string, index *int) string {
	if (*index) >= len(*tokens)-1 {
		return ""
	}
	return (*tokens)[*index+1]
}
