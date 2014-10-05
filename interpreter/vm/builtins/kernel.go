package builtins

type kernel struct {
	valueStub
}

func NewGlobalKernelClass() Class {
	k := &kernel{}
	k.initialize()
	k.class = NewClassValue()
	return k
}

func (kernel *kernel) String() string {
	return "Kernel"
}

func (kernel *kernel) Name() string {
	return "Kernel"
}

func (kernel *kernel) New(args ...Value) Value {
	return nil
}
