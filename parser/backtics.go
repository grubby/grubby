package parser

func lexBacktics(l *StatefulRubyLexer) stateFn {
	l.start += 1 // ignore the opening backtic

	isEscaped := false
	for r := l.peek(); r != '`' && !isEscaped; r = l.next() {
		isEscaped = (r == '\\')

		if r == eof {
			l.emit(tokenTypeError)
			return lexSomething
		}
	}

	l.pos -= 1 // ignore closing backtic
	l.emit(tokenTypeSubshell)
	l.accept("`")
	l.ignore()

	return lexSomething
}
