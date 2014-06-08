package parser

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/grubby/grubby/ast"
)

var (
	funcNameRegexp = regexp.MustCompile("^\\s*def (.+)$")
	funcEndRegexp  = regexp.MustCompile("^\\s*end\\s*$")
)

func ParseFunctionDefinition(lines []string) (int, ast.Node, error) {
	var (
		index  int
		closed = false
	)

	matches := funcNameRegexp.FindStringSubmatch(lines[0])
	if len(matches) != 2 {
		err := errors.New(fmt.Sprintf("Could not match function name in line '%s'", lines[0]))
		return 0, nil, err
	}
	name := matches[1]

	if name == "" {
		err := errors.New(fmt.Sprintf("Failed to find expected method name in line: '%s'", lines[0]))
		return 0, nil, err
	}

	node := ast.FuncDecl{Name: name}
	for i, line := range lines[1:] {
		if funcEndRegexp.MatchString(line) {
			index = i + 1
			closed = true
			break
		}

		node.Body = append(node.Body, Parse(line).Statement)
	}

	if !closed {
		err := errors.New(fmt.Sprintf("Func '%s', was defined but did not see an 'end' statement", name))
		return 0, nil, err
	}

	return index, node, nil
}
