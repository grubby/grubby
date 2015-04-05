package testhelpers

import (
	"errors"
	"fmt"

	"github.com/grubby/grubby/interpreter/vm/builtins"

	. "github.com/onsi/gomega/types"
)

type haveMethodMatcher struct {
	methodName   string
	matchedValue interface{}
}

func HaveMethod(name string) GomegaMatcher {
	return &haveMethodMatcher{methodName: name}
}

func (matcher *haveMethodMatcher) Match(actual interface{}) (bool, error) {
	matcher.matchedValue = actual

	val, ok := actual.(builtins.Value)
	if !ok {
		return false, nil
	}

	method := val.Method(matcher.methodName)
	return method != nil, nil
}

func (matcher *haveMethodMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to have method '%s', but it did not", matcher.matchedValue, matcher.methodName)
}

func (matcher *haveMethodMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to not have method '%s', but it did", matcher.matchedValue, matcher.methodName)
}

// instance methods
type haveInstanceMethodMatcher struct {
	methodName   string
	matchedValue interface{}
}

func HaveInstanceMethod(name string) GomegaMatcher {
	return &haveInstanceMethodMatcher{methodName: name}
}

func (matcher *haveInstanceMethodMatcher) Match(actual interface{}) (bool, error) {
	matcher.matchedValue = actual

	val, ok := actual.(builtins.Module)
	if !ok {
		return false, nil
	}

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

// private methods

type havePrivateMethodMatcher struct {
	methodName   string
	matchedValue interface{}
}

func HavePrivateMethod(name string) GomegaMatcher {
	return &havePrivateMethodMatcher{methodName: name}
}

func (matcher *havePrivateMethodMatcher) Match(actual interface{}) (bool, error) {
	matcher.matchedValue = actual

	val, ok := actual.(builtins.Module)
	if !ok {
		return false, nil
	}

	for _, method := range val.Methods() {
		if method.Name() == matcher.methodName && method.IsPrivate() {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("Expected to find a private method '%s', but could not find it", matcher.methodName))
}

func (matcher *havePrivateMethodMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%#v' to have a private method '%s', but it did not", matcher.matchedValue, matcher.methodName)
}

func (matcher *havePrivateMethodMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%v' to not have a private method '%s', but it did", matcher.matchedValue, matcher.methodName)
}

// private instance methods
type havePrivateInstanceMethodMatcher struct {
	methodName   string
	matchedValue interface{}
}

func HavePrivateInstanceMethod(name string) GomegaMatcher {
	return &havePrivateInstanceMethodMatcher{methodName: name}
}

func (matcher *havePrivateInstanceMethodMatcher) Match(actual interface{}) (bool, error) {
	matcher.matchedValue = actual

	val, ok := actual.(builtins.Module)
	if !ok {
		return false, nil
	}

	for _, method := range val.InstanceMethods() {
		if method.Name() == matcher.methodName && method.IsPrivate() {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("Expected to have a private instance method '%s' but could not find it", matcher.methodName))
}

func (matcher *havePrivateInstanceMethodMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to have method '%s', but it did not", matcher.matchedValue, matcher.methodName)
}

func (matcher *havePrivateInstanceMethodMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%s' to not have method '%s', but it did", matcher.matchedValue, matcher.methodName)
}
