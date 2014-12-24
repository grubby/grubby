package parser

func lexBacktics(l StatefulRubyLexer) stateFn {
	l.moveCurrentTokenStartIndex(1) // ignore the opening backtic

	isEscaped := false
	for r := l.peek(); r != '`' && !isEscaped; r = l.next() {
		isEscaped = (r == '\\')

		if r == eof {
			l.emit(tokenTypeError)
			return lexSomething
		}
	}

	l.moveCurrentPositionIndex(-1)
	l.emit(tokenTypeSubshell)
	l.accept("`")
	l.ignore()

	return lexSomething
}
