package builtins

import "errors"

type symbolClass struct {
	valueStub
	classStub
}

func NewSymbolClass(provider ClassProvider) Class {
	s := &symbolClass{}
	s.initialize()
	s.class = provider.ClassWithName("Class")
	s.superClass = provider.ClassWithName("Object")

	return s
}

func (c *symbolClass) String() string {
	return "Symbol"
}

func (c *symbolClass) Name() string {
	return "Symbol"
}

func (c *symbolClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Symbol:Class")
}

type symbolValue struct {
	value string
	valueStub
}

func NewSymbol(val string, provider ClassProvider) Value {
	s := &symbolValue{value: val}
	s.class = provider.ClassWithName("Symbol")
	s.initialize()
	return s
}

func (symbolValue *symbolValue) String() string {
	return symbolValue.value
}

func (symbolValue *symbolValue) Name() string {
	return symbolValue.value
}
