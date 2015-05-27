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

func NewFileClass(provider Provider) Class {
	f := &fileClass{}
	f.initialize()
	f.setStringer(f.String)
	f.class = provider.ClassProvider().ClassWithName("Class")
	f.superClass = provider.ClassProvider().ClassWithName("IO")

	f.SetConstant("ALT_SEPARATOR", provider.SingletonProvider().SingletonWithName("nil"))
	f.SetConstant("FNM_SYSCASE", NewFixnum(0, provider))

	f.AddMethod(NewNativeMethod("expand_path", provider, func(self Value, block Block, args ...Value) (Value, error) {
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

		return NewString(path, provider), nil
	}))
	f.AddMethod(NewNativeMethod("dirname", provider, func(self Value, block Block, args ...Value) (Value, error) {
		filename := args[0].(*StringValue).RawString()

		return NewString(filepath.Base(filename), provider), nil
	}))
	f.AddMethod(NewNativeMethod("exist?", provider, func(self Value, block Block, args ...Value) (Value, error) {

		filepath := args[0].(*StringValue).RawString()
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return provider.SingletonProvider().SingletonWithName("false"), nil
		} else {
			return provider.SingletonProvider().SingletonWithName("true"), nil
		}
	}))
	f.AddMethod(NewNativeMethod("join", provider, func(self Value, block Block, args ...Value) (Value, error) {
		pieces := make([]string, len(args))
		for _, str := range args {
			pieces = append(pieces, str.(*StringValue).RawString())
		}

		return NewString(filepath.Join(pieces...), provider), nil
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

func (file *fileClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, nil
}
