package builtins

func NewComparableModule(provider ClassProvider) Module {
	m := NewModule("Comparable", provider)
	m.AddMethod(NewNativeMethod("<", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod("<=", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod("==", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod(">=", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod(">", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewNativeMethod("between?", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))

	return m
}
