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
}

type UserDefinedClassInstance struct {
	valueStub

	provider Provider
	attrs    map[string]Value
}

func (i *UserDefinedClassInstance) String() string {
	return fmt.Sprintf("<%s:%p>", i.Class().String(), i)
}

func NewUserDefinedClass(name, superclassName string, provider Provider) Class {
	c := &UserDefinedClass{
		name: name,
	}
	c.initialize()
	c.setStringer(c.String)
	c.class = provider.ClassProvider().ClassWithName("Class")

	if superclassName == "" {
		superclassName = "Object"
	}

	superclass := provider.ClassProvider().ClassWithName(superclassName)
	if superclass == nil {
		panic(fmt.Sprintf("Class '%s' not found", superclassName))
	}
	c.superClass = superclass

	// FIXME : in reality, the singleton class for us should be
	// the singleton class of our superclass, but this is a short-term fix
	for _, method := range superclass.eigenclassMethods() {
		c.AddMethod(method)
	}

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

	//FIXME : these should be on module
	c.AddMethod(NewNativeMethod("attr_accessor", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			provider.MethodProvider().AddMethod(symbol.Name(), self, func(self Value, block Block, args ...Value) (Value, error) {
				value, ok := self.GetAttribute(symbol.Name())
				if !ok {
					return provider.SingletonProvider().SingletonWithName("nil"), nil
				}

				return value, nil
			})
			provider.MethodProvider().AddMethod(symbol.Name()+"=", self, func(self Value, block Block, args ...Value) (Value, error) {
				self.SetAttribute(symbol.Name(), args[0])
				return nil, nil
			})
		}

		return nil, nil
	}))
	c.AddMethod(NewNativeMethod("attr_reader", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			provider.MethodProvider().AddMethod(symbol.Name(), self, func(self Value, block Block, args ...Value) (Value, error) {
				value, ok := self.GetAttribute(symbol.Name())
				if !ok {
					return provider.SingletonProvider().SingletonWithName("nil"), nil
				}

				return value, nil
			})
		}

		return nil, nil
	}))
	c.AddMethod(NewNativeMethod("attr_writer", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			provider.MethodProvider().AddMethod(symbol.Name()+"=", self, func(self Value, block Block, args ...Value) (Value, error) {
				self.SetAttribute(symbol.Name(), args[0])
				return nil, nil
			})
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
