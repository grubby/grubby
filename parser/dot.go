package parser

func lexDot(l StatefulRubyLexer) stateFn {
	if l.accept(".") {
		l.emit(tokenTypeRange)
		return lexSomething
	}

	l.emit(tokenTypeDot)
	l.acceptRun(whitespace + newline)
	l.ignore()

	if l.accept(validMethodNameRunes) {
		l.acceptRun(validMethodNameRunes)
		l.emit(tokenTypeReference)
	}

	return lexSomething
}
