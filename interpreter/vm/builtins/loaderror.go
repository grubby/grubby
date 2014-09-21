package builtins

import (
	"fmt"
)

type loadError struct {
	filename  string
	callstack string
	valueStub
}

func NewLoadError(name string, callstack string) *loadError {
	return &loadError{filename: name, callstack: callstack}
}

func (err *loadError) String() string {
	return "LoadError"
}

func (err *loadError) Error() string {
	return fmt.Sprintf("LoadError: cannot load such file -- %s\n%s", err.filename, err.callstack)
}
