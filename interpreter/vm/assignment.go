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
		ivar := assignment.LHS.(ast.InstanceVariable)
		context.SetInstanceVariable(ivar.Name, returnValue)
	case ast.ClassVariable:
		classVar := assignment.LHS.(ast.ClassVariable)
		context.SetClassVariable(classVar.Name, returnValue)
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

func interpretConditionalAssignmentInContext(
	vm *vm,
	conditionalAssignment ast.ConditionalAssignment,
	context Value,
) (Value, error) {

	returnValue, err := vm.executeWithContext(context, conditionalAssignment.RHS)
	if err != nil {
		return nil, err
	}

	switch conditionalAssignment.LHS.(type) {
	case ast.BareReference:
		ref := conditionalAssignment.LHS.(ast.BareReference)
		if vm.ObjectSpace[ref.Name] != nil && vm.ObjectSpace[ref.Name].IsTruthy() {
			return vm.ObjectSpace[ref.Name], nil
		}

		vm.ObjectSpace[ref.Name] = returnValue
	case ast.GlobalVariable:
		globalVar := conditionalAssignment.LHS.(ast.GlobalVariable)
		if vm.CurrentGlobals[globalVar.Name] != nil && vm.CurrentGlobals[globalVar.Name].IsTruthy() {
			return vm.CurrentGlobals[globalVar.Name], nil
		}

		vm.CurrentGlobals[globalVar.Name] = returnValue
	case ast.InstanceVariable:
		ivar := conditionalAssignment.LHS.(ast.InstanceVariable)
		existingIvar := context.GetInstanceVariable(ivar.Name)
		if existingIvar != nil && existingIvar.IsTruthy() {
			return existingIvar, nil
		}

		context.SetInstanceVariable(ivar.Name, returnValue)
	case ast.ClassVariable:
		classVar := conditionalAssignment.LHS.(ast.ClassVariable)
		existingClassVar := context.GetClassVariable(classVar.Name)
		if existingClassVar != nil && existingClassVar.IsTruthy() {
			return existingClassVar, nil
		}

		context.SetClassVariable(classVar.Name, returnValue)
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
				panic(fmt.Sprintf("unexpected nil target when looking up module %s to set constant %s", vm.currentModuleName, conditionalAssignment.LHS.(ast.Constant).Name))
			}
		}

		name := conditionalAssignment.LHS.(ast.Constant).Name
		constant, err := target.Constant(name)
		if err == nil && constant.IsTruthy() {
			return constant, nil
		}

		target.SetConstant(name, returnValue)
	case ast.Class:
		asClass := conditionalAssignment.LHS.(ast.Class)
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

			constant, err := target.(Module).Constant(asClass.Name)
			if err == nil && constant.IsTruthy() {
				return constant, nil
			}
			target.(Module).SetConstant(asClass.Name, returnValue)
		}
	default:
		panic(fmt.Sprintf("unimplemented conditionalAssignment failure: %#v", conditionalAssignment.LHS))
	}

	return returnValue, nil
}
