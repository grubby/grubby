package parser_test

import (
	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parsing ruby files", func() {
	FDescribe("parsing Fixnums", func() {
		It(`parses the integer "9001"`, func() {
			node := parser.Parse("9001")
			Expect(node).To(Equal(ast.Block{
				Statements: []ast.Node{ast.ConstantInt{Value: 9001}},
			}))
		})

		It(`parses the float "3.14"`, func() {
			node := parser.Parse("3.14")
			Expect(node).To(Equal(ast.Block{
				Statements: []ast.Node{ast.ConstantFloat{Value: 3.14}},
			}))
		})
	})

	FDescribe("parsing strings", func() {
		It(`parses the string "hello world"`, func() {
			node := parser.Parse("'hello world'")
			Expect(node).To(Equal(ast.Block{
				Statements: []ast.Node{ast.SimpleString{Value: "hello world"}},
			}))
		})

		It("can parse multiline strings", func() {
			nodes := parser.Parse(`
'hello
world'
`).Statements

			Expect(nodes).To(Equal([]ast.Node{
				ast.SimpleString{Value: "hello\nworld"},
			}))
		})
	})

	FIt("parses the bare reference `foo`", func() {
		statements := parser.Parse("foo").Statements
		Expect(len(statements)).To(Equal(1))

		ref, ok := statements[0].(ast.BareReference)
		Expect(ok).To(BeTrue(), "expected to receive a bare reference")
		Expect(ref.Name).To(Equal("foo"))
	})

	FIt("parses symbols", func() {
		node := parser.Parse(":bar")
		Expect(node.Statements).To(Equal([]ast.Node{
			ast.Symbol{Name: "bar"},
		}))
	})

	FDescribe("call expressions", func() {
		It("parses a simple call expression", func() {
			statements := parser.Parse("puts()").Statements
			Expect(len(statements)).To(Equal(1))

			callExpr, ok := statements[0].(ast.CallExpression)
			Expect(ok).To(BeTrue(), "expected to receive a call expression")
			Expect(callExpr.Func).To(Equal("puts"))
		})

		It("parses a call expression with a single arg", func() {
			statements := parser.Parse(`puts('hai')`).Statements
			callExpr, ok := statements[0].(ast.CallExpression)
			Expect(len(statements)).To(Equal(1))

			Expect(ok).To(BeTrue(), "expected to receive a call expression")
			Expect(len(callExpr.Args)).To(Equal(1), "expected a single arg")

			asString, ok := callExpr.Args[0].(ast.SimpleString)
			Expect(ok).To(BeTrue(), "expected first arg to be a simple string")
			Expect(asString.Value).To(Equal("hai"))
		})

		It("parses a call expression with multiple args", func() {
			statements := parser.Parse("puts(foo, bar, baz)").Statements
			Expect(len(statements)).To(Equal(1))

			callExpr, ok := statements[0].(ast.CallExpression)
			Expect(ok).To(BeTrue(), "expected to receive a call expression")
			Expect(callExpr.Args).To(Equal([]ast.Node{
				ast.BareReference{Name: "foo"},
				ast.BareReference{Name: "bar"},
				ast.BareReference{Name: "baz"},
			}))
		})
	})

	FDescribe("method definitions", func() {
		It("parses a simple method with no args", func() {
			statements := parser.Parse(`
def foo
  puts('HAI')
end
`).Statements

			Expect(statements).To(Equal([]ast.Node{
				ast.FuncDecl{
					Name: "foo",
					Args: []ast.Node{},
					Body: []ast.Node{
						ast.CallExpression{
							Func: "puts",
							Args: []ast.Node{ast.SimpleString{Value: "HAI"}},
						},
					},
				},
			}))
		})

		It("parses a method with args", func() {
			statements := parser.Parse(`
def foo_2(bar, baz, biz)
  puts(bar)
end
`).Statements

			Expect(statements).To(Equal([]ast.Node{
				ast.FuncDecl{
					Name: "foo_2",
					Args: []ast.Node{
						ast.BareReference{Name: "bar"},
						ast.BareReference{Name: "baz"},
						ast.BareReference{Name: "biz"},
					},
					Body: []ast.Node{
						ast.CallExpression{
							Func: "puts",
							Args: []ast.Node{ast.BareReference{Name: "bar"}},
						},
					},
				},
			}))
		})
	})

	Describe("classes", func() {
		FIt("parses a simple class definition", func() {
			statements := parser.Parse(`
class MyClass
end
`).Statements

			Expect(statements).To(Equal([]ast.Node{
				ast.ClassDefn{Name: "MyClass"},
			}))
		})

		FIt("parses a class that inherits from another class", func() {
			statements := parser.Parse(`
class MyClass < Object
end
`).Statements

			Expect(statements).To(Equal([]ast.Node{
				ast.ClassDefn{
					Name:       "MyClass",
					SuperClass: "Object",
				},
			}))
		})

		It("parses a class with a namespace", func() {
			statements := parser.Parse(`
class MyNamespace::MyClass
end
`).Statements

			Expect(statements).To(Equal([]ast.Node{
				ast.ClassDefn{
					Name:      "MyClass",
					Namespace: "MyNamespace",
				},
			}))
		})

		PIt("parses a class with instance methods", func() {

		})
	})
})
