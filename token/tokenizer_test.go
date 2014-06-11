package token_test

import (
	. "github.com/grubby/grubby/token"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tokenizing", func() {
	It("returns a single string back", func() {
		Expect(Tokenize("'hello world'")).To(Equal([]string{"'hello world'"}))
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

		Expect(Tokenize(input)).To(Equal([]string{
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

		Expect(Tokenize(input)).To(Equal([]string{
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
})
