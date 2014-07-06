package parser

import (
	"fmt"
	"math"
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
		intvals := []int{int(c - '0')}

		c = rune(l.s[l.pos])
		for unicode.IsDigit(c) {
			intvals = append(intvals, int(c-'0'))

			l.pos += 1
			if l.pos == len(l.s) {
				l.pos--
				break
			}
			c = rune(l.s[l.pos])
		}

		pow := -1
		for i := len(intvals) - 1; i >= 0; i-- {
			// for i := 0; i < len(intvals); i++ {
			pow += 1
			inc := intvals[i] * int(math.Pow(10, float64(pow)))
			lval.intval += inc
		}

		return DIGIT
	}

	return int(c)
}

func (l *rubyLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
