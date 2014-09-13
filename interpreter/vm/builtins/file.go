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
		arg1 := args[0].(*StringValue).String()

		if arg1[0] == '~' {
			arg1 = os.Getenv("HOME") + arg1[1:]
		}

		var path string
		if len(args) == 2 {
			path = filepath.Join(args[1].(*StringValue).String(), arg1)
		} else {
			path, _ = filepath.Abs(arg1)
		}

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
