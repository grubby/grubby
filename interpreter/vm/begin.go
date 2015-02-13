package vm

import (
	"github.com/grubby/grubby/ast"

	. "github.com/grubby/grubby/interpreter/vm/builtins"
)

func interpretBeginInContext(
	vm *vm,
	begin ast.Begin,
	context Value,
) (Value, error) {
	_, err := vm.executeWithContext(context, begin.Body...)

	if err != nil {
		matchingRescue := false

		rubyErr, ok := err.(Value)
		if !ok {
			panic(context)
			return nil, err
		}

		for _, rescue := range begin.Rescue {
			if matchingRescue {
				break
			}

			r := rescue.(ast.Rescue)
			for _, exceptionClass := range r.Exception.Classes {
				if exceptionClass.Name == rubyErr.String() {
					_, err = vm.executeWithContext(context, r.Body...)
					if err == nil {
						matchingRescue = true
						break
					}
				}
			}
		}
	}

	return nil, err
}
