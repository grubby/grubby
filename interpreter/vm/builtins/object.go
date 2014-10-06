package builtins

type objectClass struct {
	valueStub
	instanceMethods []Method
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

func (obj *objectClass) AddInstanceMethod(method Method) {
	obj.instanceMethods = append(obj.instanceMethods, method)
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
