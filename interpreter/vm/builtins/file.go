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

func NewFileClass(provider ClassProvider, singletonProvider SingletonProvider) Class {
	f := &fileClass{}
	f.initialize()
	f.setStringer(f.String)
	f.class = provider.ClassWithName("Class")
	f.superClass = provider.ClassWithName("IO")

	f.SetConstant("ALT_SEPARATOR", singletonProvider.SingletonWithName("nil"))
	f.SetConstant("FNM_SYSCASE", NewFixnum(0, provider, singletonProvider))

	f.AddMethod(NewNativeMethod("expand_path", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		arg1 := args[0].(*StringValue).RawString()

		if arg1[0] == '~' {
			arg1 = os.Getenv("HOME") + arg1[1:]
		}

		var path string
		if len(args) == 2 {
			path = filepath.Join(args[1].(*StringValue).RawString(), arg1)
		} else {
			path, _ = filepath.Abs(arg1)
		}

		return NewString(path, provider, singletonProvider), nil
	}))

	f.AddMethod(NewNativeMethod("dirname", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		filename := args[0].(*StringValue).RawString()

		return NewString(filepath.Base(filename), provider, singletonProvider), nil
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

func (file *fileClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, nil
}
