package builtins

type classStub struct {
	superClass        Class
	_included_modules []Module
	_classVars        map[string]Value

	moduleStub
}

func (classStub *classStub) SuperClass() Class {
	return classStub.superClass
}

func (classStub *classStub) Include(module Module) {
	classStub._included_modules = append(classStub._included_modules, module)
}

func (classStub *classStub) includedModules() []Module {
	return classStub._included_modules
}

func (classStub *classStub) classVariable(name string) Value {
	if classStub._classVars == nil {
		classStub._classVars = make(map[string]Value)
	}

	return classStub._classVars[name]
}

func (classStub *classStub) setClassVariable(name string, value Value) {
	if classStub._classVars == nil {
		classStub._classVars = make(map[string]Value)
	}

	classStub._classVars[name] = value
}
