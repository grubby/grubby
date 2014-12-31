package builtins

func NewComparableModule(provider ClassProvider) Module {
	m := NewModule("Comparable", provider)
	m.AddMethod(NewMethod("<", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod("<=", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod("==", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod(">=", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod(">", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod("between?", provider, func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))

	return m
}
