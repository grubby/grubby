package builtins

import "fmt"

type method struct {
	name            string
	methods         []Method
	private_methods []Method
}

type Method interface {
	Name() string
	Value
}

func NewMethod(name string) Method {
	return &method{name: name}
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

func (method *method) AddMethod(Method) {

}

func (method *method) AddPrivateMethod(Method) {

}

func (method *method) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}
