package builtins

import "fmt"

type method struct {
	name string
	body func(...Value) (Value, error)
	valueStub
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

func (method *method) Execute(args ...Value) (Value, error) {
	return method.body(args...)
}

func (method *method) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}
