package builtins

import "fmt"

type method struct {
	name            string
	methods         []Method
	private_methods []Method
	body            func(...Value) (Value, error)
}

type Method interface {
	Name() string
	Value
	Execute(args ...Value) (Value, error)
}

func NewMethod(name string, body func(...Value) (Value, error)) Method {
	return &method{name: name, body: body}
}

func (method *method) Name() string {
	return method.name
}

func (method *method) Methods() []Method {
	return method.methods
}

func (method *method) PrivateMethods() []Method {
	return method.private_methods
}

func (method *method) AddMethod(m Method) {
	method.methods = append(method.methods, m)
}

func (method *method) AddPrivateMethod(m Method) {
	method.private_methods = append(method.private_methods, m)
}

func (method *method) Execute(args ...Value) (Value, error) {
	return method.body(args...)
}

func (method *method) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}

func (method *method) Class() Class {
	return nil
}
