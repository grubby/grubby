package builtins

type ArrayClass struct {
	valueStub
	instanceMethods []Method
}

func NewArrayClass() Class {
	a := &ArrayClass{}
	a.class = NewClassValue()
	a.initialize()
	return a
}

func (klass *ArrayClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *ArrayClass) New(args ...Value) Value {
	a := &Array{}
	a.initialize()
	a.class = klass

	a.AddMethod(NewMethod("shift", func(self Value, args ...Value) (Value, error) {
		if len(a.members) == 0 {
			return Nil(), nil
		}

		val := a.members[0]
		a.members = a.members[1:]
		return val, nil
	}))

	a.AddMethod(NewMethod("unshift", func(self Value, args ...Value) (Value, error) {
		a.members = append([]Value{args[0]}, a.members[0:]...)
		return a, nil
	}))

	a.AddMethod(NewMethod("include?", func(self Value, args ...Value) (Value, error) {
		for _, m := range a.members {
			if m == args[0] {
				return NewTrueClass().(Class).New(), nil
			}
		}

		return NewFalseClass().(Class).New(), nil
	}))

	return a
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
