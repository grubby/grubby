package builtins

type processClass struct {
	valueStub
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

type processValue struct {
	valueStub
}

func (class *processClass) New(args ...Value) Value {
	p := &processClass{}
	p.initialize()
	p.class = class

	return p
}
