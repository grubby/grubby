package token_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTokenizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grubby Token Suite")
}
