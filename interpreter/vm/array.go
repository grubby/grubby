package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretArrayInContext(
	vm *vm,
	arrayNode ast.Array,
	context Value,
) (Value, error) {

	arrayValue, _ := vm.CurrentClasses["Array"].New(vm, vm)
	array := arrayValue.(*Array)

	for _, node := range arrayNode.Nodes {
		value, err := vm.executeWithContext(context, node)
		if err != nil {
			return nil, err
		}

		array.Append(value)
	}

	return array, nil
}
