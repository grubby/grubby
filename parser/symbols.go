package parser

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphaNumericUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const validMethodNameRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!?"

func lexSymbol(l *StatefulRubyLexer) stateFn {
	if !l.accept(alpha + "_") {
		if l.accept(":") {
			l.emit(tokenTypeDoubleColon)
		} else {
			l.emit(tokenTypeColon)
		}

		return lexAnything
	}

	l.acceptRun(alphaNumericUnderscore)
	l.emit(tokenTypeSymbol)
	return lexAnything
}
