package builtins

type stringValue struct {
	value           string
	methods         []Method
	private_methods []Method
}

func NewString(val string) Value {
	return &stringValue{value: val}
}

func (stringValue *stringValue) Methods() []Method {
	return stringValue.methods
}

func (stringValue *stringValue) PrivateMethods() []Method {
	return stringValue.private_methods
}

func (stringValue *stringValue) AddMethod(m Method) {
	stringValue.methods = append(stringValue.methods, m)
}

func (stringValue *stringValue) AddPrivateMethod(m Method) {
	stringValue.private_methods = append(stringValue.private_methods, m)
}

func (stringValue *stringValue) String() string {
	return stringValue.value
}
