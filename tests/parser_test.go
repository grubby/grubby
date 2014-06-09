package tests_test

import (
	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parsing ruby files", func() {
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

	Describe("parsing strings", func() {
		It(`parses the string "hello world"`, func() {
			node := parser.Parse("'hello world'")
			Expect(node).To(Equal(ast.Block{
				Statement: ast.SimpleString{Value: "hello world"},
			}))
		})

		It("can parse multiline strings", func() {
			node := parser.Parse(`'hello
world'`)
			Expect(node.Statement).To(Equal(
				ast.SimpleString{Value: "hello\nworld"},
			))
		})
	})

	It("parses the bare reference `foo`", func() {
		statement := parser.Parse("foo").Statement
		ref, ok := statement.(ast.BareReference)
		Expect(ok).To(BeTrue(), "expected to receive a bare reference")
		Expect(ref.Name).To(Equal("foo"))
	})

	It("parses symbols", func() {
		node := parser.Parse(":bar")
		Expect(node.Statement).To(Equal(ast.Symbol{
			Name: "bar",
		}))
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

	Describe("method definitions", func() {
		It("parses a simple method with no args", func() {
			statement := parser.Parse(`
def foo
  puts('HAI')
end
`).Statement

			Expect(statement).To(Equal(ast.FuncDecl{
				Name: "foo",
				Args: []ast.Node{},
				Body: []ast.Node{
					ast.CallExpression{
						Func: "puts",
						Args: []ast.Node{ast.SimpleString{Value: "HAI"}},
					},
				},
			}))
		})

		It("parses a method with args", func() {
			statement := parser.Parse(`
def foo_2(bar)
  puts(bar)
end
`).Statement

			Expect(statement).To(Equal(ast.FuncDecl{
				Name: "foo_2",
				Args: []ast.Node{ast.BareReference{Name: "bar"}},
				Body: []ast.Node{
					ast.CallExpression{
						Func: "puts",
						Args: []ast.Node{ast.BareReference{Name: "bar"}},
					},
				},
			}))
		})
	})
})
