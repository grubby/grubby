package builtins

import "os"

type kernel struct {
	valueStub
	classStub
	instanceMethods []Method
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

func (kernel *kernel) AddInstanceMethod(m Method) {
	kernel.instanceMethods = append(kernel.instanceMethods, m)
}

// FIXME: this should be in a module_stub
func (kernel *kernel) InstanceMethods() []Method {
	return kernel.instanceMethods
}
