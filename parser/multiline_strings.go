package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/grubby/grubby/ast"
)

var terminatingRegexp = regexp.MustCompile("^\\s*[^']+'$")

func ParseMultilineString(lines []string) (int, ast.SimpleString, error) {
	var (
		closed      = false
		stmt        ast.SimpleString
		index       int
		line        string
		joinedLines = []string{}
	)

	for index, line = range lines {
		joinedLines = append(joinedLines, line)

		if terminatingRegexp.MatchString(line) {
			closed = true
			break
		}
	}

	if !closed {
		errMsg := fmt.Sprintf("Multiline string did not terminate\nStarted at '" + lines[0])
		return 0, stmt, errors.New(errMsg)
	}

	j := strings.Join(joinedLines, "\n")

	stmt = ast.SimpleString{Value: j[1 : len(j)-1]}

	return index, stmt, nil
}
