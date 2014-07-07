package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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
		return l.parseDigitsOrFail(lval, c)
	}

	return int(c)
}

func (l *rubyLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func (l *rubyLex) parseDigitsOrFail(lval *RubySymType, c rune) int {
	integerDigits := []int{int(c - '0')}

	if l.pos < len(l.s) {
		c = rune(l.s[l.pos])
		for unicode.IsDigit(c) {
			integerDigits = append(integerDigits, int(c-'0'))

			l.pos += 1
			if l.pos == len(l.s) {
				l.pos--
				break
			}
			c = rune(l.s[l.pos])
		}

		// check if we have a float
		if c == '.' {
			l.pos++
			return l.parseFloatOrFail(lval, integerDigits)
		}
	}

	pow := -1
	for i := len(integerDigits) - 1; i >= 0; i-- {
		pow += 1
		inc := integerDigits[i] * int(math.Pow(10, float64(pow)))
		lval.intval += inc
	}

	return DIGIT
}

func (l *rubyLex) parseFloatOrFail(lval *RubySymType, intDigits []int) int {
	if l.pos < len(l.s) {
		fractionalDigits := []int{}

		c := rune(l.s[l.pos])
		for unicode.IsDigit(c) {
			fractionalDigits = append(fractionalDigits, int(c-'0'))

			l.pos++
			if l.pos == len(l.s) {
				l.pos--
				break
			}
			c = rune(l.s[l.pos])
		}

		if len(fractionalDigits) == 0 {
			return -1 // does not compute
		}

		mantissa := ""
		for _, f := range fractionalDigits {
			mantissa += fmt.Sprintf("%d", f)
		}

		integer := ""
		for _, d := range intDigits {
			integer += fmt.Sprintf("%d", d)
		}

		var err error
		float := strings.Join([]string{integer, mantissa}, ".")
		lval.floatval, err = strconv.ParseFloat(float, 64)

		if err != nil {
			panic("FREAKOUT: " + err.Error())
		}
	}

	return FLOAT
}

func parseIntSliceAsInteger(digits []int) int {
	pow := -1
	val := 0
	for i := len(digits) - 1; i >= 0; i-- {
		pow += 1
		inc := digits[i] * int(math.Pow(10, float64(pow)))
		val += inc
	}

	return val
}
