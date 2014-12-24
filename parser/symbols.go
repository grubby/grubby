package parser

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphaNumericUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const validMethodNameRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!?"

func lexSymbol(l StatefulRubyLexer) stateFn {
	if !l.accept(alpha + "_@\"") {
		if l.accept(":") {
			l.acceptRun(alphaNumericUnderscore)

			for l.currentIndex()+2 < l.lengthOfInput() && l.slice(l.currentIndex(), l.currentIndex()+2) == "::" {
				l.moveCurrentPositionIndex(2)
				l.acceptRun(alphaNumericUnderscore)
			}

			l.emit(tokenTypeNamespaceResolvedModule)
		} else {
			l.emit(tokenTypeColon)
		}

		return lexSomething
	}

	// skip past the initial colon
	l.moveCurrentTokenStartIndex(1)

	// some dynamic symbols can start with " and '
	if l.slice(l.currentIndex()-1, l.currentIndex()) == "\"" {
		var (
			r    rune
			prev rune
		)

		l.ignore() // ignore : and opening quote

		for {
			prev = r
			switch r = l.next(); {
			case r == '#':
				if l.accept("{") {
					// check that we close the #{} template if present
					for innerR := l.next(); innerR != '}'; innerR = l.next() {
						if innerR == eof {
							l.emit(tokenTypeError)
						}
					}
				}
			case r == '"' && prev != '\\':
				l.moveCurrentPositionIndex(-1)
				l.emit(tokenTypeSymbol)
				l.next()
				l.ignore()
				return lexSomething
			case r == eof:
				l.emit(tokenTypeError)
				return lexSomething
			}
		}
	}

	l.accept("@")
	l.acceptRun(alphaNumericUnderscore + "?!")
	l.emit(tokenTypeSymbol)
	return lexSomething
}
