package builtins

type processModule struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewProcessModule(provider ClassProvider) Module {
	f := &processModule{}
	f.initialize()
	f.class = provider.ClassWithName("Module")
	return f
}

func (c *processModule) String() string {
	return "Process"
}

func (c *processModule) Name() string {
	return "Process"
}

func (c *processModule) AddInstanceMethod(method Method) {
	c.instanceMethods = append(c.instanceMethods, method)
}

type processValue struct {
	valueStub
}

func (module *processModule) New(args ...Value) Value {
	return nil
}

// FIXME: this should be in a module_stub
func (module *processModule) InstanceMethods() []Method {
	return module.instanceMethods
}
