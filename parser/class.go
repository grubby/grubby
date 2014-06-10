package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/grubby/grubby/ast"
)

var (
	classDefnRegex = regexp.MustCompile("^\\s*class ([A-Z][a-zA-Z_0-9]*)\\s*(<\\s[A-Z][a-zA-Z_0-9]*)?")
)

func ParseClassDefinition(lines []string) (int, ast.Node, error) {
	matches := classDefnRegex.FindStringSubmatch(lines[0])
	if len(matches) < 2 {
		err := errors.New(fmt.Sprintf("Could not match class name in line '%s'", lines[0]))
		return 0, nil, err
	}

	node := ast.ClassDefn{Name: matches[1]}

	if len(matches) > 2 && len(matches[2]) > 0 {
		superClass := matches[2]
		node.SuperClass = strings.TrimSpace(superClass[strings.Index(superClass, "<")+1:])
	}

	return 1, node, nil
}
