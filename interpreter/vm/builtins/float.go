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
	class.setStringer(class.String)
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

type FloatValue struct {
	value float64
	valueStub
}

func NewFloat(val float64, provider ClassProvider) Value {
	f := &FloatValue{value: val}
	f.class = provider.ClassWithName("Float")
	f.initialize()
	f.setStringer(f.String)
	return f
}

func (FloatValue *FloatValue) ValueAsFloat() float64 {
	return FloatValue.value
}

func (FloatValue *FloatValue) String() string {
	return fmt.Sprintf("%d", FloatValue.value)
}
