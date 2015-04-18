package builtins

import "errors"

type argumentErrorClass struct {
	valueStub
	classStub
}

func NewArgumentErrorClass(provider Provider) Class {
	n := &argumentErrorClass{}
	n.initialize()
	n.setStringer(n.String)
	n.class = provider.ClassProvider().ClassWithName("Class")
	n.superClass = provider.ClassProvider().ClassWithName("Exception")
	return n
}

func (n *argumentErrorClass) String() string {
	return "ArgumentError"
}

func (n *argumentErrorClass) Name() string {
	return "ArgumentError"
}

type argumentErrorInstance struct {
	valueStub
}

func (class *argumentErrorClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, errors.New("unimplemented")
}

func (n *argumentErrorInstance) String() string {
	return ""
}

func (n *argumentErrorInstance) IsTruthy() bool {
	return true
}
