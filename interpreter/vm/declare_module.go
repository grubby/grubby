package vm

import (
	"strings"

	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretModuleDeclarationInContext(
	vm *vm,
	moduleNode ast.ModuleDecl,
	context Value,
) (Value, error) {

	originalName := vm.currentModuleName

	defer func() {
		vm.currentModuleName = originalName
		vm.methodDeclarationMode = public
	}()

	var fullModuleName string
	if vm.currentModuleName != "" {
		fullModuleName = strings.Join([]string{vm.currentModuleName, moduleNode.FullName()}, "::")
	} else {
		fullModuleName = moduleNode.FullName()
	}

	theModule := NewModule(moduleNode.Name, vm, vm)
	vm.CurrentModules[fullModuleName] = theModule

	vm.currentModuleName = fullModuleName
	_, err := vm.executeWithContext(theModule, moduleNode.Body...)
	if err != nil {
		return nil, err
	}

	return theModule, nil
}
