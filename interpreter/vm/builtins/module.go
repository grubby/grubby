package builtins

// abstract module interface
type Module interface {
	Name() string
	AddInstanceMethod(Method)
	InstanceMethods() []Method

	Value
}

// globlal Module class
type ModuleClass struct {
	valueStub
	instanceMethods []Method
}

func NewModuleClass() Class {
	c := &ModuleClass{}
	c.initialize()
	c.instanceMethods = make([]Method, 0)
	c.class = c
	return c
}

func (c ModuleClass) New(args ...Value) Value {
	return nil
}

func (c ModuleClass) Name() string {
	return "Module"
}

func (c ModuleClass) String() string {
	return "Module"
}

func (c *ModuleClass) AddInstanceMethod(m Method) {
	c.instanceMethods = append(c.instanceMethods, m)
}

func (c *ModuleClass) InstanceMethods() []Method {
	return c.instanceMethods
}

// user defined module type
type RubyModule struct {
	name string
	valueStub

	includedModules []Value
	instanceMethods []Method
}

func NewModule(name string) Module {
	c := &RubyModule{
		name:            name,
		includedModules: make([]Value, 0),
		instanceMethods: make([]Method, 0),
	}
	c.initialize()
	c.class = NewModuleClass() // FIXME: should only be one in existence
	c.AddMethod(NewMethod("include", func(self Value, args ...Value) (Value, error) {
		for _, module := range args {
			c.includedModules = append(c.includedModules, module)
		}

		return c, nil
	}))

	return c
}

func (m RubyModule) Name() string {
	return m.name
}

func (m *RubyModule) AddInstanceMethod(method Method) {
	m.instanceMethods = append(m.instanceMethods, method)
}

func (m *RubyModule) InstanceMethods() []Method {
	return m.instanceMethods
}
