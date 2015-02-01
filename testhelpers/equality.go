package testhelpers

import (
	"fmt"

	"github.com/grubby/grubby/interpreter/vm/builtins"

	. "github.com/onsi/gomega/types"
)

type equalRubyString struct {
	expectedString string
}

func EqualRubyString(str string) GomegaMatcher {
	return &equalRubyString{expectedString: str}
}

func (matcher *equalRubyString) Match(actual interface{}) (bool, error) {
	val, ok := actual.(builtins.Value)
	if !ok {
		return false, nil
	}

	asStr, ok := val.(*builtins.StringValue)
	if !ok {
		return false, nil
	}

	return asStr.RawString() == matcher.expectedString, nil
}

func (matcher *equalRubyString) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%#v' to be the ruby string '%s', but it did not", actual, matcher.expectedString)
}

func (matcher *equalRubyString) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected '%#v' to not equal '%s', but it did", actual, matcher.expectedString)
}
