package builtins

type ObjectClass struct {
	valueStub
	classStub
	instanceMethods []Method

	provider ClassProvider
}

func NewGlobalObjectClass(provider ClassProvider) Class {
	o := &ObjectClass{}
	o.initialize()
	o.provider = provider
	return o
}

func (c *ObjectClass) SetSuperClass() {
	class := c.provider.ClassWithName("Class")
	if class == nil {
		panic("Expected Class class to exist")
	}

	superClass := c.provider.ClassWithName("BasicObject")
	if superClass == nil {
		panic("Expected BasicObject class to exist")
	}

	c.class = class
	c.superClass = superClass
}

func (obj *ObjectClass) String() string {
	return "Object"
}

func (obj *ObjectClass) Name() string {
	return "Object"
}

func (obj *ObjectClass) AddInstanceMethod(method Method) {
	obj.instanceMethods = append(obj.instanceMethods, method)
}

type object struct {
	valueStub
}

func (obj *ObjectClass) New(provider ClassProvider, args ...Value) (Value, error) {
	o := &object{}
	o.initialize()
	o.class = obj

	return o, nil
}
