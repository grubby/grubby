package builtins

type HashClass struct {
	valueStub
	classStub

	provider ClassProvider

	instanceMethods []Method
}

func NewHashClass(provider ClassProvider) Class {
	a := &HashClass{}
	a.initialize()
	a.setStringer(a.String)
	a.class = provider.ClassWithName("Class")
	a.superClass = provider.ClassWithName("Object")
	a.provider = provider
	return a
}

func (klass *HashClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *HashClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	a := &Hash{}
	a.initialize()
	a.setStringer(a.String)
	a.class = klass
	a.hash = make(map[Value]Value)

	a.AddMethod(NewNativeMethod("keys", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		o, _ := klass.provider.ClassWithName("Array").New(provider, singletonProvider)
		keys := o.(*Array)
		for key := range self.(*Hash).hash {
			keys.Append(key)
		}

		return keys, nil
	}))
	a.AddMethod(NewNativeMethod("values", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		o, _ := klass.provider.ClassWithName("Array").New(provider, singletonProvider)
		values := o.(*Array)
		for key := range self.(*Hash).hash {
			values.Append(self.(*Hash).hash[key])
		}

		return values, nil
	}))

	a.AddMethod(NewNativeMethod("[]=", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		self.(*Hash).hash[args[0]] = args[1]
		return args[1], nil
	}))

	return a, nil
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
