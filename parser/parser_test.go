package parser_test

import (
	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/grubby/grubby/parser/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("goyacc parser", func() {
	var (
		lexer parser.RubyLexer
	)

	JustBeforeEach(func() {
		parser.DebugStatements = []string{}
		parser.Statements = make([]ast.Node, 0)
	})

	Describe("parsing nothing at all", func() {
		It("succeeds without any errors", func() {
			lexer = parser.NewBetterLexer("")
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.BetterRubyLexer).LastError).To(BeNil())
		})
	})

	Context("when the code parsed is syntactically valid", func() {
		JustBeforeEach(func() {
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.BetterRubyLexer).LastError).To(BeNil())
		})

		Describe("parsing an integer", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer("5")
			})

			It("returns a ConstantInt struct representing the value", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ConstantInt{Value: 5},
				}))
			})
		})

		Describe("parsing a float", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer("123.4567")
			})

			It("returns a ConstantFloat struct representing the value", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ConstantFloat{Value: 123.4567},
				}))
			})
		})

		Describe("strings", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer("'hello world'")
			})

			It("returns a SimpleString struct", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.SimpleString{Value: "'hello world'"},
				}))
			})
		})

		Describe("symbols", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(":foo")
			})

			It("returns an ast.Symbol", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Symbol{Name: "foo"},
				}))
			})
		})

		Describe("parsing multiple lines", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(`
:foo
:bar`)
			})

			It("returns multiple nodes", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Symbol{Name: "foo"},
					ast.Symbol{Name: "bar"},
				}))
			})
		})

		Describe("variables", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer("foo")
			})

			It("returns a bare reference", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.BareReference{Name: "foo"},
				}))
			})
		})

		Describe("call expressions", func() {
			Context("without parens", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("puts 'foo'")
				})

				It("returns a call expression with one arg", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "puts"},
							Args: []ast.Node{ast.SimpleString{Value: "'foo'"}},
						},
					}))
				})
			})

			Context("with parens", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("puts('foo', 'bar', 'baz')")
				})

				It("returns a call expression with args", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "puts"},
							Args: []ast.Node{
								ast.SimpleString{Value: "'foo'"},
								ast.SimpleString{Value: "'bar'"},
								ast.SimpleString{Value: "'baz'"},
							},
						},
					}))
				})
			})

			Context("without args", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("puts()")
				})

				It("returns a call expression without args", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "puts"},
							Args: []ast.Node{},
						},
					}))
				})
			})
		})

		Describe("method definitions", func() {
			Context("without parameters", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`
def something
  puts 'hai'
end
`)
				})

				It("returns a function declaration with the body set", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Name: ast.BareReference{Name: "something"},
							Args: []ast.Node{},
							Body: []ast.Node{
								ast.CallExpression{
									Func: ast.BareReference{Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Value: "'hai'"}},
								},
							},
						},
					}))
				})
			})

			Context("with parameters surrounded by parens", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`
def multi_put(str1, str2)
  puts str1
  puts str2
end
`)
				})

				It("returns a function declaration with the args set", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Name: ast.BareReference{Name: "multi_put"},
							Args: []ast.Node{
								ast.BareReference{Name: "str1"},
								ast.BareReference{Name: "str2"},
							},
							Body: []ast.Node{
								ast.CallExpression{
									Func: ast.BareReference{Name: "puts"},
									Args: []ast.Node{ast.BareReference{Name: "str1"}},
								},
								ast.CallExpression{
									Func: ast.BareReference{Name: "puts"},
									Args: []ast.Node{ast.BareReference{Name: "str2"}},
								},
							},
						},
					}))
				})
			})

			Context("with parameters but no parens", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`
def multi_put str1, str2
  puts str1, str2
end
`)
				})

				It("returns a function declaration with the args set", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Name: ast.BareReference{Name: "multi_put"},
							Args: []ast.Node{
								ast.BareReference{Name: "str1"},
								ast.BareReference{Name: "str2"},
							},
							Body: []ast.Node{
								ast.CallExpression{
									Func: ast.BareReference{Name: "puts"},
									Args: []ast.Node{
										ast.BareReference{Name: "str1"},
										ast.BareReference{Name: "str2"},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("comments", func() {
			Context("on a single line", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("#ceci n'est pas un ligne de code")
				})

				It("is ignored", func() {
					Expect(parser.Statements).To(BeEmpty())
				})
			})

			Context("at the end of a line of code", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`
5#this is a comment
12 # this is also a comment
`)
				})

				It("reads the line of code, ignoring the comment", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.ConstantInt{Value: 5},
						ast.ConstantInt{Value: 12},
					}))
				})
			})
		})

		Describe("classes", func() {
			Context("without any frills, bells, or whistles", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`
class Foo
  puts 'hai'
end
`)
				})

				It("returns a ClassDefn with the body set", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.ClassDecl{
							Name: "Foo",
							Body: []ast.Node{
								ast.CallExpression{
									Func: ast.BareReference{Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Value: "'hai'"}},
								},
							},
						},
					}))
				})
			})

			Context("with a superclass", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`
class Foo < Bar
end
`)
				})

				It("returns a class with the given superclass", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.ClassDecl{
							Name:       "Foo",
							SuperClass: ast.Class{Name: "Bar"},
							Body:       []ast.Node{},
						},
					}))
				})
			})

			Context("with namespaces", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`
class Foo::Biz::Bar < Foo::Biz::Baz
end
`)
				})

				It("returns a class with the given namespace", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.ClassDecl{
							Name:       "Bar",
							SuperClass: ast.Class{Name: "Baz", Namespace: "Foo::Biz"},
							Namespace:  "Foo::Biz",
							Body:       []ast.Node{},
						},
					}))
				})
			})
		})

		Describe("modules", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(`
module Foo::Bar::Baz
puts 'tumescent-wasty'
end
`)
			})

			It("returns a module declaration", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ModuleDecl{
						Name:      "Baz",
						Namespace: "Foo::Bar",
						Body: []ast.Node{
							ast.CallExpression{
								Func: ast.BareReference{Name: "puts"},
								Args: []ast.Node{ast.SimpleString{Value: "'tumescent-wasty'"}},
							},
						},
					},
				}))
			})
		})

		Describe("assignment to a variable", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(`foo = 5`)
			})

			It("returns an assignment expression", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Assignment{
						LHS: ast.BareReference{Name: "foo"},
						RHS: ast.ConstantInt{Value: 5},
					},
				}))
			})
		})

		Describe("booleans", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(`
true
false
`)
			})

			It("returns a Boolean", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Boolean{Value: true},
					ast.Boolean{Value: false},
				}))
			})
		})

		Describe("unary operators", func() {
			Describe("unary NOT", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`!true`)
				})

				It("returns a Negation expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Negation{
							Target: ast.Boolean{Value: true},
						},
					}))
				})
			})

			Describe("unary COMPLEMENT", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("~false")
				})

				It("returns a Complement expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Complement{
							Target: ast.Boolean{Value: false},
						},
					}))
				})
			})

			Describe("unary plus", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("+foo")
				})

				It("returns a Positive expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Positive{
							Target: ast.BareReference{Name: "foo"},
						},
					}))
				})
			})

			Describe("unary minus", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("-867.5309")
				})

				It("returns a Negative expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Negative{
							Target: ast.ConstantFloat{Value: 867.5309},
						},
					}))
				})
			})
		})

		Describe("binary operators", func() {
			Describe("+", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("5 + 12")
				})

				It("returns a Addition expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Addition{
							LHS: ast.ConstantInt{Value: 5},
							RHS: ast.ConstantInt{Value: 12},
						},
					}))
				})
			})

			Describe("-", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("555 - 123")
				})

				It("returns a Subtraction expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Subtraction{
							LHS: ast.ConstantInt{Value: 555},
							RHS: ast.ConstantInt{Value: 123},
						},
					}))
				})
			})

			Describe("*", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("321 * 123")
				})

				It("returns a Multiplication expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Multiplication{
							LHS: ast.ConstantInt{Value: 321},
							RHS: ast.ConstantInt{Value: 123},
						},
					}))
				})
			})
		})

		XDescribe("arrays", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer("[1,2,3,   4,5,6 ]")
			})

			It("returns an Array node", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Array{
						Nodes: []ast.Node{
							ast.ConstantInt{Value: 1},
							ast.ConstantInt{Value: 2},
							ast.ConstantInt{Value: 3},
							ast.ConstantInt{Value: 4},
							ast.ConstantInt{Value: 5},
							ast.ConstantInt{Value: 6},
						},
					},
				}))
			})
		})

		XDescribe("hashes", func() {
			Context("with hashrockets", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer("{:foo => bar}")
				})

				It("returns a Hash node", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Hash{
							Pairs: []ast.HashKeyValuePair{
								{
									Key:   ast.Symbol{Name: "foo"},
									Value: ast.BareReference{Name: "bar"},
								},
							},
						},
					}))
				})
			})

			Context("with 1.9 'key: value' pairs", func() {
				BeforeEach(func() {
					lexer = parser.NewBetterLexer(`{
key: value,
foo: bar,
}`)
				})

				It("returns a Hash node with symbol keys", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Hash{
							Pairs: []ast.HashKeyValuePair{
								{
									Key:   ast.Symbol{Name: "key"},
									Value: ast.BareReference{Name: "value"},
								}, {
									Key:   ast.Symbol{Name: "foo"},
									Value: ast.BareReference{Name: "bar"},
								},
							},
						},
					}))
				})
			})
		})

		XDescribe("globals", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer("$LOAD_PATH")
			})

			It("should be parsed as a GlobalVariable", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.GlobalVariable{Name: "LOAD_PATH"},
				}))
			})
		})

		XDescribe("instance variables", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(`
@foo = :bar
@FOO = :baz
`)
			})

			It("should be parsed as an InstanceVariable", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Assignment{
						LHS: ast.InstanceVariable{Name: "foo"},
						RHS: ast.Symbol{Name: "bar"},
					},
					ast.Assignment{
						LHS: ast.InstanceVariable{Name: "FOO"},
						RHS: ast.Symbol{Name: "baz"},
					},
				}))
			})
		})

		XDescribe("class variables", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(`
@@foo = :bar
@@FOO = :baz
`)
			})

			It("should be parsed as an ClassVariable", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Assignment{
						LHS: ast.ClassVariable{Name: "foo"},
						RHS: ast.Symbol{Name: "bar"},
					},
					ast.Assignment{
						LHS: ast.ClassVariable{Name: "FOO"},
						RHS: ast.Symbol{Name: "baz"},
					},
				}))
			})
		})
	})

	// FIXME: don't mark this as pending
	XDescribe("optional whitespace is optional", func() {
		BeforeEach(func() {
			lexer = parser.NewBetterLexer(`
class Foo<Bar
	 1+1
   5    +    5
	 puts      'foo'
	 puts(     'foo',    'bar', 'baz'    )

	 def    something
	           puts 'whatever'
	   end

	 def something2(  foo  ,   bar   )
	end

  abc    =   123

  !  true
  ~    true
  +   5
  -    123
end
`)
		})

		It("parses just fine", func() {
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.RubyLex).LastError).To(BeNil())
		})
	})

	XDescribe("invalid syntax", func() {
		JustBeforeEach(func() {
			Expect(parser.RubyParse(lexer)).ToNot(BeSuccessful())
			Expect(lexer.(*parser.RubyLex).LastError).ToNot(BeNil())
		})

		Context("given a class name that starts with a lowercase character", func() {
			BeforeEach(func() {
				lexer = parser.NewBetterLexer(`
class foo
end
`)
			})

			It("fails and returns a useful parse error", func() {
				Expect(parser.Statements).To(BeEmpty())
			})
		})
	})
})
