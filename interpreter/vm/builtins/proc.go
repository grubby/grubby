package builtins

type ProcClass struct {
	valueStub
	classStub
	instanceMethods []Method

	singletonProvider SingletonProvider
}

func NewProcClass(provider Provider) Class {
	proc := &ProcClass{}
	proc.class = provider.ClassProvider().ClassWithName("Class")
	proc.superClass = provider.ClassProvider().ClassWithName("Object")
	proc.initialize()
	proc.setStringer(proc.String)

	return proc
}

func (klass *ProcClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, nil
}

func (proc *ProcClass) Name() string {
	return "Proc"
}

func (proc *ProcClass) String() string {
	return "Proc"
}

func NewProcInstance(methodName string, provider Provider) *Proc {
	proc := &Proc{methodName: methodName}
	proc.class = provider.ClassProvider().ClassWithName("Proc")
	proc.provider = provider
	return proc
}

type Proc struct {
	valueStub

	methodName string
	provider   Provider
}

func (proc *Proc) Call(args ...Value) (Value, error) {
	method := args[0].Method(proc.methodName)
	if method == nil {
		return nil, NewNoMethodError(proc.methodName, args[0].String(), args[0].Class().String(), proc.provider.StackProvider().CurrentStack())
	}

	return method.Execute(args[0], nil)
}

func (proc *Proc) setContext(newContext Value) {

}
