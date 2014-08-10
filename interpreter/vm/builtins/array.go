package builtins

type Array struct {
	members         []Value
	methods         []Method
	private_methods []Method
}

func NewArrayClass() Value {
	return &Array{}
}

func (array *Array) Methods() []Method {
	return array.methods
}

func (array *Array) PrivateMethods() []Method {
	return array.private_methods
}

func (array *Array) AddMethod(m Method) {
	array.methods = append(array.methods, m)
}

func (array *Array) AddPrivateMethod(m Method) {
	array.private_methods = append(array.private_methods, m)
}

func (array *Array) String() string {
	return "Array"
}

func (array *Array) New() Value {
	return array
}

func (array *Array) Class() Class {
	return array
}

func (array *Array) SuperClass() Class {
	return nil
}

func (array *Array) Append(v Value) {
	array.members = append(array.members, v)
}

func (array *Array) Members() []Value {
	return array.members
}
