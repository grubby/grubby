package builtins

import "errors"

type standardErrorClass struct {
	valueStub
	classStub
}

func NewStandardErrorClass(provider Provider) Class {
	n := &standardErrorClass{}
	n.initialize()
	n.setStringer(n.String)
	n.class = provider.ClassProvider().ClassWithName("Class")
	n.superClass = provider.ClassProvider().ClassWithName("Exception")
	return n
}

func (n *standardErrorClass) String() string {
	return "StandardError"
}

func (n *standardErrorClass) Name() string {
	return "StandardError"
}

type standardErrorInstance struct {
	valueStub
}

func (class *standardErrorClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, errors.New("unimplemented")
}

func (n *standardErrorInstance) String() string {
	return ""
}

func (n *standardErrorInstance) IsTruthy() bool {
	return true
}
