package builtins

import (
	"errors"
	"fmt"
)

type moduleStub struct {
	instanceMethods          map[string]Method
	privateInstanceMethods   map[string]Method
	protectedInstanceMethods map[string]Method
	constants                map[string]Value

	methodVisibility MethodVisibility
}

type MethodVisibility int

const (
	Public MethodVisibility = iota
	Private
	Protected
)

func (m *moduleStub) InstanceMethod(name string) (Method, error) {
	method := m.instanceMethods[name]
	if method == nil {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return method, nil
}

func (m *moduleStub) AddInstanceMethod(method Method) {
	if m.instanceMethods == nil {
		m.instanceMethods = make(map[string]Method)
	}

	m.instanceMethods[method.Name()] = method
}

func (m *moduleStub) RemoveInstanceMethod(method Method) {
	delete(m.instanceMethods, method.Name())
}

func (m *moduleStub) InstanceMethods() []Method {
	methods := make([]Method, 0, len(m.instanceMethods))
	for _, method := range m.instanceMethods {
		methods = append(methods, method)
	}

	return methods
}

func (m *moduleStub) PrivateInstanceMethod(name string) (Method, error) {
	method := m.privateInstanceMethods[name]
	if method == nil {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return method, nil
}

func (m *moduleStub) AddPrivateInstanceMethod(method Method) {
	if m.privateInstanceMethods == nil {
		m.privateInstanceMethods = make(map[string]Method)
	}

	m.privateInstanceMethods[method.Name()] = method
}

func (m *moduleStub) PrivateInstanceMethods() []Method {
	methods := make([]Method, 0, len(m.privateInstanceMethods))
	for _, method := range m.privateInstanceMethods {
		methods = append(methods, method)
	}

	return methods
}

func (m *moduleStub) ProtectedInstanceMethod(name string) (Method, error) {
	method := m.protectedInstanceMethods[name]
	if method == nil {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return method, nil
}

func (m *moduleStub) AddProtectedInstanceMethod(method Method) {
	if m.privateInstanceMethods == nil {
		m.privateInstanceMethods = make(map[string]Method)
	}

	m.privateInstanceMethods[method.Name()] = method
}

func (m *moduleStub) ProtectedInstanceMethods() []Method {
	methods := make([]Method, 0, len(m.privateInstanceMethods))
	for _, method := range m.privateInstanceMethods {
		methods = append(methods, method)
	}

	return methods
}

func (m *moduleStub) Constant(name string) (Value, error) {
	value, ok := m.constants[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("constant '%s' does not exist", name))
	}

	return value, nil
}

func (m *moduleStub) SetConstant(name string, value Value) {
	if m.constants == nil {
		m.constants = make(map[string]Value)
	}

	m.constants[name] = value
}

func (m *moduleStub) Constants() []Value {
	constants := []Value{}
	for _, c := range m.constants {
		constants = append(constants, c)
	}

	return constants
}

func (m *moduleStub) ConstantsWithNames() map[string]Value {
	return m.constants
}

func (m *moduleStub) ActiveVisibility() MethodVisibility {
	return m.methodVisibility
}

func (m *moduleStub) SetActiveVisibility(visibility MethodVisibility) {
	m.methodVisibility = visibility
}
