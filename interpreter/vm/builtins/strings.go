package builtins

type StringValue struct {
	value string
	valueStub
}

func NewString(val string) Value {
	return &StringValue{value: val}
}

func (stringValue *StringValue) String() string {
	return stringValue.value
}
