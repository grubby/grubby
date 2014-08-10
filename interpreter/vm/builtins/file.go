package builtins

import (
	"os"
	"path/filepath"
)

type file struct {
	valueStub
}

func NewFileClass() Value {
	f := &file{}
	f.initialize()
	f.AddMethod(NewMethod("expand_path", func(args ...Value) (Value, error) {
		arg := args[0].(*StringValue).String()

		if arg[0] == '~' {
			arg = os.Getenv("HOME") + arg[1:]
		}

		path, _ := filepath.Abs(arg)
		return NewString(path), nil
	}))

	return f
}

func (file *file) String() string {
	return "File"
}

func (file *file) New() Value {
	return file
}
