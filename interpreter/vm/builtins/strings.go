package builtins

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type StringClass struct {
	valueStub
	classStub

	provider Provider
}

func NewStringClass(provider Provider) Class {
	s := &StringClass{}
	s.initialize()
	s.setStringer(s.String)

	s.provider = provider
	s.class = provider.ClassProvider().ClassWithName("Class")
	s.superClass = provider.ClassProvider().ClassWithName("Object")

	s.AddMethod(NewNativeMethod("+", provider, func(self Value, block Block, args ...Value) (Value, error) {
		arg := args[0].(*StringValue)
		selfAsStr := self.(*StringValue)
		return NewString(selfAsStr.value+arg.value, provider), nil
	}))
	s.AddMethod(NewNativeMethod("==", provider, func(self Value, block Block, args ...Value) (Value, error) {
		asStr, ok := args[0].(*StringValue)
		if !ok {
			return provider.SingletonProvider().SingletonWithName("false"), nil
		}

		selfAsStr := self.(*StringValue)
		if selfAsStr.value == asStr.value {
			return provider.SingletonProvider().SingletonWithName("true"), nil
		} else {
			return provider.SingletonProvider().SingletonWithName("false"), nil
		}
	}))
	s.AddMethod(NewNativeMethod("<<", provider, func(self Value, block Block, args ...Value) (Value, error) {
		arg := args[0].(*StringValue)
		selfAsStr := self.(*StringValue)
		if selfAsStr.frozen {
			return nil, errors.New("RuntimeError: can't modify frozen String")
		}

		selfAsStr.value += arg.value
		return selfAsStr, nil
	}))
	s.AddMethod(NewNativeMethod("to_i", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsStr := self.(*StringValue)

		intValue, _ := strconv.ParseInt(selfAsStr.value, 0, 64)
		return NewFixnum(intValue, provider), nil
	}))
	s.AddMethod(NewNativeMethod("freeze", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsStr := self.(*StringValue)
		selfAsStr.frozen = true
		return selfAsStr, nil
	}))
	s.AddMethod(NewNativeMethod("intern", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsStr := self.(*StringValue)
		maybeSymbol := provider.SingletonProvider().SymbolWithName(selfAsStr.value)
		if maybeSymbol != nil {
			return maybeSymbol, nil
		}

		symbolFromString := NewSymbol(selfAsStr.value, provider)
		provider.SingletonProvider().AddSymbol(symbolFromString)
		return symbolFromString, nil
	}))
	s.AddMethod(NewNativeMethod("split", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsStr := self.(*StringValue)
		separator := args[0].(*StringValue)

		val, err := provider.ClassProvider().ClassWithName("Array").New(provider)
		if err != nil {
			return nil, err
		}

		array := val.(*Array)
		pieces := strings.Split(selfAsStr.value, separator.value)
		for _, piece := range pieces {
			str := NewString(piece, provider)
			array.Append(str)
		}

		return array, nil
	}))

	return s
}

func (c *StringClass) String() string {
	return "String"
}

func (c *StringClass) Name() string {
	return "String"
}

func (class *StringClass) New(provider Provider, args ...Value) (Value, error) {
	str := &StringValue{}
	str.initialize()
	str.setStringer(str.String)
	str.setPrettyPrinter(str.PrettyPrint)
	str.class = class

	return str, nil
}

type StringValue struct {
	value string
	valueStub
	frozen bool
}

func (s *StringValue) String() string {
	return fmt.Sprintf("%s", s.value)
}

func (s *StringValue) PrettyPrint() string {
	return fmt.Sprintf("\"%s\"", s.value)
}

func (s *StringValue) RawString() string {
	return s.value
}

func NewString(str string, provider Provider) Value {
	s, _ := provider.ClassProvider().ClassWithName("String").New(provider)
	s.(*StringValue).value = str
	return s
}
