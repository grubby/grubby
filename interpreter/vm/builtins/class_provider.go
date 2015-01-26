package builtins

import (
	"github.com/grubby/grubby/ast"
)

type ArgEvaluator interface {
	EvaluateArgInContext(ast.Node, Value) (Value, error)
}

type ClassProvider interface {
	ClassWithName(string) Class
}

type SingletonProvider interface {
	SingletonWithName(string) Value
	NewSingletonWithName(string, Value)
}
