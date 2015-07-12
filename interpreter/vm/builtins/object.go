package builtins

import "fmt"

type ObjectClass struct {
	valueStub
	classStub

	provider Provider
}

func NewGlobalObjectClass(provider Provider) Class {
	o := &ObjectClass{}
	o.initialize()
	o.setStringer(o.String)
	o.provider = provider

	o.AddMethod(NewNativeMethod("==", provider, func(self Value, block Block, args ...Value) (Value, error) {
		if self == args[0] {
			return provider.SingletonProvider().SingletonWithName("true"), nil
		} else {
			return provider.SingletonProvider().SingletonWithName("false"), nil
		}
	}))

	o.AddMethod(NewNativeMethod("=~", provider, func(self Value, block Block, args ...Value) (Value, error) {
		// intended to be implemented by subclasses
		return provider.SingletonProvider().SingletonWithName("nil"), nil
	}))

	return o
}

func (c *ObjectClass) SetSuperClass() {
	class := c.provider.ClassProvider().ClassWithName("Class")
	if class == nil {
		panic("Expected Class class to exist")
	}

	superClass := c.provider.ClassProvider().ClassWithName("BasicObject")
	if superClass == nil {
		panic("Expected BasicObject class to exist")
	}

	c.class = class
	c.superClass = superClass
}

func (obj *ObjectClass) String() string {
	return "Object"
}

func (obj *ObjectClass) Name() string {
	return "Object"
}

type object struct {
	valueStub
}

func (o *object) String() string {
	return fmt.Sprintf("<%s:%p>", o.Class().String(), o)
}

func (obj *ObjectClass) New(provider Provider, args ...Value) (Value, error) {
	o := &object{}
	o.initialize()
	o.setStringer(o.String)
	o.class = obj

	return o, nil
}
