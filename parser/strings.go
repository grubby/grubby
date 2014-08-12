package parser

func lexSingleQuoteString(l *BetterRubyLexer) stateFn {
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
