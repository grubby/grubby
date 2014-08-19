package parser

import (
	"unicode"
	"unicode/utf8"
)

func lexReference(l *StatefulRubyLexer) stateFn {
	l.acceptRun(alphaNumericUnderscore)

	switch l.input[l.start:l.pos] {
	case "def":
		l.emit(tokenTypeDEF)
	case "do":
		l.emit(tokenTypeDO)
	case "end":
		l.emit(tokenTypeEND)
	case "if":
		l.emit(tokenTypeIF)
	case "else":
		l.emit(tokenTypeELSE)
	case "elsif":
		l.emit(tokenTypeELSIF)
	case "class":
		l.emit(tokenTypeCLASS)
	case "module":
		l.emit(tokenTypeMODULE)
	case "true":
		l.emit(tokenTypeTRUE)
	case "false":
		l.emit(tokenTypeFALSE)
	case "__FILE__":
		l.emit(tokenType__FILE__)
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
