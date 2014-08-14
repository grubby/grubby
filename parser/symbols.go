package parser

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphaNumericUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func lexSymbol(l *StatefulRubyLexer) stateFn {
	if !l.accept(alpha + "_") {
		l.emit(tokenTypeColon)
		for l.accept(":") {
			l.emit(tokenTypeColon)
		}

		return lexAnything
	}

	l.acceptRun(alphaNumericUnderscore)
	l.emit(tokenTypeSymbol)
	return lexAnything
}
