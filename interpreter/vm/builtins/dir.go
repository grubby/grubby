package builtins

import "os"

func NewDirClass(provider Provider) Class {
	class := NewGenericClass("Dir", "Object", provider)

	class.AddMethod(NewNativeMethod("pwd", provider, func(self Value, block Block, args ...Value) (Value, error) {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		return NewString(dir, provider), nil
	}))

	return class
}
