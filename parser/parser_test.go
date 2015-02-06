package parser_test

import (
	"fmt"

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
			Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).ToNot(HaveOccurred())
		})
	})

	Context("when the code parsed is syntactically valid", func() {
		JustBeforeEach(func() {
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).ToNot(HaveOccurred())
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

				Context("with string interpolation inside string interpolation", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
"#{@tag}#{ "(#{@comment})" if @comment }:#{escape @description}"
`)
					})

					It("returns an interpolated string", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.InterpolatedString{
								Line:  1,
								Value: `#{@tag}#{ "(#{@comment})" if @comment }:#{escape @description}`,
							},
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

			Describe("ascii character literals", func() {
				for asciiValue := 33; asciiValue <= 126; asciiValue++ {
					func(ascii int) {
						Context(fmt.Sprintf("?%s", string(ascii)), func() {
							BeforeEach(func() {
								lexer = parser.NewLexer("?" + string(ascii))
							})

							It("parses as a character literal", func() {
								Expect(parser.Statements).To(Equal([]ast.Node{
									ast.CharacterLiteral{Value: string(ascii)},
								}))
							})
						})
					}(asciiValue)
				}
			})

			Describe("heredoc", func() {
				Context("as an arg to a method call", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
foo(<<-EOS, 'a', 'b', 'c', __LINE__ + 1)
something
goes
here
EOS
`)
					})

					It("should parse the heredoc around the method call", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.CallExpression{
								Line: 1,
								Func: ast.BareReference{Line: 1, Name: "foo"},
								Args: []ast.Node{
									ast.InterpolatedString{Line: 2, Value: "something\ngoes\nhere"},
									ast.SimpleString{Line: 1, Value: "a"},
									ast.SimpleString{Line: 1, Value: "b"},
									ast.SimpleString{Line: 1, Value: "c"},
									ast.CallExpression{
										Line:   1,
										Target: ast.LineNumberConstReference{Line: 1},
										Func:   ast.BareReference{Line: 1, Name: "+"},
										Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 1}},
									},
								},
							},
						}))
					})
				})

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
							ast.InterpolatedString{Line: 2, Value: "spheniscomorphic-monoptic"},
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
							ast.InterpolatedString{Line: 2, Value: "resenter-postvenous\n  FOO"},
						}))
					})
				})
			})

			Context("with % notation", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
%[ain't life grand]
%<this is the worst>
%{GET OUT}
%(DONT EVER WRITE CODE LIKE THIS)
`)
				})

				It("is parsed as an interpolated string", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.InterpolatedString{Line: 1, Value: "ain't life grand"},
						ast.InterpolatedString{Line: 2, Value: "this is the worst"},
						ast.InterpolatedString{Line: 3, Value: "GET OUT"},
						ast.InterpolatedString{Line: 4, Value: "DONT EVER WRITE CODE LIKE THIS"},
					}))
				})
			})
		})

		Describe("symbols", func() {
			Describe("generated dynamically with strings using interpolation", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`instance_variable_get :"@#{symbol}"`)
				})

				It("is parsed as a regular Symbol with the interpolation template still present", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "instance_variable_get"},
							Args: []ast.Node{ast.Symbol{Name: "@#{symbol}"}},
						},
					}))
				})
			})

			Context("without any special characters", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(":foo")
				})

				It("returns an ast.Symbol", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Symbol{Name: "foo"},
					}))
				})
			})

			Context("with an @ following the :", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(":@foo")
				})

				It("is parsed just like a regular symbol", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{ast.Symbol{Name: "@foo"}}))
				})
			})

			Context("with special characters", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(":foo!bar?")
				})

				It("is parsed as a symbol", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Symbol{Name: "foo!bar?"},
					}))
				})
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
					ast.Symbol{Line: 1, Name: "foo"},
					ast.Symbol{Line: 2, Name: "bar"},
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

			Context("__LINE__", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("__LINE__")
				})

				It("returns a line number reference", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.LineNumberConstReference{},
					}))
				})

				Describe("as an object", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer("__LINE__ + 1")
					})

					It("can have methods called on it", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.CallExpression{
								Target: ast.LineNumberConstReference{},
								Func:   ast.BareReference{Name: "+"},
								Args:   []ast.Node{ast.ConstantInt{Value: 1}},
							},
						}))
					})
				})
			})
		})

		Describe("case statements", func() {
			Context("without a value to compare against", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
case
  when truthy_method?
    0
  when falsey_method?
    1
end
`)
				})

				It("should be parsed as an ast.SwitchStatement", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.SwitchStatement{
							Line: 1,
							Cases: []ast.SwitchCase{
								ast.SwitchCase{
									Conditions: []ast.Node{
										ast.CallExpression{
											Line: 2,
											Func: ast.BareReference{Line: 2, Name: "truthy_method?"},
										},
									},
									Body: []ast.Node{ast.ConstantInt{Line: 3, Value: 0}},
								},
								ast.SwitchCase{
									Conditions: []ast.Node{
										ast.CallExpression{
											Line: 4,
											Func: ast.BareReference{Line: 4, Name: "falsey_method?"},
										},
									},
									Body: []ast.Node{ast.ConstantInt{Line: 5, Value: 1}},
								},
							},
						},
					}))
				})
			})

			Context("with a value to compare against", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
case some_integer
when 1, 3
  puts 'even'
when 2, 4
  puts 'odd'
when ?^
  puts 'a single character (^) literal'
when ?:
  puts 'a single character (:) literal'
else
  puts 'whoops'
end
`)
				})

				It("should be parsed as an ast.SwitchStatement", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.SwitchStatement{
							Line:      1,
							Condition: ast.BareReference{Line: 1, Name: "some_integer"},
							Cases: []ast.SwitchCase{
								ast.SwitchCase{
									Conditions: []ast.Node{
										ast.ConstantInt{Line: 2, Value: 1},
										ast.ConstantInt{Line: 2, Value: 3},
									},
									Body: []ast.Node{
										ast.CallExpression{
											Line: 3,
											Func: ast.BareReference{Line: 3, Name: "puts"},
											Args: []ast.Node{ast.SimpleString{Line: 3, Value: "even"}},
										},
									},
								},
								ast.SwitchCase{
									Conditions: []ast.Node{
										ast.ConstantInt{Line: 4, Value: 2},
										ast.ConstantInt{Line: 4, Value: 4},
									},
									Body: []ast.Node{
										ast.CallExpression{
											Line: 5,
											Func: ast.BareReference{Line: 5, Name: "puts"},
											Args: []ast.Node{ast.SimpleString{Line: 5, Value: "odd"}},
										},
									},
								},
								ast.SwitchCase{
									Conditions: []ast.Node{
										ast.CharacterLiteral{Line: 6, Value: "^"},
									},
									Body: []ast.Node{
										ast.CallExpression{
											Line: 7,
											Func: ast.BareReference{Line: 7, Name: "puts"},
											Args: []ast.Node{ast.SimpleString{Line: 7, Value: "a single character (^) literal"}},
										},
									},
								},
								ast.SwitchCase{
									Conditions: []ast.Node{
										ast.CharacterLiteral{Line: 8, Value: ":"},
									},
									Body: []ast.Node{
										ast.CallExpression{
											Line: 9,
											Func: ast.BareReference{Line: 9, Name: "puts"},
											Args: []ast.Node{ast.SimpleString{Line: 9, Value: "a single character (:) literal"}},
										},
									},
								},
							},
							Else: []ast.Node{
								ast.CallExpression{
									Line: 11,
									Func: ast.BareReference{Line: 11, Name: "puts"},
									Args: []ast.Node{
										ast.SimpleString{Line: 11, Value: "whoops"},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("procs", func() {
			Context("created with curly braces", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("Proc.new { Mock.verify_count }")
				})

				It("is parsed as a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.BareReference{Name: "Proc"},
							Func:   ast.BareReference{Name: "new"},
							Args:   []ast.Node{},
							OptionalBlock: ast.Block{
								Body: []ast.Node{
									ast.CallExpression{
										Target: ast.BareReference{Name: "Mock"},
										Func:   ast.BareReference{Name: "verify_count"},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("call expressions", func() {
			Context("with a value that should be converted to a proc", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("describe(&blocks); explain(&:it_well)")
				})

				It("converts the argument to a proc", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "describe"},
							Args: []ast.Node{
								ast.CallExpression{
									Target: ast.BareReference{Name: "blocks"},
									Func:   ast.BareReference{Name: "to_proc"},
								},
							},
						},
						ast.CallExpression{
							Func: ast.BareReference{Name: "explain"},
							Args: []ast.Node{
								ast.CallExpression{
									Target: ast.Symbol{Name: "it_well"},
									Func:   ast.BareReference{Name: "to_proc"},
								},
							},
						},
					}))
				})
			})

			Context("with inline assignment of binary operators", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
a += 5
b -= [1]
c /= 'x'
d *= nil
`)
				})

				It("is parsed with the operator as the method name", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.BareReference{Line: 1, Name: "a"},
							Func:   ast.BareReference{Line: 1, Name: "+="},
							Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 5}},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.BareReference{Line: 2, Name: "b"},
							Func:   ast.BareReference{Line: 2, Name: "-="},
							Args: []ast.Node{ast.Array{
								Line:  2,
								Nodes: []ast.Node{ast.ConstantInt{Line: 2, Value: 1}},
							}},
						},
						ast.CallExpression{
							Line:   3,
							Target: ast.BareReference{Line: 3, Name: "c"},
							Func:   ast.BareReference{Line: 3, Name: "/="},
							Args:   []ast.Node{ast.SimpleString{Line: 3, Value: "x"}},
						},
						ast.CallExpression{
							Line:   4,
							Target: ast.BareReference{Line: 4, Name: "d"},
							Func:   ast.BareReference{Line: 4, Name: "*="},
							Args:   []ast.Node{ast.Nil{Line: 4}},
						},
					}))
				})
			})

			Context("with arguments split across newlines", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
method_with_lots_of_args('foo',
                         'bar',
                         'baz',
                         &buz)
`)
				})

				It("should be parsed as though the newlines were not present", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line: 1,
							Func: ast.BareReference{Line: 1, Name: "method_with_lots_of_args"},
							Args: []ast.Node{
								ast.SimpleString{Line: 1, Value: "foo"},
								ast.SimpleString{Line: 2, Value: "bar"},
								ast.SimpleString{Line: 3, Value: "baz"},
								ast.CallExpression{
									Line:   4,
									Func:   ast.BareReference{Line: 4, Name: "to_proc"},
									Target: ast.BareReference{Line: 4, Name: "buz"},
								},
							},
						},
					}))
				})
			})

			Context("with args and a block, but no parens", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
Signal.trap "INT", "TERM" do
  MSpec.actions :abort
end
`)
				})

				It("is parsed correctly", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.BareReference{Line: 1, Name: "Signal"},
							Func:   ast.BareReference{Line: 1, Name: "trap"},
							Args: []ast.Node{
								ast.InterpolatedString{Line: 1, Value: "INT"},
								ast.InterpolatedString{Line: 1, Value: "TERM"},
							},
							OptionalBlock: ast.Block{
								Line: 1,
								Body: []ast.Node{
									ast.CallExpression{
										Line:   2,
										Target: ast.BareReference{Line: 2, Name: "MSpec"},
										Func:   ast.BareReference{Line: 2, Name: "actions"},
										Args:   []ast.Node{ast.Symbol{Line: 2, Name: "abort"}},
									},
								},
							},
						},
					}))
				})
			})

			Context("with a block passed to a call expression targeting a group", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
    (@repeat || 1).times do
      yield
    end
`)
				})

				It("is parsed as a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line: 1,
							Target: ast.Group{
								Line: 1,
								Body: []ast.Node{
									ast.CallExpression{
										Line:   1,
										Target: ast.InstanceVariable{Line: 1, Name: "repeat"},
										Func:   ast.BareReference{Line: 1, Name: "||"},
										Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 1}},
									},
								},
							},
							Func: ast.BareReference{Line: 1, Name: "times"},
							Args: []ast.Node{},
							OptionalBlock: ast.Block{
								Line: 1,
								Body: []ast.Node{ast.Yield{Line: 2}},
							},
						},
					}))
				})
			})

			Context("with args and a block split across newlines", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
method_with_lots_of_args('foo',
                         'bar',
                         'baz') do |foo|
  puts foo
end
`)
				})

				It("should be parsed as though the newlines were not present", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line: 1,
							Func: ast.BareReference{Line: 1, Name: "method_with_lots_of_args"},
							Args: []ast.Node{
								ast.SimpleString{Line: 1, Value: "foo"},
								ast.SimpleString{Line: 2, Value: "bar"},
								ast.SimpleString{Line: 3, Value: "baz"},
							},
							OptionalBlock: ast.Block{
								Line: 3,
								Args: []ast.Node{ast.BareReference{Line: 3, Name: "foo"}},
								Body: []ast.Node{
									ast.CallExpression{
										Line: 4,
										Func: ast.BareReference{Line: 4, Name: "puts"},
										Args: []ast.Node{ast.BareReference{Line: 4, Name: "foo"}},
									},
								},
							},
						},
					}))
				})
			})

			Context("with a proc argument", func() {
				Context("inside parens", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer("takes_a_proc('foo', 'bar', &baz)")
					})

					It("is parsed with an ast.ProcArg argument", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.CallExpression{
								Func: ast.BareReference{Name: "takes_a_proc"},
								Args: []ast.Node{
									ast.SimpleString{Value: "foo"},
									ast.SimpleString{Value: "bar"},
									ast.CallExpression{
										Func:   ast.BareReference{Name: "to_proc"},
										Target: ast.BareReference{Name: "baz"},
									},
								},
							},
						}))
					})
				})

				Context("without parens", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer("takes_a_proc 'foo', 'bar', &baz")
					})

					It("is parsed with an ast.ProcArg argument", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.CallExpression{
								Func: ast.BareReference{Name: "takes_a_proc"},
								Args: []ast.Node{
									ast.SimpleString{Value: "foo"},
									ast.SimpleString{Value: "bar"},
									ast.CallExpression{
										Target: ast.BareReference{Name: "baz"},
										Func:   ast.BareReference{Name: "to_proc"},
									},
								},
							},
						}))
					})
				})
			})

			Context("with a static method", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("foo = String(bar)")
				})

				It("is parsed like a regular method call", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							LHS: ast.BareReference{Name: "foo"},
							RHS: ast.CallExpression{
								Func: ast.BareReference{Name: "String"},
								Args: []ast.Node{ast.BareReference{Name: "bar"}},
							},
						},
					}))
				})
			})

			Context("chained together with a block at the end", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("MSpec.retrieve(:files).inject(0) { |max, f| f.size > max ? f.size : max }")
				})

				It("is parsed as nested call expressions", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.CallExpression{
								Target: ast.BareReference{Name: "MSpec"},
								Func:   ast.BareReference{Name: "retrieve"},
								Args: []ast.Node{
									ast.Symbol{Name: "files"},
								},
							},
							Func: ast.BareReference{Name: "inject"},
							Args: []ast.Node{
								ast.ConstantInt{Value: 0},
							},
							OptionalBlock: ast.Block{
								Args: []ast.Node{
									ast.BareReference{Name: "max"},
									ast.BareReference{Name: "f"},
								},
								Body: []ast.Node{
									ast.Ternary{
										Condition: ast.CallExpression{
											Target: ast.CallExpression{
												Target: ast.BareReference{Name: "f"},
												Func:   ast.BareReference{Name: "size"},
											},
											Func: ast.BareReference{Name: ">"},
											Args: []ast.Node{ast.BareReference{Name: "max"}},
										},
										True: ast.CallExpression{
											Target: ast.BareReference{Name: "f"},
											Func:   ast.BareReference{Name: "size"},
										},
										False: ast.BareReference{Name: "max"},
									},
								},
							},
						},
					}))
				})
			})

			Context("chained together with dots and parentheses", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("SpecVersion.new(String(other)).to_i")
				})

				It("is parsed as a nested call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.CallExpression{
								Target: ast.BareReference{Name: "SpecVersion"},
								Func:   ast.BareReference{Name: "new"},
								Args: []ast.Node{
									ast.CallExpression{
										Func: ast.BareReference{Name: "String"},
										Args: []ast.Node{ast.BareReference{Name: "other"}},
									},
								},
							},
							Func: ast.BareReference{Name: "to_i"},
						},
					}))
				})
			})

			Context("with an expression wrapped in parentheses", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`('hello %s world' % ['cruel']).inspect`)
				})

				It("is parsed with the expression group as the target", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.Group{
								Body: []ast.Node{
									ast.CallExpression{
										Target: ast.SimpleString{Value: "hello %s world"},
										Func:   ast.BareReference{Name: "%"},
										Args: []ast.Node{
											ast.Array{
												Nodes: []ast.Node{
													ast.SimpleString{Value: "cruel"},
												},
											},
										},
									},
								},
							},
							Func: ast.BareReference{Name: "inspect"},
							Args: []ast.Node{},
						},
					}))
				})
			})

			Context("with methods containing ! and ?", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("5.even?;5.taint!; block_given?; search_and_destroy!")
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
						ast.CallExpression{
							Func: ast.BareReference{Name: "block_given?"},
						},
						ast.CallExpression{
							Func: ast.BareReference{Name: "search_and_destroy!"},
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
							Line:   1,
							Target: ast.ConstantInt{Line: 1, Value: 5},
							Func:   ast.BareReference{Line: 1, Name: "!"},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.ConstantInt{Line: 2, Value: 123},
							Func:   ast.BareReference{Line: 2, Name: "abc"},
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

			Context("on an object with args in parentheses", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
foo.send(:catalecta_coassistant)
ARGV.shift
`)
				})

				It("correctly sets the target", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.BareReference{Line: 1, Name: "foo"},
							Func:   ast.BareReference{Line: 1, Name: "send"},
							Args:   []ast.Node{ast.Symbol{Line: 1, Name: "catalecta_coassistant"}},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.BareReference{Line: 2, Name: "ARGV"},
							Func:   ast.BareReference{Line: 2, Name: "shift"},
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

		Describe("array splat", func() {
			Context("in a call expression", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("foo(*bar)")
				})

				It("marks the argument for array deferencing", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "foo"},
							Args: []ast.Node{
								ast.StarSplat{Value: ast.BareReference{Name: "bar"}},
							},
						},
					}))
				})
			})
		})

		Describe("method definitions", func() {
			Context("for setter methods", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
def foo=(bar)
end

def self.bar=(foo)
end
`)
				})

				It("should be parsed with the equals sign in the name", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Line:   1,
							Target: nil,
							Name:   ast.BareReference{Line: 1, Name: "foo="},
							Args: []ast.Node{
								ast.MethodParam{
									Name:    ast.BareReference{Line: 1, Name: "bar"},
									IsSplat: false,
								},
							},
							Body: []ast.Node{},
						},
						ast.FuncDecl{
							Line:   4,
							Target: ast.Self{Line: 4},
							Name:   ast.BareReference{Line: 4, Name: "bar="},
							Args: []ast.Node{
								ast.MethodParam{
									Name:    ast.BareReference{Line: 4, Name: "foo"},
									IsSplat: false,
								},
							},
							Body: []ast.Node{},
						},
					}))
				})
			})

			Context("on an object at runtime", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
obj = Object.new
def obj.start
end
`)
				})

				It("is parsed as a method declaration", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							Line: 1,
							LHS:  ast.BareReference{Line: 1, Name: "obj"},
							RHS: ast.CallExpression{
								Line:   1,
								Target: ast.BareReference{Line: 1, Name: "Object"},
								Func:   ast.BareReference{Line: 1, Name: "new"},
							},
						},
						ast.FuncDecl{
							Line:   2,
							Target: ast.BareReference{Line: 2, Name: "obj"},
							Name:   ast.BareReference{Line: 2, Name: "start"},
							Body:   []ast.Node{},
							Args:   []ast.Node{},
						},
					}))
				})
			})

			Context("with special names", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
def <=>
  self.to_i <=> other
end
`)
				})

				It("parses them as binary operators", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "<=>"},
							Args: []ast.Node{},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Target: ast.CallExpression{
										Line:   2,
										Target: ast.Self{Line: 2},
										Func:   ast.BareReference{Line: 2, Name: "to_i"},
									},
									Func: ast.BareReference{Line: 2, Name: "<=>"},
									Args: []ast.Node{ast.BareReference{Line: 2, Name: "other"}},
								},
							},
						},
					}))
				})
			})

			Context("with splat-args", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
def on(*args)
end
`)
				})

				It("marks the param as being splatted", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "on"},
							Args: []ast.Node{
								ast.MethodParam{
									Name:    ast.BareReference{Line: 1, Name: "args"},
									IsSplat: true,
								},
							},
							Body: []ast.Node{},
						},
					}))
				})
			})

			Context("with a named proc parameter", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
def test(*args, &block)
end
`)
				})

				It("marks the param as a proc", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "test"},
							Args: []ast.Node{
								ast.MethodParam{
									Name:    ast.BareReference{Line: 1, Name: "args"},
									IsSplat: true,
								},
								ast.MethodParam{
									Name:   ast.BareReference{Line: 1, Name: "block"},
									IsProc: true,
								},
							},
							Body: []ast.Node{},
						},
					}))
				})
			})

			Context("with ? or ! in the name", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
def foo!
  raise
end

def foo?
  false
end
`)
				})

				It("is parsed with the correct method name", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "foo!"},
							Args: []ast.Node{},
							Body: []ast.Node{
								ast.BareReference{Line: 2, Name: "raise"},
							},
						},
						ast.FuncDecl{
							Line: 5,
							Name: ast.BareReference{Line: 5, Name: "foo?"},
							Args: []ast.Node{},
							Body: []ast.Node{
								ast.Boolean{Line: 6, Value: false},
							},
						},
					}))
				})
			})

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
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "something"},
							Args: []ast.Node{},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 2, Value: "hai"}},
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
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "foo"},
							Args: []ast.Node{
								ast.MethodParam{
									Name:         ast.BareReference{Line: 1, Name: "a"},
									DefaultValue: ast.ConstantInt{Line: 1, Value: 123},
								},
							},
							Body: []ast.Node{},
						},
					}))
				})
			})

			Context("with default value args and a block", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
def self.describe(mod, options=nil, &block)
end
`)
				})

				It("returns a function declaration with appropriate method parameters", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.FuncDecl{
							Line:   1,
							Target: ast.Self{Line: 1},
							Name:   ast.BareReference{Line: 1, Name: "describe"},
							Args: []ast.Node{
								ast.MethodParam{Name: ast.BareReference{Line: 1, Name: "mod"}},
								ast.MethodParam{
									Name:         ast.BareReference{Line: 1, Name: "options"},
									DefaultValue: ast.Nil{Line: 1},
								},
								ast.MethodParam{
									Name:   ast.BareReference{Line: 1, Name: "block"},
									IsProc: true,
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
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "multi_put"},
							Args: []ast.Node{
								ast.MethodParam{Name: ast.BareReference{Line: 1, Name: "str1"}},
								ast.MethodParam{Name: ast.BareReference{Line: 1, Name: "str2"}},
							},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{ast.BareReference{Line: 2, Name: "str1"}},
								},
								ast.CallExpression{
									Line: 3,
									Func: ast.BareReference{Line: 3, Name: "puts"},
									Args: []ast.Node{ast.BareReference{Line: 3, Name: "str2"}},
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
							Line: 1,
							Name: ast.BareReference{Line: 1, Name: "multi_put"},
							Args: []ast.Node{
								ast.MethodParam{Name: ast.BareReference{Line: 1, Name: "str1"}},
								ast.MethodParam{Name: ast.BareReference{Line: 1, Name: "str2"}},
							},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{
										ast.BareReference{Line: 2, Name: "str1"},
										ast.BareReference{Line: 2, Name: "str2"},
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
						ast.ConstantInt{Line: 1, Value: 5},
						ast.ConstantInt{Line: 2, Value: 12},
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
							Line: 1,
							Name: "Foo",
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 2, Value: "hai"}},
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
							Line:       1,
							Name:       "Foo",
							SuperClass: ast.Class{Line: 1, Name: "Bar"},
							Body:       []ast.Node{},
						},
					}))
				})
			})

			Context("with namespaces", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
class Foo::Bar
end
`)
				})

				It("returns a class with the given namespace", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.ClassDecl{
							Line:      1,
							Name:      "Bar",
							Namespace: "Foo",
							Body:      []ast.Node{},
						},
					}))
				})
			})

			Context("with namespaces and super classes", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
class Foo::Biz::Bar < Foo::Biz::Baz
end
`)
				})

				It("returns a class with the given namespace", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.ClassDecl{
							Line:       1,
							Name:       "Bar",
							SuperClass: ast.Class{Line: 1, Name: "Baz", Namespace: "Foo::Biz"},
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
							Line:       1,
							Name:       "Whatever",
							SuperClass: ast.Class{Line: 1, Name: "ThisShouldWork"},
							Body: []ast.Node{
								ast.FuncDecl{
									Line: 2,
									Name: ast.BareReference{Line: 2, Name: "initialize"},
									Args: []ast.Node{},
									Body: []ast.Node{
										ast.BareReference{Line: 3, Name: "super"},
										ast.CallExpression{
											Line:   5,
											Target: ast.BareReference{Line: 5, Name: "config"},
											Func:   ast.BareReference{Line: 5, Name: "[]="},
											Args: []ast.Node{
												ast.Symbol{Line: 5, Name: "files"},
												ast.Array{Line: 5, Nodes: []ast.Node{}},
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

		Describe("eigenclasses", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
class Foo
  class << self
    puts 'this is evaluated inside the eigenclass for the Foo class'
  end
end
`)
			})

			It("can be used to modify the unique singleton class for an object", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ClassDecl{
						Line: 1,
						Name: "Foo",
						Body: []ast.Node{
							ast.EigenClass{
								Line:   2,
								Target: ast.Self{Line: 2},
								Body: []ast.Node{
									ast.CallExpression{
										Line: 3,
										Func: ast.BareReference{Line: 3, Name: "puts"},
										Args: []ast.Node{
											ast.SimpleString{
												Line:  3,
												Value: "this is evaluated inside the eigenclass for the Foo class",
											},
										},
									},
								},
							},
						},
					},
				}))
			})
		})

		// this is ambiguous because it is impossible to determine if you meant
		// foo [:something] as a call expression (e.g.: foo([:an_array_literal]))
		// ... or ...
		// foo [:something] as a call expression on foo (e.g.: call .[] on `foo`)
		Describe("ambiguous [] syntax", func() {
			Context("calling [] and []= on an instance variable", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
@shared [state.to_s] = state
@shared [state.to_s] # returns 'state'
`)
				})

				It("should be parsed as a call expression for the []= operator", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.InstanceVariable{Line: 1, Name: "shared"},
							Func:   ast.BareReference{Line: 1, Name: "[]="},
							Args: []ast.Node{
								ast.CallExpression{
									Line:   1,
									Target: ast.BareReference{Line: 1, Name: "state"},
									Func:   ast.BareReference{Line: 1, Name: "to_s"},
								},
								ast.BareReference{Line: 1, Name: "state"},
							},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.InstanceVariable{Line: 2, Name: "shared"},
							Func:   ast.BareReference{Line: 2, Name: "[]"},
							Args: []ast.Node{
								ast.CallExpression{
									Line:   2,
									Target: ast.BareReference{Line: 2, Name: "state"},
									Func:   ast.BareReference{Line: 2, Name: "to_s"},
								},
							},
						},
					}))
				})
			})

			Context("calling [] or []= on a call expression", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("retrieve(:features)[feature] = true")
				})

				It("should be parsed as a call expression for the appropriate method", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.CallExpression{
								Func: ast.BareReference{Name: "retrieve"},
								Args: []ast.Node{ast.Symbol{Name: "features"}},
							},
							Func: ast.BareReference{Name: "[]="},
							Args: []ast.Node{
								ast.BareReference{Name: "feature"},
								ast.Boolean{Value: true},
							},
						},
					}))
				})
			})
		})

		Describe("slicing of an array", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("'1.5.12'.split('.')[0,n]")
			})

			It("should be parsed as a call expression with an array arg", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.CallExpression{
						Target: ast.CallExpression{
							Target: ast.SimpleString{Value: "1.5.12"},
							Func:   ast.BareReference{Name: "split"},
							Args:   []ast.Node{ast.SimpleString{Value: "."}},
						},
						Func: ast.BareReference{Name: "[]"},
						Args: []ast.Node{
							ast.ConstantInt{Value: 0}, ast.BareReference{Name: "n"},
						},
					},
				}))
			})
		})

		Describe("modules", func() {
			Describe("referred to with leading :: to indicate that it belongs to the global namespace", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
SomeModule::InThe::CurrentScope
::SomeModule::InTheGlobalNamespace
::AnotherModule
`)
				})

				It("is parsed with the IsGlobalNamespace field set", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Class{
							Line:              1,
							Name:              "CurrentScope",
							Namespace:         "SomeModule::InThe",
							IsGlobalNamespace: false,
						},
						ast.Class{
							Line:              2,
							Name:              "InTheGlobalNamespace",
							Namespace:         "SomeModule",
							IsGlobalNamespace: true,
						},
						ast.Class{
							Line:              3,
							Name:              "AnotherModule",
							Namespace:         "",
							IsGlobalNamespace: true,
						},
					}))
				})
			})

			Context("with namespaces", func() {
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
							Line:      1,
							Name:      "Baz",
							Namespace: "Foo::Bar",
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 2, Value: "tumescent-wasty"}},
								},
							},
						},
					}))
				})
			})
		})

		Describe("preventing line breaks from terminating an expression", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
# all whitespace is ignored after the / until a nonwhitespace char is seen
foo  \
   # comments here are ignored, as are newlines
   \
   .inspect
`)
			})

			It("can be done using a backslash at the end of a line", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.CallExpression{
						Line:   2,
						Target: ast.BareReference{Line: 2, Name: "foo"},
						Func:   ast.BareReference{Line: 5, Name: "inspect"},
					},
				}))
			})
		})

		Describe("conditional assignment", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
a ||= 'aftergrass-Dowieite'
@options[:shared] ||= false
`)
			})

			It("returns a ConditionalAssignment expression", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ConditionalAssignment{
						Line: 1,
						LHS:  ast.BareReference{Line: 1, Name: "a"},
						RHS:  ast.SimpleString{Line: 1, Value: "aftergrass-Dowieite"},
					},
					ast.ConditionalAssignment{
						Line: 2,
						LHS: ast.CallExpression{
							Line:   2,
							Target: ast.InstanceVariable{Line: 2, Name: "options"},
							Func:   ast.BareReference{Line: 2, Name: "[]"},
							Args:   []ast.Node{ast.Symbol{Line: 2, Name: "shared"}},
						},
						RHS: ast.Boolean{Line: 2, Value: false},
					},
				}))
			})
		})

		Describe("assignment", func() {
			Context("inside of a method call", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("puts(hey = 'so what')")
				})

				It("is parsed as a call expression with assignment as an arg", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Func: ast.BareReference{Name: "puts"},
							Args: []ast.Node{
								ast.Assignment{
									LHS: ast.BareReference{Name: "hey"},
									RHS: ast.SimpleString{Value: "so what"},
								},
							},
						},
					}))
				})
			})

			Context("to multiple keys in a hash", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
HASH['first_key']    =
HASH['second_key'] = [:something]
`)
				})

				It("is parsed as nested call expressions", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.BareReference{Line: 1, Name: "HASH"},
							Func:   ast.BareReference{Line: 1, Name: "[]="},
							Args: []ast.Node{
								ast.SimpleString{Line: 1, Value: "first_key"},
								ast.CallExpression{
									Line:   2,
									Target: ast.BareReference{Line: 2, Name: "HASH"},
									Func:   ast.BareReference{Line: 2, Name: "[]="},
									Args: []ast.Node{
										ast.SimpleString{Line: 2, Value: "second_key"},
										ast.Array{
											Line:  2,
											Nodes: []ast.Node{ast.Symbol{Line: 2, Name: "something"}},
										},
									},
								},
							},
						},
					}))
				})
			})

			Context("to a single variable", func() {
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

			Context("to multiple variables", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("foo, bar = [1,2,3]")
				})

				It("returns an assignment expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							LHS: ast.Array{
								Nodes: []ast.Node{
									ast.BareReference{Name: "foo"},
									ast.BareReference{Name: "bar"},
								},
							},
							RHS: ast.Array{
								Nodes: []ast.Node{
									ast.ConstantInt{Value: 1},
									ast.ConstantInt{Value: 2},
									ast.ConstantInt{Value: 3},
								},
							},
						},
					}))
				})
			})

			Context("to multiple variables, with a splat", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("target, *actions = clause.split(/([=+-])/)")
				})

				It("is parsed as an assignment expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							LHS: ast.Array{Nodes: []ast.Node{
								ast.BareReference{Name: "target"},
								ast.StarSplat{Value: ast.BareReference{Name: "actions"}},
							}},
							RHS: ast.CallExpression{
								Target: ast.BareReference{Name: "clause"},
								Func:   ast.BareReference{Name: "split"},
								Args:   []ast.Node{ast.Regex{Value: "([=+-])"}},
							},
						},
					}))
				})
			})

			Context("to multiple instance or class variables", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
@foo, @bar = [1, 2]
@@foo, @@bar = [3, 4]
`)
				})

				It("should be parsed as a multiple assignment as well", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							Line: 1,
							LHS: ast.Array{
								Line: 1,
								Nodes: []ast.Node{
									ast.InstanceVariable{Line: 1, Name: "foo"},
									ast.InstanceVariable{Line: 1, Name: "bar"},
								},
							},
							RHS: ast.Array{
								Line: 1,
								Nodes: []ast.Node{
									ast.ConstantInt{Line: 1, Value: 1},
									ast.ConstantInt{Line: 1, Value: 2},
								},
							},
						},
						ast.Assignment{
							Line: 2,
							LHS: ast.Array{
								Line: 2,
								Nodes: []ast.Node{
									ast.ClassVariable{Line: 2, Name: "foo"},
									ast.ClassVariable{Line: 2, Name: "bar"},
								},
							},
							RHS: ast.Array{
								Line: 2,
								Nodes: []ast.Node{
									ast.ConstantInt{Line: 2, Value: 3},
									ast.ConstantInt{Line: 2, Value: 4},
								},
							},
						},
					}))
				})
			})

			Context("to indices in an array or hash", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("array[i], array[r] = array[r], array[i]")
				})

				It("be parsed as an assignment expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							LHS: ast.Array{
								Nodes: []ast.Node{
									ast.CallExpression{
										Target: ast.BareReference{Name: "array"},
										Func:   ast.BareReference{Name: "[]="},
										Args:   []ast.Node{ast.BareReference{Name: "i"}},
									},
									ast.CallExpression{
										Target: ast.BareReference{Name: "array"},
										Func:   ast.BareReference{Name: "[]="},
										Args:   []ast.Node{ast.BareReference{Name: "r"}},
									},
								},
							},
							RHS: ast.Array{
								Nodes: []ast.Node{
									ast.CallExpression{
										Target: ast.BareReference{Name: "array"},
										Func:   ast.BareReference{Name: "[]="},
										Args:   []ast.Node{ast.BareReference{Name: "r"}},
									},
									ast.CallExpression{
										Target: ast.BareReference{Name: "array"},
										Func:   ast.BareReference{Name: "[]="},
										Args:   []ast.Node{ast.BareReference{Name: "i"}},
									},
								},
							},
						},
					}))
				})
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
					ast.Boolean{Line: 1, Value: true},
					ast.Boolean{Line: 2, Value: false},
				}))
			})
		})

		Describe("unary operators", func() {
			Describe("unary NOT", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`!!true`)
				})

				It("returns a Negation expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Negation{
							Target: ast.Negation{
								Target: ast.Boolean{Value: true},
							},
						},
					}))
				})
			})

			Describe("unary COMPLEMENT", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("~~false")
				})

				It("returns a Complement expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Complement{
							Target: ast.Complement{
								Target: ast.Boolean{Value: false},
							},
						},
					}))
				})
			})

			Describe("unary plus", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("++foo")
				})

				It("returns a Positive expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Positive{
							Target: ast.Positive{
								Target: ast.BareReference{Name: "foo"},
							},
						},
					}))
				})
			})

			Describe("unary minus", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("--867.5309")
				})

				It("returns a Negative expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Negative{
							Target: ast.Negative{
								Target: ast.ConstantFloat{Value: 867.5309},
							},
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
					lexer = parser.NewLexer(`
321 / 123
(abc) / (xyz)
`)
				})

				It("returns a division call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.ConstantInt{Line: 1, Value: 321},
							Func:   ast.BareReference{Line: 1, Name: "/"},
							Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 123}},
						},
						ast.CallExpression{
							Line: 2,
							Target: ast.Group{
								Line: 2,
								Body: []ast.Node{ast.BareReference{Line: 2, Name: "abc"}},
							},
							Func: ast.BareReference{Line: 2, Name: "/"},
							Args: []ast.Node{
								ast.Group{
									Line: 2,
									Body: []ast.Node{ast.BareReference{Line: 2, Name: "xyz"}},
								},
							},
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
							Line:   1,
							Target: ast.ConstantInt{Line: 1, Value: 321},
							Func:   ast.BareReference{Line: 1, Name: "<<"},
							Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 123}},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.ConstantInt{Line: 2, Value: 555},
							Func:   ast.BareReference{Line: 2, Name: ">>"},
							Args:   []ast.Node{ast.ConstantInt{Line: 2, Value: 666}},
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
File.lchmod mode & 01777, path
`)
				})

				It("returns a call expressions for those methods", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.ConstantInt{Line: 1, Value: 1},
							Func:   ast.BareReference{Line: 1, Name: "&"},
							Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 0}},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.ConstantInt{Line: 2, Value: 1},
							Func:   ast.BareReference{Line: 2, Name: "|"},
							Args:   []ast.Node{ast.ConstantInt{Line: 2, Value: 0}},
						},
						ast.CallExpression{
							Line:   3,
							Target: ast.ConstantInt{Line: 3, Value: 1},
							Func:   ast.BareReference{Line: 3, Name: "^"},
							Args:   []ast.Node{ast.ConstantInt{Line: 3, Value: 5}},
						},
						ast.CallExpression{
							Line:   4,
							Target: ast.BareReference{Line: 4, Name: "File"},
							Func:   ast.BareReference{Line: 4, Name: "lchmod"},
							Args: []ast.Node{
								ast.CallExpression{
									Line:   4,
									Target: ast.BareReference{Line: 4, Name: "mode"},
									Func:   ast.BareReference{Line: 4, Name: "&"},
									Args:   []ast.Node{ast.ConstantInt{Line: 4, Value: 1777}},
								},
								ast.BareReference{Line: 4, Name: "path"},
							},
						},
					}))
				})
			})

			Describe("operators for sorting", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
0 < 1
1 > 0
5 <= 55
12 >= 21
22 <=> 22
`)
				})

				It("is parsed as an call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.ConstantInt{Line: 1, Value: 0},
							Func:   ast.BareReference{Line: 1, Name: "<"},
							Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 1}},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.ConstantInt{Line: 2, Value: 1},
							Func:   ast.BareReference{Line: 2, Name: ">"},
							Args:   []ast.Node{ast.ConstantInt{Line: 2, Value: 0}},
						},
						ast.CallExpression{
							Line:   3,
							Target: ast.ConstantInt{Line: 3, Value: 5},
							Func:   ast.BareReference{Line: 3, Name: "<="},
							Args:   []ast.Node{ast.ConstantInt{Line: 3, Value: 55}},
						},
						ast.CallExpression{
							Line:   4,
							Target: ast.ConstantInt{Line: 4, Value: 12},
							Func:   ast.BareReference{Line: 4, Name: ">="},
							Args:   []ast.Node{ast.ConstantInt{Line: 4, Value: 21}},
						},
						ast.CallExpression{
							Line:   5,
							Target: ast.ConstantInt{Line: 5, Value: 22},
							Func:   ast.BareReference{Line: 5, Name: "<=>"},
							Args:   []ast.Node{ast.ConstantInt{Line: 5, Value: 22}},
						},
					}))
				})
			})

			Describe("binary boolean operators", func() {
				Context("with simple types on the left and right side", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
1 && 0
1 || 0
1 and 0
1 or 0
`)
					})

					// FIXME: these first two should NOT be call expressions
					It("is parsed as a call expression", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.CallExpression{
								Line:   1,
								Target: ast.ConstantInt{Line: 1, Value: 1},
								Func:   ast.BareReference{Line: 1, Name: "&&"},
								Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 0}},
							},
							ast.CallExpression{
								Line:   2,
								Target: ast.ConstantInt{Line: 2, Value: 1},
								Func:   ast.BareReference{Line: 2, Name: "||"},
								Args:   []ast.Node{ast.ConstantInt{Line: 2, Value: 0}},
							},
							ast.WeakLogicalAnd{
								Line: 3,
								LHS:  ast.ConstantInt{Line: 3, Value: 1},
								RHS:  ast.ConstantInt{Line: 3, Value: 0},
							},
							ast.WeakLogicalOr{
								Line: 4,
								LHS:  ast.ConstantInt{Line: 4, Value: 1},
								RHS:  ast.ConstantInt{Line: 4, Value: 0},
							},
						}))
					})
				})

				Context("with complex types on the left and right side", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`retrieve(:features)[feature] || false`)
					})

					It("is parsed as a call expression", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.CallExpression{
								Target: ast.CallExpression{
									Target: ast.CallExpression{
										Func: ast.BareReference{Name: "retrieve"},
										Args: []ast.Node{ast.Symbol{Name: "features"}},
									},
									Func: ast.BareReference{Name: "[]"},
									Args: []ast.Node{ast.BareReference{Name: "feature"}},
								},
								Func: ast.BareReference{Name: "||"},
								Args: []ast.Node{ast.Boolean{Value: false}},
							},
						}))
					})
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
							Line:   1,
							Target: ast.ConstantInt{Line: 1, Value: 1},
							Func:   ast.BareReference{Line: 1, Name: "=="},
							Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 1}},
						},
						ast.CallExpression{
							Line:   2,
							Target: ast.ConstantInt{Line: 2, Value: 1},
							Func:   ast.BareReference{Line: 2, Name: "==="},
							Args:   []ast.Node{ast.ConstantInt{Line: 2, Value: 1}},
						},
						ast.CallExpression{
							Line:   3,
							Target: ast.ConstantInt{Line: 3, Value: 1},
							Func:   ast.BareReference{Line: 3, Name: "!="},
							Args:   []ast.Node{ast.ConstantInt{Line: 3, Value: 1}},
						},
						ast.CallExpression{
							Line:   4,
							Target: ast.ConstantInt{Line: 4, Value: 1},
							Func:   ast.BareReference{Line: 4, Name: "=~"},
							Args:   []ast.Node{ast.ConstantInt{Line: 4, Value: 1}},
						},
						ast.CallExpression{
							Line:   5,
							Target: ast.ConstantInt{Line: 5, Value: 1},
							Func:   ast.BareReference{Line: 5, Name: "!~"},
							Args:   []ast.Node{ast.ConstantInt{Line: 5, Value: 1}},
						},
					}))
				})
			})
		})

		Describe("arrays", func() {
			Context("with newlines between elements", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`[1,2,3,
4,
5,
6
]`)
				})

				It("is parsed as an Array", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Array{Nodes: []ast.Node{
							ast.ConstantInt{Value: 1},
							ast.ConstantInt{Value: 2},
							ast.ConstantInt{Value: 3},
							ast.ConstantInt{Line: 1, Value: 4},
							ast.ConstantInt{Line: 2, Value: 5},
							ast.ConstantInt{Line: 3, Value: 6},
						}},
					}))
				})
			})

			Context("as the target and argument of a binary operator", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("[1,2,3,4,5] - [1]")
				})

				It("is parsed as a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.Array{Nodes: []ast.Node{
								ast.ConstantInt{Value: 1},
								ast.ConstantInt{Value: 2},
								ast.ConstantInt{Value: 3},
								ast.ConstantInt{Value: 4},
								ast.ConstantInt{Value: 5},
							}},
							Func: ast.BareReference{Name: "-"},
							Args: []ast.Node{
								ast.Array{Nodes: []ast.Node{
									ast.ConstantInt{Value: 1},
								}},
							},
						},
					}))
				})
			})

			Context("of simple built-in types", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("[1,2,3,   4,5,6 ]")
				})

				It("is parsed as an Array", func() {
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

			Context("of named variables", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("[cladosiphonic, capillitial, bicarbureted, argentose]")
				})

				It("is parsed as an Array", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Array{
							Nodes: []ast.Node{
								ast.BareReference{Name: "cladosiphonic"},
								ast.BareReference{Name: "capillitial"},
								ast.BareReference{Name: "bicarbureted"},
								ast.BareReference{Name: "argentose"},
							},
						},
					}))
				})
			})

			Context("of arrays", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("[[], [], []]")
				})

				It("is parsed as an Array", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Array{Nodes: []ast.Node{
							ast.Array{Nodes: []ast.Node{}},
							ast.Array{Nodes: []ast.Node{}},
							ast.Array{Nodes: []ast.Node{}},
						}},
					}))
				})
			})
		})

		Describe("hashes", func() {
			Context("with each key-value pair on a newline and no comma on the last line", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
# these two hashes are the same
{
  :foo => :bar,
  :bar => :foo
}

{
  :foo => :bar,
  :bar => :foo,
}

`)
				})

				It("returns a Hash node", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Hash{
							Line: 2,
							Pairs: []ast.HashKeyValuePair{
								{
									Key:   ast.Symbol{Line: 3, Name: "foo"},
									Value: ast.Symbol{Line: 3, Name: "bar"},
								},
								{
									Key:   ast.Symbol{Line: 4, Name: "bar"},
									Value: ast.Symbol{Line: 4, Name: "foo"},
								},
							},
						},
						ast.Hash{
							Line: 7,
							Pairs: []ast.HashKeyValuePair{
								{
									Key:   ast.Symbol{Line: 8, Name: "foo"},
									Value: ast.Symbol{Line: 8, Name: "bar"},
								},
								{
									Key:   ast.Symbol{Line: 9, Name: "bar"},
									Value: ast.Symbol{Line: 9, Name: "foo"},
								},
							},
						},
					}))
				})
			})

			Context("without anything inside", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("{}")
				})

				It("is parsed as a hash", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Hash{},
					}))
				})
			})

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

			Describe("assigning a value via a setter", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("Sharpware.nasality = 'cladosiphonic-capillitial'")
				})

				It("is parsed as a call expression", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Target: ast.BareReference{Name: "Sharpware"},
							Func:   ast.BareReference{Name: "nasality="},
							Args: []ast.Node{
								ast.SimpleString{Value: "cladosiphonic-capillitial"},
							},
						},
					}))
				})
			})

			Describe("assigning a value to a key", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("hash[:key] = :value")
				})

				It("is parsed as a call expression", func() {
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
							Line: 0,
							Pairs: []ast.HashKeyValuePair{
								{
									Key:   ast.Symbol{Line: 1, Name: "key"},
									Value: ast.BareReference{Line: 1, Name: "value"},
								}, {
									Key:   ast.Symbol{Line: 2, Name: "foo"},
									Value: ast.BareReference{Line: 2, Name: "bar"},
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
						Line: 1,
						LHS:  ast.InstanceVariable{Line: 1, Name: "foo"},
						RHS:  ast.Symbol{Line: 1, Name: "bar"},
					},
					ast.Assignment{
						Line: 2,
						LHS:  ast.InstanceVariable{Line: 2, Name: "FOO"},
						RHS:  ast.Symbol{Line: 2, Name: "baz"},
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
						Line: 1,
						LHS:  ast.ClassVariable{Line: 1, Name: "foo"},
						RHS:  ast.Symbol{Line: 1, Name: "bar"},
					},
					ast.Assignment{
						Line: 2,
						LHS:  ast.ClassVariable{Line: 2, Name: "FOO"},
						RHS:  ast.Symbol{Line: 2, Name: "baz"},
					},
				}))
			})
		})

		Describe("lambdas", func() {
			Context("without any frills", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("something = lambda { puts 'hai'; exit }")
				})

				It("is parsed as an ast.Lambda", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							LHS: ast.BareReference{Name: "something"},
							RHS: ast.Lambda{
								Body: ast.Block{
									Body: []ast.Node{
										ast.CallExpression{
											Func: ast.BareReference{Name: "puts"},
											Args: []ast.Node{ast.SimpleString{Value: "hai"}},
										},
										ast.BareReference{Name: "exit"},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("a method with rescue statements at the end", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
def samsonic_obey
  raise 'whoopsie-daisy'
rescue Nope
  puts 'unlikely'
rescue NotThisEither => e
  puts 'contrived'
rescue Errno::ENOTEMPTY, Errno::ENOENT
  puts 'less contrived'
rescue
  puts 'aww yisss'
end
`)
			})

			It("is parsed as rescue statements", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.FuncDecl{
						Line: 1,
						Name: ast.BareReference{Line: 1, Name: "samsonic_obey"},
						Args: []ast.Node{},
						Body: []ast.Node{
							ast.CallExpression{
								Line: 2,
								Func: ast.BareReference{Line: 2, Name: "raise"},
								Args: []ast.Node{ast.SimpleString{Line: 2, Value: "whoopsie-daisy"}},
							},
						},
						Rescues: []ast.Node{
							ast.Rescue{
								Line: 3,
								Exception: ast.RescueException{
									Classes: []ast.Class{{Line: 3, Name: "Nope"}},
								},
								Body: []ast.Node{ast.CallExpression{
									Line: 4,
									Func: ast.BareReference{Line: 4, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 4, Value: "unlikely"}},
								}},
							},
							ast.Rescue{
								Line: 5,
								Exception: ast.RescueException{
									Classes: []ast.Class{{Line: 5, Name: "NotThisEither"}},
									Var:     ast.BareReference{Line: 5, Name: "e"},
								},
								Body: []ast.Node{ast.CallExpression{
									Line: 6,
									Func: ast.BareReference{Line: 6, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 6, Value: "contrived"}},
								}},
							},
							ast.Rescue{
								Line: 7,
								Exception: ast.RescueException{
									Classes: []ast.Class{
										{
											Line:      7,
											Name:      "ENOTEMPTY",
											Namespace: "Errno",
										}, {
											Line:      7,
											Name:      "ENOENT",
											Namespace: "Errno",
										},
									},
								},
								Body: []ast.Node{ast.CallExpression{
									Line: 8,
									Func: ast.BareReference{Line: 8, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 8, Value: "less contrived"}},
								}},
							},
							ast.Rescue{
								Line:      9,
								Exception: ast.RescueException{},
								Body: []ast.Node{ast.CallExpression{
									Line: 10,
									Func: ast.BareReference{Line: 10, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 10, Value: "aww yisss"}},
								}},
							},
						},
					},
				}))
			})
		})

		Describe("the 'alias' keyword", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
module Foo
  alias wat zomg
  alias seriously? yes_srsly!
end
`)
			})

			It("does not allow commas and converts its 'args' to symbols", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.ModuleDecl{
						Line: 1,
						Name: "Foo",
						Body: []ast.Node{
							ast.Alias{
								Line: 2,
								To:   ast.Symbol{Line: 2, Name: "wat"},
								From: ast.Symbol{Line: 2, Name: "zomg"},
							},
							ast.Alias{
								Line: 3,
								To:   ast.Symbol{Line: 3, Name: "seriously?"},
								From: ast.Symbol{Line: 3, Name: "yes_srsly!"},
							},
						},
					},
				}))
			})
		})

		Describe("the retry keyword", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
begin
  retry if falsey_method()
  retry
end
`)
			})

			It("should be parsed as a Retry node", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Begin{
						Line: 1,
						Body: []ast.Node{
							ast.IfBlock{
								Line: 2,
								Condition: ast.CallExpression{
									Line: 2,
									Args: []ast.Node{},
									Func: ast.BareReference{Line: 2, Name: "falsey_method"},
								},
								Body: []ast.Node{ast.Retry{Line: 2}},
							},
							ast.Retry{Line: 3},
						},
						Rescue: []ast.Node{},
					},
				}))
			})
		})

		Describe("with conditional returns inside a method", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
def method_with_conditional_return
  names.each do |name|
    return Kernel.load(name) if File.exist?(File.expand_path(name))
    return unless false
  end
end
`)
			})

			It("should be parsed correctly", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.FuncDecl{
						Line: 1,
						Name: ast.BareReference{Line: 1, Name: "method_with_conditional_return"},
						Args: []ast.Node{},
						Body: []ast.Node{
							ast.CallExpression{
								Line:   2,
								Target: ast.BareReference{Line: 2, Name: "names"},
								Func:   ast.BareReference{Line: 2, Name: "each"},
								Args:   []ast.Node{},
								OptionalBlock: ast.Block{
									Line: 2,
									Args: []ast.Node{ast.BareReference{Line: 2, Name: "name"}},
									Body: []ast.Node{
										ast.IfBlock{
											Line: 3,
											Condition: ast.CallExpression{
												Line:   3,
												Target: ast.BareReference{Line: 3, Name: "File"},
												Func:   ast.BareReference{Line: 3, Name: "exist?"},
												Args: []ast.Node{
													ast.CallExpression{
														Line:   3,
														Target: ast.BareReference{Line: 3, Name: "File"},
														Func:   ast.BareReference{Line: 3, Name: "expand_path"},
														Args:   []ast.Node{ast.BareReference{Line: 3, Name: "name"}},
													},
												},
											},
											Body: []ast.Node{
												ast.Return{
													Line: 3,
													Value: ast.CallExpression{
														Line:   3,
														Target: ast.BareReference{Line: 3, Name: "Kernel"},
														Func:   ast.BareReference{Line: 3, Name: "load"},
														Args:   []ast.Node{ast.BareReference{Line: 3, Name: "name"}},
													},
												},
											},
										},
										ast.IfBlock{
											Line: 4,
											Condition: ast.Negation{
												Line:   4,
												Target: ast.Boolean{Line: 4, Value: false},
											},
											Body: []ast.Node{
												ast.Return{Line: 4},
											},
										},
									},
								},
							},
						},
					},
				}))
			})
		})

		Describe("with expressions that return a value", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
def with_a_block()
  yield 'nonrestricted-consonantize'
  return 'albitite-compotor', 'catalecta-coassistant', 'semiannual-pomfret'
end
`)
			})

			It("should parse the expressions correctly", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.FuncDecl{
						Line: 1,
						Name: ast.BareReference{Line: 1, Name: "with_a_block"},
						Args: []ast.Node{},
						Body: []ast.Node{
							ast.Yield{
								Line: 2,
								Value: ast.SimpleString{
									Line:  2,
									Value: "nonrestricted-consonantize",
								},
							},
							ast.Return{
								Line: 3,
								Value: ast.Nodes{
									ast.SimpleString{Line: 3, Value: "albitite-compotor"},
									ast.SimpleString{Line: 3, Value: "catalecta-coassistant"},
									ast.SimpleString{Line: 3, Value: "semiannual-pomfret"},
								},
							},
						},
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
							Line: 1,
							Func: ast.BareReference{Line: 1, Name: "function_that_takes_a_block"},
							Args: []ast.Node{},
							OptionalBlock: ast.Block{
								Line: 1,
								Body: []ast.Node{
									ast.CallExpression{
										Line: 2,
										Func: ast.BareReference{Line: 2, Name: "puts"},
										Args: []ast.Node{
											ast.SimpleString{Line: 2, Value: "semiannual-pomfret"},
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
with_a_block do |with, some, args|
  'aww yiss'
end
`)
				})

				It("is parsed as an ast.Block with args", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line: 1,
							Func: ast.BareReference{Line: 1, Name: "with_a_block"},
							Args: []ast.Node{},
							OptionalBlock: ast.Block{
								Line: 1,
								Args: []ast.Node{
									ast.BareReference{Line: 1, Name: "with"},
									ast.BareReference{Line: 1, Name: "some"},
									ast.BareReference{Line: 1, Name: "args"},
								},
								Body: []ast.Node{ast.SimpleString{Line: 2, Value: "aww yiss"}},
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
							Args:   []ast.Node{},
							OptionalBlock: ast.Block{
								Args: []ast.Node{ast.BareReference{Name: "foo"}},
								Body: []ast.Node{
									ast.CallExpression{
										Func: ast.BareReference{Name: "puts"},
										Args: []ast.Node{ast.BareReference{Name: "foo"}},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("ranges", func() {
			Context("of integers", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("-1..-5")
				})

				It("should be parsed as a Range", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Range{
							Start: ast.Negative{Target: ast.ConstantInt{Value: 1}},
							End:   ast.Negative{Target: ast.ConstantInt{Value: 5}},
						},
					}))
				})
			})
		})

		Describe("% notation", func() {
			Describe("for regular expressions", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("%r(string/)")
				})

				It("should parse it as an ast.Regex", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Regex{Value: "string/"},
					}))
				})
			})
		})

		Describe("regex literals", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("/^foo.*bar$/")
			})

			It("is parsed as an ast.Regex", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Regex{Value: "^foo.*bar$"},
				}))
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
							Condition: ast.Negation{Target: ast.Boolean{Value: false}},
							Body:      []ast.Node{ast.ConstantInt{Value: 5}},
						},
					}))
				})
			})

			Context("with a comparision of call expressions", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
unless target[0..1] == config[:config_ext]
end
`)
				})

				It("is parsed as an IfBlock", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Line: 1,
							Condition: ast.Negation{
								Line: 1,
								Target: ast.CallExpression{
									Line: 1,
									Target: ast.CallExpression{
										Line:   1,
										Target: ast.BareReference{Line: 1, Name: "target"},
										Func:   ast.BareReference{Line: 1, Name: "[]"},
										Args: []ast.Node{
											ast.Range{
												Line:  1,
												Start: ast.ConstantInt{Line: 1, Value: 0},
												End:   ast.ConstantInt{Line: 1, Value: 1},
											},
										},
									},
									Func: ast.BareReference{Line: 1, Name: "=="},
									Args: []ast.Node{
										ast.CallExpression{
											Line:   1,
											Target: ast.BareReference{Line: 1, Name: "config"},
											Func:   ast.BareReference{Line: 1, Name: "[]"},
											Args:   []ast.Node{ast.Symbol{Line: 1, Name: "config_ext"}},
										},
									},
								},
							},
							Body: []ast.Node{},
						},
					}))
				})
			})

			Context("inside another unless", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
unless @zomg
  no_wai = 9999 unless yes_wai
  yes_wai = 9999 unless no_wai
end
`)
				})

				It("is parsed as a nested IfBlock", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Line: 1,
							Condition: ast.Negation{
								Line:   1,
								Target: ast.InstanceVariable{Line: 1, Name: "zomg"},
							},
							Body: []ast.Node{
								ast.IfBlock{
									Line: 2,
									Condition: ast.Negation{
										Line:   2,
										Target: ast.BareReference{Line: 2, Name: "yes_wai"},
									},
									Body: []ast.Node{
										ast.Assignment{
											Line: 2,
											LHS:  ast.BareReference{Line: 2, Name: "no_wai"},
											RHS:  ast.ConstantInt{Line: 2, Value: 9999},
										},
									},
								},
								ast.IfBlock{
									Line: 3,
									Condition: ast.Negation{
										Line:   3,
										Target: ast.BareReference{Line: 3, Name: "no_wai"},
									},
									Body: []ast.Node{
										ast.Assignment{
											Line: 3,
											LHS:  ast.BareReference{Line: 3, Name: "yes_wai"},
											RHS:  ast.ConstantInt{Line: 3, Value: 9999},
										},
									},
								},
							},
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
							Line:      1,
							Condition: ast.Boolean{Line: 1, Value: false},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 2, Value: "Romanize-whereover"}},
								},
							},
						},
					}))
				})
			})

			Context("at the end of an expression", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
exit(1) if false
foo = :bar if something.truthy_method
raise OptionError, "description" if args.size < 2
`)
				})

				It("is parsed as an ast.IfBlock", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.IfBlock{
							Line:      1,
							Condition: ast.Boolean{Line: 1, Value: false},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 1,
									Func: ast.BareReference{Line: 1, Name: "exit"},
									Args: []ast.Node{ast.ConstantInt{Line: 1, Value: 1}},
								},
							},
						},
						ast.IfBlock{
							Line: 2,
							Condition: ast.CallExpression{
								Line:   2,
								Target: ast.BareReference{Line: 2, Name: "something"},
								Func:   ast.BareReference{Line: 2, Name: "truthy_method"},
							},
							Body: []ast.Node{
								ast.Assignment{
									Line: 2,
									LHS:  ast.BareReference{Line: 2, Name: "foo"},
									RHS:  ast.Symbol{Line: 2, Name: "bar"},
								},
							},
						},
						ast.IfBlock{
							Line: 3,
							Condition: ast.CallExpression{
								Line: 3,
								Target: ast.CallExpression{
									Line:   3,
									Target: ast.BareReference{Line: 3, Name: "args"},
									Func:   ast.BareReference{Line: 3, Name: "size"},
								},
								Func: ast.BareReference{Line: 3, Name: "<"},
								Args: []ast.Node{ast.ConstantInt{Line: 3, Value: 2}},
							},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 3,
									Func: ast.BareReference{Line: 3, Name: "raise"},
									Args: []ast.Node{
										ast.BareReference{Line: 3, Name: "OptionError"},
										ast.InterpolatedString{Line: 3, Value: "description"},
									},
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
							Line:      1,
							Condition: ast.Boolean{Line: 1, Value: false},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{ast.SimpleString{Line: 2, Value: "Romanize-whereover"}},
								},
							},
							Else: []ast.Node{
								ast.IfBlock{
									Line:      3,
									Condition: ast.Boolean{Line: 3, Value: true},
									Body: []ast.Node{
										ast.CallExpression{
											Line: 4,
											Func: ast.BareReference{Line: 4, Name: "puts"},
											Args: []ast.Node{
												ast.SimpleString{Line: 4, Value: "Kiplingese-disinvolve"},
											},
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
							Line:      1,
							Condition: ast.Boolean{Line: 1, Value: false},
							Body: []ast.Node{
								ast.SimpleString{Line: 2, Value: "purifier-cartouche"},
							},
							Else: []ast.Node{
								ast.IfBlock{
									Line:      3,
									Condition: ast.Boolean{Line: 3, Value: false},
									Body: []ast.Node{
										ast.SimpleString{Line: 4, Value: "bronchophthisis-hypersurface"},
									},
								},
								ast.IfBlock{
									Line:      5,
									Condition: ast.Boolean{Line: 5, Value: false},
									Body: []ast.Node{
										ast.SimpleString{Line: 6, Value: "sharpware-nasality"},
									},
								},
								ast.IfBlock{
									Line:      7,
									Condition: ast.Boolean{Line: 7, Value: true},
									Body: []ast.Node{
										ast.SimpleString{Line: 8, Value: "Osmeridae-harpylike"},
									},
								},
							},
						},
					}))
				})
			})
		})

		Describe("an expression that can fail followed by rescue", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("value = can_raise() rescue 'whoops'")
			})

			It("uses the provided value as a fall back for assignment", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Assignment{
						LHS: ast.BareReference{Name: "value"},
						RHS: ast.RescueModifier{
							Statement: ast.CallExpression{
								Args: []ast.Node{},
								Func: ast.BareReference{Name: "can_raise"},
							},
							Rescue: ast.SimpleString{Value: "whoops"},
						},
					},
				}))
			})
		})

		Describe("rescuing without a class, and capturing the exception thrown", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
begin
rescue => wat
end
`)
			})

			It("should be parsed as a BeginBlock struct", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Begin{
						Line: 1,
						Body: []ast.Node{},
						Rescue: []ast.Node{
							ast.Rescue{
								Line: 2,
								Body: []ast.Node{},
								Exception: ast.RescueException{
									Var: ast.BareReference{Line: 2, Name: "wat"},
								},
							},
						},
					},
				}))
			})
		})

		Describe("an else clause for begin / rescue / else / end", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
begin
  'this always happens'
rescue
  'only when an exception occurs in the begin scope'
else
  'this only happens if there were no exceptions'
end
`)
			})

			It("should be parsed as a BeginBlock struct", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Begin{
						Line: 1,
						Body: []ast.Node{ast.SimpleString{Line: 2, Value: "this always happens"}},
						Rescue: []ast.Node{
							ast.Rescue{
								Line: 3,
								Body: []ast.Node{
									ast.SimpleString{
										Line:  4,
										Value: "only when an exception occurs in the begin scope",
									},
								}},
						},
						Else: []ast.Node{
							ast.SimpleString{Line: 6, Value: "this only happens if there were no exceptions"},
						},
					},
				}))
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
						Line: 1,
						Body: []ast.Node{
							ast.CallExpression{
								Line: 2,
								Func: ast.BareReference{Line: 2, Name: "foo"},
								Args: []ast.Node{},
							},
						},
						Rescue: []ast.Node{
							ast.Rescue{
								Line: 3,
								Body: []ast.Node{
									ast.CallExpression{
										Line: 4,
										Func: ast.BareReference{Line: 4, Name: "bar"},
										Args: []ast.Node{},
									},
								},
							},
							ast.Rescue{
								Line: 5,
								Exception: ast.RescueException{
									Classes: []ast.Class{{Line: 5, Name: "LoadError"}},
								},
								Body: []ast.Node{
									ast.CallExpression{
										Line: 6,
										Func: ast.BareReference{Line: 6, Name: "biz"},
										Args: []ast.Node{},
									},
								},
							},
							ast.Rescue{
								Line: 7,
								Exception: ast.RescueException{
									Var:     ast.BareReference{Line: 7, Name: "e"},
									Classes: []ast.Class{{Line: 7, Name: "Exception"}},
								},
								Body: []ast.Node{
									ast.CallExpression{
										Line: 8,
										Func: ast.BareReference{Line: 8, Name: "baz"},
										Args: []ast.Node{},
									},
								},
							},
						},
					},
				}))
			})
		})

		Describe("ternary ?", func() {
			Context("as the right hand side of an assignment expression", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
s = short ? short.dup : "  "
`)
				})

				It("is parsed correctly", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Assignment{
							Line: 1,
							LHS:  ast.BareReference{Line: 1, Name: "s"},
							RHS: ast.Ternary{
								Line:      1,
								Condition: ast.BareReference{Line: 1, Name: "short"},
								True: ast.CallExpression{
									Line:   1,
									Target: ast.BareReference{Line: 1, Name: "short"},
									Func:   ast.BareReference{Line: 1, Name: "dup"},
								},
								False: ast.InterpolatedString{Line: 1, Value: "  "},
							},
						},
					}))
				})
			})

			Context("with simple values", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("true ? 'string' : 5")
				})

				It("should be parsed as a Ternary node", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Ternary{
							Condition: ast.Boolean{Value: true},
							True:      ast.SimpleString{Value: "string"},
							False:     ast.ConstantInt{Value: 5},
						},
					}))
				})
			})

			Context("with call expressions", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer("what() ? mm.hmm : thats_cool()")
				})

				It("should be parsed as a Ternary node", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Ternary{
							Condition: ast.CallExpression{
								Func: ast.BareReference{Name: "what"},
								Args: []ast.Node{},
							},
							True: ast.CallExpression{
								Target: ast.BareReference{Name: "mm"},
								Func:   ast.BareReference{Name: "hmm"},
							},
							False: ast.CallExpression{
								Func: ast.BareReference{Name: "thats_cool"},
								Args: []ast.Node{},
							},
						},
					}))
				})
			})
		})

		Describe("semicolons", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(";;a; b; c;;")
			})

			It("are treated as a line separator", func() {
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
('hello %s world' % 'grubby')
`)
			})

			It("is treated as a grouping for a set of expressions", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.Group{
						Line: 1,
						Body: []ast.Node{ast.ConstantInt{Line: 1, Value: 1}},
					},
					ast.Group{
						Line: 2,
						Body: []ast.Node{
							ast.SimpleString{Line: 2, Value: "hey"},
							ast.SimpleString{Line: 2, Value: "this"},
							ast.SimpleString{Line: 2, Value: "works!"},
						},
					},
					ast.Group{
						Line: 3,
						Body: []ast.Node{
							ast.CallExpression{
								Line:   3,
								Target: ast.Array{Line: 3, Nodes: []ast.Node{}},
								Func:   ast.BareReference{Line: 3, Name: "unshift"},
							},
						},
					},
					ast.Group{
						Line: 4,
						Body: []ast.Node{
							ast.CallExpression{
								Line:   4,
								Target: ast.SimpleString{Line: 4, Value: "hello %s world"},
								Func:   ast.BareReference{Line: 4, Name: "%"},
								Args:   []ast.Node{ast.SimpleString{Line: 4, Value: "grubby"}},
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
						Line:      1,
						Name:      "Bar",
						Namespace: "Foo",
					},
					ast.CallExpression{
						Line: 2,
						Target: ast.Class{
							Line:      2,
							Name:      "Baz",
							Namespace: "Foo::Bar",
						},
						Func: ast.BareReference{Line: 2, Name: "method_call"},
					},
				}))
			})
		})

		Describe("memoizing methods", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
def memoized_func
  unless @value
    @value = expensive_function_call()
  end

  @value
end
`)
			})

			It("should be parsed correctly", func() {
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.FuncDecl{
						Line: 1,
						Name: ast.BareReference{Line: 1, Name: "memoized_func"},
						Args: []ast.Node{},
						Body: []ast.Node{
							ast.IfBlock{
								Line: 2,
								Condition: ast.Negation{
									Line:   2,
									Target: ast.InstanceVariable{Line: 2, Name: "value"},
								},
								Body: []ast.Node{
									ast.Assignment{
										Line: 3,
										LHS:  ast.InstanceVariable{Line: 3, Name: "value"},
										RHS: ast.CallExpression{
											Line: 3,
											Func: ast.BareReference{Line: 3, Name: "expensive_function_call"},
											Args: []ast.Node{},
										},
									},
								},
							},
							ast.InstanceVariable{Line: 6, Name: "value"},
						},
					},
				}))
			})
		})

		Describe("loops", func() {
			Context("with blocks", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
5.times do
  next
end
`)
				})

				It("is parsed as a block, but the 'next' keyword is valid here", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.CallExpression{
							Line:   1,
							Target: ast.ConstantInt{Line: 1, Value: 5},
							Func:   ast.BareReference{Line: 1, Name: "times"},
							Args:   []ast.Node{},
							OptionalBlock: ast.Block{
								Line: 1,
								Body: []ast.Node{ast.Next{}},
							},
						},
					}))
				})
			})

			Context("with an until statement", func() {
				BeforeEach(func() {
					lexer = parser.NewLexer(`
until 1 == 2
  puts 'INFINITE LOOP AHOY!!!!1'
end
`)
				})

				It("is parsed as a Loop", func() {
					Expect(parser.Statements).To(Equal([]ast.Node{
						ast.Loop{
							Line: 1,
							Condition: ast.Negation{
								Line: 1,
								Target: ast.CallExpression{
									Line:   1,
									Target: ast.ConstantInt{Line: 1, Value: 1},
									Func:   ast.BareReference{Line: 1, Name: "=="},
									Args:   []ast.Node{ast.ConstantInt{Line: 1, Value: 2}},
								},
							},
							Body: []ast.Node{
								ast.CallExpression{
									Line: 2,
									Func: ast.BareReference{Line: 2, Name: "puts"},
									Args: []ast.Node{
										ast.SimpleString{Line: 2, Value: "INFINITE LOOP AHOY!!!!1"},
									},
								},
							},
						},
					}))
				})
			})

			Describe("while statements", func() {
				Context("at the end of an expression", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
'check this out' while false

begin
  puts 'whaaat'
end while true
`)
					})

					It("are valid loops too", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.Loop{
								Line:      1,
								Condition: ast.Boolean{Line: 1, Value: false},
								Body: []ast.Node{
									ast.SimpleString{Line: 1, Value: "check this out"},
								},
							},
							ast.Loop{
								Line:      5,
								Condition: ast.Boolean{Line: 5, Value: true},
								Body: []ast.Node{
									ast.Begin{
										Line: 3,
										Body: []ast.Node{
											ast.CallExpression{
												Line: 4,
												Func: ast.BareReference{Line: 4, Name: "puts"},
												Args: []ast.Node{ast.SimpleString{Line: 4, Value: "whaaat"}},
											},
										},
										Rescue: []ast.Node{},
									},
								},
							},
						}))
					})
				})

				Context("with a trailing end keyword", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
while foo = bar.baz
  puts 'welp'
  break if false
  next if false
end
`)
					})

					It("is parsed into a Loop struct", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.Loop{
								Line: 1,
								Condition: ast.Assignment{
									Line: 1,
									LHS:  ast.BareReference{Line: 1, Name: "foo"},
									RHS: ast.CallExpression{
										Line:   1,
										Target: ast.BareReference{Line: 1, Name: "bar"},
										Func:   ast.BareReference{Line: 1, Name: "baz"},
									},
								},
								Body: []ast.Node{
									ast.CallExpression{
										Line: 2,
										Func: ast.BareReference{Line: 2, Name: "puts"},
										Args: []ast.Node{ast.SimpleString{Line: 2, Value: "welp"}},
									},
									ast.IfBlock{
										Line:      3,
										Condition: ast.Boolean{Line: 3, Value: false},
										Body:      []ast.Node{ast.Break{}},
									},
									ast.IfBlock{
										Line:      4,
										Condition: ast.Boolean{Line: 4, Value: false},
										Body:      []ast.Node{ast.Next{}},
									},
								},
							},
						}))
					})
				})

				Context("with a deeply nested next keyword", func() {
					BeforeEach(func() {
						lexer = parser.NewLexer(`
while true
  if false
    if false
      next
    end
  end
end
`)
					})

					It("parses correctly", func() {
						Expect(parser.Statements).To(Equal([]ast.Node{
							ast.Loop{
								Line:      1,
								Condition: ast.Boolean{Line: 1, Value: true},
								Body: []ast.Node{
									ast.IfBlock{
										Line:      2,
										Condition: ast.Boolean{Line: 2, Value: false},
										Body: []ast.Node{
											ast.IfBlock{
												Line:      3,
												Condition: ast.Boolean{Line: 3, Value: false},
												Body: []ast.Node{
													ast.Next{},
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
			Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).To(BeNil())
			Expect(len(parser.Statements)).To(Equal(4))
		})
	})

	Describe("more complicated expressions", func() {
		Context("bit shift, ternary and trailing 'if'", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
s << (short ? ", " : "  ") if long
`)
			})

			It("is parsed correctly", func() {
				Expect(parser.RubyParse(lexer)).To(BeSuccessful())
				Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).To(BeNil())
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.IfBlock{
						Line:      1,
						Condition: ast.BareReference{Line: 1, Name: "long"},
						Body: []ast.Node{
							ast.CallExpression{
								Line:   1,
								Target: ast.BareReference{Line: 1, Name: "s"},
								Func:   ast.BareReference{Line: 1, Name: "<<"},
								Args: []ast.Node{
									ast.Group{
										Line: 1,
										Body: []ast.Node{
											ast.Ternary{
												Line:      1,
												Condition: ast.BareReference{Line: 1, Name: "short"},
												True:      ast.InterpolatedString{Line: 1, Value: ", "},
												False:     ast.InterpolatedString{Line: 1, Value: "  "},
											},
										},
									},
								},
							},
						},
					},
				}))
			})
		})

		Context("conditionals around assignment to a call expression", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
unless option = match?(opt)
end`)
			})

			It("is parsed correctly", func() {
				Expect(parser.RubyParse(lexer)).To(BeSuccessful())
				Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).To(BeNil())
				Expect(parser.Statements).To(Equal([]ast.Node{
					ast.IfBlock{
						Line: 1,
						Condition: ast.Negation{
							Line: 1,
							Target: ast.Assignment{
								Line: 1,
								LHS:  ast.BareReference{Line: 1, Name: "option"},
								RHS: ast.CallExpression{
									Line: 1,
									Func: ast.BareReference{Line: 1, Name: "match?"},
									Args: []ast.Node{ast.BareReference{Line: 1, Name: "opt"}},
								},
							},
						},
						Body: []ast.Node{},
					},
				}))
			})
		})

		Describe("line numbers", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
:clog                    #1
'hemihypalgesia'         #2
"prematrimonial"         #3
/unrevived/              #4
SANNAITE                 #5
@stickers                #6
@@Penny                  #7
$commercialism           #8
__LINE__                 #9
__FILE__                 #10
Choletelin::Gatter       #11
Incross::Distant.swollen #12
lambda { p 'bakula' }    #13
!'ethnically'            #14
~'viridigenous'          #15
+5                       #16
-6                       #17
true && false            #18
true || false            #19
1 + 1                    #20
1 - 1                    #21
1 * 1                    #22
1 / 1                    #23
0.12                     #24
123                      #25
                         #26
def whippet              #27
  return 'cheesecake'    #28
end                      #29
                         #30
class Turbinidae         #31
end                      #32
                         #33
module Weighage          #34
end                      #35
                         #36
class << self            #37
end                      #38
                         #39
a = 'unfightable'        #40
a, b = 'poisonful'       #41
a ||= 'dastardliness'    #42
                         #43
if a                     #44
  puts 'presecure'       #45
end                      #46
                         #47
begin                    #48
rescue Bananas           #49
end                      #50
                         #51
5.times do               #52
  yield 'butterman'      #53
end                      #54
                         #55
while true               #56
  puts 'disloyalist'     #57
end                      #58
                         #59
case something           #60
  when true              #61
    puts 'reascension'   #62
end                      #63
                         #64
0..12                    #65
                         #66
true ? 'a' : 'b'         #67
                         #68
alias dazing holomorphy  #69
                         #70
5 & 11                   #71
5 | 11                   #72
self                     #73
nil                      #74
`)
			})

			JustBeforeEach(func() {
				Expect(parser.RubyParse(lexer)).To(BeSuccessful())
				Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).To(BeNil())
				Expect(len(parser.Statements)).To(Equal(44))
			})

			It("sets the correct line number for terminal statements", func() {
				Expect(parser.Statements[0].LineNumber()).To(Equal(1))
				Expect(parser.Statements[1].LineNumber()).To(Equal(2))
				Expect(parser.Statements[2].LineNumber()).To(Equal(3))
				Expect(parser.Statements[3].LineNumber()).To(Equal(4))
				Expect(parser.Statements[4].LineNumber()).To(Equal(5))
				Expect(parser.Statements[5].LineNumber()).To(Equal(6))
				Expect(parser.Statements[6].LineNumber()).To(Equal(7))
				Expect(parser.Statements[7].LineNumber()).To(Equal(8))
				Expect(parser.Statements[8].LineNumber()).To(Equal(9))
				Expect(parser.Statements[9].LineNumber()).To(Equal(10))
				Expect(parser.Statements[10].LineNumber()).To(Equal(11))
				Expect(parser.Statements[11].LineNumber()).To(Equal(12))
				Expect(parser.Statements[12].LineNumber()).To(Equal(13))
				Expect(parser.Statements[13].LineNumber()).To(Equal(14))
				Expect(parser.Statements[14].LineNumber()).To(Equal(15))
				Expect(parser.Statements[15].LineNumber()).To(Equal(16))
				Expect(parser.Statements[16].LineNumber()).To(Equal(17))
				Expect(parser.Statements[17].LineNumber()).To(Equal(18))
				Expect(parser.Statements[18].LineNumber()).To(Equal(19))
				Expect(parser.Statements[19].LineNumber()).To(Equal(20))
				Expect(parser.Statements[20].LineNumber()).To(Equal(21))
				Expect(parser.Statements[21].LineNumber()).To(Equal(22))
				Expect(parser.Statements[22].LineNumber()).To(Equal(23))
				Expect(parser.Statements[23].LineNumber()).To(Equal(24))
				Expect(parser.Statements[24].LineNumber()).To(Equal(25))
				Expect(parser.Statements[25].LineNumber()).To(Equal(27))
				Expect(parser.Statements[26].LineNumber()).To(Equal(31))
				Expect(parser.Statements[27].LineNumber()).To(Equal(34))
				Expect(parser.Statements[28].LineNumber()).To(Equal(37))
				Expect(parser.Statements[29].LineNumber()).To(Equal(40))
				Expect(parser.Statements[30].LineNumber()).To(Equal(41))
				Expect(parser.Statements[31].LineNumber()).To(Equal(42))
				Expect(parser.Statements[32].LineNumber()).To(Equal(44))
				Expect(parser.Statements[33].LineNumber()).To(Equal(48))
				Expect(parser.Statements[34].LineNumber()).To(Equal(52))
				Expect(parser.Statements[35].LineNumber()).To(Equal(56))
				Expect(parser.Statements[36].LineNumber()).To(Equal(60))
				Expect(parser.Statements[37].LineNumber()).To(Equal(65))
				Expect(parser.Statements[38].LineNumber()).To(Equal(67))
				Expect(parser.Statements[39].LineNumber()).To(Equal(69))
				Expect(parser.Statements[40].LineNumber()).To(Equal(71))
				Expect(parser.Statements[41].LineNumber()).To(Equal(72))
				Expect(parser.Statements[42].LineNumber()).To(Equal(73))
				Expect(parser.Statements[43].LineNumber()).To(Equal(74))
			})
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
        foo ||
          bar
        foo &&
          bar


      	end

  abc    =   123

  !  true
  ~    true
  +   5
  -    123

  a = 5 or
    false
  b = a and
    true
