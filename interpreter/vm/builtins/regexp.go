package builtins

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type regexpClass struct {
	valueStub
	classStub
	instanceMethods []Method
}

func NewRegexpClass(provider Provider) Class {
	c := &regexpClass{}
	c.class = provider.ClassProvider().ClassWithName("Class")
	c.superClass = provider.ClassProvider().ClassWithName("Object")
	c.initialize()
	c.setStringer(c.String)

	c.AddMethod(NewNativeMethod("quote", provider, func(self Value, block Block, args ...Value) (Value, error) {
		argAsString := args[0].(*StringValue).value

		quoted := regexp.QuoteMeta(argAsString)
		return NewString(strings.Replace(quoted, " ", "\\ ", -1), provider), nil
	}))
	c.AddMethod(NewNativeMethod("match", provider, func(self Value, block Block, args ...Value) (Value, error) {
		selfAsRegexp := self.(*Regexp)
		str := args[0].(*StringValue)

		regex, err := regexp.Compile(selfAsRegexp.expression)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("something wrong with your regexp, bub -- %s", selfAsRegexp.expression))
		}

		if regex.MatchString(str.String()) {
			return provider.SingletonProvider().SingletonWithName("true"), nil
		} else {
			return provider.SingletonProvider().SingletonWithName("false"), nil
		}
	}))

	return c
}

func (regexp *regexpClass) Name() string {
	return "Regexp"
}

func (regexp *regexpClass) String() string {
	return "Regexp"
}

func (klass *regexpClass) New(provider Provider, args ...Value) (Value, error) {
	a := &Regexp{expression: args[0].String()}
	a.initialize()
	a.setStringer(a.String)
	a.class = klass

	return a, nil
}

type Regexp struct {
	valueStub
	expression string
}

func (regexp *Regexp) String() string {
	return regexp.expression
}

func NewRegexp(provider Provider, expr string) Value {
	a := &Regexp{expression: expr}
	a.initialize()
	a.setStringer(a.String)
	a.class = provider.ClassProvider().ClassWithName("Regexp")

	return a
}
