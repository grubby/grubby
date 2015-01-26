package builtins

import "errors"

type numericClass struct {
	valueStub
	classStub
}

func NewNumericClass(provider ClassProvider) Class {
	class := &numericClass{}
	class.initialize()
	class.class = provider.ClassWithName("Class")
	class.superClass = provider.ClassWithName("Object")

	return class
}

func (c *numericClass) String() string {
	return "Numeric"
}

func (c *numericClass) Name() string {
	return "Numeric"
}

func (c *numericClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Numeric:Class")
}
