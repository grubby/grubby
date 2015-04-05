package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretCallExpressionInContext(
	vm *vm,
	callExpr ast.CallExpression,
	context Value,
) (Value, error) {
	var (
		err         error
		target      Value
		method      Method
		returnValue Value
	)

	if callExpr.Target != nil {
		target, err = vm.executeWithContext(context, callExpr.Target)
		if err != nil {
			return nil, err
		}
	} else {
		target = context
	}

	if target == nil {
		nilValue := vm.singletons["nil"]
		return nil, NewNoMethodError(callExpr.Func.Name, nilValue.String(), nilValue.Class().String(), vm.stack.String())
	}

	method = target.Method(callExpr.Func.Name)
	if method == nil {
		return nil, NewNoMethodError(callExpr.Func.Name, target.PrettyPrint(), target.Class().String(), vm.CurrentStack())
	}

	args := []Value{}
	for _, astArgument := range callExpr.Args {
		arg, err := vm.executeWithContext(context, astArgument)
		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	vm.stack.Unshift(method.Name(), vm.currentFilename, callExpr.LineNumber())
	didShift := false
	defer func() {
		if didShift == false {
			vm.stack.Shift()
		}
	}()

	var block Block
	if callExpr.OptionalBlock.Provided() {
		blockValue, err := vm.executeWithContext(context, callExpr.OptionalBlock)

		if err != nil {
			return nil, err
		}

		block = blockValue.(Block)
	}

	returnValue, err = method.Execute(target, block, args...)
	vm.stack.Shift()
	didShift = true

	return returnValue, err
}
