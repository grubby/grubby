package vm

import (
	"strings"

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
	})).OrSome(gomads.Maybe(func() interface{} {
		if vm.currentModuleName == "" {
			return nil
		}

		parts := strings.Split(vm.currentModuleName, "::")
		count := len(parts) - 1
		for index, _ := range parts {
			namespace := append(parts[0:(count-index)], name)
			nameToLookup := strings.Join(namespace, "::")

			maybe := gomads.Maybe(func() interface{} {
				return vm.CurrentClasses[nameToLookup]
			}).OrSome(gomads.Maybe(func() interface{} {
				return vm.CurrentModules[nameToLookup]
			}))

			something, ok := maybe.Value().(Value)
			if ok {
				return something
			}
		}

		return nil
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
