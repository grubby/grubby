package builtins

func NewComparableModule() Module {
	m := NewModule("Comparable")
	m.AddMethod(NewMethod("<", func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod("<=", func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod("==", func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod(">=", func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod(">", func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))
	m.AddMethod(NewMethod("between?", func(self Value, args ...Value) (Value, error) {
		return nil, nil
	}))

	return m
}
