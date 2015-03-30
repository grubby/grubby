package builtins

import "fmt"

type Method interface {
	Value

	Name() string
	IsPrivate() bool
	IsPublic() bool
	IsProtected() bool

	Execute(self Value, block Block, args ...Value) (Value, error)

	SetVisibility(MethodVisibility)
}

type nativeMethod struct {
	valueStub
	name              string
	visibility        MethodVisibility
	body              func(self Value, block Block, args ...Value) (Value, error)
	classProvider     ClassProvider
	singletonProvider SingletonProvider
}

func NewNativeMethod(name string, classProvider ClassProvider, singletonProvider SingletonProvider, body func(self Value, block Block, args ...Value) (Value, error)) Method {
	return newNativeMethod(name, Public, classProvider, singletonProvider, body)
}

func NewNativePrivateMethod(name string, classProvider ClassProvider, singletonProvider SingletonProvider, body func(self Value, block Block, args ...Value) (Value, error)) Method {
	return newNativeMethod(name, Private, classProvider, singletonProvider, body)
}

func newNativeMethod(name string, visibility MethodVisibility, provider ClassProvider, singletonProvider SingletonProvider, body func(self Value, block Block, args ...Value) (Value, error)) Method {
	m := &nativeMethod{
		name:              name,
		body:              body,
		visibility:        visibility,
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
	return method.visibility == Private
}

func (method *nativeMethod) IsPublic() bool {
	return method.visibility == Public
}

func (method *nativeMethod) IsProtected() bool {
	return method.visibility == Protected
}

func (method *nativeMethod) Execute(self Value, block Block, args ...Value) (Value, error) {
	return method.body(self, block, args...)
}

func (method *nativeMethod) String() string {
	return fmt.Sprintf("#Method: FIXME(ClassNameGoesHere)#%s", method.name)
}

func (method *nativeMethod) SetVisibility(visibility MethodVisibility) {
	method.visibility = visibility
}