end

with_a_block { |foo| puts foo.inspect } # comment goes here
func.with_a_block { |foo | puts foo.inspect    } # all the comments # yep
`)
		})

		It("parses just fine", func() {
			Expect(parser.RubyParse(lexer)).To(BeSuccessful())
			Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).To(BeNil())
		})
	})

	Describe("invalid ruby", func() {
		JustBeforeEach(func() {
			Expect(parser.RubyParse(lexer)).ToNot(BeSuccessful())
			Expect(lexer.(*parser.ConcreteStatefulRubyLexer).LastError).To(HaveOccurred())
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

		PContext("when the 'next' keyword is outside of a loop or block", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("next")
			})

			It("fails to parse", func() {
				Expect(parser.Statements).To(BeEmpty())
			})
		})

		XContext("when the 'break' keyword is outside of a loop or block", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("break")
			})

			It("fails to parse", func() {
				Expect(parser.Statements).To(BeEmpty())
			})
		})

		PContext("when the 'return' keyword is outside of a method", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("return 5")
			})

			It("fails to parse", func() {
				Expect(parser.Statements).To(BeEmpty())
			})
		})

		PContext("when a method is provided two blocks", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer(`
takes_a_block(&baz) do
  puts 'FAIL'
end
`)
			})

			It("is fails to parse", func() {
				Expect(parser.Statements).To(BeEmpty())
			})
		})

		PContext("when a splat precedes a reference outside of a call expression or method declaration", func() {
			BeforeEach(func() {
				lexer = parser.NewLexer("*[1,2,3]")
			})

			It("is fails to parse", func() {
				Expect(parser.Statements).To(BeEmpty())
			})
		})
	})
})
