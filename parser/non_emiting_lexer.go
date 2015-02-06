package parser

import "strings"

type nonEmitingLexer struct {
	lexer  StatefulRubyLexer
	Tokens []token
}

func NewNonEmitingLexer(l StatefulRubyLexer) *nonEmitingLexer {
	return &nonEmitingLexer{lexer: l}
}

func (l *nonEmitingLexer) next() (r rune) {
	return l.lexer.next()
}

func (l *nonEmitingLexer) backup() {
	l.lexer.backup()
}

func (l *nonEmitingLexer) ignore() {
	l.lexer.ignore()
}

// peek returns but does not consume
// the next rune in the input.
func (l *nonEmitingLexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

// accepts a single rune from valid
func (l *nonEmitingLexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *nonEmitingLexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *nonEmitingLexer) CurrentLineNumber() int {
	return l.lexer.CurrentLineNumber()
}

func (l *nonEmitingLexer) parsedNewLine() {
	l.lexer.parsedNewLine()
}

func (l *nonEmitingLexer) emit(t tokenType) {
	l.Tokens = append(l.Tokens, token{
		typ:   t,
		value: l.lexer.currentSlice(),
		line:  l.lexer.CurrentLineNumber(),
	})
	l.lexer.ignore()
}

func (l *nonEmitingLexer) lastToken() token {
	return l.Tokens[len(l.Tokens)-1]
}

func (l *nonEmitingLexer) slice(start, end int) string {
	return l.lexer.slice(start, end)
}

func (l *nonEmitingLexer) currentSlice() string {
	return l.lexer.currentSlice()
}

func (l *nonEmitingLexer) moveCurrentTokenStartIndex(val int) {
	l.lexer.moveCurrentTokenStartIndex(val)
}

func (l *nonEmitingLexer) moveCurrentPositionIndex(val int) {
	l.lexer.moveCurrentPositionIndex(val)
}

func (l *nonEmitingLexer) setCurrentPositionIndex(val int) {
	l.lexer.setCurrentPositionIndex(val)
}

func (l *nonEmitingLexer) startIndex() int {
	return l.lexer.startIndex()
}

func (l *nonEmitingLexer) currentIndex() int {
	return l.lexer.currentIndex()
}

func (l *nonEmitingLexer) lengthOfInput() int {
	return l.lexer.lengthOfInput()
}

func (l *nonEmitingLexer) Error(error string) {
	l.lexer.Error(error)
}

func (l *nonEmitingLexer) Lex(lval *RubySymType) int {
	return l.lexer.Lex(lval)
}
