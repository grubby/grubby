package builtins

// abstract class interface
type Class interface {
	New(args ...Value) Value
	Value
}

// globlal Class class
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

func (c ClassValue) Name() string {
	return "Class"
}

func (c ClassValue) String() string {
	return "Class"
}

// user defined class type
type UserDefinedClass struct {
	name string
	valueStub
}

func NewUserDefinedClass(name string) Class {
	c := &UserDefinedClass{name: name}
	c.initialize()
	c.class = NewClassValue()
	return c
}

func (c UserDefinedClass) New(args ...Value) Value {
	return nil
}

func (c UserDefinedClass) String() string {
	return c.name
}
