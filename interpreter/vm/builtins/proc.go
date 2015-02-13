package builtins

type ProcClass struct {
	valueStub
	classStub
	instanceMethods []Method

	singletonProvider SingletonProvider
}

func NewProcClass(classProvider ClassProvider, singletonProvider SingletonProvider) Class {
	proc := &ProcClass{}
	proc.class = classProvider.ClassWithName("Class")
	proc.superClass = classProvider.ClassWithName("Object")
	proc.initialize()
	proc.setStringer(proc.String)

	return proc
}

func (klass *ProcClass) New(classProvider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	return nil, nil
}

func (proc *ProcClass) Name() string {
	return "Proc"
}

func (proc *ProcClass) String() string {
	return "Proc"
}

func NewProcInstance(methodName string, classProvider ClassProvider) *Proc {
	proc := &Proc{methodName: methodName}
	proc.class = classProvider.ClassWithName("Proc")
	return proc
}

type Proc struct {
	valueStub

	methodName string
}

func (proc *Proc) Call(args ...Value) (Value, error) {
	method, err := args[0].Method(proc.methodName)
	if err != nil {
		return nil, err
	}

	return method.Execute(args[0], nil)
}
