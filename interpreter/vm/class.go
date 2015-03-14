package vm

import (
	"github.com/grubby/grubby/ast"
	"github.com/tjarratt/gomads"

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
	maybeTheValue := gomads.Maybe(func() interface{} {
		return vm.CurrentClasses[name]
	}).OrSome(gomads.Maybe(func() interface{} {
		return vm.CurrentModules[name]
	})).OrSome(gomads.Maybe(func() interface{} {
		module, ok := vm.CurrentClasses[class.Namespace]
		if !ok {
			return nil
		}

		value, err := module.Constant(class.Name)
		if err != nil {
			return nil
		}

		return value
	})).OrSome(gomads.Maybe(func() interface{} {
		module, ok := vm.CurrentModules[class.Namespace]
		if !ok {
			return nil
		}

		value, err := module.Constant(class.Name)
		if err != nil {
			return nil
		}

		return value
	})).OrSome(
		NewNameError(
			name,
			context.String(),
			context.Class().String(),
			vm.stack.String(),
		),
	)

	value, ok = maybeTheValue.Value().(Value)

	if ok {
		return value, nil
	} else {
		return nil, value.(error)
	}
}
