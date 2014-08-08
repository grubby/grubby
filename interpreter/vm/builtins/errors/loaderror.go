package errors

import "fmt"

type loadError struct {
	filename string
}

func NewLoadError(name string) *loadError {
	return &loadError{filename: name}
}

func (err *loadError) Error() string {
	return fmt.Sprintf("LoadError: cannot load such file -- %s", err.filename)
}
