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

		l.acceptRun(whitespace)
		l.ignore()

		if (len(l.input) > l.start+5) && l.input[l.start:l.start+5] == "self." {
			l.acceptRun("self")
			l.emit(tokenTypeReference)
			l.accept(".")
			l.emit(tokenTypeDot)
		}

		if l.accept(validMethodNameRunes) {
			l.acceptRun(validMethodNameRunes)
			l.accept("=")
			l.emit(tokenTypeReference)
		}
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
	case "unless":
		l.emit(tokenTypeUNLESS)
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
	case "__LINE__":
		l.emit(tokenType__LINE__)
	case "__ENCODING__":
		l.emit(tokenType__ENCODING__)
	case "for":
		l.emit(tokenTypeFOR)
	case "while":
		l.emit(tokenTypeWHILE)
	case "until":
		l.emit(tokenTypeUNTIL)
	case "begin":
		l.emit(tokenTypeBEGIN)
	case "rescue":
		l.emit(tokenTypeRESCUE)
	case "ensure":
		l.emit(tokenTypeENSURE)
	case "break":
		l.emit(tokenTypeBREAK)
	case "next":
		l.emit(tokenTypeNEXT)
	case "redo":
		l.emit(tokenTypeREDO)
	case "retry":
		l.emit(tokenTypeRETRY)
	case "return":
		l.emit(tokenTypeRETURN)
	case "yield":
		l.emit(tokenTypeYIELD)
	case "and":
		l.emit(tokenTypeAND)
	case "or":
		l.emit(tokenTypeOR)
	case "lambda":
		l.emit(tokenTypeLAMBDA)
	case "case":
		l.emit(tokenTypeCASE)
	case "when":
		l.emit(tokenTypeWHEN)
	default:
		r, _ := utf8.DecodeRuneInString(l.input[l.start:])

		if l.accept("?!") {
			l.emit(tokenTypeMethodName)
		} else if unicode.IsUpper(r) {
			l.emit(tokenTypeCapitalizedReference)
		} else {
			l.emit(tokenTypeReference)
		}
	}
	return lexSomething
}
