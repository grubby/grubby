package tests_test

import (
	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("babby's first parser", func() {
	Describe("parsing Fixnums", func() {
		It(`parses the integer "9001"`, func() {
			node := parser.Parse("9001")
			Expect(node).To(Equal(ast.Block{
				Statement: ast.ConstantInt{Value: 9001},
			}))
		})

		It(`parses the float "3.14"`, func() {
			node := parser.Parse("3.14")
			Expect(node).To(Equal(ast.Block{
				Statement: ast.ConstantFloat{Value: 3.14},
			}))
		})
	})

	It(`parses the string "hello world"`, func() {
		node := parser.Parse("'hello world'")
		Expect(node).To(Equal(ast.Block{
			Statement: ast.SimpleString{Value: "hello world"},
		}))
	})

	It("parses the bare reference `foo`", func() {
		statement := parser.Parse("foo").Statement
		ref, ok := statement.(ast.BareReference)
		Expect(ok).To(BeTrue(), "expected to receive a bare reference")
		Expect(ref.Name).To(Equal("foo"))
	})

	Describe("call expressions", func() {
		It("parses a simple call expression", func() {
			statement := parser.Parse("puts()").Statement
			_, ok := statement.(ast.CallExpression)
			Expect(ok).To(BeTrue(), "expected to receive a call expression")
		})
	})
})
