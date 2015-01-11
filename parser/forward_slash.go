package parser

import "fmt"

func lexSlash(l StatefulRubyLexer) stateFn {
	switch l.lastToken().typ {
	case tokenTypeInteger:
		parseAsOperator(l)
	case tokenTypeFloat:
		parseAsOperator(l)
	case tokenTypeString:
		parseAsOperator(l)
	case tokenTypeDoubleQuoteString:
		parseAsOperator(l)
	case tokenTypeCharacter:
		parseAsOperator(l)
	case tokenTypeSymbol:
		parseAsOperator(l)
	case tokenTypeReference:
		parseAsOperator(l)
	case tokenTypeCapitalizedReference:
		parseAsOperator(l)
	case tokenTypeGlobal:
		parseAsOperator(l)
	case tokenTypeLParen:
		parseAsRegex(l)
	case tokenTypeRParen:
		parseAsOperator(l)
	case tokenTypeComma:
		parseAsRegex(l)
	case tokenTypeNewline:
		parseAsRegex(l)
	case tokenTypeDEF:
		parseAsRegex(l)
	case tokenTypeDO:
		parseAsRegex(l)
	case tokenTypeEND:
		parseAsRegex(l)
	case tokenTypeIF:
		parseAsRegex(l)
	case tokenTypeELSE:
		parseAsRegex(l)
	case tokenTypeELSIF:
		parseAsRegex(l)
	case tokenTypeUNLESS:
		parseAsRegex(l)
	case tokenTypeTRUE:
		parseAsOperator(l)
	case tokenTypeFALSE:
		parseAsOperator(l)
	case tokenTypeLessThan:
		parseAsRegex(l)
	case tokenTypeGreaterThan:
		parseAsRegex(l)
	case tokenTypeColon:
		parseAsRegex(l)
	case tokenTypeSemicolon:
		parseAsRegex(l)
	case tokenTypeEqual:
		parseAsRegex(l)
	case tokenTypeBang:
		parseAsRegex(l)
	case tokenTypeTilde:
		parseAsRegex(l)
	case tokenTypeUnaryPlus:
		parseAsRegex(l)
	case tokenTypeBinaryPlus:
		parseAsRegex(l)
	case tokenTypeBinaryMinus:
		parseAsRegex(l)
	case tokenTypeUnaryMinus:
		parseAsRegex(l)
	case tokenTypeStar:
		parseAsRegex(l)
	case tokenTypeLBracket:
		parseAsRegex(l)
	case tokenTypeRBracket:
		parseAsRegex(l)
	case tokenTypeLBrace:
		parseAsRegex(l)
	case tokenTypeRBrace:
		parseAsRegex(l)
	case tokenType__FILE__:
		parseAsOperator(l)
	case tokenType__LINE__:
		parseAsOperator(l)
	case tokenTypeDot:
		parseAsOperator(l)
	case tokenTypePipe:
		parseAsRegex(l)
	case tokenTypeSubshell:
		parseAsOperator(l)
	case tokenTypeOperator:
		parseAsRegex(l)
	case tokenTypeBEGIN:
		parseAsRegex(l)
	case tokenTypeRESCUE:
		parseAsRegex(l)
	case tokenTypeENSURE:
		parseAsRegex(l)
	case tokenTypeBREAK:
		parseAsRegex(l)
	case tokenTypeNEXT:
		parseAsRegex(l)
	case tokenTypeREDO:
		parseAsRegex(l)
	case tokenTypeRETRY:
		parseAsRegex(l)
	case tokenTypeRETURN:
		parseAsRegex(l)
	case tokenTypeYIELD:
		parseAsRegex(l)
	case tokenTypeQuestionMark:
		parseAsRegex(l)
	case tokenTypeMethodName:
		parseAsRegex(l)
	case tokenTypeWHILE:
		parseAsRegex(l)
	case tokenTypeAND:
		parseAsRegex(l)
	case tokenTypeOR:
		parseAsRegex(l)
	case tokenTypeLAMBDA:
		parseAsRegex(l)
	case tokenTypeCASE:
		parseAsRegex(l)
	case tokenTypeWHEN:
		parseAsRegex(l)
	case tokenTypeOrEquals:
		parseAsRegex(l)
	case tokenTypeRange:
		parseAsRegex(l)
	case tokenTypeError:
		parseAsRegex(l)
	case tokenTypeSELF:
		parseAsOperator(l)
	case tokenTypeNIL:
		parseAsOperator(l)
	default:
		panic(fmt.Sprintf("Unknown node preceding '/' :: '%#v'", l.lastToken))
	}

	return lexSomething
}

func parseAsRegex(l StatefulRubyLexer) {
	l.ignore() // ignore opening '/'

	var r, prev rune
	shouldBreak := false

	for {
		prev = r
		switch r = l.next(); {
		case r == '/' && prev != '\\':
			l.backup()
			l.emit(tokenTypeRegex)
			l.accept("/")
			l.ignore() // ignore closing slash
			shouldBreak = true
		case r == eof:
			l.emit(tokenTypeError)
			shouldBreak = true
		}

		if shouldBreak {
			break
		}
	}
}

func parseAsOperator(l StatefulRubyLexer) {
	if l.accept("=") {
		l.emit(tokenTypeOperator)
	} else {
		l.emit(tokenTypeForwardSlash)
	}
}
