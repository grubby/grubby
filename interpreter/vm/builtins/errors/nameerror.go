package errors

import "fmt"

type nameError struct {
	filename  string
	context   string
	className string
}

func NewNameError(name, context, className string) *nameError {
	return &nameError{
		filename:  name,
		context:   context,
		className: className,
	}
}

func (err *nameError) Error() string {
	return fmt.Sprintf("NameError: undefined local variable or method '%s' for %s:%s", err.filename, err.context, err.className)
}
