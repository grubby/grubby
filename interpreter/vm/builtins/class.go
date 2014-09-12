package builtins

type Class interface {
	New() Value
	Value
}

type ClassValue struct {
	valueStub
}

func NewClassValue() Class {
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
