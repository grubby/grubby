package builtins

// abstract class interface
type Class interface {
	New(args ...Value) Value
	Name() string

	AddInstanceMethod(Method)

	Value
}

// globlal Class class
type ClassValue struct {
	valueStub
	instanceMethods []Method
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

func (c ClassValue) AddInstanceMethod(m Method) {
	c.instanceMethods = append(c.instanceMethods, m)
}

// user defined class type
type UserDefinedClass struct {
	name string
	valueStub

	includedModules []Value
	instanceMethods map[string]Method
}

type UserDefinedClassInstance struct {
	valueStub
}

func NewUserDefinedClass(name string) Class {
	c := &UserDefinedClass{
		name:            name,
		includedModules: make([]Value, 0),
		instanceMethods: make(map[string]Method),
	}
	c.initialize()
	c.class = NewClassValue()
	c.AddMethod(NewMethod("include", func(self Value, args ...Value) (Value, error) {
		for _, module := range args {
			c.includedModules = append(c.includedModules, module)
		}

		return c, nil
	}))

	return c
}

func (c *UserDefinedClass) New(args ...Value) Value {
	instance := &UserDefinedClassInstance{}
	instance.initialize()
	instance.class = c

	for _, m := range c.instanceMethods {
		instance.AddMethod(m)
	}

	for _, module := range c.includedModules {
		for _, method := range module.(Module).InstanceMethods() {
			instance.AddMethod(method)
		}
	}

	return instance
}

func (c UserDefinedClass) Name() string {
	return c.name
}

func (c UserDefinedClass) String() string {
	return c.name
}

func (c *UserDefinedClass) AddInstanceMethod(m Method) {
	c.instanceMethods[m.Name()] = m
}
