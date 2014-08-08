package builtins

import "fmt"

type intValue struct {
	value           int
	methods         []Method
	private_methods []Method
}

func NewInt(val int) Value {
	return &intValue{value: val}
}

func (intValue *intValue) Methods() []Method {
	return intValue.methods
}

func (intValue *intValue) PrivateMethods() []Method {
	return intValue.private_methods
}

func (intValue *intValue) AddMethod(m Method) {
	intValue.methods = append(intValue.methods, m)
}

func (intValue *intValue) AddPrivateMethod(m Method) {
	intValue.private_methods = append(intValue.private_methods, m)
}

func (intValue *intValue) String() string {
	return fmt.Sprintf("%d", intValue.value)
}

func (intValue *intValue) Class() Class {
	return nil
}
