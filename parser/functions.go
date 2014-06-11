package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/grubby/grubby/ast"
)

var (
	funcNameRegexp = regexp.MustCompile("^([a-zA-Z_][a-zA-Z_0-9]*)$")
)

func ParseFunctionDefinition(t string, block *ast.Block, index *int, tokens *[]string) ast.FuncDecl {

	if t != "def" {
		panic("expected to be parsing a function definiton. Instead passed: '" + t + "'")
	}

	nextToken := (*tokens)[(*index)+1]
	if !funcNameRegexp.MatchString(nextToken) {
		panic(fmt.Sprintf("Failed to find expected method name in line: '%s'", nextToken))
	}
	name := nextToken

	(*index) += 2 // have processed def && name tokens

	// keep walking args
	parsedArgs := []ast.Node{}
	if (*tokens)[*index] == "(" {
		for j := *index + 1; j < len(*tokens); j++ {
			tok := (*tokens)[j]
			if tok == ")" {
				*index = j + 1
				break
			}

			if strings.HasSuffix(tok, ",") {
				tok = tok[0 : len(tok)-1]
			}
			parsedArgs = append(parsedArgs, Parse(tok).Statements...)
		}
	}

	decl := ast.FuncDecl{
		Name: name,
		Args: parsedArgs,
	}

	foundEnd := false
	newBlock := ast.Block{Statements: []ast.Node{}}

	// walk the body of the function until we find an unexpected 'END'
	for *index < len(*tokens) {
		err := parseNextTokens(&newBlock, index, tokens)
		if _, ok := err.(*UnexpectedEndError); ok {
			foundEnd = true
			decl.Body = newBlock.Statements
			break
		}
	}

	if !foundEnd {
		panic("Expected to find an end to function '" + decl.Name + "'")
	}

	return decl
}
