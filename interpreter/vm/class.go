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
		err   error
	)

	name := class.FullName()
	value, ok = vm.CurrentClasses[name]
	if !ok {
		value, ok = vm.CurrentModules[name]
		if !ok {
			module, ok := vm.CurrentClasses[class.Namespace]
			if !ok {
				module, ok := vm.CurrentModules[class.Namespace]
				if !ok {
					return nil, NewNameError(name, context.String(), context.Class().String(), vm.stack.String())
				} else {
					value, err = module.Constant(class.Name)
					if err != nil {
						return nil, NewNameError(name, context.String(), context.Class().String(), vm.stack.String())
					}
				}
			} else {
				value, err = module.Constant(class.Name)
				if err != nil {
					return nil, NewNameError(name, context.String(), context.Class().String(), vm.stack.String())
				}
			}
		}
	}

	return value, nil
}
