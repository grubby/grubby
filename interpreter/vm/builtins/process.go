package builtins

import (
	"fmt"
	"os"
)

func NewProcessModule(provider Provider) Module {
	module := NewGenericModule("Process", provider)

	module.AddMethod(NewNativeMethod("pid", provider, func(self Value, block Block, args ...Value) (Value, error) {
		return NewString(fmt.Sprintf("%d", os.Getpid()), provider), nil
	}))

	return module
}
