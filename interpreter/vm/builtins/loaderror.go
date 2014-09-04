package builtins

import (
	"fmt"
)

type loadError struct {
	filename string
	valueStub
}

func NewLoadError(name string) *loadError {
	return &loadError{filename: name}
}

func (err *loadError) String() string {
	return "LoadError"
}

func (err *loadError) Error() string {
	return fmt.Sprintf("LoadError: cannot load such file -- %s", err.filename)
}
