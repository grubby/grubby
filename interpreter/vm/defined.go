package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretDefinedKeyword(vm *vm, defined ast.Defined, context Value) (Value, error) {
	_, err := vm.executeWithContext(context, defined.Node)
	if err == nil {
		return NewString("garbage", vm), nil
	} else {
		return vm.SingletonWithName("nil"), nil
	}
}
