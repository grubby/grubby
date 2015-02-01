package builtins

import (
	"errors"
	"fmt"
)

// this type repesents the shared behavior and data of all Ruby Values
// all values will need to store the methods that are defined on them
// (in addition to their class, and other information)
type valueStub struct {
	eigenclass_methods map[string]Method
	private_methods    map[string]Method
	class              Class

	stringer func() string

	instance_variables map[string]Value
}

func (valueStub *valueStub) initialize() {
	valueStub.eigenclass_methods = make(map[string]Method)
	valueStub.private_methods = make(map[string]Method)
	valueStub.instance_variables = make(map[string]Value)
}

// Method Lookup //

/*
  1. Methods defined in the object's singleton class (i.e. the object itself)
  2. Modules mixed into the singleton class in reverse order of inclusion
  3. Methods defined by the object's class
	4. Modules included into the object's class in reverse order of inclusion
	5. Methods defined by the object's superclass, i.e. inherited methods
  6. Once BasicObject is reached, start at 1 with "method_missing" method
  7. Fail. Loudly.
*/

func (valueStub *valueStub) Method(name string) (Method, error) {
	//	  1. Methods defined in the object's singleton class (i.e. the object itself)
	m, ok := valueStub.eigenclass_methods[name]
	if ok {
		return m, nil
	}

	//    2. Modules mixed into the singleton class in reverse order of inclusion
	// FIXME: respect step 2 here

	//	  3. Methods defined by the object's class
	m, ok = valueStub.class.eigenclassMethods()[name]
	if ok {
		return m, nil
	}

	//		4. Modules included into the object's class in reverse order of inclusion
	// FIXME: this should be reversed (should be fixed in Include method)
	for _, module := range valueStub.class.includedModules() {
		m, ok := module.eigenclassMethods()[name]
		if ok {
			return m, nil
		}
	}

	//		5. Methods defined by the object's superclass, i.e. inherited methods
	super := valueStub.Class().SuperClass()
	for super != nil {
		m, ok := super.eigenclassMethods()[name]
		if ok {
			return m, nil
		}

		super = super.SuperClass()
	}

	return nil, NewNoMethodError(name, valueStub.String(), valueStub.Class().String(), "")
}

func (valueStub *valueStub) PrivateMethod(name string) (Method, error) {
	m, ok := valueStub.private_methods[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("method: '%s' does not exist", name))
	}

	return m, nil
}

func (valueStub *valueStub) Methods() []Method {
	values := make([]Method, 0, len(valueStub.eigenclass_methods))
	for _, m := range valueStub.eigenclass_methods {
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
	valueStub.eigenclass_methods[m.Name()] = m
}

func (valueStub *valueStub) RemoveMethod(m Method) {
	delete(valueStub.eigenclass_methods, m.Name())
}

func (valueStub *valueStub) AddPrivateMethod(m Method) {
	valueStub.private_methods[m.Name()] = m
}

func (valueStub *valueStub) String() string {
	return valueStub.stringer()
}

func (valueStub *valueStub) setStringer(stringer func() string) {
	valueStub.stringer = stringer
}

func (valueStub *valueStub) Class() Class {
	return valueStub.class
}

func (valueStub *valueStub) eigenclassMethods() map[string]Method {
	return valueStub.eigenclass_methods
}

func (valueStub *valueStub) GetInstanceVariable(name string) Value {
	return valueStub.instance_variables[name]
}

func (valueStub *valueStub) SetInstanceVariable(name string, value Value) {
	valueStub.instance_variables[name] = value
}

func (v *valueStub) IsTruthy() bool {
	return true
}
