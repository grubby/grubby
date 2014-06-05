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
			callExpr, ok := statement.(ast.CallExpression)
			Expect(ok).To(BeTrue(), "expected to receive a call expression")
			Expect(callExpr.Func).To(Equal("puts"))
		})

		It("parses a call expression with a single arg", func() {
			statement := parser.Parse(`puts('hai')`).Statement
			callExpr, ok := statement.(ast.CallExpression)
			Expect(ok).To(BeTrue(), "expected to receive a call expression")
			Expect(len(callExpr.Args)).To(Equal(1), "expected a single arg")

			asString, ok := callExpr.Args[0].(ast.SimpleString)
			Expect(ok).To(BeTrue(), "expected first arg to be a simple string")
			Expect(asString.Value).To(Equal("hai"))
		})

		It("parses a call expression with multiple args", func() {
			statement := parser.Parse("puts(foo, bar, baz)").Statement
			callExpr, ok := statement.(ast.CallExpression)
			Expect(ok).To(BeTrue(), "expected to receive a call expression")
			Expect(len(callExpr.Args)).To(Equal(3), "expected three arguments")
		})
	})
})
