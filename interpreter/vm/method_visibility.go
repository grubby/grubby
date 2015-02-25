package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

type methodVisibility int

const (
	public methodVisibility = iota
	private
	protected
)

func interpretMethodsWithVisibility(vm *vm, methods []ast.Symbol, mode methodVisibility, context Value) (Value, error) {
	if len(methods) == 0 {
		vm.methodDeclarationMode = mode
		return nil, nil
	} else {

		contextAsModule := context.(Module)

		for _, methodName := range methods {
			method, err := contextAsModule.InstanceMethod(methodName.Name)
			if err != nil {
				return nil, err
			}

			contextAsModule.RemoveInstanceMethod(method)
			contextAsModule.AddPrivateInstanceMethod(method)
		}
	}

	return nil, nil
}
