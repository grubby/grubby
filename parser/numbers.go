package parser

const digits = "0123456789"

func lexNumber(l StatefulRubyLexer) stateFn {
	if l.accept("0") && l.peek() == 'x' {
		l.accept("x")
		l.acceptRun(alphaNumeric)
		l.emit(tokenTypeInteger)
		return lexSomething
	} else {
		l.acceptRun(digits)
		if l.accept(".") {
			if l.accept(digits) {
				l.acceptRun(digits)
				l.emit(tokenTypeFloat)
				return lexSomething
			} else {
				l.backup()
				l.emit(tokenTypeInteger)
				return lexSomething
			}
		}
	}

	l.emit(tokenTypeInteger)
	return lexSomething
}
