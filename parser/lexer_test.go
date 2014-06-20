package parser_test

import (
	. "github.com/grubby/grubby/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tokenizing", func() {
	var lexer Lexer
	BeforeEach(func() {
		lexer = NewLexer()
	})

	It("returns a single string back", func() {
		Expect(lexer.Tokenize("'hello world'")).To(Equal([]string{"'hello world'"}))
	})

	It("understands multiline strings", func() {
		input := `
'hello
world'`

		Expect(lexer.Tokenize(input)).To(Equal([]string{"'hello\nworld'"}))
	})

	XDescribe("collections", func() {
		It("has tokens for opening and closing array brackets", func() {
			input := `
[1,2,3]
`

			Expect(lexer.Tokenize(input)).To(Equal([]string{
				"[",
				"1,",
				"2,",
				"3",
				"]",
			}))
		})
	})

	It("returns a chunk of ruby code split on delimiters", func() {
		input := `
require 'foo'

class Whatever < Foo
  def self.something
    puts 'HAI'
  end

  def something; puts 'BAI'; end;
end
`

		Expect(lexer.Tokenize(input)).To(Equal([]string{
			"require",
			"'foo'",
			"class",
			"Whatever",
			"<",
			"Foo",
			"def",
			"self.something",
			"puts",
			"'HAI'",
			"end",
			"def",
			"something",
			"puts",
			"'BAI'",
			"end",
			"end",
		}))
	})

	It("splits arguments inside parens", func() {
		input := `
def foo
  puts('HAI', 'AND', 'BAI')
end
`

		Expect(lexer.Tokenize(input)).To(Equal([]string{
			"def",
			"foo",
			"puts",
			"(",
			"'HAI',",
			"'AND',",
			"'BAI'",
			")",
			"end",
		}))
	})

	Describe("RubyLexer interface", func() {
		var symbol *RubySymType

		BeforeEach(func() {
			symbol = &RubySymType{}
		})

		It("identifies a single digit", func() {
			lexer.Tokenize("666")

			Expect(lexer.Lex(symbol)).To(Equal(DIGIT), "Expected to lex a DIGIT")
			Expect(symbol.Val).To(Equal(666))
		})

		It("identifies a float", func() {
			lexer.Tokenize("6.66")

			Expect(lexer.Lex(symbol)).To(Equal(FLOAT), "Expected to lex a FLOAT")
			Expect(symbol.FloatVal).To(Equal(6.66))
		})
	})
})
