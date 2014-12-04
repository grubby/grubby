package parser

func lexPercentSign(l *StatefulRubyLexer) stateFn {
	stringType := tokenTypeDoubleQuoteString
	if l.accept("r") {
		stringType = tokenTypeRegex
		l.start += 1
	}

	if l.accept("`~!@#$%^&*-_=+()[]{}<>\\|;:'\",./?") {
		delimiter := closingDelimiter(l.input[l.start+1 : l.pos])

		l.ignore()
		var r, prev rune
		for {
			prev = r
			switch r = l.next(); {
			case string(r) == delimiter && prev != '\\':
				l.backup()
				l.emit(stringType)
				l.next()
				l.ignore() // ignore closing delimiter
				return lexAnything
			case r == eof:
				l.emit(tokenTypeError)
				return lexAnything
			}
		}
	} else {
		l.emit(tokenTypeOperator)
	}

	return lexAnything
}

func closingDelimiter(openingDelimiter string) string {
	switch openingDelimiter {
	case "{":
		return "}"
	case "(":
		return ")"
	case "<":
		return ">"
	case "[":
		return "]"
	default:
		return openingDelimiter
	}
}
