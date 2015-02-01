package builtins

import "errors"

type symbolClass struct {
	valueStub
	classStub
}

func NewSymbolClass(provider ClassProvider) Class {
	s := &symbolClass{}
	s.initialize()
	s.setStringer(s.String)
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

type SymbolValue struct {
	value string
	valueStub
}

func NewSymbol(val string, provider ClassProvider) Value {
	s := &SymbolValue{value: val}
	s.class = provider.ClassWithName("Symbol")
	s.initialize()
	s.setStringer(s.String)
	return s
}

func (SymbolValue *SymbolValue) String() string {
	return SymbolValue.value
}

func (SymbolValue *SymbolValue) Name() string {
	return SymbolValue.value
}
