package tests_test

import (
	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("babby's first parser", func() {
	Describe("parsing Fixnums", func() {
		It(`parses the string "9001"`, func() {
			node := parser.Parse("9001")
			Expect(node).To(Equal(ast.Block{
				Statement: ast.ConstantInt{Value: 9001},
			}))
		})

		It(`parses the string "3.14"`, func() {
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
})
