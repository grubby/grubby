package builtins

type StringClass struct {
	valueStub
	classStub

	provider ClassProvider
}

func NewStringClass(provider ClassProvider) Class {
	s := &StringClass{}
	s.initialize()
	s.provider = provider
	s.class = provider.ClassWithName("Class")
	s.superClass = provider.ClassWithName("Object")
	return s
}

func (c *StringClass) String() string {
	return "String"
}

func (c *StringClass) Name() string {
	return "String"
}

func (class *StringClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	str := &StringValue{}
	str.initialize()
	str.class = class
	str.AddMethod(NewNativeMethod("+", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		arg := args[0].(*StringValue)
		return NewString(str.value+arg.value, provider, singletonProvider), nil
	}))
	str.AddMethod(NewNativeMethod("==", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		asStr, ok := args[0].(*StringValue)
		if !ok {
			return singletonProvider.SingletonWithName("false"), nil
		}

		selfAsStr := self.(*StringValue)
		if selfAsStr.value == asStr.value {
			return singletonProvider.SingletonWithName("true"), nil
		} else {
			return singletonProvider.SingletonWithName("false"), nil
		}
	}))

	return str, nil
}

type StringValue struct {
	value string
	valueStub
}

func (stringValue *StringValue) String() string {
	return stringValue.value
}

func NewString(str string, provider ClassProvider, singletonProvider SingletonProvider) Value {
	s, _ := provider.ClassWithName("String").New(provider, singletonProvider)
	s.(*StringValue).value = str
	return s
}
