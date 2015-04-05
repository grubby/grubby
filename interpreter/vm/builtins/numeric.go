package builtins

import "errors"

type numericClass struct {
	valueStub
	classStub
}

func NewNumericClass(provider Provider) Class {
	class := &numericClass{}
	class.initialize()
	class.setStringer(class.String)
	class.class = provider.ClassProvider().ClassWithName("Class")
	class.superClass = provider.ClassProvider().ClassWithName("Object")

	return class
}

func (c *numericClass) String() string {
	return "Numeric"
}

func (c *numericClass) Name() string {
	return "Numeric"
}

func (c *numericClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Numeric:Class")
}
