package builtins

import (
	"errors"
	"fmt"
)

type fixnumClass struct {
	valueStub
	classStub
}

func NewFixnumClass(provider ClassProvider, singletonProvider SingletonProvider) Class {
	class := &fixnumClass{}
	class.initialize()
	class.setStringer(class.String)
	class.class = provider.ClassWithName("Class")
	class.superClass = provider.ClassWithName("Integer")

	class.AddMethod(NewNativeMethod("even?", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		asFixnum := self.(*fixnumInstance)

		if asFixnum.value%2 == 0 {
			return singletonProvider.SingletonWithName("true"), nil
		} else {
			return singletonProvider.SingletonWithName("false"), nil
		}
	}))

	class.AddMethod(NewNativeMethod("+", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		arg, ok := args[0].(*fixnumInstance)
		if !ok {
			return nil, errors.New(fmt.Sprintf("TypeError: %s can't be coerced into Fixnum", args[0].Class().String()))
		}
		asFixnum := self.(*fixnumInstance)
		return NewFixnum(asFixnum.value+arg.value, provider, singletonProvider), nil
	}))

	return class
}

func (c *fixnumClass) String() string {
	return "Fixnum"
}

func (c *fixnumClass) Name() string {
	return "Fixnum"
}

func (c *fixnumClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Fixnum:Class")
}

type fixnumInstance struct {
	value int64
	valueStub
}

func NewFixnum(val int64, provider ClassProvider, singletonProvider SingletonProvider) Value {
	name := fmt.Sprintf("%d", val)
	singleton := singletonProvider.SingletonWithName(name)
	if singleton == nil {
		i := &fixnumInstance{value: val}
		i.class = provider.ClassWithName("Fixnum")
		i.initialize()
		i.setStringer(i.String)

		singletonProvider.NewSingletonWithName(name, i)
		return i
	} else {
		return singleton
	}
}

func (fixnumInstance *fixnumInstance) Value() int64 {
	return fixnumInstance.value
}

func (fixnumInstance *fixnumInstance) String() string {
	return fmt.Sprintf("%d", fixnumInstance.value)
}
