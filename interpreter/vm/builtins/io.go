package builtins

type ioClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewIOClass(provider Provider) Class {
	i := &ioClass{}
	i.initialize()
	i.setStringer(i.String)
	i.class = provider.ClassProvider().ClassWithName("Class")
	i.superClass = provider.ClassProvider().ClassWithName("Object")
	return i
}

func (io *ioClass) Name() string {
	return "IO"
}

func (io *ioClass) String() string {
	return "IO"
}

func (io *ioClass) AddInstanceMethod(m Method) {
	io.instanceMethods = append(io.instanceMethods, m)
}

func (io *ioClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, nil
}
