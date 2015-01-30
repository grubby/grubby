package builtins

import (
	"errors"
	"fmt"
)

// abstract module interface
type Module interface {
	Name() string
	AddInstanceMethod(Method)
	InstanceMethods() []Method
	InstanceMethod(string) (Method, error)

	Value
}

// globlal Module class
type ModuleClass struct {
	valueStub
	classStub
}

func NewModuleClass(classProvider ClassProvider, singletonProvider SingletonProvider) Class {
	c := &ModuleClass{}
	c.initialize()
	c.class = classProvider.ClassWithName("Class")
	c.superClass = classProvider.ClassWithName("Object")

	c.AddMethod(NewNativeMethod("private_class_method", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		methodName, ok := args[0].(*symbolValue)
		if !ok {
			return nil, errors.New(fmt.Sprintf("TypeError: %v is not a symbol", args[0]))
		}

		method, err := self.Method(methodName.String())
		if err != nil {
			panic(fmt.Sprintf("%#v", self))
			return nil, err
		}

		self.RemoveMethod(method)
		self.AddPrivateMethod(method)

		return methodName, nil
	}))

	return c
}

func (c ModuleClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, nil
}

func (c ModuleClass) Name() string {
	return "Module"
}

func (c ModuleClass) String() string {
	return "Module"
}

// user defined module type
type RubyModule struct {
	name string
	valueStub
	moduleStub

	includedModules []Value
}

func NewModule(name string, provider ClassProvider, singletonProvider SingletonProvider) Module {
	c := &RubyModule{
		name:            name,
		includedModules: make([]Value, 0),
	}
	c.initialize()
	c.class = provider.ClassWithName("Module")

	c.AddMethod(NewNativeMethod("include", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, module := range args {
			c.includedModules = append(c.includedModules, module)
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("extend", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, module := range args {
			for _, method := range module.Methods() {
				self.AddMethod(method)
			}
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("module_function", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("expected exactly one arg")
		}

		symbol, ok := args[0].(*symbolValue)
		if !ok {
			return nil, errors.New("expected method name to be a symbol")
		}

		instanceMethod, err := self.(*RubyModule).InstanceMethod(symbol.String())
		if err != nil {
			return nil, err
		}

		self.(*RubyModule).AddMethod(instanceMethod)
		// FIXME: this should mark the original instance method as private
		return self, nil
	}))

	return c
}

func (m RubyModule) Name() string {
	return m.name
}

func (m *RubyModule) String() string {
	return fmt.Sprintf("%s:Module", m.name)
}
