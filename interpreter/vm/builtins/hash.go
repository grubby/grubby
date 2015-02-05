package builtins

import (
	"fmt"
	"strings"
)

type HashClass struct {
	valueStub
	classStub

	provider ClassProvider

	instanceMethods []Method
}

func NewHashClass(provider ClassProvider, singletonProvider SingletonProvider) Class {
	class := &HashClass{}
	class.initialize()
	class.setStringer(class.String)
	class.class = provider.ClassWithName("Class")
	class.superClass = provider.ClassWithName("Object")
	class.provider = provider

	class.AddMethod(NewNativeMethod("keys", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		o, _ := provider.ClassWithName("Array").New(provider, singletonProvider)
		keys := o.(*Array)
		for key := range self.(*Hash).hash {
			keys.Append(key)
		}

		return keys, nil
	}))
	class.AddMethod(NewNativeMethod("values", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		o, _ := provider.ClassWithName("Array").New(provider, singletonProvider)
		values := o.(*Array)
		for key := range self.(*Hash).hash {
			values.Append(self.(*Hash).hash[key])
		}

		return values, nil
	}))

	class.AddMethod(NewNativeMethod("[]=", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		self.(*Hash).hash[args[0]] = args[1]
		return args[1], nil
	}))

	class.AddMethod(NewNativeMethod("[]", provider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsHash := self.(*Hash)
		value, ok := selfAsHash.hash[args[0]]

		if !ok {
			return singletonProvider.SingletonWithName("nil"), nil
		} else {
			return value, nil
		}
	}))

	return class
}

func (klass *HashClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *HashClass) New(provider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	hash := &Hash{}
	hash.initialize()
	hash.setStringer(hash.String)
	hash.class = klass
	hash.hash = make(map[Value]Value)

	return hash, nil
}

func (hash *HashClass) Name() string {
	return "Hash"
}

func (hash *HashClass) String() string {
	return "Hash"
}

type Hash struct {
	hash map[Value]Value
	valueStub
}

func (hash *Hash) String() string {
	pieces := []string{}
	for key, value := range hash.hash {
		pieces = append(pieces, fmt.Sprintf("%s => %s", key.String(), value.String()))
	}

	return fmt.Sprintf("{%s}", strings.Join(pieces, ", "))
}

func (hash *Hash) Add(key, value Value) {
	hash.hash[key] = value
}
