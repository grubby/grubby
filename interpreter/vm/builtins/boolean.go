package builtins

type trueClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewTrueClass(provider ClassProvider) Class {
	o := &trueClass{}
	o.initialize()
	o.class = provider.ClassWithName("Class")
	o.superClass = provider.ClassWithName("Object")
	return o
}

func (t *trueClass) AddInstanceMethod(m Method) {
	t.instanceMethods = append(t.instanceMethods, m)
}

func (t *trueClass) String() string {
	return "True"
}

func (t *trueClass) Name() string {
	return "True"
}

type trueInstance struct {
	valueStub
}

func (obj *trueClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	o := &trueInstance{}
	o.initialize()
	o.class = obj

	return o, nil
}

type falseClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewFalseClass(provider ClassProvider) Class {
	o := &falseClass{}
	o.initialize()
	o.class = provider.ClassWithName("Class")
	o.superClass = provider.ClassWithName("Object")
	return o
}

func (f *falseClass) AddInstanceMethod(m Method) {
	f.instanceMethods = append(f.instanceMethods, m)
}

func (f *falseClass) String() string {
	return "False"
}

func (f *falseClass) Name() string {
	return "False"
}

type falseInstance struct {
	valueStub
}

func (obj *falseClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	o := &falseInstance{}
	o.initialize()
	o.class = obj

	return o, nil
}

func (f *falseInstance) IsTruthy() bool {
	return false
}
