package parser

func lexLessThan(l *StatefulRubyLexer) stateFn {
	if l.accept("<") {
		heredocEndsAtFirstColumn := true
		if l.accept("-") {
			l.start += 3
			heredocEndsAtFirstColumn = false
		} else {
			l.start += 2
		}

		if l.accept(alphaNumericUnderscore) {
			l.acceptRun(alphaNumericUnderscore)
			heredocIdentifier := l.input[l.start:l.pos]
			l.accept("\n")
			l.ignore()

			for {
				r := l.next()
				if r == eof {
					l.emit(tokenTypeError)
				}

				if r == '\n' {
					beginningOfLine := l.pos
					if !heredocEndsAtFirstColumn {
						l.acceptRun(whitespace)
					}

					beginningOfHeredoc := l.pos

					if l.accept(alphaNumericUnderscore) {
						l.acceptRun(alphaNumericUnderscore)
						if l.input[beginningOfHeredoc:l.pos] == heredocIdentifier {
							l.pos = beginningOfLine - 1
							l.emit(tokenTypeDoubleQuoteString)
							l.accept("\n")
							l.acceptRun(whitespace)
							l.acceptRun(alphaNumericUnderscore)
							l.ignore()
							break
						}
					}
				}
			}
		} else {
			l.start -= 2
			l.emit(tokenTypeOperator)
		}
	} else if l.accept("=") {
		l.accept(">")
		l.emit(tokenTypeOperator)
	} else {
		l.emit(tokenTypeLessThan)
	}

	return lexAnything
}
