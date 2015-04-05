package vm

import (
	"github.com/grubby/grubby/ast"
	"github.com/tjarratt/gomads"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretBareReferenceInContext(
	vm *vm,
	ref ast.BareReference,
	context Value,
) (Value, error) {

	var name string = ref.Name
	var returnErr error

	maybe := gomads.Maybe(func() interface{} {
		m, err := vm.localVariableStack.Retrieve(name)
		if err == nil {
			return m
		} else {
			return nil
		}
	}).OrSome(gomads.Maybe(func() interface{} {
		m, ok := vm.ObjectSpace[name]
		if ok {
			return m
		} else {
			return nil
		}
	})).OrSome(gomads.Maybe(func() interface{} {
		m, ok := vm.CurrentClasses[name]
		if ok {
			return m
		} else {
			return nil
		}
	})).OrSome(gomads.Maybe(func() interface{} {
		m, ok := vm.CurrentModules[name]
		if ok {
			return m
		} else {
			return nil
		}
	})).OrSome(gomads.Maybe(func() interface{} {
		maybeMethod := context.Method(name)
		if maybeMethod == nil {
			return nil
		}

		value, err := maybeMethod.Execute(context, nil)
		if err != nil {
			returnErr = err
			return nil
		} else {
			return value
		}
	}))

	if returnErr != nil {
		return nil, returnErr
	}

	value, ok := maybe.Value().(Value)
	if ok {
		return value, nil
	} else {
		return nil, NewNameError(
			name,
			context.String(),
			context.Class().String(),
			vm.stack.String(),
		)
	}
}
