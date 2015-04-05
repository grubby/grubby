package builtins

type trueClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewTrueClass(provider Provider) Class {
	o := &trueClass{}
	o.initialize()
	o.setStringer(o.String)
	o.class = provider.ClassProvider().ClassWithName("Class")
	o.superClass = provider.ClassProvider().ClassWithName("Object")
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

func (t *trueInstance) String() string {
	return "true"
}

func (obj *trueClass) New(provider Provider, args ...Value) (Value, error) {
	o := &trueInstance{}
	o.initialize()
	o.setStringer(o.String)
	o.class = obj

	return o, nil
}

type falseClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewFalseClass(provider Provider) Class {
	o := &falseClass{}
	o.initialize()
	o.setStringer(o.String)
	o.class = provider.ClassProvider().ClassWithName("Class")
	o.superClass = provider.ClassProvider().ClassWithName("Object")
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

func (f *falseInstance) String() string {
	return "false"
}

func (obj *falseClass) New(provider Provider, args ...Value) (Value, error) {
	o := &falseInstance{}
	o.initialize()
	o.setStringer(o.String)
	o.class = obj

	return o, nil
}

func (f *falseInstance) IsTruthy() bool {
	return false
}
