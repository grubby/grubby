package builtins

import (
	"errors"
	"fmt"
	"strings"
)

// abstract module interface
type Module interface {
	Name() string
	AddInstanceMethod(Method)
	RemoveInstanceMethod(Method)

	InstanceMethods() []Method
	InstanceMethod(string) (Method, error)

	AddPrivateInstanceMethod(Method)
	PrivateInstanceMethods() []Method
	PrivateInstanceMethod(string) (Method, error)

	AddProtectedInstanceMethod(Method)
	ProtectedInstanceMethods() []Method
	ProtectedInstanceMethod(string) (Method, error)

	Constants() []Value
	ConstantsWithNames() map[string]Value
	Constant(string) (Value, error)
	SetConstant(string, Value)

	Value

	ActiveVisibility() MethodVisibility
	SetActiveVisibility(MethodVisibility)
}

type Evaluator interface {
	EvaluateStringInContext(string, Value) (Value, error)
	EvaluateStringInContextAndNewStack(string, Value) (Value, error)
}

// globlal Module class
type ModuleClass struct {
	valueStub
	classStub
}

func NewModuleClass(classProvider ClassProvider, singletonProvider SingletonProvider, evaluator Evaluator) Class {
	c := &ModuleClass{}
	c.initialize()
	c.setStringer(c.String)
	c.class = classProvider.ClassWithName("Class")
	c.superClass = classProvider.ClassWithName("Object")

	c.AddMethod(NewNativeMethod("private_class_method", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		methodName, ok := args[0].(*SymbolValue)
		if !ok {
			return nil, errors.New(fmt.Sprintf("TypeError: %v is not a symbol", args[0]))
		}

		method, err := self.Method(methodName.value)
		if err != nil {
			return nil, err
		}

		self.RemoveMethod(method)

		native, ok := method.(*nativeMethod)
		if ok {
			native.private = true
			self.AddMethod(native)
		} else if rubyMethod, ok := method.(*RubyMethod); ok {
			rubyMethod.private = true
			self.AddMethod(rubyMethod)
		} else {
			panic(fmt.Sprintf("unknown method type: %T", method))
		}

		return methodName, nil
	}))

	c.AddMethod(NewNativeMethod("public", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)
		if len(args) == 0 {
			selfAsModule.SetActiveVisibility(Public)
			return singletonProvider.SingletonWithName("nil"), nil
		}

		for _, arg := range args {
			argAsSym, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New(fmt.Sprintf("TypeError: '%s' is not a symbol", arg))
			}

			method, err := selfAsModule.InstanceMethod(argAsSym.value)
			if err != nil {
				return nil, err
			}

			selfAsModule.RemoveInstanceMethod(method)
			selfAsModule.AddInstanceMethod(method)
		}

		return singletonProvider.SingletonWithName("nil"), nil
	}))
	c.AddMethod(NewNativeMethod("private", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)

		if len(args) == 0 {
			selfAsModule.SetActiveVisibility(Private)
			return singletonProvider.SingletonWithName("nil"), nil
		}

		for _, arg := range args {
			argAsSym, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New(fmt.Sprintf("TypeError: '%s' is not a symbol", arg))
			}

			method, err := selfAsModule.InstanceMethod(argAsSym.value)
			if err != nil {
				return nil, err
			}

			selfAsModule.RemoveInstanceMethod(method)
			selfAsModule.AddPrivateInstanceMethod(method)
		}

		return singletonProvider.SingletonWithName("nil"), nil
	}))
	c.AddMethod(NewNativeMethod("protected", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)
		if len(args) == 0 {
			selfAsModule.SetActiveVisibility(Protected)
			return singletonProvider.SingletonWithName("nil"), nil
		}

		for _, arg := range args {
			argAsSym, ok := arg.(*SymbolValue)
			if !ok {
				return nil, errors.New(fmt.Sprintf("TypeError: '%s' is not a symbol", arg))
			}

			method, err := selfAsModule.InstanceMethod(argAsSym.value)
			if err != nil {
				return nil, err
			}

			selfAsModule.RemoveInstanceMethod(method)
			selfAsModule.AddProtectedInstanceMethod(method)
		}

		return singletonProvider.SingletonWithName("nil"), nil
	}))

	c.AddMethod(NewNativeMethod("module_eval", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		input := args[0].(*StringValue).RawString()
		input = strings.Replace(input, "\\n", "\n", -1)
		return evaluator.EvaluateStringInContextAndNewStack(input, self)
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

	includedModules []Module
}

func NewModule(name string, classProvider ClassProvider, singletonProvider SingletonProvider) Module {
	c := &RubyModule{
		name:            name,
		includedModules: make([]Module, 0),
	}
	c.initialize()
	c.setStringer(c.String)
	c.class = classProvider.ClassWithName("Module")

	c.AddMethod(NewNativeMethod("include", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)
		for _, val := range args {
			moduleToInclude, ok := val.(Module)
			if !ok {
				return nil, errors.New("TypeError: wrong argument type (expected Module)")
			}

			for name, constant := range moduleToInclude.ConstantsWithNames() {
				selfAsModule.SetConstant(name, constant)
			}

			c.includedModules = append(c.includedModules, moduleToInclude)

		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("extend", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, module := range args {
			for _, method := range module.Methods() {
				self.AddMethod(method)
			}
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("module_function", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("expected exactly one arg")
		}

		symbol, ok := args[0].(*SymbolValue)
		if !ok {
			return nil, errors.New("expected method name to be a symbol")
		}

		instanceMethod, err := self.(*RubyModule).InstanceMethod(symbol.value)
		if err != nil {
			instanceMethod, err = self.(*RubyModule).PrivateInstanceMethod(symbol.value)
		}

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
	return m.name
}
