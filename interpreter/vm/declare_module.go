package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretModuleDeclarationInContext(
	vm *vm,
	moduleNode ast.ModuleDecl,
	context Value,
) (Value, error) {

	defer func() {
		vm.methodDeclarationMode = public
	}()

	theModule := NewModule(moduleNode.Name, vm, vm)
	vm.CurrentModules[moduleNode.Name] = theModule

	_, err := vm.executeWithContext(theModule, moduleNode.Body...)
	if err != nil {
		return nil, err
	}

	return theModule, nil
}
