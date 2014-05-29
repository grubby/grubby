package tests_test

import (
	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("babby's first parser", func() {
	It("parses the string `1`", func() {
		node := parser.Parse("9001")
		Expect(node).To(Equal(ast.Block{
			Declaration: ast.ConstantInt{Value: 9001},
		}))
	})
})
