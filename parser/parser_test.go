package parser_test

import (
	"github.com/grubby/grubby/ast"
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("goyacc parser", func() {
	var (
		lexer parser.RubyLexer
	)

	JustBeforeEach(func() {
		parser.Statements = make([]ast.Node, 0)
		Expect(parser.RubyParse(lexer)).To(Equal(0))
	})

	Describe("parsing an integer", func() {
		BeforeEach(func() {
			lexer = parser.NewLexer("5\n")
		})

		It("works, mostly", func() {
			Expect(parser.Statements).To(Equal([]ast.Node{
				ast.ConstantInt{Value: 5},
			}))
		})
	})

	Describe("parsing a float", func() {
		BeforeEach(func() {
			lexer = parser.NewLexer("123.4567\n")
		})

		It("works, mostly", func() {
			Expect(parser.Statements).To(Equal([]ast.Node{
				ast.ConstantFloat{Value: 123.4567},
			}))
		})
	})
})
