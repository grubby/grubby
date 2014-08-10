package builtins

type file struct {
	methods         []Method
	private_methods []Method
}

func NewFileClass() Value {
	return &file{}
}

func (file *file) Methods() []Method {
	return file.methods
}

func (file *file) PrivateMethods() []Method {
	return file.private_methods
}

func (file *file) AddMethod(m Method) {
	file.methods = append(file.methods, m)
}

func (file *file) AddPrivateMethod(m Method) {
	file.private_methods = append(file.private_methods, m)
}

func (file *file) String() string {
	return "File"
}

func (file *file) New() Value {
	return file
}

func (file *file) Class() Class {
	return file
}

func (file *file) SuperClass() Class {
	return nil
}
