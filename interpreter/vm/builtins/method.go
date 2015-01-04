package builtins

import "fmt"

type method struct {
	name string
	body func(Value, ...Value) (Value, error)
	valueStub
}

type Method interface {
	Name() string
	Value
	Execute(self Value, args ...Value) (Value, error)
}

func AddMethod(provider ClassProvider, target Value, name string, body func(Value, ...Value) (Value, error)) {
	m := &method{name: name, body: body}
	m.class = provider.ClassWithName("Method")
	m.initialize()
	target.AddMethod(m)
}

func NewMethod(name string, provider ClassProvider, body func(Value, ...Value) (Value, error)) Method {
	m := &method{name: name, body: body}
	m.class = provider.ClassWithName("Method")
	m.initialize()
	return m
}

func (method *method) Name() string {
	return method.name
}

func (method *method) Execute(self Value, args ...Value) (Value, error) {
	return method.body(self, args...)
}

func (method *method) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}
