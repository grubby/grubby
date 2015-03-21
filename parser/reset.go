package parser

import "github.com/grubby/grubby/ast"

func Reset() {
	Statements = []ast.Node{}
	DebugStatements = []string{}
}
