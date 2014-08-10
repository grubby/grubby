package builtins

import (
	"errors"
	"fmt"
)

// this type repesents the shared behavior and data of all Ruby Values
// all values will need to store the methods that are defined on them
// (in addition to their class, and other information)
type valueStub struct {
	methods         map[string]Method
	private_methods map[string]Method
	class           Class
}

func (valueStub *valueStub) initialize() {
	valueStub.methods = make(map[string]Method)
	valueStub.private_methods = make(map[string]Method)
}

func (valueStub *valueStub) Method(name string) (Method, error) {
	m, ok := valueStub.methods[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return m, nil
}

func (valueStub *valueStub) PrivateMethod(name string) (Method, error) {
	m, ok := valueStub.private_methods[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return m, nil
}

func (valueStub *valueStub) Methods() []Method {
	values := make([]Method, len(valueStub.methods))
	for _, m := range valueStub.methods {
		values = append(values, m)
	}

	return values
}

func (valueStub *valueStub) PrivateMethods() []Method {
	values := make([]Method, len(valueStub.private_methods))
	for _, m := range valueStub.private_methods {
		values = append(values, m)
	}

	return values
}

func (valueStub *valueStub) AddMethod(m Method) {
	valueStub.methods[m.Name()] = m
}

func (valueStub *valueStub) AddPrivateMethod(m Method) {
	valueStub.private_methods[m.Name()] = m
}

func (valueStub *valueStub) String() string {
	return "ValueStub"
}

func (valueStub *valueStub) Class() Class {
	return valueStub.class
}
