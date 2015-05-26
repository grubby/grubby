package builtins

import (
	"os"
)

func NewENVConstant(provider Provider) Value {
	object, _ := provider.ClassProvider().ClassWithName("Object").New(provider)
	object.AddMethod(NewNativeMethod("[]", provider, func(self Value, block Block, args ...Value) (Value, error) {
		key := args[0].(*StringValue).value
		return NewString(os.Getenv(key), provider), nil
	}))

	return object
}
