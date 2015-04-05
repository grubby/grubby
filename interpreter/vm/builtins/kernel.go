package builtins

import "os"

type kernel struct {
	valueStub
	classStub
	moduleStub
}

func NewGlobalKernelModule(provider Provider) Module {
	k := &kernel{}
	k.initialize()
	k.setStringer(k.String)
	k.class = provider.ClassProvider().ClassWithName("Module")

	k.AddMethod(NewNativeMethod("puts", provider, func(self Value, block Block, args ...Value) (Value, error) {
		for _, arg := range args {
			var stringified string
			to_s := arg.Method("to_s")
			if to_s != nil {
				value, err := to_s.Execute(arg, nil)
				if err != nil {
					return nil, err
				}
				stringified = value.String()
			} else {
				stringified = arg.String()
			}

			os.Stdout.Write([]byte(stringified + "\n"))
		}

		return nil, nil
	}))

	k.AddMethod(NewNativeMethod("singleton_methods", provider, func(self Value, block Block, args ...Value) (Value, error) {
		methodsArray, err := provider.ClassProvider().ClassWithName("Array").New(provider)
		if err != nil {
			return nil, err
		}

		for _, method := range self.eigenclassMethods() {
			if !method.IsPublic() {
				continue
			}

			symbol := provider.SingletonProvider().SymbolWithName(method.Name())

			if symbol == nil {
				symbol = NewSymbol(method.Name(), provider)
				provider.SingletonProvider().AddSymbol(symbol)
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
