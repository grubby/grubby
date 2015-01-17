package builtins

import (
	"errors"
	"fmt"
)

type moduleStub struct {
	instanceMethods map[string]Method
}

func (m *moduleStub) InstanceMethod(name string) (Method, error) {
	method := m.instanceMethods[name]
	if method == nil {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return method, nil
}

func (m *moduleStub) AddInstanceMethod(method Method) {
	m.instanceMethods[method.Name()] = method
}

func (m *moduleStub) InstanceMethods() []Method {
	methods := make([]Method, 0, len(m.instanceMethods))
	for _, method := range m.instanceMethods {
		methods = append(methods, method)
	}

	return methods
}
