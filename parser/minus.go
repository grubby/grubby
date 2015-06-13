package parser

import "fmt"

func lexMinus(l StatefulRubyLexer) stateFn {
	if l.accept("=") {
		l.emit(tokenTypeOperator)
		return lexSomething
	}

	switch l.lastToken().typ {
	case tokenTypeInteger:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeFloat:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeString:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeDoubleQuoteString:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeCharacter:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeSymbol:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeReference:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeCapitalizedReference:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeGlobal:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeLParen:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeRParen:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeComma:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeNewline:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeDEF:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeDO:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeEND:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeIF:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeELSE:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeELSIF:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeUNLESS:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeTRUE:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeFALSE:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeLessThan:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeGreaterThan:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeColon:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeSemicolon:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeEqual:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeBang:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeTilde:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeUnaryMinus:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeBinaryMinus:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeBinaryPlus:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeUnaryPlus:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeStar:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeLBracket:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeRBracket:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeLBrace:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeRBrace:
		l.emit(tokenTypeBinaryMinus)
	case tokenType__FILE__:
		l.emit(tokenTypeBinaryMinus)
	case tokenType__LINE__:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeDot:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypePipe:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeSubshell:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeOperator:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeHashRocket:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeBEGIN:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeRESCUE:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeENSURE:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeBREAK:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeNEXT:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeREDO:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeRETRY:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeRETURN:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeYIELD:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeQuestionMark:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeMethodName:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeWHILE:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeAND:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeOR:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeLAMBDA:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeCASE:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeWHEN:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeOrEquals:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeAndEquals:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeRange:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeExclusiveRange:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeError:
		l.emit(tokenTypeUnaryMinus)
	case tokenTypeSELF:
		l.emit(tokenTypeBinaryMinus)
	case tokenTypeNIL:
		l.emit(tokenTypeBinaryMinus)
	default:
		panic(fmt.Sprintf("Unknown node preceding '-' :: '%#v'", l.lastToken()))
	}

	return lexSomething
}
