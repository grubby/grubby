package parser

func lexComment(l *StatefulRubyLexer) stateFn {
	for r := l.next(); r != '\n' && r != eof; r = l.next() {
	}

	l.backup()
	l.ignore()
	return lexSomething
}
