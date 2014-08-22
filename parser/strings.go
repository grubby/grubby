package parser

func lexSingleQuoteString(l *StatefulRubyLexer) stateFn {
	var (
		r    rune
		prev rune
	)

	for {
		prev = r
		switch r = l.next(); {
		case r == '\'' && prev != '\'':
			l.emit(tokenTypeString)
			return lexAnything
		case r == eof:
			l.emit(tokenTypeError)
			return lexAnything
		}
	}

	return lexAnything
}

func lexDoubleQuoteString(l *StatefulRubyLexer) stateFn {
	var (
		r    rune
		prev rune
	)

	l.start += 1

	for {
		prev = r
		switch r = l.next(); {
		case r == '"' && prev != '\\':
			l.pos -= 1
			l.emit(tokenTypeDoubleQuoteString)
			l.next()
			l.ignore()
			return lexAnything
		case r == eof:
			l.emit(tokenTypeError)
			return lexAnything
		}
	}

	return lexAnything
}
