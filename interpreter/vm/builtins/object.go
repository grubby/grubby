package builtins

type object struct {
	methods         []Method
	private_methods []Method
}

func NewGlobalObject() Value {
	return &object{}
}

func (obj *object) Methods() []Method {
	return obj.methods
}

func (obj *object) PrivateMethods() []Method {
	return obj.private_methods
}

func (object *object) AddMethod(Method) {

}

func (object *object) AddPrivateMethod(Method) {

}
