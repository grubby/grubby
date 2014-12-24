package parser

import "fmt"

func lexSlash(l StatefulRubyLexer) stateFn {
	parseAsRegex := func() {
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

	parseAsOperator := func() {
		if l.accept("=") {
			l.emit(tokenTypeOperator)
		} else {
			l.emit(tokenTypeForwardSlash)
		}
	}

	switch l.lastToken().typ {
	case tokenTypeInteger:
		parseAsOperator()
	case tokenTypeFloat:
		parseAsOperator()
	case tokenTypeString:
		parseAsOperator()
	case tokenTypeDoubleQuoteString:
		parseAsOperator()
	case tokenTypeCharacter:
		parseAsOperator()
	case tokenTypeSymbol:
		parseAsOperator()
	case tokenTypeReference:
		parseAsOperator()
	case tokenTypeCapitalizedReference:
		parseAsOperator()
	case tokenTypeGlobal:
		parseAsOperator()
	case tokenTypeLParen:
		parseAsRegex()
	case tokenTypeRParen:
		parseAsOperator()
	case tokenTypeComma:
		parseAsRegex()
	case tokenTypeNewline:
		parseAsRegex()
	case tokenTypeDEF:
		parseAsRegex()
	case tokenTypeDO:
		parseAsRegex()
	case tokenTypeEND:
		parseAsRegex()
	case tokenTypeIF:
		parseAsRegex()
	case tokenTypeELSE:
		parseAsRegex()
	case tokenTypeELSIF:
		parseAsRegex()
	case tokenTypeUNLESS:
		parseAsRegex()
	case tokenTypeTRUE:
		parseAsOperator()
	case tokenTypeFALSE:
		parseAsOperator()
	case tokenTypeLessThan:
		parseAsRegex()
	case tokenTypeGreaterThan:
		parseAsRegex()
	case tokenTypeColon:
		parseAsRegex()
	case tokenTypeSemicolon:
		parseAsRegex()
	case tokenTypeEqual:
		parseAsRegex()
	case tokenTypeBang:
		parseAsRegex()
	case tokenTypeTilde:
		parseAsRegex()
	case tokenTypePlus:
		parseAsRegex()
	case tokenTypeMinus:
		parseAsRegex()
	case tokenTypeStar:
		parseAsRegex()
	case tokenTypeLBracket:
		parseAsRegex()
	case tokenTypeRBracket:
		parseAsRegex()
	case tokenTypeLBrace:
		parseAsRegex()
	case tokenTypeRBrace:
		parseAsRegex()
	case tokenType__FILE__:
		parseAsOperator()
	case tokenTypeDot:
		parseAsOperator()
	case tokenTypePipe:
		parseAsRegex()
	case tokenTypeSubshell:
		parseAsOperator()
	case tokenTypeOperator:
		parseAsRegex()
	case tokenTypeBEGIN:
		parseAsRegex()
	case tokenTypeRESCUE:
		parseAsRegex()
	case tokenTypeENSURE:
		parseAsRegex()
	case tokenTypeBREAK:
		parseAsRegex()
	case tokenTypeNEXT:
		parseAsRegex()
	case tokenTypeREDO:
		parseAsRegex()
	case tokenTypeRETRY:
		parseAsRegex()
	case tokenTypeRETURN:
		parseAsRegex()
	case tokenTypeYIELD:
		parseAsRegex()
	case tokenTypeQuestionMark:
		parseAsRegex()
	case tokenTypeMethodName:
		parseAsRegex()
	case tokenTypeWHILE:
		parseAsRegex()
	case tokenTypeAND:
		parseAsRegex()
	case tokenTypeOR:
		parseAsRegex()
	case tokenTypeLAMBDA:
		parseAsRegex()
	case tokenTypeCASE:
		parseAsRegex()
	case tokenTypeWHEN:
		parseAsRegex()
	case tokenTypeOrEquals:
		parseAsRegex()
	case tokenTypeRange:
		parseAsRegex()
	case tokenTypeError:
		parseAsRegex()
	default:
		panic(fmt.Sprintf("Unknown node preceding '/' :: '%#v'", l.lastToken))
	}

	return lexSomething
}
