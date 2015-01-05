package builtins

type ArrayClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewArrayClass(provider ClassProvider) Class {
	a := &ArrayClass{}
	a.class = provider.ClassWithName("Class")
	a.superClass = provider.ClassWithName("Object")
	a.initialize()
	return a
}

func (klass *ArrayClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *ArrayClass) New(provider ClassProvider, args ...Value) (Value, error) {
	a := &Array{}
	a.initialize()
	a.class = klass

	a.AddMethod(NewNativeMethod("shift", provider, func(self Value, args ...Value) (Value, error) {
		if len(a.members) == 0 {
			// FIXME: this should return the singleton for nil
			return provider.ClassWithName("Nil").New(provider)
		}

		val := a.members[0]
		a.members = a.members[1:]
		return val, nil
	}))

	a.AddMethod(NewNativeMethod("unshift", provider, func(self Value, args ...Value) (Value, error) {
		a.members = append([]Value{args[0]}, a.members[0:]...)
		return a, nil
	}))

	a.AddMethod(NewNativeMethod("include?", provider, func(self Value, args ...Value) (Value, error) {
		for _, m := range a.members {
			if m == args[0] {
				// FIXME: needs a singleton provider? (someway to inject singletons)
				return provider.ClassWithName("True").New(provider)
			}
		}

		// FIXME: needs singletons, as above
		return provider.ClassWithName("False").New(provider)
	}))

	return a, nil
}

func (array *ArrayClass) Name() string {
	return "Array"
}

func (array *ArrayClass) String() string {
	return "Array"
}

type Array struct {
	valueStub
	members []Value
}

func (array *Array) Append(v Value) {
	array.members = append(array.members, v)
}

func (array *Array) Members() []Value {
	return array.members
}

func (array *Array) String() string {
	return "Array"
}
