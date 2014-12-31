package builtins

type ioClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewIOClass(provider ClassProvider) Class {
	i := &ioClass{}
	i.initialize()
	i.class = provider.ClassWithName("Class")
	i.superClass = provider.ClassWithName("Object")
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

func (io *ioClass) New(provider ClassProvider, args ...Value) Value {
	return nil
}
