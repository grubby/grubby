package builtins

import (
	"errors"
	"fmt"
)

type ArrayClass struct {
	valueStub
	classStub
	instanceMethods []Method

	singletonProvider SingletonProvider
}

func NewArrayClass(classProvider ClassProvider, singletonProvider SingletonProvider) Class {
	a := &ArrayClass{}
	a.class = classProvider.ClassWithName("Class")
	a.superClass = classProvider.ClassWithName("Object")
	a.initialize()
	a.setStringer(a.String)

	a.AddMethod(NewNativeMethod("shift", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		if len(a.members) == 0 {
			return singletonProvider.SingletonWithName("nil"), nil
		}

		val := a.members[0]
		a.members = a.members[1:]
		return val, nil
	}))

	a.AddMethod(NewNativeMethod("unshift", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		a.members = append([]Value{args[0]}, a.members[0:]...)
		return a, nil
	}))

	a.AddMethod(NewNativeMethod("include?", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		for _, m := range a.members {
			if m == args[0] {
				return singletonProvider.SingletonWithName("true"), nil
			}
		}

		return singletonProvider.SingletonWithName("false"), nil
	}))

	a.AddMethod(NewNativeMethod("-", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		argAsArray, ok := args[0].(*Array)
		if !ok {
			return nil, errors.New(fmt.Sprintf("TypeError: no implicit conversion of %s into Array", args[0].Class().String()))
		}

		selfAsArray := self.(*Array)
		indicesToRemove := []int{}
		for _, otherMember := range argAsArray.members {
			for index, member := range selfAsArray.members {
				equalMethod, err := member.Method("==")
				if err != nil {
					return nil, err
				}
				equal, err := equalMethod.Execute(member, block, otherMember)
				if err != nil {
					return nil, err
				}

				if equal.IsTruthy() {
					indicesToRemove = append(indicesToRemove, index)
				}
			}
		}

		for _, indexToRemove := range indicesToRemove {
			a.members = append(a.members[:indexToRemove], a.members[indexToRemove+1:]...)
		}

		return self, nil
	}))

	a.AddMethod(NewNativeMethod("select", classProvider, singletonProvider, func(self Value, block Block, args ...Value) (Value, error) {
		arr, _ := classProvider.ClassWithName("Array").New(classProvider, singletonProvider)
		filteredArray := arr.(*Array)
		selfAsArray := self.(*Array)

		for _, element := range selfAsArray.members {
			result, err := block.Call(element)
			if err != nil {
				return nil, err
			}

			if result.IsTruthy() {
				filteredArray.members = append(filteredArray.members, element)
			}
		}

		return filteredArray, nil
	}))

	return a
}

func (klass *ArrayClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *ArrayClass) New(classProvider ClassProvider, singletonProvider SingletonProvider, args ...Value) (Value, error) {
	a := &Array{}
	a.initialize()
	a.setStringer(a.String)
	a.class = klass

	return a, nil
}

func (array *ArrayClass) Name() string {
	return "Array"
}

func (array *ArrayClass) String() string {
	return "Array"
}

type Array struct {
	valueStub
	members []Value
}

func (array *Array) Append(v Value) {
	array.members = append(array.members, v)
}

func (array *Array) Members() []Value {
	return array.members
}

func (array *Array) String() string {
	return "Array"
}
