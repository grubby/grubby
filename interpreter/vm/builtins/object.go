package builtins

type objectClass struct {
	valueStub
}

func NewGlobalObjectClass() Class {
	o := &objectClass{}
	o.initialize()
	o.class = NewClassValue() // FIXME: this should be set to the global reference
	return o
}

func (obj *objectClass) String() string {
	return "Object"
}

func (obj *objectClass) Name() string {
	return "Object"
}

type object struct {
	valueStub
}

func (obj *objectClass) New(args ...Value) Value {
	o := &object{}
	o.initialize()
	o.class = obj

	return o
}
