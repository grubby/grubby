package vm

import (
	"strings"

	"github.com/grubby/grubby/ast"
	"github.com/tjarratt/gomads"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretConstantInContext(
	vm *vm,
	constantNode ast.Constant,
	context Value,
) (Value, error) {

	maybeTarget := gomads.Maybe(func() interface{} {
		if vm.currentModuleName == "" {
			return vm.CurrentClasses["Object"]
		} else {
			return nil
		}
	}).OrSome(gomads.Maybe(func() interface{} {
		return vm.CurrentClasses[vm.currentModuleName]
	})).OrSome(gomads.Maybe(func() interface{} {
		return vm.CurrentModules[vm.currentModuleName]
	}))

	target, ok := maybeTarget.Value().(Module)

	maybeConstant := gomads.Maybe(func() interface{} {
		constant, err := target.Constant(constantNode.Name)
		if err == nil {
			return constant
		} else {
			return nil
		}
	}).OrSome(gomads.Maybe(func() interface{} {
		return vm.CurrentClasses[constantNode.Name]
	})).OrSome(gomads.Maybe(func() interface{} {
		return vm.CurrentModules[constantNode.Name]
	})).OrSome(gomads.Maybe(func() interface{} {
		if vm.currentModuleName == "" {
			return nil
		}

		parts := strings.Split(vm.currentModuleName, "::")
		count := len(parts) - 1
		for index, _ := range parts {
			namespace := append(parts[0:(count-index)], constantNode.Name)
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
	}))

	constant, ok := maybeConstant.Value().(Value)
	if ok {
		return constant, nil
	} else {
		return nil, NewNameError(
			constantNode.Name,
			context.String(),
			context.Class().String(),
			vm.stack.String(),
		)
	}

}
