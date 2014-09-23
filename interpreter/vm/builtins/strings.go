package builtins

type StringClass struct {
	valueStub
}

func NewStringClass() Class {
	s := &StringClass{}
	s.initialize()
	s.class = NewClassValue() // FIXME: should be created once globally
	return s
}

func (class *StringClass) New(args ...Value) Value {
	str := &StringValue{}
	str.initialize()
	str.class = class
	return str
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
