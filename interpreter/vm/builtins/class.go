package builtins

type Class interface {
	New() Value
	String() string
}

type ClassValue struct {
	valueStub
}

func NewClassValue() Value {
	c := &ClassValue{}
	c.initialize()
	c.class = c
	return c
}

func (c ClassValue) New() Value {
	return nil
}

func (c ClassValue) String() string {
	return "Class"
}
