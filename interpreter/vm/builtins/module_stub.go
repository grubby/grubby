package builtins

import (
	"errors"
	"fmt"
)

type moduleStub struct {
	instanceMethods        map[string]Method
	privateInstanceMethods map[string]Method
}

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
