package builtins

import (
	"os"
	"path/filepath"
)

type fileClass struct {
	valueStub
}

func NewFileClass() Value {
	f := &fileClass{}
	f.initialize()
	f.class = NewClassValue() // FIXME: this should be a global reference
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

func (file *fileClass) String() string {
	return "File"
}

func (file *fileClass) New() Value {
	return file
}
