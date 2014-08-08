package builtins

type object struct {
	methods         []Method
	private_methods []Method
}

func NewGlobalObjectClass() Value {
	return &object{}
}

func (obj *object) Methods() []Method {
	return obj.methods
}

func (obj *object) PrivateMethods() []Method {
	return obj.private_methods
}

func (object *object) AddMethod(m Method) {
	object.methods = append(object.methods, m)
}

func (object *object) AddPrivateMethod(m Method) {
	object.private_methods = append(object.private_methods, m)
}

func (obj *object) String() string {
	return "Object"
}

func (obj *object) New() Value {
	return obj
}

func (obj *object) Class() Class {
	return obj
}

func (obj *object) SuperClass() Class {
	return nil
}
