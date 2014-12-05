package parser

func lexSingleQuoteString(l *StatefulRubyLexer) stateFn {
	var (
		r    rune
		prev rune
	)

	l.ignore() // ignore single quote

	for {
		prev = r
		switch r = l.next(); {
		case r == '\'' && prev != '\\':
			l.backup()
			l.emit(tokenTypeString)
			l.accept("'")
			l.ignore()
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
		case r == '#':
			if l.accept("{") {
				lexUntilClosingMatchingBraces('{', '}')(l)
			}
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

func lexUntilClosingMatchingBraces(openingBrace, closingBrace rune) func(*StatefulRubyLexer) {
	return func(l *StatefulRubyLexer) {
		for {
			switch r := l.next(); {
			case r == openingBrace:
				lexUntilClosingMatchingBraces(openingBrace, closingBrace)(l)
			case r == closingBrace:
				return
			case r == eof:
				l.emit(tokenTypeError)
				return
			}
		}
	}
}
