package parser

func lexComment(l *BetterRubyLexer) stateFn {
	for r := l.next(); r != '\n' && r != eof; r = l.next() {
	}

	l.backup()
	l.ignore()
	return lexAnything
}
