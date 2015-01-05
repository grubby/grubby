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

	Value
}

// globlal Module class
type ModuleClass struct {
	valueStub
	classStub
	instanceMethods map[string]Method
}

func NewModuleClass(provider ClassProvider) Class {
	c := &ModuleClass{}
	c.initialize()
	c.instanceMethods = make(map[string]Method, 0)
	c.class = provider.ClassWithName("Class")
	c.superClass = provider.ClassWithName("Object")

	return c
}

func (c ModuleClass) New(provider ClassProvider, args ...Value) Value {
	return nil
}

func (c ModuleClass) Name() string {
	return "Module"
}

func (c ModuleClass) String() string {
	return "Module"
}

func (c *ModuleClass) AddInstanceMethod(m Method) {
	c.instanceMethods[m.Name()] = m
}

func (c *ModuleClass) InstanceMethods() []Method {
	methods := []Method{}
	for _, m := range c.instanceMethods {
		methods = append(methods, m)
	}

	return methods
}

// user defined module type
type RubyModule struct {
	name string
	valueStub

	includedModules []Value
	instanceMethods map[string]Method
}

func NewModule(name string, provider ClassProvider) Module {
	c := &RubyModule{
		name:            name,
		includedModules: make([]Value, 0),
		instanceMethods: make(map[string]Method, 0),
	}
	c.initialize()
	c.class = provider.ClassWithName("Module")

	c.AddMethod(NewNativeMethod("include", provider, func(self Value, args ...Value) (Value, error) {
		for _, module := range args {
			c.includedModules = append(c.includedModules, module)
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("module_function", provider, func(self Value, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("expected exactly one arg")
		}

		symbol, ok := args[0].(*SymbolValue)
		if !ok {
			return nil, errors.New("expected method name to be a symbol")
		}

		var instanceMethod Method
		for name, m := range self.(*RubyModule).instanceMethods {
			if name == symbol.String() {
				instanceMethod = m
				break
			}
		}

		if instanceMethod == nil {
			return nil, errors.New(fmt.Sprintf("method '%s' does not exist", symbol.String()))
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

func (m *RubyModule) AddInstanceMethod(method Method) {
	m.instanceMethods[method.Name()] = method
}

func (m *RubyModule) InstanceMethods() []Method {
	methods := []Method{}
	for _, m := range m.instanceMethods {
		methods = append(methods, m)
	}

	return methods
}

func (m *RubyModule) InstanceMethod(name string) (Method, error) {
	method := m.instanceMethods[name]
	if method == nil {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return method, nil
}

func (m *RubyModule) String() string {
	return fmt.Sprintf("%s:Module", m.name)
}
