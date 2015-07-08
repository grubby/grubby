package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretSwitchStatement(vm *vm, switchStmt ast.SwitchStatement, context Value) (Value, error) {

	conditionToMatch, err := vm.executeWithContext(context, switchStmt.Condition)
	if err != nil {
		return nil, err
	}

	for _, caseStmt := range switchStmt.Cases {
		for _, conditionNode := range caseStmt.Conditions {
			condition, err := vm.executeWithContext(context, conditionNode)
			if err != nil {
				return nil, err
			}

			method := condition.Method("===")
			if method == nil {
				return nil, NewNoMethodError("===", condition.String(), condition.Class().String(), vm.stack.String())
			}

			matches, err := method.Execute(condition, nil, conditionToMatch)
			if matches.IsTruthy() {
				return vm.executeWithContext(context, caseStmt.Body...)
			}
		}
	}

	if switchStmt.Else != nil {
		return vm.executeWithContext(context, switchStmt.Else...)
	}

	return vm.SingletonWithName("nil"), nil
}
