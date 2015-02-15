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

	var (
		value Value
		ok    bool
	)

	name := class.FullName()
	value, ok = vm.CurrentClasses[name]
	if !ok {
		value, ok = vm.CurrentModules[name]
		if !ok {
			return nil, NewNameError(name, context.String(), context.Class().String(), vm.stack.String())
		}
	}

	return value, nil
}
