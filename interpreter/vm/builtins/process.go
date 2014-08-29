package builtins

type process struct {
	valueStub
}

func NewProcessClass() Value {
	f := &process{}
	f.initialize()
	f.class = NewClassValue().(Class)
	return f
}
