package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretClassInContext(
	vm *vm,
	class ast.Class,
	context Value,
) (Value, error) {

	className := class.FullName()
	value, ok := vm.CurrentClasses[className]
	if !ok {
		return nil, NewNameError(className, context.String(), context.Class().String(), vm.stack.String())
	}

	return value, nil
}
