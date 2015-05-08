package vm

import (
	"errors"
	"fmt"
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

	theClass, ok := vm.CurrentClasses[fullClassName]
	if !ok {
		theClass = NewUserDefinedClass(classNode.Name, classNode.SuperClass.FullName(), vm)
		vm.CurrentClasses[fullClassName] = theClass
	} else {
		if classNode.SuperClass.FullName() != "" {
			return nil, errors.New(fmt.Sprintf("TypeError: superclass mismatch for class %s", classNode.Name))
		}
	}

	vm.currentModuleName = fullClassName
	_, err := vm.executeWithContext(theClass, classNode.Body...)
	if err != nil {
		return nil, err
	}

	return theClass, nil
}
