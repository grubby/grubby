package parser

import (
	"unicode"
	"unicode/utf8"
)

func lexReference(l StatefulRubyLexer) stateFn {
	l.acceptRun(alphaNumericUnderscore)

	defaultCase := func() {
		r, _ := utf8.DecodeRuneInString(l.slice(l.startIndex(), l.startIndex()+1))

		if l.accept("?!") {
			l.emit(tokenTypeMethodName)
		} else if unicode.IsUpper(r) {
			l.emit(tokenTypeCapitalizedReference)
		} else {
			l.emit(tokenTypeReference)
		}
	}

	switch l.currentSlice() {
	case "def":
		l.emit(tokenTypeDEF)

		l.acceptRun(whitespace)
		l.ignore()

		if l.lengthOfInput() > l.startIndex()+5 && l.slice(l.startIndex(), l.startIndex()+5) == "self." {
			l.acceptRun("self")
			l.emit(tokenTypeSELF)
			l.accept(".")
			l.emit(tokenTypeDot)
		}

		if l.accept(validMethodNameRunes) {
			l.acceptRun(validMethodNameRunes)
			l.accept("=")
			l.emit(tokenTypeReference)
		} else if l.accept("<") {
			if l.accept("=") {
				l.accept(">")
			} else {
				l.accept("<")
			}
			l.emit(tokenTypeOperator)
		} else if l.accept(">") {
			l.accept("=")
			l.emit(tokenTypeOperator)
		} else if l.accept("=") {
			l.acceptRun("=")
			l.emit(tokenTypeOperator)
		}
	case "defined":
		if l.accept("?") {
			l.emit(tokenTypeDEFINED)
		} else {
			defaultCase()
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
	case "self":
		l.emit(tokenTypeSELF)
	case "nil":
		l.emit(tokenTypeNIL)
	case "alias":
		l.emit(tokenTypeALIAS)
		readCommaDelimitedSymbolsUntilEOL(l)

	default:
		defaultCase()
	}
	return lexSomething
}

func readSymbolsUntilEOL(l StatefulRubyLexer) {
	l.acceptRun(whitespace)
	l.ignore()
	for l.accept(alphaNumericUnderscore) {
		l.acceptRun(alphaNumericUnderscore)
		l.accept("?!")
		l.emit(tokenTypeSymbol)
		l.acceptRun(whitespace)
		l.ignore()
	}
}

func readCommaDelimitedSymbolsUntilEOL(l StatefulRubyLexer) {
	l.acceptRun(whitespace)
	l.ignore()
	for l.accept(alphaNumericUnderscore) {
		l.acceptRun(alphaNumericUnderscore)
		l.accept("?!")
		l.emit(tokenTypeSymbol)
		l.acceptRun(whitespace)
		l.ignore()

		l.accept(",")
		l.acceptRun(whitespace)
		l.ignore()
	}
}
