package vm

import (
	"fmt"
	"strings"

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
	case ast.Constant:
		var target Module
		if vm.currentModuleName == "" {
			target = vm.CurrentClasses["Object"]
		} else {
			target = vm.CurrentClasses[vm.currentModuleName]
			if target == nil {
				target = vm.CurrentModules[vm.currentModuleName]
			}

			if target == nil {
				panic(fmt.Sprintf("unexpected nil target when looking up module %s to set constant %s", vm.currentModuleName, assignment.LHS.(ast.Constant).Name))
			}
		}

		target.SetConstant(assignment.LHS.(ast.Constant).Name, returnValue)
	case ast.Class:
		asClass := assignment.LHS.(ast.Class)
		if asClass.Namespace != "" {
			var target Value
			namespace := asClass.Namespace
			for _, nameToLookup := range strings.Split(namespace, "::") {
				class, err := vm.GetClass(nameToLookup)
				if err != nil {
					module, err := vm.GetModule(nameToLookup)
					if err != nil {
						return nil, err
					}
					target = module
				} else {
					target = class
				}
			}

			target.(Module).SetConstant(asClass.Name, returnValue)
		}
	default:
		panic(fmt.Sprintf("unimplemented assignment failure: %#v", assignment.LHS))
	}

	return returnValue, nil
}
