package builtins

import (
	"errors"
	"fmt"

	"github.com/grubby/grubby/ast"
)

type methodArg struct {
	Name  string
	Value Value
}

type RubyMethod struct {
	valueStub
	name       string
	lineNumber int
	visibility MethodVisibility

	args []ast.MethodParam

	body func(self Value, method *RubyMethod) (Value, error)

	invocationArgs  []methodArg
	unevaluatedBody []ast.Node

	provider      Provider
	evaluator     ArgEvaluator
	classProvider ClassProvider
	stackProvider StackProvider
}

func NewRubyMethod(
	name string,
	lineNumber int,
	args []ast.MethodParam,
	rubyBody []ast.Node,
	provider Provider,
	evaluator ArgEvaluator,
	body func(self Value, method *RubyMethod) (Value, error),
) Method {
	m := &RubyMethod{
		name:            name,
		lineNumber:      lineNumber,
		body:            body,
		args:            args,
		visibility:      Public,
		provider:        provider,
		classProvider:   provider.ClassProvider(),
		stackProvider:   provider.StackProvider(),
		evaluator:       evaluator,
		unevaluatedBody: rubyBody,
	}
	m.class = provider.ClassProvider().ClassWithName("Method")
	m.initialize()
	m.setStringer(m.String)
	return m
}

func NewPrivateRubyMethod(
	name string,
	lineNumber int,
	args []ast.MethodParam,
	rubyBody []ast.Node,
	provider Provider,
	evaluator ArgEvaluator,
	body func(self Value, method *RubyMethod) (Value, error),
) Method {
	m := &RubyMethod{
		name:            name,
		lineNumber:      lineNumber,
		body:            body,
		args:            args,
		visibility:      Private,
		evaluator:       evaluator,
		provider:        provider,
		classProvider:   provider.ClassProvider(),
		stackProvider:   provider.StackProvider(),
		unevaluatedBody: rubyBody,
	}
	m.class = provider.ClassProvider().ClassWithName("Method")
	m.initialize()
	m.setStringer(m.String)
	return m
}

func (method *RubyMethod) Name() string {
	return method.name
}

func (method *RubyMethod) IsPrivate() bool {
	return method.visibility == Private
}

func (method *RubyMethod) IsPublic() bool {
	return method.visibility == Public
}

func (method *RubyMethod) IsProtected() bool {
	return method.visibility == Protected
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

		if arg.IsSplat {
			argValue, err = method.classProvider.ClassWithName("Array").New(method.provider)
			if err != nil {
				return nil, err
			}

			for i := index; i < len(args); i++ {
				argValue.(*Array).Append(args[i])
			}
		} else {
			if index >= len(args) {
				if arg.DefaultValue != nil {
					argValue, err = method.evaluator.EvaluateArgInContext(arg.DefaultValue, self)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, errors.New(fmt.Sprintf("expected to invoke a method '%s' on '%s' with %d args, but we were only provided %d", method, self, len(method.args), len(args)))
				}
			} else {
				argValue = args[index]
			}
		}

		argument := methodArg{
			Name:  arg.Name,
			Value: argValue,
		}
		method.invocationArgs = append(method.invocationArgs, argument)
	}

	method.stackProvider.UnshiftStackFrame(method.name, "fixme -- method name goes here", method.lineNumber)
	defer method.stackProvider.ShiftStackFrame()
	defer func() { method.invocationArgs = nil }()

	return method.body(self, method)
}

// FIXME: in order to fix this, the method needs to know what it is attached to
func (method *RubyMethod) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}

func (method *RubyMethod) Visibility() MethodVisibility {
	return method.visibility
}

func (method *RubyMethod) SetVisibility(visibility MethodVisibility) {
	method.visibility = visibility
}

func (method *RubyMethod) methodBody() func(self Value, block Block, args ...Value) (Value, error) {
	return method.Execute
}
