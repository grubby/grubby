package builtins

type StringClass struct {
	valueStub
	classStub

	provider ClassProvider

	instanceMethods []Method
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

func (class *StringClass) New(provider ClassProvider, args ...Value) Value {
	str := &StringValue{}
	str.initialize()
	str.class = class
	str.AddMethod(NewMethod("+", provider, func(self Value, args ...Value) (Value, error) {
		arg := args[0].(*StringValue)
		return NewString(str.value+arg.value, class.provider), nil
	}))

	return str
}

func (class *StringClass) AddInstanceMethod(method Method) {
	class.instanceMethods = append(class.instanceMethods, method)
}

type StringValue struct {
	value string
	valueStub
}

func (stringValue *StringValue) String() string {
	return stringValue.value
}

func NewString(str string, provider ClassProvider) Value {
	s := provider.ClassWithName("String").New(provider)
	s.(*StringValue).value = str
	return s
}
