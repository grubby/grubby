package builtins

type kernel struct {
	valueStub
}

func NewGlobalKernelClass() Value {
	k := &kernel{}
	k.initialize()
	return k
}

func (kernel *kernel) String() string {
	return "Kernel"
}
