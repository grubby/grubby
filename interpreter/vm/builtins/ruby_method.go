package builtins

import (
	"fmt"

	"github.com/grubby/grubby/ast"
)

type methodArg struct {
	Name  string
	Value Value
}

type RubyMethod struct {
	valueStub
	name     string
	argNames []string

	body func(self Value, method *RubyMethod) (Value, error)

	invocationArgs  []methodArg
	unevaluatedBody []ast.Node
}

func NewRubyMethod(name string, argNames []string, rubyBody []ast.Node, provider ClassProvider, body func(self Value, method *RubyMethod) (Value, error)) Method {
	m := &RubyMethod{
		name:            name,
		argNames:        argNames,
		body:            body,
		unevaluatedBody: rubyBody,
	}
	m.class = provider.ClassWithName("Method")
	m.initialize()
	return m
}

func (method *RubyMethod) Name() string {
	return method.name
}

func (method *RubyMethod) Args() []methodArg {
	if method.invocationArgs == nil {
		panic("asked for args when a method was not being invoked")
	}

	return method.invocationArgs
}

func (method *RubyMethod) Body() []ast.Node {
	return method.unevaluatedBody
}

func (method *RubyMethod) Execute(self Value, args ...Value) (Value, error) {
	method.invocationArgs = make([]methodArg, 0, len(args))
	for index, arg := range args {
		arg := methodArg{
			Name:  method.argNames[index],
			Value: arg,
		}
		method.invocationArgs = append(method.invocationArgs, arg)
	}
	defer func() {
		method.invocationArgs = nil
	}()

	return method.body(self, method)
}

// FIXME: in order to fix this, the method needs to know "self"
func (method *RubyMethod) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}
