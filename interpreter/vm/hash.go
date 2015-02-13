package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretHashInContext(
	vm *vm,
	hashNode ast.Hash,
	context Value,
) (Value, error) {

	hashValue, _ := vm.CurrentClasses["Hash"].New(vm, vm)
	hash := hashValue.(*Hash)
	for _, keyPair := range hashNode.Pairs {
		key, err := vm.executeWithContext(context, keyPair.Key)
		if err != nil {
			return nil, err
		}

		val, err := vm.executeWithContext(context, keyPair.Value)
		if err != nil {
			return nil, err
		}

		hash.Add(key, val)
	}

	return hash, nil
}
