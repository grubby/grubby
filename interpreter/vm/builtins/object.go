package builtins

type ObjectClass struct {
	valueStub
	classStub

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

type object struct {
	valueStub
}

func (obj *ObjectClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	o := &object{}
	o.initialize()
	o.class = obj

	return o, nil
}
