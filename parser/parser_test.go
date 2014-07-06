package parser_test

import (
	"github.com/grubby/grubby/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("goyacc parser", func() {
	var (
		line  string
		lexer parser.RubyLexer
	)

	JustBeforeEach(func() {
		lexer = parser.NewLexer(line)
		parser.Statements = []interface{}{}
		Expect(parser.RubyParse(lexer)).To(Equal(0))
	})

	Describe("parsing simple types", func() {
		BeforeEach(func() {
			line = "5\n"
		})

		It("works, mostly", func() {
			Expect(parser.Statements).To(Equal([]interface{}{5}))
		})
	})
})
