package builtins

import (
	"errors"
	"fmt"
)

type symbolClass struct {
	valueStub
	classStub
}

func NewSymbolClass(classProvider ClassProvider, singletonProvider SingletonProvider) Class {
	s := &symbolClass{}
	s.initialize()
	s.setStringer(s.String)
	s.class = classProvider.ClassWithName("Class")
	s.superClass = classProvider.ClassWithName("Object")

	s.AddMethod(NewNativeMethod("to_proc", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsSymbol := self.(*SymbolValue)
		return NewProcInstance(selfAsSymbol.value, classProvider), nil
	}))
	s.AddMethod(NewNativeMethod("to_s", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsSymbol := self.(*SymbolValue)
		return NewString(selfAsSymbol.value, classProvider, singletonProvider), nil
	}))

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
	return fmt.Sprintf(":%s", SymbolValue.value)
}

func (SymbolValue *SymbolValue) Name() string {
	return SymbolValue.value
}
