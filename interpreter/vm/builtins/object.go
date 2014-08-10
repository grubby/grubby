package builtins

type objectClass struct {
	valueStub
}

func NewGlobalObjectClass() Value {
	return &objectClass{}
}

func (obj *objectClass) String() string {
	return "Object"
}

type object struct {
	valueStub
}

func (obj *objectClass) New() Value {
	o := &object{}
	o.initialize()
	o.class = obj

	return o
}
