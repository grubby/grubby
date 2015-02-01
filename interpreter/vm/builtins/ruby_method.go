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
	name string

	args []ast.MethodParam

	body func(self Value, method *RubyMethod) (Value, error)

	invocationArgs  []methodArg
	unevaluatedBody []ast.Node

	evaluator ArgEvaluator
}

func NewRubyMethod(
	name string,
	args []ast.MethodParam,
	rubyBody []ast.Node,
	provider ClassProvider,
	evaluator ArgEvaluator,
	body func(self Value, method *RubyMethod) (Value, error),
) Method {
	m := &RubyMethod{
		name:            name,
		body:            body,
		args:            args,
		evaluator:       evaluator,
		unevaluatedBody: rubyBody,
	}
	m.class = provider.ClassWithName("Method")
	m.initialize()
	m.setStringer(m.String)
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

func (method *RubyMethod) Execute(self Value, block Block, args ...Value) (Value, error) {
	method.invocationArgs = make([]methodArg, 0, len(args))
	for index, arg := range method.args {

		var (
			argValue Value
			err      error
		)

		if index >= len(args) {
			if arg.DefaultValue != nil {
				argValue, err = method.evaluator.EvaluateArgInContext(arg.DefaultValue, self)
				if err != nil {
					return nil, err
				}
			} else {
				panic("whoops")
			}
		} else {
			argValue = args[index]
		}

		argument := methodArg{
			Name:  arg.Name.Name,
			Value: argValue,
		}
		method.invocationArgs = append(method.invocationArgs, argument)
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
