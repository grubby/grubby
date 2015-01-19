package builtins

type NilClass struct {
	valueStub
	classStub
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

type nilInstance struct {
	valueStub
}

func (class *NilClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	n := &nilInstance{}
	n.initialize()
	n.class = class

	return n, nil
}

func (n *nilInstance) String() string {
	return ""
}

func (n *nilInstance) IsTruthy() bool {
	return false
}
