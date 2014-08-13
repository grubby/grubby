package parser

import (
	"unicode"
	"unicode/utf8"
)

func lexReference(l *BetterRubyLexer) stateFn {
	l.acceptRun(alphaNumericUnderscore)

	switch l.input[l.start:l.pos] {
	case "def":
		l.emit(tokenTypeDEF)
	case "end":
		l.emit(tokenTypeEND)
	case "class":
		l.emit(tokenTypeCLASS)
	case "module":
		l.emit(tokenTypeMODULE)
	case "true":
		l.emit(tokenTypeTRUE)
	case "false":
		l.emit(tokenTypeFALSE)
	default:
		r, _ := utf8.DecodeRuneInString(l.input[l.start:])
		if unicode.IsUpper(r) {
			l.emit(tokenTypeCapitalizedReference)
		} else {
			l.emit(tokenTypeReference)
		}
	}
	return lexAnything
}
