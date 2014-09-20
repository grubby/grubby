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
			lexer = parser.NewLexer("")
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.StatefulRubyLexer).LastError).To(BeNil())
		})
	})

	Context("when the code parsed is syntactically valid", func() {
		JustBeforeEach(func() {
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.StatefulRubyLexer).LastError).To(BeNil())
		})

		Describe("parsing an integer", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("5")
			})

			It("returns a ConstantInt struct representing the value", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ConstantInt{Value: 5},
				}))
			})
		})

		Describe("parsing a float", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("123.4567")
			})

			It("returns a ConstantFloat struct representing the value", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ConstantFloat{Value: 123.4567},
				}))
			})
		})

		Describe("backtics", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("`echo 'this wont work on windows'`")
			})

			It("returns a Subshell struct", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Subshell{Command: "echo 'this wont work on windows'"},
				}))
			})
		})

		Describe("strings", func() {
			Context("with single quotes", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("'hello world'")
				})

				It("returns a SimpleString struct", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.SimpleString{Value: "hello world"},
					}))
				})
			})

			Context("with escaped single quotes", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("'hello \\' world'")
				})

				It("returns a SimpleString struct", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.SimpleString{Value: "hello \\' world"},
					}))
				})
			})

			Context("with double quotes", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`"pianic-#{foo}-vespid"`)
				})

				It("returns a InterpolatedString struct", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.InterpolatedString{Value: "pianic-#{foo}-vespid"},
					}))
				})

				Context("with escaped double quotes", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`"pianic-\"-vespid"`)
					})

					It("returns a InterpolatedString struct", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.InterpolatedString{Value: `pianic-\"-vespid`},
						}))
					})
				})

				Context("with double quotes inside the interpolation", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`"Raj-#{5 * " "}-Corin"`)
					})

					It("returns a InterpolatedString struct", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.InterpolatedString{Value: `Raj-#{5 * " "}-Corin`},
						}))
					})
				})
			})

			Describe("heredoc", func() {
				Context("with a dash", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
<<-FOO
spheniscomorphic-monoptic
   FOO
`)
					})

					It("returns an interpolated string, ending when it finds the identifier", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.InterpolatedString{Value: "spheniscomorphic-monoptic"},
						}))
					})
				})

				Context("without dash", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
<<FOO
resenter-postvenous
  FOO
FOO
`)
					})

					It("returns an interpolated string, ending only when it finds the identifier at column 0", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.InterpolatedString{Value: "resenter-postvenous\n  FOO"},
						}))
					})
				})
			})
		})

		Describe("symbols", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(":foo")
			})

			It("returns an ast.Symbol", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Symbol{Name: "foo"},
				}))
			})
		})

		Describe("parsing multiple lines", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
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
			Context("that the user has defined", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("foo")
				})

				It("returns a bare reference", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.BareReference{Name: "foo"},
					}))
				})
			})

			Context("__FILE__", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("__FILE__")
				})

				It("returns a file name reference", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FileNameConstReference{},
					}))
				})
			})
		})

		Describe("call expressions", func() {
			Context("with methods containing ! and ?", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("5.even?;5.taint!")
				})

				It("is parsed just fine", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 5},
							Func:   ast.BareReference{Name: "even?"},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 5},
							Func:   ast.BareReference{Name: "taint!"},
						},
					}))
				})
			})

			Context("with a dot", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
5.!
123.abc()
`)
				})

				It("is parsed as a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 5},
							Func:   ast.BareReference{Name: "!"},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 123},
							Func:   ast.BareReference{Name: "abc"},
							Args:   []ast.Node{},
						},
					}))
				})
			})

			Context("without parens", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("puts 'foo'")
				})

				It("returns a call expression with one arg", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "puts"},
							Args: []ast.Node{ast.SimpleString{Value: "foo"}},
						},
					}))
				})
			})

			Context("with parens", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("puts('foo', 'bar', 'baz')")
				})

				It("returns a call expression with args", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "puts"},
							Args: []ast.Node{
								ast.SimpleString{Value: "foo"},
								ast.SimpleString{Value: "bar"},
								ast.SimpleString{Value: "baz"},
							},
						},
					}))
				})
			})

			Context("without args", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("puts()")
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

			Context("on an object", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
foo.send(:catalecta_coassistant)
ARGV.shift
`)
				})

				It("correctly sets the target", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.BareReference{Name: "foo"},
							Func:   ast.BareReference{Name: "send"},
							Args:   []ast.Node{ast.Symbol{Name: "catalecta_coassistant"}},
						},
						ast.CallExpression{
							Target: ast.BareReference{Name: "ARGV"},
							Func:   ast.BareReference{Name: "shift"},
						},
					}))
				})
			})

			Context("a call expression with mixed arguments", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("File.expand_path('../../lib', __FILE__)")
				})

				It("parsed correctly", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.BareReference{Name: "File"},
							Func:   ast.BareReference{Name: "expand_path"},
							Args: []ast.Node{
								ast.SimpleString{Value: "../../lib"},
								ast.FileNameConstReference{},
							},
						},
					}))
				})
			})

			Context("with a call expression as an argument", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("$:.unshift File.expand_path('../../lib', __FILE__)")
				})

				It("applies the correct args to each call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func:   ast.BareReference{Name: "unshift"},
							Target: ast.GlobalVariable{Name: ":"},
							Args: []ast.Node{
								ast.CallExpression{
									Target: ast.BareReference{Name: "File"}, // TODO: should be a class
									Func:   ast.BareReference{Name: "expand_path"},
									Args: []ast.Node{
										ast.SimpleString{Value: "../../lib"},
										ast.FileNameConstReference{},
									},
								},
							},
						},
					}))
				})
			})

			Context("very nested call expressions", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("$:.unshift File.expand_path(File.dirname(__FILE__) + '/../lib')")
				})

				It("is parsed with the correct args on the correct methods", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.GlobalVariable{Name: ":"},
							Func:   ast.BareReference{Name: "unshift"},
							Args: []ast.Node{
								ast.CallExpression{
									Target: ast.BareReference{Name: "File"},
									Func:   ast.BareReference{Name: "expand_path"},
									Args: []ast.Node{
										ast.CallExpression{
											Target: ast.CallExpression{
												Target: ast.BareReference{Name: "File"},
												Func:   ast.BareReference{Name: "dirname"},
												Args:   []ast.Node{ast.FileNameConstReference{}},
											},
											Func: ast.BareReference{Name: "+"},
											Args: []ast.Node{ast.SimpleString{Value: "/../lib"}},
										},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("method definitions", func() {
			Context("without parameters", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
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
									Args: []ast.Node{ast.SimpleString{Value: "hai"}},
								},
							},
						},
					}))
				})
			})

			Context("with parameters with default values", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
def foo(a = 123)
end
`)
				})

				It("returns a function declaration with the default values set", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Name: ast.BareReference{"foo"},
							Args: []ast.Node{
								ast.MethodParam{
									Name:         ast.BareReference{Name: "a"},
									DefaultValue: ast.ConstantInt{Value: 123},
								},
							},
							Body: []ast.Node{},
						},
					}))
				})
			})

			Context("with parameters surrounded by parens", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
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
								ast.MethodParam{Name: ast.BareReference{Name: "str1"}},
								ast.MethodParam{Name: ast.BareReference{Name: "str2"}},
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
					lexer = parser.NewLexer(`
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
								ast.MethodParam{Name: ast.BareReference{Name: "str1"}},
								ast.MethodParam{Name: ast.BareReference{Name: "str2"}},
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
					lexer = parser.NewLexer("#ceci n'est pas un ligne de code")
				})

				It("is ignored", func() {
					Expect(parser.Statements).To(BeEmpty())
				})
			})

			Context("at the end of a line of code", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
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
					lexer = parser.NewLexer(`
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
									Args: []ast.Node{ast.SimpleString{Value: "hai"}},
								},
							},
						},
					}))
				})
			})

			Context("with a superclass", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
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
					lexer = parser.NewLexer(`
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

			Context("with a method defined", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
class Whatever < ThisShouldWork
  def initialize
    super

    config[:files] = []
  end
end
`)
				})

				It("should parse a class with a single method defined", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.ClassDecl{
							Name:       "Whatever",
							SuperClass: ast.Class{Name: "ThisShouldWork"},
							Body: []ast.Node{
								ast.FuncDecl{
									Name: ast.BareReference{Name: "initialize"},
									Args: []ast.Node{},
									Body: []ast.Node{
										ast.BareReference{Name: "super"},
										ast.CallExpression{
											Target: ast.BareReference{Name: "config"},
											Func:   ast.BareReference{Name: "[]="},
											Args: []ast.Node{
												ast.Symbol{Name: "files"},
												ast.Array{Nodes: []ast.Node{}},
											},
										},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("modules", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
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
								Args: []ast.Node{ast.SimpleString{Value: "tumescent-wasty"}},
							},
						},
					},
				}))
			})
		})

		Describe("assignment to a variable", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`foo = 5`)
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
				lexer = parser.NewLexer(`
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
					lexer = parser.NewLexer(`!true`)
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
					lexer = parser.NewLexer("~false")
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
					lexer = parser.NewLexer("+foo")
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
					lexer = parser.NewLexer("-867.5309")
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
					lexer = parser.NewLexer("5 + 12")
				})

				It("returns a CallExpression with the '+' method", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 5},
							Func:   ast.BareReference{Name: "+"},
							Args:   []ast.Node{ast.ConstantInt{Value: 12}},
						},
					}))
				})
			})

			Describe("-", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("555 - 123")
				})

				It("returns a subtraction call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 555},
							Func:   ast.BareReference{Name: "-"},
							Args:   []ast.Node{ast.ConstantInt{Value: 123}},
						},
					}))
				})
			})

			Describe("*", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("321 * 123")
				})

				It("returns a multiplication call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 321},
							Func:   ast.BareReference{Name: "*"},
							Args:   []ast.Node{ast.ConstantInt{Value: 123}},
						},
					}))
				})
			})

			Describe("/", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("321 / 123")
				})

				It("returns a division call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 321},
							Func:   ast.BareReference{Name: "/"},
							Args:   []ast.Node{ast.ConstantInt{Value: 123}},
						},
					}))
				})
			})

			Describe("%", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("321 % 123")
				})

				It("returns a modulo call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 321},
							Func:   ast.BareReference{Name: "%"},
							Args:   []ast.Node{ast.ConstantInt{Value: 123}},
						},
					}))
				})
			})

			Describe("**", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("321 ** 123")
				})

				It("returns a modulo call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 321},
							Func:   ast.BareReference{Name: "**"},
							Args:   []ast.Node{ast.ConstantInt{Value: 123}},
						},
					}))
				})
			})

			Describe("<< and >>", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
321 << 123
555 >> 666
`)
				})

				It("returns a shovel call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 321},
							Func:   ast.BareReference{Name: "<<"},
							Args:   []ast.Node{ast.ConstantInt{Value: 123}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 555},
							Func:   ast.BareReference{Name: ">>"},
							Args:   []ast.Node{ast.ConstantInt{Value: 666}},
						},
					}))
				})
			})

			Describe("bitwise binary operators", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
1 & 0
1 | 0
1 ^ 5
`)
				})

				It("returns a call expressions for those methods", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "&"},
							Args:   []ast.Node{ast.ConstantInt{Value: 0}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "|"},
							Args:   []ast.Node{ast.ConstantInt{Value: 0}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "^"},
							Args:   []ast.Node{ast.ConstantInt{Value: 5}},
						},
					}))
				})
			})

			Describe("operators for sorting", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
5 <= 55
12 >= 21
22 <=> 22
`)
				})

				It("is parsed as an call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 5},
							Func:   ast.BareReference{Name: "<="},
							Args:   []ast.Node{ast.ConstantInt{Value: 55}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 12},
							Func:   ast.BareReference{Name: ">="},
							Args:   []ast.Node{ast.ConstantInt{Value: 21}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 22},
							Func:   ast.BareReference{Name: "<=>"},
							Args:   []ast.Node{ast.ConstantInt{Value: 22}},
						},
					}))
				})
			})

			Describe("binary boolean operators", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
1 && 0
1 || 0
`)
				})

				It("is parsed as a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "&&"},
							Args:   []ast.Node{ast.ConstantInt{Value: 0}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "||"},
							Args:   []ast.Node{ast.ConstantInt{Value: 0}},
						},
					}))
				})
			})

			Describe("equality operators", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
1 == 1
1 === 1
1 != 1
1 =~ 1
1 !~ 1
`)
				})

				It("is parsed as a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "=="},
							Args:   []ast.Node{ast.ConstantInt{Value: 1}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "==="},
							Args:   []ast.Node{ast.ConstantInt{Value: 1}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "!="},
							Args:   []ast.Node{ast.ConstantInt{Value: 1}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "=~"},
							Args:   []ast.Node{ast.ConstantInt{Value: 1}},
						},
						ast.CallExpression{
							Target: ast.ConstantInt{Value: 1},
							Func:   ast.BareReference{Name: "!~"},
							Args:   []ast.Node{ast.ConstantInt{Value: 1}},
						},
					}))
				})
			})
		})

		Describe("arrays", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("[1,2,3,   4,5,6 ]")
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

		Describe("hashes", func() {
			Context("with hashrockets", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("{:foo => bar}")
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

			Describe("assigning a value to a key", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("hash[:key] = :value")
				})

				It("returns a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.BareReference{Name: "hash"},
							Func:   ast.BareReference{Name: "[]="},
							Args: []ast.Node{
								ast.Symbol{Name: "key"},
								ast.Symbol{Name: "value"},
							},
						},
					}))
				})
			})

			Describe("retrieving a value for a given key", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("hash[:key]")
				})

				It("returns a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.BareReference{Name: "hash"},
							Func:   ast.BareReference{Name: "[]"},
							Args:   []ast.Node{ast.Symbol{Name: "key"}},
						},
					}))
				})
			})

			Context("with 1.9 'key: value' pairs", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`{
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

		Describe("globals", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("$LOAD_PATH; $0")
			})

			It("should be parsed as a GlobalVariable", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.GlobalVariable{Name: "LOAD_PATH"},
					ast.GlobalVariable{Name: "0"},
				}))
			})
		})

		Describe("instance variables", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
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

		Describe("class variables", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
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

		Describe("blocks", func() {
			Context("without any args", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
function_that_takes_a_block do
  puts 'semiannual-pomfret'
end
`)
				})

				It("is parsed as an ast.Block", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "function_that_takes_a_block"},
							Args: []ast.Node{
								ast.Block{
									Body: []ast.Node{
										ast.CallExpression{
											Func: ast.BareReference{Name: "puts"},
											Args: []ast.Node{ast.SimpleString{Value: "semiannual-pomfret"}},
										},
									},
								},
							},
						},
					}))
				})
			})

			Context("with args", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
with_a_block do |and, with, some, args|
'aww yiss'
end
`)
				})

				It("is parsed as an ast.Block with args", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "with_a_block"},
							Args: []ast.Node{
								ast.Block{
									Args: []ast.Node{
										ast.BareReference{Name: "and"},
										ast.BareReference{Name: "with"},
										ast.BareReference{Name: "some"},
										ast.BareReference{Name: "args"},
									},
									Body: []ast.Node{ast.SimpleString{Value: "aww yiss"}},
								},
							},
						},
					}))
				})
			})

			Context("with curly braces", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("with.a_block {|foo| puts foo}")
				})

				It("is parsed as an ast.Block with args", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.BareReference{Name: "with"},
							Func:   ast.BareReference{Name: "a_block"},
							Args: []ast.Node{
								ast.Block{
									Args: []ast.Node{ast.BareReference{Name: "foo"}},
									Body: []ast.Node{
										ast.CallExpression{
											Func: ast.BareReference{Name: "puts"},
											Args: []ast.Node{ast.BareReference{Name: "foo"}},
										},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("unless", func() {
			Context("at the end of an expression", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("5 unless false")
				})

				It("is parsed as an IfBlock", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Condition: ast.Negation{Target: ast.Boolean{false}},
							Body:      []ast.Node{ast.ConstantInt{Value: 5}},
						},
					}))
				})
			})
		})

		Describe("if else blocks", func() {
			Context("without an else", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
if false
  puts 'Romanize-whereover'
end`)
				})

				It("is parsed as an ast.IfBlock struct", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Condition: ast.Boolean{Value: false},
							Body: []ast.Node{
								ast.CallExpression{
									Func: ast.BareReference{Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Value: "Romanize-whereover"}},
								},
							},
						},
					}))
				})
			})

			Context("at the end of an expression", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("exit(1) if false")
				})

				It("is parsed as an ast.IfBlock", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Condition: ast.Boolean{Value: false},
							Body: []ast.Node{
								ast.CallExpression{
									Func: ast.BareReference{Name: "exit"},
									Args: []ast.Node{ast.ConstantInt{Value: 1}},
								},
							},
						},
					}))
				})
			})

			Context("with an else", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
if false
  puts 'Romanize-whereover'
else
  puts 'Kiplingese-disinvolve'
end`)
				})

				It("is parsed as an ast.IfBlock struct", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Condition: ast.Boolean{Value: false},
							Body: []ast.Node{
								ast.CallExpression{
									Func: ast.BareReference{Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Value: "Romanize-whereover"}},
								},
							},
							Else: []ast.Node{
								ast.IfBlock{
									Condition: ast.Boolean{Value: true},
									Body: []ast.Node{
										ast.CallExpression{
											Func: ast.BareReference{Name: "puts"},
											Args: []ast.Node{ast.SimpleString{Value: "Kiplingese-disinvolve"}},
										},
									},
								},
							},
						},
					}))
				})
			})

			Context("with multiple elsif conditions", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
if false
  'purifier-cartouche'
elsif false
  'bronchophthisis-hypersurface'
elsif false
  'sharpware-nasality'
else
  'Osmeridae-harpylike'
end
`)
				})

				It("returns a very nested set of conditions to check", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Condition: ast.Boolean{Value: false},
							Body: []ast.Node{
								ast.SimpleString{Value: "purifier-cartouche"},
							},
							Else: []ast.Node{
								ast.IfBlock{
									Condition: ast.Boolean{Value: false},
									Body: []ast.Node{
										ast.SimpleString{Value: "bronchophthisis-hypersurface"},
									},
								},
								ast.IfBlock{
									Condition: ast.Boolean{Value: false},
									Body: []ast.Node{
										ast.SimpleString{Value: "sharpware-nasality"},
									},
								},
								ast.IfBlock{
									Condition: ast.Boolean{Value: true},
									Body: []ast.Node{
										ast.SimpleString{Value: "Osmeridae-harpylike"},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("begin followed by several rescue statements", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
begin
  foo()
rescue
  bar()
rescue LoadError
  biz()
rescue Exception => e
  baz()
end
`)
			})

			It("is parsed as a BeginBlock struct", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Begin{
						Body: []ast.Node{
							ast.CallExpression{
								Func: ast.BareReference{Name: "foo"},
								Args: []ast.Node{},
							},
						},
						Rescue: []ast.Node{
							ast.Rescue{
								Body: []ast.Node{
									ast.CallExpression{
										Func: ast.BareReference{Name: "bar"},
										Args: []ast.Node{},
									},
								},
							},
							ast.Rescue{
								Exception: ast.RescueException{
									Class: ast.BareReference{"LoadError"},
								},
								Body: []ast.Node{
									ast.CallExpression{
										Func: ast.BareReference{Name: "biz"},
										Args: []ast.Node{},
									},
								},
							},
							ast.Rescue{
								Exception: ast.RescueException{
									Var:   ast.BareReference{Name: "e"},
									Class: ast.BareReference{Name: "Exception"},
								},
								Body: []ast.Node{
									ast.CallExpression{
										Func: ast.BareReference{Name: "baz"},
										Args: []ast.Node{},
									},
								},
							},
						},
					},
				}))
			})
		})

		Describe("semicolons", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(";;a; b; c;;")
			})

			It("is treated as a line separator", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.BareReference{Name: "a"},
					ast.BareReference{Name: "b"},
					ast.BareReference{Name: "c"},
				}))
			})
		})

		Describe("parentheses", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
(1)
('hey'; 'this'; 'works!')
([].unshift)
`)
			})

			It("is treated as a grouping for a set of expressions", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Group{
						Body: []ast.Node{ast.ConstantInt{Value: 1}},
					},
					ast.Group{
						Body: []ast.Node{
							ast.SimpleString{Value: "hey"},
							ast.SimpleString{Value: "this"},
							ast.SimpleString{Value: "works!"},
						},
					},
					ast.Group{
						Body: []ast.Node{
							ast.CallExpression{
								Target: ast.Array{Nodes: []ast.Node{}},
								Func:   ast.BareReference{Name: "unshift"},
							},
						},
					},
				}))
			})
		})

		Describe("modules", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
Foo::Bar
Foo::Bar::Baz.method_call
`)
			})

			It("can be used to refer to classes and methods inside modules", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Class{
						Name:      "Bar",
						Namespace: "Foo",
					},
					ast.CallExpression{
						Target: ast.Class{
							Name:      "Baz",
							Namespace: "Foo::Bar",
						},
						Func: ast.BareReference{Name: "method_call"},
					},
				}))
			})
		})
	})

	Describe("a normal file you might parse", func() {
		BeforeEach(func() {
			lexer = parser.NewLexer(`
#!/usr/bin/env ruby

$:.unshift File.expand_path('../../lib', __FILE__)

require 'mspec/commands/mspec-run'
require 'mspec'

MSpecRun.main
`)
		})

		It("is parsed just fine", func() {
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.StatefulRubyLexer).LastError).To(BeNil())
		})
	})

	Describe("having tons of optional whitespace", func() {
		BeforeEach(func() {
			lexer = parser.NewLexer(`
class Foo<Bar
	 1+1
   5    +    5
	 puts      'foo'
	 puts(     'foo',    'bar', 'baz'    )

	 def    something
	           puts 'whatever'
               puts 'fasciculated-stripe'
	  end

	    def something2(  foo  ,   bar   )


      	end

  abc    =   123

  !  true
  ~    true
  +   5
  -    123
end

with_a_block { |foo| puts foo.inspect } # comment goes here
func.with_a_block { |foo | puts foo.inspect    } # all the comments # yep
`)
		})

		It("parses just fine", func() {
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.StatefulRubyLexer).LastError).To(BeNil())
		})
	})

	Describe("invalid syntax", func() {
		JustBeforeEach(func() {
			Expect(parser.RubyParse(lexer)).ToNot(BeSuccessful())
			Expect(lexer.(*parser.StatefulRubyLexer).LastError).ToNot(BeNil())
		})

		Context("given a class name that starts with a lowercase character", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
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
