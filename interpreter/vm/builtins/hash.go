package builtins

type HashClass struct {
	valueStub
	instanceMethods []Method
}

func NewHashClass() Class {
	a := &HashClass{}
	a.class = NewClassValue()
	a.initialize()
	return a
}

func (klass *HashClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *HashClass) New(args ...Value) Value {
	a := &Hash{}
	a.initialize()
	a.class = klass
	a.hash = make(map[Value]Value)

	a.AddMethod(NewMethod("keys", func(self Value, args ...Value) (Value, error) {
		keys := NewArrayClass().New().(*Array)
		for key := range self.(*Hash).hash {
			keys.Append(key)
		}

		return keys, nil
	}))
	a.AddMethod(NewMethod("values", func(self Value, args ...Value) (Value, error) {
		values := NewArrayClass().New().(*Array)
		for key := range self.(*Hash).hash {
			values.Append(self.(*Hash).hash[key])
		}

		return values, nil
	}))

	a.AddMethod(NewMethod("[]=", func(self Value, args ...Value) (Value, error) {
		self.(*Hash).hash[args[0]] = args[1]
		return args[1], nil
	}))

	return a
}

func (hash *HashClass) Name() string {
	return "Hash"
}

func (hash *HashClass) String() string {
	return "Hash"
}

type Hash struct {
	hash map[Value]Value
	valueStub
}

func (hash *Hash) String() string {
	return "Hash"
}

func (hash *Hash) Add(key, value Value) {
	hash.hash[key] = value
}
