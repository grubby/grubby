package builtins

type BasicObjectClass struct {
	valueStub
	classStub

	provider Provider
}

func NewBasicObjectClass(provider Provider) Class {
	o := &BasicObjectClass{}
	o.initialize()
	o.setStringer(o.String)
	o.provider = provider
	return o
}

func (c *BasicObjectClass) SetSuperClass() {
	class := c.provider.ClassProvider().ClassWithName("Class")
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

func (obj *BasicObjectClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, nil
}
