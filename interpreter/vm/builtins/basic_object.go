package builtins

type BasicObjectClass struct {
	valueStub
	classStub
	instanceMethods []Method

	provider ClassProvider
}

func NewBasicObjectClass(provider ClassProvider) Class {
	o := &BasicObjectClass{}
	o.initialize()
	o.provider = provider
	return o
}

func (c *BasicObjectClass) SetSuperClass() {
	class := c.provider.ClassWithName("Class")
	if class == nil {
		panic("Expected Class class to exist")
	}
	c.class = class
}

func (obj *BasicObjectClass) String() string {
	return "BasicObject"
}

func (obj *BasicObjectClass) Name() string {
	return "BasicObject"
}

func (obj *BasicObjectClass) AddInstanceMethod(method Method) {
	obj.instanceMethods = append(obj.instanceMethods, method)
}

func (obj *BasicObjectClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, nil
}
