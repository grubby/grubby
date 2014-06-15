package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/grubby/grubby/ast"
)

var (
	namespaceRegexp = regexp.MustCompile("^([A-Z][a-zA-Z_0-9]*::)+([A-Z][a-zA-Z_0-9]*)")
	classNameRegex  = regexp.MustCompile("^(?:(?:[A-Z][a-zA-Z_0-9]*::))*([A-Z][a-zA-Z_0-9]*)\\s*(<\\s[A-Z][a-zA-Z_0-9]*)?")
)

func ParseClassDefinition(t string, block *ast.Block, index *int, tokens *[]string) ast.ClassDefn {
	if t != "class" {
		panic("Expected to be parsing a class definition. Received: '" + t + "'")
	}

	*index += 1 // move past the class token
	nextToken := (*tokens)[*index]
	if !classNameRegex.MatchString(nextToken) {
		panic(fmt.Sprintf("Failed to find class name in '%s'", nextToken))
	}

	name := nextToken
	node := ast.ClassDefn{Name: name, Body: []ast.Node{}}

	matches := namespaceRegexp.FindStringSubmatch(nextToken)
	if len(matches) > 2 {
		name := matches[len(matches)-1]
		namespaces := strings.Replace(matches[0], name, "", 1)
		node.Namespace = namespaces[:len(namespaces)-2]
		node.Name = name
	}

	*index += 1
	nextToken = (*tokens)[(*index)]

	if nextToken == "<" {
		*index += 2
		nextToken = (*tokens)[*index-1]
		if !classNameRegex.MatchString(nextToken) {
			panic(fmt.Sprintf("Expected super class '%s' to be a class name", nextToken))
		}

		node.SuperClass = nextToken
	}

	foundEnd := false
	newBlock := ast.Block{Statements: []ast.Node{}}

	for *index < len(*tokens) {
		err := parseNextTokens(&newBlock, index, tokens)
		if _, ok := err.(*UnexpectedEndError); ok {
			foundEnd = true
			node.Body = append(node.Body, newBlock.Statements...)
			break
		}
	}

	if !foundEnd {
		panic("Expected to find an end to class '" + node.Name + "'")
	}

	*index += 1 // move past 'end'
	return node
}
