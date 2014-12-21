package parser

func lexLessThan(l StatefulRubyLexer) stateFn {
	if l.accept("<") {
		heredocEndsAtFirstColumn := true
		if l.accept("-") {
			l.moveCurrentTokenStartIndex(3)
			heredocEndsAtFirstColumn = false
		} else {
			l.moveCurrentTokenStartIndex(2)
		}

		if l.accept(alphaNumericUnderscore) {
			l.acceptRun(alphaNumericUnderscore)
			heredocIdentifier := l.currentSlice()

			//            Were You Aware???
			//            =================
			// ruby heredocs start **after** a newline
			// this means that you can put otherwise valid ruby
			// ***AFTER*** the heredoc identifier (on the same line)
			// so we must continue lexing until we see a newline
			// then read the heredoc until we see the closing identifier
			readNewline := false
			nonEmitingLexer := NewNonEmitingLexer(l)
			for readNewline == false {
				switch r := l.next(); {
				case r == '\n':
					readNewline = true
					break
				default:
					l.backup()
					stateFn := lexSomething(nonEmitingLexer)
					if stateFn != nil {
						stateFn(nonEmitingLexer)
					}
				}
			}

			l.accept("\n")
			l.ignore()

			for {
				r := l.next()
				if r == eof {
					l.emit(tokenTypeError)
				}

				if r == '\n' {
					beginningOfLine := l.currentIndex()
					if !heredocEndsAtFirstColumn {
						l.acceptRun(whitespace)
					}

					beginningOfHeredoc := l.currentIndex()

					if l.accept(alphaNumericUnderscore) {
						l.acceptRun(alphaNumericUnderscore)
						if l.slice(beginningOfHeredoc, l.currentIndex()) == heredocIdentifier {
							l.setCurrentPositionIndex(beginningOfLine - 1)
							l.emit(tokenTypeDoubleQuoteString)
							l.accept("\n")
							l.acceptRun(whitespace)
							l.acceptRun(alphaNumericUnderscore)
							l.ignore()

							// quickly emit the rest of the tokens that preceded the heredoc
							concreteLexer := nonEmitingLexer.lexer.(*ConcreteStatefulRubyLexer)
							for _, t := range nonEmitingLexer.Tokens {
								concreteLexer.emitToken(t)
							}

							break
						}
					}
				}
			}
		} else {
			l.moveCurrentTokenStartIndex(-2)
			l.emit(tokenTypeOperator)
		}
	} else if l.accept("=") {
		l.accept(">")
		l.emit(tokenTypeOperator)
	} else {
		l.emit(tokenTypeLessThan)
	}

	return lexSomething
}
