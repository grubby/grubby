package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretTernaryInContext(
	vm *vm,
	ternary ast.Ternary,
	context Value,
) (Value, error) {

	value, err := vm.executeWithContext(context, ternary.Condition)
	if err != nil {
		return nil, err
	} else {
		if value.IsTruthy() {
			return vm.executeWithContext(context, ternary.True)
		} else {
			return vm.executeWithContext(context, ternary.False)
		}
	}
}
