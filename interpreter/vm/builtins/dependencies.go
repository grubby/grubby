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

	SymbolWithName(string) Value
	AddSymbol(Value)
}

type StackProvider interface {
	CurrentStack() string
}

type Provider interface {
	ArgEvaluator() ArgEvaluator
	ClassProvider() ClassProvider
	SingletonProvider() SingletonProvider
	StackProvider() StackProvider
}
