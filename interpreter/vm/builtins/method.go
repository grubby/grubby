package builtins

import "fmt"

type Method interface {
	Value
	Name() string
	Execute(self Value, block Block, args ...Value) (Value, error)
}

type nativeMethod struct {
	valueStub
	name              string
	body              func(self Value, block Block, args ...Value) (Value, error)
	classProvider     ClassProvider
	singletonProvider SingletonProvider
}

func NewNativeMethod(name string, provider ClassProvider, singletonProvider SingletonProvider, body func(self Value, block Block, args ...Value) (Value, error)) Method {
	m := &nativeMethod{
		name:              name,
		body:              body,
		classProvider:     provider,
		singletonProvider: singletonProvider,
	}
	m.class = provider.ClassWithName("Method")
	m.initialize()
	m.setStringer(m.String)
	return m
}
func (method *nativeMethod) Name() string {
	return method.name
}

func (method *nativeMethod) Execute(self Value, block Block, args ...Value) (Value, error) {
	return method.body(self, block, args...)
}

func (method *nativeMethod) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}
