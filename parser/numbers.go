package parser

const digits = "0123456789"

func lexNumber(l *StatefulRubyLexer) stateFn {
	l.acceptRun(digits)
	if l.accept(".") {
		if l.accept(digits) {
			l.acceptRun(digits)
			l.emit(tokenTypeFloat)
		} else if l.accept(alphaNumericUnderscore + "!") {
			l.backup()
			l.backup()
			l.emit(tokenTypeInteger)
			l.accept(".")
			l.emit(tokenTypeDot)
			return lexReference
		}
		return lexAnything
	}

	l.emit(tokenTypeInteger)
	return lexAnything
}
