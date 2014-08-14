package parser

const digits = "0123456789"

func lexNumber(l *StatefulRubyLexer) stateFn {
	l.acceptRun(digits)
	if l.peek() == '.' {
		l.accept(".")
		l.acceptRun(digits)
		l.emit(tokenTypeFloat)
		return lexAnything
	}

	l.emit(tokenTypeInteger)
	return lexAnything
}
