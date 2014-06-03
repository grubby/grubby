package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/grubby/grubby/ast"
)

var integerRegexp = regexp.MustCompile("[0-9]+")

func Parse(input string) ast.Block {
	block := ast.Block{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if integerRegexp.MatchString(line) {
			val, _ := strconv.Atoi(line)
			block.Statement = ast.ConstantInt{Value: val}
		}
	}

	return block
}
