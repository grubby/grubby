package builtins

import "fmt"

type Method interface {
	Value

	Name() string
	IsPrivate() bool

	Execute(self Value, block Block, args ...Value) (Value, error)
}

type nativeMethod struct {
	valueStub
	name              string
	private           bool
	body              func(self Value, block Block, args ...Value) (Value, error)
	classProvider     ClassProvider
	singletonProvider SingletonProvider
}

func NewNativeMethod(name string, classProvider ClassProvider, singletonProvider SingletonProvider, body func(self Value, block Block, args ...Value) (Value, error)) Method {
	return newNativeMethod(name, false, classProvider, singletonProvider, body)
}

func NewNativePrivateMethod(name string, classProvider ClassProvider, singletonProvider SingletonProvider, body func(self Value, block Block, args ...Value) (Value, error)) Method {
	return newNativeMethod(name, true, classProvider, singletonProvider, body)
}

func newNativeMethod(name string, private bool, provider ClassProvider, singletonProvider SingletonProvider, body func(self Value, block Block, args ...Value) (Value, error)) Method {
	m := &nativeMethod{
		name:              name,
		body:              body,
		private:           private,
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

func (method *nativeMethod) IsPrivate() bool {
	return method.private
}

func (method *nativeMethod) Execute(self Value, block Block, args ...Value) (Value, error) {
	return method.body(self, block, args...)
}

func (method *nativeMethod) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}
