package builtins

import "fmt"

type FloatValue struct {
	value float64
	valueStub
}

func NewFloat(val float64) Value {
	f := &FloatValue{value: val}
	f.initialize()
	return f
}

func (floatValue *FloatValue) ValueAsFloat() float64 {
	return floatValue.value
}

func (floatValue *FloatValue) String() string {
	return fmt.Sprintf("%d", floatValue.value)
}
