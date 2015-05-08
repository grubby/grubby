package builtins

import "fmt"

type nameError struct {
	filename  string
	context   string
	className string
	callStack string
	valueStub
}

func NewNameError(name, context, className string, callStack string) *nameError {
	return &nameError{
		filename:  name,
		context:   context,
		className: className,
		callStack: callStack,
	}
}

func (err *nameError) Error() string {
	return fmt.Sprintf("NameError: undefined local variable or method '%s' for %s:%s\n%s", err.filename, err.context, err.className, err.callStack)
}

func (err *nameError) String() string {
	return "NameError"
}
