package builtins

type genericModule struct {
	name string

	valueStub
	moduleStub
}

func NewGenericModule(name string, provider Provider) Module {
	m := &genericModule{name: name}
	m.initialize()
	m.setStringer(m.String)

	m.class = provider.ClassProvider().ClassWithName("Module")

	return m
}

func (c *genericModule) String() string {
	return c.name
}

func (c *genericModule) Name() string {
	return c.name
}

func (c *genericModule) New(provider Provider, args ...Value) (Value, error) {
	panic("unimplemented")
}
