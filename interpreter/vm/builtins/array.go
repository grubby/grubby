package builtins

import (
	"errors"
	"fmt"
	"strings"
)

type ArrayClass struct {
	valueStub
	classStub
	instanceMethods []Method

	singletonProvider SingletonProvider
}

func NewArrayClass(provider Provider) Class {
	a := &ArrayClass{}
	a.class = provider.ClassProvider().ClassWithName("Class")
	a.superClass = provider.ClassProvider().ClassWithName("Object")
	a.initialize()
	a.setStringer(a.String)

	a.AddMethod(NewNativeMethod("shift", provider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		if len(a.members) == 0 {
			return provider.SingletonProvider().SingletonWithName("nil"), nil
		}

		val := a.members[0]
		a.members = a.members[1:]
		return val, nil
	}))

	a.AddMethod(NewNativeMethod("unshift", provider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		a.members = append([]Value{args[0]}, a.members[0:]...)
		return a, nil
	}))

	a.AddMethod(NewNativeMethod("include?", provider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		for _, m := range a.members {
			if m == args[0] {
				return provider.SingletonProvider().SingletonWithName("true"), nil
			}
		}

		return provider.SingletonProvider().SingletonWithName("false"), nil
	}))

	a.AddMethod(NewNativeMethod("-", provider, func(self Value, block Block, args ...Value) (Value, error) {
		a := self.(*Array)
		argAsArray, ok := args[0].(*Array)
		if !ok {
			return nil, errors.New(fmt.Sprintf("TypeError: no implicit conversion of %s into Array", args[0].Class().String()))
		}

		selfAsArray := self.(*Array)
		indicesToRemove := map[int]bool{}
		for _, otherMember := range argAsArray.members {
			for index, member := range selfAsArray.members {
				equalMethod := member.Method("==")
				if equalMethod == nil {
					return nil, NewNoMethodError("==", member.String(), member.Class().String(), provider.StackProvider().CurrentStack())
				}
				equal, err := equalMethod.Execute(member, block, otherMember)
				if err != nil {
					return nil, err
				}

				if equal.IsTruthy() {
					indicesToRemove[index] = true
				}
			}
		}

		newMembers := []Value{}
		for index, element := range a.members {
			_, exists := indicesToRemove[index]
			if !exists {
				newMembers = append(newMembers, element)
			}
		}

		a.members = newMembers
		return self, nil
	}))

	a.AddMethod(NewNativeMethod("select", provider, func(self Value, block Block, args ...Value) (Value, error) {
		arr, _ := provider.ClassProvider().ClassWithName("Array").New(provider)
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
	a.AddMethod(NewNativeMethod("map", provider, func(self Value, block Block, args ...Value) (Value, error) {
		arr, _ := provider.ClassProvider().ClassWithName("Array").New(provider)
		newArray := arr.(*Array)
		selfAsArray := self.(*Array)

		var mapper Block

		if len(args) == 1 {
			mapper = args[0].(*Proc)
		} else {
			mapper = block
		}

		for _, element := range selfAsArray.members {
			result, err := mapper.Call(element)
			if err != nil {
				return nil, err
			}

			newArray.members = append(newArray.members, result)
		}

		return newArray, nil
	}))
	a.AddMethod(NewNativeMethod("each", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsArray := self.(*Array)
		for _, element := range selfAsArray.members {
			_, err := block.Call(element)
			if err != nil {
				return nil, err
			}
		}

		return selfAsArray, nil
	}))
	a.AddMethod(NewNativeMethod("join", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsArray := self.(*Array)
		separator := args[0].(*StringValue).value
		pieces := make([]string, len(selfAsArray.members))

		for index, element := range selfAsArray.members {
			pieces[index] = element.String()
		}

		return NewString(strings.Join(pieces, separator), provider), nil
	}))
	a.AddMethod(NewNativeMethod("any?", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsArray := self.(*Array)
		for _, member := range selfAsArray.members {
			value, err := block.Call(member)
			if err != nil {
				return nil, err
			}

			if value.IsTruthy() {
				return provider.SingletonProvider().SingletonWithName("true"), nil
			}
		}

		return provider.SingletonProvider().SingletonWithName("false"), nil
	}))

	return a
}

func (klass *ArrayClass) AddInstanceMethod(m Method) {
	klass.instanceMethods = append(klass.instanceMethods, m)
}

func (klass *ArrayClass) New(provider Provider, args ...Value) (Value, error) {
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
