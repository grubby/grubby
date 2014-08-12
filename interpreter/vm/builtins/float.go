package builtins

import "fmt"

type FloatValue struct {
	value float64
	valueStub
}

func NewFloat(val float64) Value {
	return &FloatValue{value: val}
}

func (floatValue *FloatValue) ValueAsFloat() float64 {
	return floatValue.value
}

func (floatValue *FloatValue) String() string {
	return fmt.Sprintf("%d", floatValue.value)
}
