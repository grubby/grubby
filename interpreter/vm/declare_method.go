package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretMethodDeclarationInContext(
	vm *vm,
	funcNode ast.FuncDecl,
	context Value,
) (Value, error) {

	var returnValue Value

	method := NewRubyMethod(
		funcNode.MethodName(),
		funcNode.MethodArgs(),
		funcNode.Body,
		vm,
		vm,
		func(self Value, method *RubyMethod) (Value, error) {
			vm.localVariableStack.Unshift()
			defer vm.localVariableStack.Shift()

			for _, arg := range method.Args() {
				vm.localVariableStack.Store(arg.Name, arg.Value)
			}

			return vm.executeWithContext(self, method.Body()...)
		})
	returnValue = method

	if context == vm.ObjectSpace["main"] && funcNode.Target == nil {
		method = NewPrivateRubyMethod(
			funcNode.MethodName(),
			funcNode.MethodArgs(),
			funcNode.Body,
			vm,
			vm,
			func(self Value, method *RubyMethod) (Value, error) {
				vm.localVariableStack.Unshift()
				defer vm.localVariableStack.Shift()

				for _, arg := range method.Args() {
					vm.localVariableStack.Store(arg.Name, arg.Value)
				}

				return vm.executeWithContext(self, method.Body()...)
			})
		returnValue = method
		vm.CurrentModules["Kernel"].AddMethod(method)
	} else if vm.inEigenclassBlock {
		context.AddMethod(method)
	} else {
		switch funcNode.Target.(type) {
		case ast.Self:
			context.AddMethod(method)
		case nil:
			switch context.(Module).ActiveVisibility() {
			case Public:
				method.SetVisibility(Public)
			case Private:
				method.SetVisibility(Private)
			case Protected:
				method.SetVisibility(Protected)
			}

			context.(Module).AddInstanceMethod(method)
		default:
			value, err := vm.executeWithContext(context, funcNode.Target)
			if err != nil {
				return nil, err
			}

			value.AddMethod(method)
		}
	}

	return returnValue, nil
}
