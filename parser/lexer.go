package parser

import (
	"fmt"
	"unicode"
)

type rubyLex struct {
	s   string
	pos int
}

func NewLexer(str string) *rubyLex {
	return &rubyLex{s: str}
}

func (l *rubyLex) Lex(lval *RubySymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c - '0')
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c - 'a')
		return LETTER
	}

	return int(c)
}

func (l *rubyLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
