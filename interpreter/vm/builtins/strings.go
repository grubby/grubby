package builtins

type StringClass struct {
	valueStub
	instanceMethods []Method
}

func NewStringClass() Class {
	s := &StringClass{}
	s.initialize()
	s.class = NewClassValue() // FIXME: should be created once globally
	return s
}

func (c *StringClass) String() string {
	return "String"
}

func (c *StringClass) Name() string {
	return "String"
}

func (class *StringClass) New(args ...Value) Value {
	str := &StringValue{}
	str.initialize()
	str.class = class
	str.AddMethod(NewMethod("+", func(args ...Value) (Value, error) {
		arg := args[0].(*StringValue)
		return NewString(str.value + arg.value), nil
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

func NewString(str string) Value {
	s := NewStringClass().New() // FIXME: needs a reference to global class cache
	s.(*StringValue).value = str
	return s
}
