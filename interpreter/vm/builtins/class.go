package builtins

// abstract class interface
type Class interface {
	New(args ...Value) Value
	Name() string
	Value
}

// globlal Class class
type ClassValue struct {
	valueStub
	instanceMethods        []Method
	privateInstanceMethods []Method
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

	instance_methods map[string]Method
}

type UserDefinedClassInstance struct {
	valueStub
}

func NewUserDefinedClass(name string) Class {
	c := &UserDefinedClass{
		name:             name,
		instance_methods: make(map[string]Method),
	}
	c.initialize()
	c.class = NewClassValue()
	return c
}

func (c *UserDefinedClass) New(args ...Value) Value {
	instance := &UserDefinedClassInstance{}
	instance.initialize()
	instance.class = c

	for _, m := range c.instance_methods {
		instance.AddMethod(m)
	}

	return instance
}

func (c UserDefinedClass) Name() string {
	return c.name
}

func (c UserDefinedClass) String() string {
	return c.name
}

func (c *UserDefinedClass) AddMethod(m Method) {
	c.instance_methods[m.Name()] = m
}
