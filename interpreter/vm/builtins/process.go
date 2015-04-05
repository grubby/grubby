package builtins

type processModule struct {
	valueStub
	classStub
	moduleStub
}

func NewProcessModule(provider Provider) Module {
	f := &processModule{}
	f.initialize()
	f.setStringer(f.String)
	f.class = provider.ClassProvider().ClassWithName("Module")
	return f
}

func (c *processModule) String() string {
	return "Process"
}

func (c *processModule) Name() string {
	return "Process"
}
