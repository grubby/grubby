package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretClassDeclarationInContext(
	vm *vm,
	classNode ast.ClassDecl,
	context Value,
) (Value, error) {

	defer func() {
		vm.methodDeclarationMode = public
	}()

	theClass := NewUserDefinedClass(classNode.Name, vm, vm)
	vm.CurrentClasses[classNode.FullName()] = theClass

	_, err := vm.executeWithContext(theClass, classNode.Body...)
	if err != nil {
		return nil, err
	}

	return theClass, nil
}
