package parser

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphaNumericUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func lexSymbol(l *BetterRubyLexer) stateFn {
	l.accept(alpha + "_")
	l.acceptRun(alphaNumericUnderscore)
	l.emit(tokenTypeSymbol)
	return lexAnything
}
