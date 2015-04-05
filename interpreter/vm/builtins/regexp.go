package builtins

import (
	"errors"
	"regexp"
	"strings"
)

type regexpClass struct {
	valueStub
	classStub
}

func NewRegexpClass(provider Provider) Class {
	class := &regexpClass{}
	class.initialize()
	class.setStringer(class.String)
	class.class = provider.ClassProvider().ClassWithName("Class")
	class.superClass = provider.ClassProvider().ClassWithName("Object")

	class.AddMethod(NewNativeMethod("quote", provider, func(self Value, block Block, args ...Value) (Value, error) {
		argAsString := args[0].(*StringValue).value

		quoted := regexp.QuoteMeta(argAsString)
		return NewString(strings.Replace(quoted, " ", "\\ ", -1), provider), nil
	}))

	return class
}

func (c *regexpClass) String() string {
	return "Regexp"
}

func (c *regexpClass) Name() string {
	return "Regexp"
}

func (c *regexpClass) New(provider Provider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Regexp:Class")
}
