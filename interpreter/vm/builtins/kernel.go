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
	k.setStringer(k.String)
	k.class = provider.ClassWithName("Module")

	k.AddMethod(NewNativeMethod("puts", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			os.Stdout.Write([]byte(arg.String() + "\n"))
		}

		return nil, nil
	}))

	k.AddMethod(NewNativeMethod("singleton_methods", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		methodsArray, err := provider.ClassWithName("Array").New(provider, singletonProvider)
		if err != nil {
			return nil, err
		}

		for _, method := range self.eigenclassMethods() {
			symbol := singletonProvider.SymbolWithName(method.Name())

			if symbol == nil {
				symbol = NewSymbol(method.Name(), provider)
				singletonProvider.AddSymbol(symbol)
			}

			methodsArray.(*Array).Append(symbol)
		}

		return methodsArray, nil
	}))

	return k
}

func (kernel *kernel) String() string {
	return "Kernel"
}

func (kernel *kernel) Name() string {
	return "Kernel"
}
