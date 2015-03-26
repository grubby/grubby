package vm

import (
	"strings"

	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretClassDeclarationInContext(
	vm *vm,
	classNode ast.ClassDecl,
	context Value,
) (Value, error) {

	originalName := vm.currentModuleName

	defer func() {
		vm.currentModuleName = originalName
	}()

	var fullClassName string
	if vm.currentModuleName != "" {
		fullClassName = strings.Join([]string{vm.currentModuleName, classNode.FullName()}, "::")
	} else {
		fullClassName = classNode.FullName()
	}

	theClass := NewUserDefinedClass(classNode.Name, vm, vm)
	vm.CurrentClasses[fullClassName] = theClass

	vm.currentModuleName = fullClassName
	_, err := vm.executeWithContext(theClass, classNode.Body...)
	if err != nil {
		return nil, err
	}

	return theClass, nil
}
