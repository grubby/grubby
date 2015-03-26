package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretBareReferenceInContext(
	vm *vm,
	ref ast.BareReference,
	context Value,
) (Value, error) {
	var (
		returnValue Value
		returnErr   error
	)

	name := ref.Name
	maybe, err := vm.localVariableStack.Retrieve(name)
	if err == nil {
		returnValue = maybe
	} else {
		maybe, ok := vm.ObjectSpace[name]
		if ok {
			returnValue = maybe
		} else {
			maybe, ok := vm.CurrentClasses[name]
			if ok {
				returnValue = maybe
			} else {
				maybe, ok := vm.CurrentModules[name]
				if ok {
					returnValue = maybe
				} else {
					maybeMethod, err := context.Method(name)
					if err == nil {
						returnValue, returnErr = maybeMethod.Execute(context, nil)
					} else {
						returnValue = nil
						returnErr = NewNameError(name, context.String(), context.Class().String(), vm.stack.String())
					}
				}
			}
		}
	}

	return returnValue, returnErr
}
