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

func NewModuleClass(provider Provider, evaluator Evaluator) Class {
	c := &ModuleClass{}
	c.initialize()
	c.setStringer(c.String)
	c.class = provider.ClassProvider().ClassWithName("Class")
	c.superClass = provider.ClassProvider().ClassWithName("Object")

	c.AddMethod(NewNativeMethod("private_class_method", provider, func(self Value, block Block, args ...Value) (Value, error) {
		methodName, ok := args[0].(*SymbolValue)
		if !ok {
			return nil, errors.New(fmt.Sprintf("TypeError: %v is not a symbol", args[0]))
		}

		method := self.Method(methodName.value)
		if method == nil {
			return nil, NewNoMethodError(methodName.value, self.String(), self.Class().String(), provider.StackProvider().CurrentStack())
		}

		self.RemoveMethod(method)

		native, ok := method.(*nativeMethod)
		if ok {
			native.visibility = Private
			self.AddMethod(native)
		} else if rubyMethod, ok := method.(*RubyMethod); ok {
			rubyMethod.visibility = Private
			self.AddMethod(rubyMethod)
		} else {
			panic(fmt.Sprintf("unknown method type: %T", method))
		}

		return methodName, nil
	}))

	c.AddMethod(NewNativeMethod("public", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)
		if len(args) == 0 {
			selfAsModule.SetActiveVisibility(Public)
			return provider.SingletonProvider().SingletonWithName("nil"), nil
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

			method.SetVisibility(Public)
		}

		return provider.SingletonProvider().SingletonWithName("nil"), nil
	}))
	c.AddMethod(NewNativeMethod("private", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)

		if len(args) == 0 {
			selfAsModule.SetActiveVisibility(Private)
			return provider.SingletonProvider().SingletonWithName("nil"), nil
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

			method.SetVisibility(Private)
		}

		return provider.SingletonProvider().SingletonWithName("nil"), nil
	}))
	c.AddMethod(NewNativeMethod("protected", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)
		if len(args) == 0 {
			selfAsModule.SetActiveVisibility(Protected)
			return provider.SingletonProvider().SingletonWithName("nil"), nil
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

			method.SetVisibility(Protected)
		}

		return provider.SingletonProvider().SingletonWithName("nil"), nil
	}))

	c.AddMethod(NewNativeMethod("module_eval", provider, func(self Value, block Block, args ...Value) (Value, error) {
		input := args[0].(*StringValue).RawString()
		input = strings.Replace(input, "\\n", "\n", -1)
		return evaluator.EvaluateStringInContextAndNewStack(input, self)
	}))

	c.AddMethod(NewNativeMethod("include", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)
		for _, val := range args {
			moduleToInclude, ok := val.(Module)
			if !ok {
				return nil, errors.New("TypeError: wrong argument type (expected Module)")
			}

			for _, method := range moduleToInclude.InstanceMethods() {
				selfAsModule.AddInstanceMethod(method)
			}

			for name, constant := range moduleToInclude.ConstantsWithNames() {
				selfAsModule.SetConstant(name, constant)
			}
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("extend", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, module := range args {
			for _, method := range module.Methods() {
				self.AddMethod(method)
			}
		}

		return c, nil
	}))

	c.AddMethod(NewNativeMethod("module_function", provider, func(self Value, block Block, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("expected exactly one arg")
		}

		symbol, ok := args[0].(*SymbolValue)
		if !ok {
			return nil, errors.New("expected method name to be a symbol")
		}

		instanceMethod, err := self.(*RubyModule).InstanceMethod(symbol.value)

		if err != nil {
			return nil, err
		}

		self.(*RubyModule).AddMethod(instanceMethod)
		instanceMethod.SetVisibility(Private)
		return self, nil
	}))

	c.AddMethod(NewNativeMethod("alias_method", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)

		firstSymbol, ok := args[0].(*SymbolValue)
		if !ok {
			return nil, errors.New("expected method name to be a symbol")
		}
		secondSymbol, ok := args[1].(*SymbolValue)
		if !ok {
			return nil, errors.New("expected method name to be a symbol")
		}

		method, err := selfAsModule.InstanceMethod(secondSymbol.Name())
		if err != nil {
			return nil, err
		}

		selfAsModule.AddInstanceMethod(NewNativeMethod(firstSymbol.Name(), provider, method.methodBody()))

		return nil, nil
	}))

	c.AddMethod(NewNativeMethod("const_defined?", provider, func(self Value, block Block, args ...Value) (Value, error) {
		objClass := provider.ClassProvider().ClassWithName("Object")
		_, err := objClass.Constant(args[0].(*SymbolValue).Name())
		if err == nil {
			return provider.SingletonProvider().SingletonWithName("true"), nil
		} else {
			return provider.SingletonProvider().SingletonWithName("false"), nil
		}
	}))
	c.AddMethod(NewNativeMethod("method_defined?", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsModule := self.(Module)
		_, err := selfAsModule.InstanceMethod(args[0].(*SymbolValue).value)

		if err == nil {
			return provider.SingletonProvider().SingletonWithName("true"), nil
		} else {
			return provider.SingletonProvider().SingletonWithName("false"), nil
		}
	}))

	return c
}

func (c ModuleClass) New(provider Provider, args ...Value) (Value, error) {
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
}

func NewModule(name string, provider Provider) Module {
	c := &RubyModule{
		name: name,
	}
	c.initialize()
	c.setStringer(c.String)
	c.class = provider.ClassProvider().ClassWithName("Module")

	return c
}

func (m RubyModule) Name() string {
	return m.name
}

func (m *RubyModule) String() string {
	return m.name
}
