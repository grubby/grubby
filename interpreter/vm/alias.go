package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretAliasInContext(
	vm *vm,
	aliasNode ast.Alias,
	context Value,
) (Value, error) {
	// FIXME: assumes that the context will be a module, but could also be a class
	contextModule := context.(Module)

	m, err := contextModule.InstanceMethod(aliasNode.From.Name)
	if err != nil {
		return nil, NewNameError(aliasNode.From.Name, contextModule.String(), contextModule.String(), vm.stack.String())
	}

	contextModule.AddInstanceMethod(NewNativeMethod(aliasNode.To.Name, vm, func(self Value, block Block, args ...Value) (Value, error) {
		return m.Execute(self, block, args...)
	}))

	return nil, nil
}
