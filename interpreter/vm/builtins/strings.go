package builtins

type StringValue struct {
	value           string
	methods         []Method
	private_methods []Method
}

func NewString(val string) Value {
	return &StringValue{value: val}
}

func (stringValue *StringValue) Methods() []Method {
	return stringValue.methods
}

func (stringValue *StringValue) PrivateMethods() []Method {
	return stringValue.private_methods
}

func (stringValue *StringValue) AddMethod(m Method) {
	stringValue.methods = append(stringValue.methods, m)
}

func (stringValue *StringValue) AddPrivateMethod(m Method) {
	stringValue.private_methods = append(stringValue.private_methods, m)
}

func (stringValue *StringValue) String() string {
	return stringValue.value
}

func (stringValue *StringValue) Class() Class {
	return nil
}
