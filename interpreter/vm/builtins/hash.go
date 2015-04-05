package builtins

import (
	"fmt"
	"strings"
)

type HashClass struct {
	valueStub
	classStub

	provider Provider

	instanceMethods []Method
}

func NewHashClass(provider Provider) Class {
	class := &HashClass{}
	class.initialize()
	class.setStringer(class.String)
	class.class = provider.ClassProvider().ClassWithName("Class")
	class.superClass = provider.ClassProvider().ClassWithName("Object")
	class.provider = provider

	class.AddMethod(NewNativeMethod("keys", provider, func(self Value, block Block, args ...Value) (Value, error) {
		o, _ := provider.ClassProvider().ClassWithName("Array").New(provider)
		keys := o.(*Array)
		for key := range self.(*Hash).hash {
			keys.Append(key)
		}

		return keys, nil
	}))
	class.AddMethod(NewNativeMethod("values", provider, func(self Value, block Block, args ...Value) (Value, error) {
		o, _ := provider.ClassProvider().ClassWithName("Array").New(provider)
		values := o.(*Array)
		for key := range self.(*Hash).hash {
			values.Append(self.(*Hash).hash[key])
		}

		return values, nil
	}))

	class.AddMethod(NewNativeMethod("each", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsHash := self.(*Hash)
		for key, value := range selfAsHash.hash {
			_, err := block.Call(key, value)
			if err != nil {
				return nil, err
			}
		}

		return selfAsHash, nil
	}))

	class.AddMethod(NewNativeMethod("[]=", provider, func(self Value, block Block, args ...Value) (Value, error) {
		self.(*Hash).hash[args[0]] = args[1]
		return args[1], nil
	}))

	class.AddMethod(NewNativeMethod("[]", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsHash := self.(*Hash)
		value, ok := selfAsHash.hash[args[0]]

		if !ok {
			return provider.SingletonProvider().SingletonWithName("nil"), nil
		} else {
			return value, nil
		}
	}))

	return class
}

func (klass *HashClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *HashClass) New(provider Provider, args ...Value) (Value, error) {
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
		pieces = append(pieces, fmt.Sprintf("%s => %s", key.String(), value.PrettyPrint()))
	}

	return fmt.Sprintf("{%s}", strings.Join(pieces, ", "))
}

func (hash *Hash) Add(key, value Value) {
	hash.hash[key] = value
}
