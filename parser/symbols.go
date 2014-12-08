package parser

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphaNumericUnderscore = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const validMethodNameRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!?"

func lexSymbol(l *StatefulRubyLexer) stateFn {
	if !l.accept(alpha + "_@\"") {
		if l.accept(":") {
			l.emit(tokenTypeDoubleColon)
		} else {
			l.emit(tokenTypeColon)
		}

		return lexAnything
	}

	// some dynamic symbols can start with " and '
	if l.input[l.pos-1:l.pos] == "\"" {
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
				l.pos -= 1
				l.emit(tokenTypeSymbol)
				l.next()
				l.ignore()
				return lexAnything
			case r == eof:
				l.emit(tokenTypeError)
				return lexAnything
			}
		}
	}

	l.accept("@")
	l.acceptRun(alphaNumericUnderscore + "?!")
	l.emit(tokenTypeSymbol)
	return lexAnything
}
