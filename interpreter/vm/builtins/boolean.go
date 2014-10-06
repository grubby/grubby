package builtins

type trueClass struct {
	valueStub
	instanceMethods []Method
}

func NewTrueClass() Class {
	o := &trueClass{}
	o.initialize()
	o.class = NewClassValue().(Class)
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

type true struct {
	valueStub
}

func (obj *trueClass) New(args ...Value) Value {
	o := &true{}
	o.initialize()
	o.class = obj

	return o
}

type falseClass struct {
	valueStub
	instanceMethods []Method
}

func NewFalseClass() Class {
	o := &falseClass{}
	o.initialize()
	o.class = NewClassValue().(Class)
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

type false struct {
	valueStub
}

func (obj *falseClass) New(args ...Value) Value {
	o := &false{}
	o.initialize()
	o.class = obj

	return o
}
