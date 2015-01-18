package builtins

type classStub struct {
	superClass        Class
	_included_modules []Module

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
