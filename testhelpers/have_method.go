package testhelpers

import (
	"errors"
	"fmt"

	"github.com/grubby/grubby/interpreter/vm/builtins"

	. "github.com/onsi/gomega/types"
)

type haveMethodMatcher struct {
	methodName   string
	matchedValue builtins.Value
}

func HaveMethod(name string) GomegaMatcher {
	return &haveMethodMatcher{methodName: name}
}

func (matcher *haveMethodMatcher) Match(actual interface{}) (bool, error) {
	val, ok := actual.(builtins.Value)
	if !ok {
		return false, nil
	}

	matcher.matchedValue = val

	_, err := val.Method(matcher.methodName)
	return err == nil, nil
}

func (matcher *haveMethodMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to have method '%s', but it did not", matcher.matchedValue, matcher.methodName)
}

func (matcher *haveMethodMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to not have method '%s', but it did", matcher.matchedValue, matcher.methodName)
}

type haveInstanceMethodMatcher struct {
	methodName   string
	matchedValue builtins.Value
}

func HaveInstanceMethod(name string) GomegaMatcher {
	return &haveInstanceMethodMatcher{methodName: name}
}

func (matcher *haveInstanceMethodMatcher) Match(actual interface{}) (bool, error) {
	val, ok := actual.(*builtins.RubyModule)
	if !ok {
		return false, nil
	}

	matcher.matchedValue = val

	for _, method := range val.InstanceMethods() {
		if method.Name() == matcher.methodName {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("no such method '%s'", matcher.methodName))
}

func (matcher *haveInstanceMethodMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to have method '%s', but it did not", matcher.matchedValue, matcher.methodName)
}

func (matcher *haveInstanceMethodMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to not have method '%s', but it did", matcher.matchedValue, matcher.methodName)
}
