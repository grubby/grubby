package parser_test

import (
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("goyacc parser", func() {
	var (
		lexer parser.RubyLexer
	)

	JustBeforeEach(func() {
		parser.Statements = []interface{}{}
		Expect(parser.RubyParse(lexer)).To(Equal(0))
	})

	Describe("parsing an integer", func() {
		BeforeEach(func() {
			lexer = parser.NewLexer("5\n")
		})

		It("works, mostly", func() {
			Expect(parser.Statements).To(Equal([]interface{}{5}))
		})
	})

	Describe("parsing a float", func() {
		BeforeEach(func() {
			lexer = parser.NewLexer("123.4567\n")
		})

		It("works, mostly", func() {
			Expect(parser.Statements).To(Equal([]interface{}{123.4567}))
		})
	})
})
