package parser

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/grubby/grubby/ast"
)

var (
	classNameRegexp = regexp.MustCompile("^\\s*class ([A-Z][a-zA-Z_0-9]*)$")
)

func ParseClassDefinition(lines []string) (int, ast.Node, error) {
	matches := classNameRegexp.FindStringSubmatch(lines[0])
	if len(matches) < 2 {
		err := errors.New(fmt.Sprintf("Could not match class name in line '%s'", lines[0]))
		return 0, nil, err
	}

	node := ast.ClassDefn{Name: matches[1]}
	return 1, node, nil
}
