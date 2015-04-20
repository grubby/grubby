package builtins

type genericClass struct {
	name string

	valueStub
	classStub
}

func NewGenericClass(name, superClass string, provider Provider) Class {
	c := &genericClass{name: name}
	c.initialize()
	c.setStringer(c.String)

	c.class = provider.ClassProvider().ClassWithName("Class")
	if superClass != "" {
		c.superClass = provider.ClassProvider().ClassWithName(superClass)
	}

	return c
}

func (c *genericClass) String() string {
	return c.name
}

func (c *genericClass) Name() string {
	return c.name
}

func (c *genericClass) New(provider Provider, args ...Value) (Value, error) {
	panic("unimplemented")
}
