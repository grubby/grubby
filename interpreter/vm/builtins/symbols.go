package builtins

import (
	"errors"
	"fmt"
)

type symbolClass struct {
	valueStub
	classStub
}

func NewSymbolClass(provider Provider) Class {
	s := &symbolClass{}
	s.initialize()
	s.setStringer(s.String)
	s.class = provider.ClassProvider().ClassWithName("Class")
	s.superClass = provider.ClassProvider().ClassWithName("Object")

	s.AddMethod(NewNativeMethod("to_proc", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsSymbol := self.(*SymbolValue)
		return NewProcInstance(selfAsSymbol.value, provider), nil
	}))
	s.AddMethod(NewNativeMethod("to_s", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsSymbol := self.(*SymbolValue)
		return NewString(selfAsSymbol.value, provider), nil
	}))

	return s
}

func (c *symbolClass) String() string {
	return "Symbol"
}

func (c *symbolClass) Name() string {
	return "Symbol"
}

func (c *symbolClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Symbol:Class")
}

type SymbolValue struct {
	value string
	valueStub
}

func NewSymbol(val string, provider Provider) Value {
	s := &SymbolValue{value: val}
	s.class = provider.ClassProvider().ClassWithName("Symbol")
	s.initialize()
	s.setStringer(s.String)
	return s
}

func (SymbolValue *SymbolValue) String() string {
	return fmt.Sprintf(":%s", SymbolValue.value)
}

func (SymbolValue *SymbolValue) Name() string {
	return SymbolValue.value
}
