package builtins

type processClass struct {
	valueStub
	instanceMethods []Method
}

func NewProcessClass() Class {
	f := &processClass{}
	f.initialize()
	f.class = NewClassValue().(Class)
	return f
}

func (c *processClass) String() string {
	return "Process"
}

func (c *processClass) Name() string {
	return "Process"
}

func (c *processClass) AddInstanceMethod(method Method) {
	c.instanceMethods = append(c.instanceMethods, method)
}

type processValue struct {
	valueStub
}

func (class *processClass) New(args ...Value) Value {
	p := &processClass{}
	p.initialize()
	p.class = class

	return p
}
