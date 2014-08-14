package builtins

type StringValue struct {
	value string
	valueStub
}

func NewString(val string) Value {
	s := &StringValue{value: val}
	s.initialize()
	return s
}

func (stringValue *StringValue) String() string {
	return stringValue.value
}
