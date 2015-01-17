package builtins

import "os"

type kernel struct {
	valueStub
	classStub
	moduleStub
}

func NewGlobalKernelModule(provider ClassProvider, singletonProvider SingletonProvider) Module {
	k := &kernel{}
	k.initialize()
	k.class = provider.ClassWithName("Module")

	k.AddMethod(NewNativeMethod("puts", provider, singletonProvider, func(self Value, args ...Value) (Value, error) {
		for _, arg := range args {
			os.Stdout.Write([]byte(arg.String() + "\n"))
		}

		return nil, nil
	}))

	return k
}

func (kernel *kernel) String() string {
	return "Kernel"
}

func (kernel *kernel) Name() string {
	return "Kernel"
}
