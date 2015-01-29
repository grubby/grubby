package parser

import "fmt"

func lexAmpersand(l StatefulRubyLexer) stateFn {
	if l.accept("&") {
		l.emit(tokenTypeOperator)
		return lexSomething
	}

	switch l.lastToken().typ {
	case tokenTypeInteger:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeFloat:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeString:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeDoubleQuoteString:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeCharacter:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeSymbol:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeReference:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeCapitalizedReference:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeGlobal:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeLParen:
		parseAsProcArg(l)
	case tokenTypeRParen:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeComma:
		parseAsProcArg(l)
	case tokenTypeNewline:
		parseAsProcArg(l)
	case tokenTypeDEF:
		parseAsProcArg(l)
	case tokenTypeDO:
		parseAsProcArg(l)
	case tokenTypeEND:
		parseAsProcArg(l)
	case tokenTypeIF:
		parseAsProcArg(l)
	case tokenTypeELSE:
		parseAsProcArg(l)
	case tokenTypeELSIF:
		parseAsProcArg(l)
	case tokenTypeUNLESS:
		parseAsProcArg(l)
	case tokenTypeTRUE:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeFALSE:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeLessThan:
		parseAsProcArg(l)
	case tokenTypeGreaterThan:
		parseAsProcArg(l)
	case tokenTypeColon:
		parseAsProcArg(l)
	case tokenTypeSemicolon:
		parseAsProcArg(l)
	case tokenTypeEqual:
		parseAsProcArg(l)
	case tokenTypeBang:
		parseAsProcArg(l)
	case tokenTypeTilde:
		parseAsProcArg(l)
	case tokenTypeUnaryPlus:
		parseAsProcArg(l)
	case tokenTypeBinaryPlus:
		parseAsProcArg(l)
	case tokenTypeBinaryMinus:
		parseAsProcArg(l)
	case tokenTypeUnaryMinus:
		parseAsProcArg(l)
	case tokenTypeStar:
		parseAsProcArg(l)
	case tokenTypeLBracket:
		parseAsProcArg(l)
	case tokenTypeRBracket:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeLBrace:
		parseAsProcArg(l)
	case tokenTypeRBrace:
		parseAsBinaryBitwiseOperator(l)
	case tokenType__FILE__:
		parseAsBinaryBitwiseOperator(l)
	case tokenType__LINE__:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeDot:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypePipe:
		parseAsProcArg(l)
	case tokenTypeSubshell:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeOperator:
		parseAsProcArg(l)
	case tokenTypeBEGIN:
		parseAsProcArg(l)
	case tokenTypeRESCUE:
		parseAsProcArg(l)
	case tokenTypeENSURE:
		parseAsProcArg(l)
	case tokenTypeBREAK:
		parseAsProcArg(l)
	case tokenTypeNEXT:
		parseAsProcArg(l)
	case tokenTypeREDO:
		parseAsProcArg(l)
	case tokenTypeRETRY:
		parseAsProcArg(l)
	case tokenTypeRETURN:
		parseAsProcArg(l)
	case tokenTypeYIELD:
		parseAsProcArg(l)
	case tokenTypeQuestionMark:
		parseAsProcArg(l)
	case tokenTypeMethodName:
		parseAsProcArg(l)
	case tokenTypeWHILE:
		parseAsProcArg(l)
	case tokenTypeAND:
		parseAsProcArg(l)
	case tokenTypeOR:
		parseAsProcArg(l)
	case tokenTypeLAMBDA:
		parseAsProcArg(l)
	case tokenTypeCASE:
		parseAsProcArg(l)
	case tokenTypeWHEN:
		parseAsProcArg(l)
	case tokenTypeOrEquals:
		parseAsProcArg(l)
	case tokenTypeRange:
		parseAsProcArg(l)
	case tokenTypeError:
		parseAsProcArg(l)
	case tokenTypeSELF:
		parseAsBinaryBitwiseOperator(l)
	case tokenTypeNIL:
		parseAsBinaryBitwiseOperator(l)
	default:
		panic(fmt.Sprintf("Unknown node preceding '&' :: '%#v'", l.lastToken))
	}

	return lexSomething
}

func parseAsProcArg(l StatefulRubyLexer) {
	l.ignore() // ignore the initial &
	l.emit(tokenTypeProcArg)
}

func parseAsBinaryBitwiseOperator(l StatefulRubyLexer) {
	l.emit(tokenTypeAmpersand)
}
