package builtins

type NilClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewNilClass(provider ClassProvider) Class {
	n := &NilClass{}
	n.initialize()
	n.class = provider.ClassWithName("Class")
	n.superClass = provider.ClassWithName("Object")
	return n
}

func (n *NilClass) String() string {
	return "NilClass"
}

func (n *NilClass) Name() string {
	return "NilClass"
}

func (n *NilClass) AddInstanceMethod(method Method) {
	n.instanceMethods = append(n.instanceMethods, method)
}

type nilInstance struct {
	valueStub
}

func (class *NilClass) New(provider ClassProvider, args ...Value) Value {
	n := &nilInstance{}
	n.initialize()
	n.class = class

	return n
}

func (n *nilInstance) String() string {
	return ""
}
