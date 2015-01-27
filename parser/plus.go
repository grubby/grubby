package parser

import "fmt"

func lexPlus(l StatefulRubyLexer) stateFn {
	if l.accept("=") {
		l.emit(tokenTypeOperator)
		return lexSomething
	}

	switch l.lastToken().typ {
	case tokenTypeInteger:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeFloat:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeString:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeDoubleQuoteString:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeCharacter:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeSymbol:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeReference:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeCapitalizedReference:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeGlobal:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeLParen:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeRParen:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeComma:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeNewline:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeDEF:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeDO:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeEND:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeIF:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeELSE:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeELSIF:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeUNLESS:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeTRUE:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeFALSE:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeLessThan:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeGreaterThan:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeColon:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeSemicolon:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeEqual:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeBang:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeTilde:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeUnaryPlus:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeBinaryPlus:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeBinaryMinus:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeUnaryMinus:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeStar:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeLBracket:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeRBracket:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeLBrace:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeRBrace:
		l.emit(tokenTypeBinaryPlus)
	case tokenType__FILE__:
		l.emit(tokenTypeBinaryPlus)
	case tokenType__LINE__:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeDot:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypePipe:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeSubshell:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeOperator:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeBEGIN:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeRESCUE:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeENSURE:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeBREAK:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeNEXT:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeREDO:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeRETRY:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeRETURN:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeYIELD:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeQuestionMark:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeMethodName:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeWHILE:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeAND:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeOR:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeLAMBDA:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeCASE:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeWHEN:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeOrEquals:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeRange:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeError:
		l.emit(tokenTypeUnaryPlus)
	case tokenTypeSELF:
		l.emit(tokenTypeBinaryPlus)
	case tokenTypeNIL:
		l.emit(tokenTypeBinaryPlus)
	default:
		panic(fmt.Sprintf("Unknown node preceding '+' :: '%#v'", l.lastToken()))
	}

	return lexSomething
}
