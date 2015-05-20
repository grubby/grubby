package builtins

import "fmt"

type noMethodError struct {
	method    string
	context   string
	className string
	callstack string
	valueStub
}

func NewNoMethodError(name, context, className, callstack string) *noMethodError {
	return &noMethodError{
		method:    name,
		context:   context,
		className: className,
		callstack: callstack,
	}
}

func (err *noMethodError) Error() string {
	return fmt.Sprintf("NoMethodError: undefined method '%s' for %s:%s\n%s", err.method, err.context, err.className, err.callstack)
}

func (err *noMethodError) String() string {
	return err.Error()
}
