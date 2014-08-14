package builtins

type objectClass struct {
	valueStub
}

func NewGlobalObjectClass() Value {
	o := &objectClass{}
	o.initialize()
	o.class = NewClassValue().(Class)
	return o
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
