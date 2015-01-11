package builtins

type ArrayClass struct {
	valueStub
	classStub
	instanceMethods []Method

	singletonProvider SingletonProvider
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

func (klass *ArrayClass) New(classProvider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	a := &Array{}
	a.initialize()
	a.class = klass

	a.AddMethod(NewNativeMethod("shift", classProvider, singletonProvider, func(self Value, args ...Value) (Value, error) {
		if len(a.members) == 0 {
			return singletonProvider.SingletonWithName("nil"), nil
		}

		val := a.members[0]
		a.members = a.members[1:]
		return val, nil
	}))

	a.AddMethod(NewNativeMethod("unshift", classProvider, singletonProvider, func(self Value, args ...Value) (Value, error) {
		a.members = append([]Value{args[0]}, a.members[0:]...)
		return a, nil
	}))

	a.AddMethod(NewNativeMethod("include?", classProvider, singletonProvider, func(self Value, args ...Value) (Value, error) {
		for _, m := range a.members {
			if m == args[0] {
				return singletonProvider.SingletonWithName("true"), nil
			}
		}

		return singletonProvider.SingletonWithName("false"), nil
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
