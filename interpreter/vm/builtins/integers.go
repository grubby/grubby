package builtins

import "fmt"

type intValue struct {
	value int
	valueStub
}

func NewInt(val int) Value {
	return &intValue{value: val}
}

func (intValue *intValue) Value() int {
	return intValue.value
}

func (intValue *intValue) String() string {
	return fmt.Sprintf("%d", intValue.value)
}
