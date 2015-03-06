package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretIfStatementInContext(
	vm *vm,
	ifBlock ast.IfBlock,
	context Value,
) (Value, error) {
	truthy := false
	switch ifBlock.Condition.(type) {
	case ast.Boolean:
		truthy = ifBlock.Condition.(ast.Boolean).Value
	default:
		value, err := vm.executeWithContext(context, ifBlock.Condition)
		if err != nil {
			return nil, err
		}

		truthy = value.IsTruthy()
	}

	if truthy {
		return vm.executeWithContext(context, ifBlock.Body...)
	} else {
		return vm.executeWithContext(context, ifBlock.Else...)
	}
}
