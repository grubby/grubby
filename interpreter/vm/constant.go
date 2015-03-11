package vm

import (
	"fmt"

	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretConstantInContext(
	vm *vm,
	constantNode ast.Constant,
	context Value,
) (Value, error) {

	var target Module
	if vm.currentModuleName == "" {
		target = vm.CurrentClasses["Object"]
	} else {
		target = vm.CurrentClasses[vm.currentModuleName]
		if target == nil {
			target = vm.CurrentModules[vm.currentModuleName]
		}

		if target == nil {
			panic(fmt.Sprintf("unexpected nil target when looking up module %s to find constant %s", vm.currentModuleName, constantNode.Name))
		}
	}

	constant, err := target.Constant(constantNode.Name)
	if err != nil {
		maybeClass, ok := vm.CurrentClasses[constantNode.Name]
		if ok {
			return maybeClass, nil
		}

		maybeModule, ok := vm.CurrentModules[constantNode.Name]
		if ok {
			return maybeModule, nil
		}

		return nil, err
	}

	return constant, nil
}
