package builtins

import (
	"errors"
	"fmt"
)

// abstract class interface
type Class interface {
	Module // includes Value

	New(provider Provider, args ...Value) (Value, error)

	SuperClass() Class

	Include(Module)

	includedModules() []Module
	classVariable(string) Value
	setClassVariable(string, Value)
}

// globlal Class class
type ClassValue struct {
	valueStub
	classStub

	provider Provider
}

func NewClassClass(provider Provider) Class {
	c := &ClassValue{}
	c.initialize()
	c.setStringer(c.String)
	c.class = c
	c.provider = provider

	c.AddMethod(NewNativeMethod("new", provider, func(self Value, block Block, args ...Value) (Value, error) {
		instance, err := self.(Class).New(provider, args...)
		if err != nil {
			return nil, err
		}

		method := instance.Method("initialize")
		if method != nil {
			_, err = method.Execute(instance, block, args...)
			if err != nil {
				panic(err)
			}
		}

		return instance, nil
	}))

	return c
}

func (c *ClassValue) SetSuperClass() {
	moduleClass := c.provider.ClassProvider().ClassWithName("Module")
	if moduleClass == nil {
		panic("Expected Module class to exist")
	}
	c.superClass = moduleClass
}

func (c ClassValue) New(provider Provider, args ...Value) (Value, error) {
	return nil, nil
}

func (c ClassValue) Name() string {
	return "Class"
}

func (c ClassValue) String() string {
	return "Class"
}

// user defined class type
type UserDefinedClass struct {
	name string
	valueStub
	classStub

	attr_readers []string
	attr_writers []string
}

type UserDefinedClassInstance struct {
	valueStub

	provider Provider
	attrs    map[string]Value
}

func (i *UserDefinedClassInstance) String() string {
	return fmt.Sprintf("<%s:%p>", i.Class().String(), i)
}

func NewUserDefinedClass(name string, provider Provider) Class {
	c := &UserDefinedClass{
		name: name,
	}
	c.initialize()
	c.setStringer(c.String)
	c.class = provider.ClassProvider().ClassWithName("Class")

	// FIXME: should be provided as an argument
	c.superClass = provider.ClassProvider().ClassWithName("Object")

	c.AddMethod(NewNativeMethod("include", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			c.Include(arg.(Module))
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("extend", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, module := range args {
			for _, method := range module.(Module).InstanceMethods() {
				self.AddMethod(method)
			}
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("attr_accessor", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			class := self.(*UserDefinedClass)
			class.attr_readers = append(class.attr_readers, symbol.Name())
			class.attr_writers = append(class.attr_writers, symbol.Name())
		}

		return nil, nil
	}))
	c.AddMethod(NewNativeMethod("attr_reader", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			class := self.(*UserDefinedClass)
			class.attr_readers = append(class.attr_readers, symbol.Name())
		}

		return nil, nil
	}))
	c.AddMethod(NewNativeMethod("attr_writer", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			class := self.(*UserDefinedClass)
			class.attr_writers = append(class.attr_writers, symbol.Name())
		}

		return nil, nil
	}))

	return c
}

func (c *UserDefinedClass) New(provider Provider, args ...Value) (Value, error) {
	instance := &UserDefinedClassInstance{}
	instance.initialize()
	instance.setStringer(instance.String)
	instance.provider = provider
	instance.attrs = make(map[string]Value)
	instance.class = c

	for _, m := range c.instanceMethods {
		instance.AddMethod(m)
	}

	for _, module := range c.includedModules() {
		for _, method := range module.(Module).InstanceMethods() {
			instance.AddMethod(method)
		}
	}

	// FIXME: these should be defined on Module
	for _, attr := range c.attr_readers {
		instance.AddMethod(NewNativeMethod(attr, provider, func(self Value, block Block, args ...Value) (Value, error) {
			this := self.(*UserDefinedClassInstance)
			value, ok := this.attrs[attr]
			if !ok {
				return provider.SingletonProvider().SingletonWithName("nil"), nil
			}

			return value, nil
		}))
	}

	for _, attr := range c.attr_writers {
		instance.AddMethod(NewNativeMethod(attr+"=", provider, func(self Value, block Block, args ...Value) (Value, error) {
			this := self.(*UserDefinedClassInstance)
			this.attrs[attr] = args[0]
			return nil, nil
		}))
	}

	method := instance.Method("initialize")
	if method != nil {
		_, err := method.Execute(instance, nil, args...)
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}

func (c UserDefinedClass) Name() string {
	return c.name
}

func (c UserDefinedClass) String() string {
	return c.name
}

func (c *UserDefinedClass) Class() Class {
	return c.class
}
