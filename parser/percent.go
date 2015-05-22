package parser

func lexPercentSign(l StatefulRubyLexer) stateFn {
	stringType := tokenTypeDoubleQuoteString

	if l.accept("r") {
		stringType = tokenTypeRegex
		l.moveCurrentTokenStartIndex(1)
	} else if l.accept("Q") {
		l.moveCurrentTokenStartIndex(1)
	} else if l.accept("w") {
		l.moveCurrentTokenStartIndex(1)
		stringType = tokenTypeArrayOfWhitespaceSeparatedStrings
	}

	if l.accept("`~!@#$%^&*-_=+()[]{}<>\\|;:'\",./?") {
		delimiter := closingDelimiter(l.currentSlice()[1:])

		l.ignore()
		var r, prev1, prev2 rune
		for {
			prev2 = prev1
			prev1 = r
			switch r = l.next(); {
			case string(r) == delimiter && (prev1 != '\\' || prev1 == '\\' && prev2 == '\\'):
				l.backup()
				l.emit(stringType)
				l.next()
				l.ignore() // ignore closing delimiter
				return lexSomething
			case r == '#' && l.accept("{"):
				lexUntilClosingMatchingBraces('{', '}')(l)
			case r == eof:
				l.emit(tokenTypeError)
				return lexSomething
			}
		}
	} else {
		l.emit(tokenTypeOperator)
	}

	return lexSomething
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
