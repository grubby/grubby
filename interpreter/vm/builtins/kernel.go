package builtins

type kernel struct {
	methods         []Method
	private_methods []Method
}

func NewGlobalKernel() Value {
	return &kernel{}
}

func (kernel *kernel) Methods() []Method {
	return kernel.methods
}

func (kernel *kernel) PrivateMethods() []Method {
	return kernel.private_methods
}

func (kernel *kernel) AddMethod(m Method) {
	kernel.methods = append(kernel.methods, m)
}

func (kernel *kernel) AddPrivateMethod(m Method) {
	kernel.private_methods = append(kernel.private_methods, m)
}
