package parser

const digits = "0123456789"

func lexNumber(l StatefulRubyLexer) stateFn {
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

	l.emit(tokenTypeInteger)
	return lexSomething
}
