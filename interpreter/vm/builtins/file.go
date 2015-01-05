package builtins

import (
	"os"
	"path/filepath"
)

type fileClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewFileClass(provider ClassProvider) Class {
	f := &fileClass{}
	f.initialize()
	f.class = provider.ClassWithName("Class")
	f.superClass = provider.ClassWithName("IO")
	f.AddMethod(NewNativeMethod("expand_path", provider, func(self Value, args ...Value) (Value, error) {
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

		return NewString(path, provider), nil
	}))

	f.AddMethod(NewNativeMethod("dirname", provider, func(self Value, args ...Value) (Value, error) {
		filename := args[0].(*StringValue).String()

		return NewString(filepath.Base(filename), provider), nil
	}))

	return f
}

func (file *fileClass) AddInstanceMethod(m Method) {
	file.instanceMethods = append(file.instanceMethods, m)
}

func (file *fileClass) Name() string {
	return "File"
}

func (file *fileClass) String() string {
	return "File"
}

func (file *fileClass) New(provider ClassProvider, args ...Value) Value {
	return nil
}
