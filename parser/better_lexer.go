package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/grubby/grubby/ast"
)

const (
	eof rune = iota
)

type token struct {
	typ   tokenType
	value string
}

type tokenType int

const (
	tokenTypeError tokenType = iota
	tokenTypeEOF
	tokenTypeInteger
	tokenTypeFloat
)

type BetterRubyLexer struct {
	input string
	start int
	pos   int
	width int // width of last rune read from input

	tokens chan token

	LastError error
}

type stateFn func(*BetterRubyLexer) stateFn

func NewBetterLexer(input string) *BetterRubyLexer {
	lexer := &BetterRubyLexer{
		input:  input,
		tokens: make(chan token),
	}

	go lexer.run()
	return lexer
}

func (lexer *BetterRubyLexer) run() {
	for state := lexAnything; state != nil; {
		state = state(lexer)
	}

	close(lexer.tokens)
}

func lexAnything(l *BetterRubyLexer) stateFn {
	for {
		switch r := l.next(); {
		case '0' <= r && r <= '9':
			l.backup()
			return lexNumber
		}

		if l.next() == eof {
			break
		}
	}

	l.emit(tokenTypeEOF)
	return nil
}

const digits = "0123456789"

func lexNumber(l *BetterRubyLexer) stateFn {
	l.acceptRun(digits)
	if l.next() == '.' {
		l.acceptRun(digits)
		l.emit(tokenTypeFloat)
		return lexAnything
	}

	l.emit(tokenTypeInteger)
	return lexAnything
}

func (l *BetterRubyLexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *BetterRubyLexer) backup() {
	l.pos -= l.width
}

// ignore skips over the pending input before this point.
func (l *BetterRubyLexer) ignore() {
	l.start = l.pos
}

// peek returns but does not consume
// the next rune in the input.
func (l *BetterRubyLexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

// accepts a single rune from valid
func (l *BetterRubyLexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *BetterRubyLexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (lexer *BetterRubyLexer) emit(t tokenType) {
	lexer.tokens <- token{
		typ:   t,
		value: lexer.input[lexer.start:lexer.pos],
	}
}

func (lexer *BetterRubyLexer) Lex(lval *RubySymType) int {
	debug("Called Lex()")
	defer func() { debug("") }()

	for token := range lexer.tokens {
		switch token.typ {
		case tokenTypeInteger:
			debug("integer: %s", token.value)
			intVal, err := strconv.Atoi(token.value)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantInt{Value: intVal}
			return NODE // Consider: should this be a different type?
		case tokenTypeFloat:
			debug("float: %s", token.value)
			floatval, err := strconv.ParseFloat(token.value, 64)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantFloat{Value: floatval}
			return NODE // as above, maybe a different type?
		case tokenTypeEOF:
			debug("EOF")
			return EOF
		default:
			panic(fmt.Sprintf("unknown token: '%s'", token))
		}
	}

	return 0
}

func (lexer *BetterRubyLexer) Error(error string) {
	lexer.LastError = errors.New(fmt.Sprintf("syntax error: %s\n", error))
}
