package builtins

import "errors"

// abstract class interface
type Class interface {
	Value

	New(provider ClassProvider, args ...Value) (Value, error)
	Name() string

	AddInstanceMethod(Method)

	SuperClass() Class

	Include(Module)

	includedModules() []Module
}

// globlal Class class
type ClassValue struct {
	valueStub
	classStub
	instanceMethods []Method

	provider ClassProvider
}

func NewClassClass(provider ClassProvider) Class {
	c := &ClassValue{}
	c.initialize()
	c.class = c
	c.provider = provider

	return c
}

func (c *ClassValue) SetSuperClass() {
	moduleClass := c.provider.ClassWithName("Module")
	if moduleClass == nil {
		panic("Expected Module class to exist")
	}
	c.superClass = moduleClass
}

func (c ClassValue) New(provider ClassProvider, args ...Value) (Value, error) {
	return nil, nil
}

func (c ClassValue) Name() string {
	return "Class"
}

func (c ClassValue) String() string {
	return "Class"
}

func (c ClassValue) AddInstanceMethod(m Method) {
	c.instanceMethods = append(c.instanceMethods, m)
}

// user defined class type
type UserDefinedClass struct {
	name string
	valueStub
	classStub

	instanceMethods map[string]Method

	attr_readers []string
	attr_writers []string
}

type UserDefinedClassInstance struct {
	valueStub

	provider ClassProvider
	attrs    map[string]Value
}

func NewUserDefinedClass(name string, provider ClassProvider) Class {
	c := &UserDefinedClass{
		name:            name,
		instanceMethods: make(map[string]Method),
	}
	c.initialize()
	c.class = provider.ClassWithName("Class")
	c.superClass = nil // FIXME: should be provided as an argument

	c.AddMethod(NewNativeMethod("include", provider, func(self Value, args ...Value) (Value, error) {
		for _, arg := range args {
			c.Include(arg.(Module))
		}

		return c, nil
	}))
	c.AddMethod(NewNativeMethod("new", provider, func(self Value, args ...Value) (Value, error) {
		instance, err := c.New(provider, args...)
		if err != nil {
			return nil, err
		}

		method, err := instance.Method("initialize")
		if err == nil {
			_, err = method.Execute(instance, args...)
			if err != nil {
				panic(err)
			}
		}

		return instance, nil
	}))
	c.AddMethod(NewNativeMethod("attr_accessor", provider, func(self Value, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*symbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			class := self.(*UserDefinedClass)
			class.attr_readers = append(class.attr_readers, symbol.Name())
			class.attr_writers = append(class.attr_writers, symbol.Name())
		}

		return nil, nil
	}))
	c.AddMethod(NewNativeMethod("attr_reader", provider, func(self Value, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*symbolValue)
			if !ok {
				return nil, errors.New("not a symbol or a string")
			}

			class := self.(*UserDefinedClass)
			class.attr_readers = append(class.attr_readers, symbol.Name())
		}

		return nil, nil
	}))
	c.AddMethod(NewNativeMethod("attr_writer", provider, func(self Value, args ...Value) (Value, error) {
		for _, arg := range args {
			symbol, ok := arg.(*symbolValue)
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

func (c *UserDefinedClass) New(provider ClassProvider, args ...Value) (Value, error) {
	instance := &UserDefinedClassInstance{}
	instance.initialize()
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
		instance.AddMethod(NewNativeMethod(attr, provider, func(self Value, args ...Value) (Value, error) {
			this := self.(*UserDefinedClassInstance)
			value, ok := this.attrs[attr]
			if !ok {
				return provider.ClassWithName("Nil").New(provider)
			}

			return value, nil
		}))
	}

	for _, attr := range c.attr_writers {
		instance.AddMethod(NewNativeMethod(attr+"=", provider, func(self Value, args ...Value) (Value, error) {
			this := self.(*UserDefinedClassInstance)
			this.attrs[attr] = args[0]
			return nil, nil
		}))
	}

	method, err := instance.Method("initialize")
	if err == nil {
		_, err = method.Execute(instance, args...)
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

func (c *UserDefinedClass) AddInstanceMethod(m Method) {
	c.instanceMethods[m.Name()] = m
}

func (c *UserDefinedClass) Class() Class {
	return c.class
}
