package builtins

func NewComparableModule(provider ClassProvider, singletonProvider SingletonProvider) Module {
	m := NewModule("Comparable", provider, singletonProvider)
	m.AddMethod(NewNativeMethod("<", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod("<=", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod("==", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod(">=", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod(">", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod("between?", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		return nil, nil
	}))

	return m
}
