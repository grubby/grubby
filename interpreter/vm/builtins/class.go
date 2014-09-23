package builtins

type Class interface {
	New(args ...Value) Value
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

func (c ClassValue) New(args ...Value) Value {
	return nil
}

func (c ClassValue) String() string {
	return "Class"
}
