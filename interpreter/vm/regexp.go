package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretRegexpInContext(vm *vm, regexpNode ast.Regex, context Value) (Value, error) {
	return NewRegexp(vm, regexpNode.Value), nil
}
