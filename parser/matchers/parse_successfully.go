package matchers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grubby/grubby/parser"
	"github.com/onsi/gomega"
)

type beSuccessfulMatcher struct{}

func BeSuccessful() gomega.OmegaMatcher {
	return beSuccessfulMatcher{}
}

func (matcher beSuccessfulMatcher) Match(actual interface{}) (bool, error) {
	switch actual.(type) {
	case int:
		return actual == 0, nil
	default:
		return false, errors.New(fmt.Sprintf("Expected %#v to be an int, but it was a %T", actual, actual))

	}
}

func (matcher beSuccessfulMatcher) FailureMessage(_ interface{}) string {
	messages := []string{
		"Expected parsing to succeed, but it failed.",
		"Debug output from parser:",
	}
	messages = append(messages, parser.DebugStatements...)
	return strings.Join(messages, "\n")
}

func (matcher beSuccessfulMatcher) NegatedFailureMessage(_ interface{}) string {
	messages := []string{
		"Expected parsing to fail but it succeeded.",
		"Debug output from parser:",
	}
	messages = append(messages, parser.DebugStatements...)
	return strings.Join(messages, "\n")
}
