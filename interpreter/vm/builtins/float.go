package builtins

import (
	"errors"
	"fmt"
)

type floatClass struct {
	valueStub
	classStub
}

func NewFloatClass(provider ClassProvider) Class {
	class := &floatClass{}
	class.initialize()
	class.class = provider.ClassWithName("Class")
	class.superClass = provider.ClassWithName("Numeric")

	return class
}

func (c *floatClass) String() string {
	return "Float"
}

func (c *floatClass) Name() string {
	return "Float"
}

func (c *floatClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Float:Class")
}

type floatValue struct {
	value float64
	valueStub
}

func NewFloat(val float64, provider ClassProvider) Value {
	f := &floatValue{value: val}
	f.class = provider.ClassWithName("Float")
	f.initialize()
	return f
}

func (floatValue *floatValue) ValueAsFloat() float64 {
	return floatValue.value
}

func (floatValue *floatValue) String() string {
	return fmt.Sprintf("%d", floatValue.value)
}
