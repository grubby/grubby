package builtins

import "errors"

type exceptionClass struct {
	valueStub
	classStub
}

func NewExceptionClass(provider Provider) Class {
	n := &exceptionClass{}
	n.initialize()
	n.setStringer(n.String)
	n.class = provider.ClassProvider().ClassWithName("Class")
	n.superClass = provider.ClassProvider().ClassWithName("Object")
	return n
}

func (n *exceptionClass) String() string {
	return "Exception"
}

func (n *exceptionClass) Name() string {
	return "Exception"
}

type exceptionInstance struct {
	valueStub
}

func (class *exceptionClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, errors.New("unimplemented")
}

func (n *exceptionInstance) String() string {
	return ""
}

func (n *exceptionInstance) IsTruthy() bool {
	return true
}
