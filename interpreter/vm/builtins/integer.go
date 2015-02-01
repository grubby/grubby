package builtins

import "errors"

type integerClass struct {
	valueStub
	classStub
}

func NewIntegerClass(provider ClassProvider) Class {
	class := &integerClass{}
	class.initialize()
	class.setStringer(class.String)
	class.class = provider.ClassWithName("Class")
	class.superClass = provider.ClassWithName("Numeric")

	return class
}

func (c *integerClass) String() string {
	return "Integer"
}

func (c *integerClass) Name() string {
	return "Integer"
}

func (c *integerClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Integer:Class")
}
