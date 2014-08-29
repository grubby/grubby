package builtins

type ArrayClass struct {
	valueStub
}

func NewArrayClass() Value {
	a := &ArrayClass{}
	a.class = NewClassValue().(Class)
	a.initialize()
	return a
}

func (klass *ArrayClass) New() Value {
	a := &Array{}
	a.initialize()
	a.class = klass

	a.AddMethod(NewMethod("shift", func(args ...Value) (Value, error) {
		if len(a.members) == 0 {
			return Nil(), nil
		}

		val := a.members[0]
		a.members = a.members[1:]
		return val, nil
	}))
	a.AddMethod(NewMethod("include?", func(args ...Value) (Value, error) {
		for _, m := range a.members {
			if m == args[0] {
				return NewTrueClass().(Class).New(), nil
			}
		}

		return NewFalseClass().(Class).New(), nil
	}))

	return a
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
