package builtins

import "fmt"

type noMethodError struct {
	method    string
	context   string
	className string
	valueStub
}

func NewNoMethodError(name, context, className string) *noMethodError {
	return &noMethodError{
		method:    name,
		context:   context,
		className: className,
	}
}

func (err *noMethodError) Error() string {
	return fmt.Sprintf("NoMethodError: undefined method '%s' for %s:%s", err.method, err.context, err.className)
}
