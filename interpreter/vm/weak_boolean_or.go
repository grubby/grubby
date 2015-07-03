package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretWeakBooleanOr(vm *vm, weakOr ast.WeakLogicalOr, context Value) (Value, error) {
	lhs, err := vm.executeWithContext(context, weakOr.LHS)
	if err != nil {
		return nil, err
	}

	if lhs.IsTruthy() {
		return lhs, nil
	}

	return vm.executeWithContext(context, weakOr.RHS)
}
