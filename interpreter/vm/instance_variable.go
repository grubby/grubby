package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretInstanceVariableInContext(
	vm *vm,
	ref ast.InstanceVariable,
	context Value,
) (Value, error) {

	return context.GetInstanceVariable(ref.Name), nil
}
