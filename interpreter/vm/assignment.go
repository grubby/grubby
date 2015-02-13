package vm

import (
	"fmt"

	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretAssignmentInContext(
	vm *vm,
	assignment ast.Assignment,
	context Value,
) (Value, error) {

	returnValue, err := vm.executeWithContext(context, assignment.RHS)
	if err != nil {
		return nil, err
	}

	switch assignment.LHS.(type) {
	case ast.BareReference:
		ref := assignment.LHS.(ast.BareReference)
		vm.ObjectSpace[ref.Name] = returnValue
	case ast.GlobalVariable:
		globalVar := assignment.LHS.(ast.GlobalVariable)
		vm.CurrentGlobals[globalVar.Name] = returnValue
	case ast.InstanceVariable:
		iVar := assignment.LHS.(ast.InstanceVariable)
		context.SetInstanceVariable(iVar.Name, returnValue)
	default:
		panic(fmt.Sprintf("unimplemented assignment failure: %#v", assignment.LHS))
	}

	return returnValue, nil
}
