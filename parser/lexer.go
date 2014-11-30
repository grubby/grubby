package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/grubby/grubby/ast"
)

var DebugStatements = []string{}

const eof = -1

type token struct {
	typ   tokenType
	value string
}

type tokenType int

const (
	tokenTypeError tokenType = iota
	tokenTypeEOF
	tokenTypeInteger
	tokenTypeFloat
	tokenTypeString
	tokenTypeDoubleQuoteString
	tokenTypeRegex
	tokenTypeCharacter
	tokenTypeSymbol
	tokenTypeReference
	tokenTypeMethodName
	tokenTypeGlobal
	tokenTypeCapitalizedReference
	tokenTypeNewline
	tokenTypeLParen
	tokenTypeRParen
	tokenTypeComma
	tokenTypeDEF
	tokenTypeDO
	tokenTypeEND
	tokenTypeIF
	tokenTypeELSE
	tokenTypeELSIF
	tokenTypeUNLESS
	tokenTypeCLASS
	tokenTypeMODULE
	tokenTypeTRUE
	tokenTypeFALSE
	tokenTypeLessThan
	tokenTypeGreaterThan
	tokenTypeColon
	tokenTypeDoubleColon
	tokenTypeSemicolon
	tokenTypeEqual
	tokenTypeBang
	tokenTypeTilde
	tokenTypePlus
	tokenTypeMinus
	tokenTypeStar
	tokenTypeLBracket
	tokenTypeRBracket
	tokenTypeLBrace
	tokenTypeRBrace
	tokenTypeDollarSign
	tokenTypeAtSign
	tokenTypeDot
	tokenTypeRange
	tokenTypePipe
	tokenTypeOrEquals
	tokenTypeForwardSlash
	tokenTypeAmpersand
	tokenTypeSubshell
	tokenTypeOperator
	tokenTypeQuestionMark
	tokenTypeFOR
	tokenTypeWHILE
	tokenTypeUNTIL
	tokenTypeBEGIN
	tokenTypeRESCUE
	tokenTypeENSURE
	tokenTypeBREAK
	tokenTypeNEXT
	tokenTypeREDO
	tokenTypeRETRY
	tokenTypeRETURN
	tokenTypeYIELD
	tokenTypeAND
	tokenTypeOR
	tokenTypeLAMBDA
	tokenTypeCASE
	tokenTypeWHEN
	tokenType__FILE__
	tokenType__LINE__
	tokenType__ENCODING__
)

type StatefulRubyLexer struct {
	input string
	start int
	pos   int
	width int // width of last rune read from input

	tokens chan token

	lastToken token
	LastError error
}

type stateFn func(*StatefulRubyLexer) stateFn

func NewLexer(input string) *StatefulRubyLexer {
	lexer := &StatefulRubyLexer{
		input:  input,
		tokens: make(chan token),
	}

	go lexer.run()
	return lexer
}

func (lexer *StatefulRubyLexer) run() {
	for state := lexAnything; state != nil; {
		state = state(lexer)
	}

	close(lexer.tokens)
}

func lexAnything(l *StatefulRubyLexer) stateFn {
	for {
		switch r := l.next(); {
		case '0' <= r && r <= '9':
			l.backup()
			return lexNumber
		case r == '\'':
			return lexSingleQuoteString
		case r == '"':
			return lexDoubleQuoteString
		case r == '?':
			// FIXME: this is not an exhaustive list of character literals
			if l.accept(validMethodNameRunes + "-^:'\"") {
				l.start += 1 // skip past the ?
				l.acceptRun(validMethodNameRunes + "-^:'\"")
				l.emit(tokenTypeCharacter)
			} else {
				l.emit(tokenTypeQuestionMark)
			}
		case r == ':':
			l.start += 1 // skip past the colon
			return lexSymbol
		case r == ';':
			l.emit(tokenTypeSemicolon)
		case r == ' ' || r == '\t':
			l.acceptRun(whitespace)
			l.ignore()
		case r == '\n':
			l.backup()
			return lexNewlines
		case ('a' <= r && r <= 'z') || r == '_' || ('A' <= r && r <= 'Z'):
			l.backup()
			return lexReference
		case r == '(':
			l.emit(tokenTypeLParen)
		case r == ')':
			l.emit(tokenTypeRParen)
		case r == ',':
			l.emit(tokenTypeComma)
		case r == '#':
			return lexComment
		case r == '<':
			return lexLessThan
		case r == '>':
			if l.accept("=") || l.accept(">") {
				l.emit(tokenTypeOperator)
			} else {
				l.emit(tokenTypeGreaterThan)
			}
		case r == '=':
			if l.accept("=") {
				l.accept("=")
				l.emit(tokenTypeOperator)
			} else if l.accept("~") {
				l.emit(tokenTypeOperator)
			} else if l.accept(">") {
				l.emit(tokenTypeOperator)
			} else {
				l.emit(tokenTypeEqual)
			}
		case r == '!':
			if l.accept("=") {
				l.emit(tokenTypeOperator)
			} else if l.accept("~") {
				l.emit(tokenTypeOperator)
			} else {
				l.emit(tokenTypeBang)
			}
		case r == '~':
			l.emit(tokenTypeTilde)
		case r == '+':
			if l.accept("=") {
				l.emit(tokenTypeOperator)
			} else {
				l.emit(tokenTypePlus)
			}
		case r == '-':
			if l.accept("=") {
				l.emit(tokenTypeOperator)
			} else {
				l.emit(tokenTypeMinus)
			}
		case r == '*':
			if l.accept("=") {
				l.emit(tokenTypeOperator)
			} else if l.accept("*") {
				l.emit(tokenTypeOperator)
			} else {
				l.emit(tokenTypeStar)
			}
		case r == '[':
			l.emit(tokenTypeLBracket)
		case r == ']':
			l.emit(tokenTypeRBracket)
		case r == '{':
			l.emit(tokenTypeLBrace)
		case r == '}':
			l.emit(tokenTypeRBrace)
		case r == '$':
			validGlobalNameRunes := alphaNumericUnderscore + ":"
			if l.accept(validGlobalNameRunes) {
				l.backup()
				l.ignore()
				l.acceptRun(validGlobalNameRunes)
				l.emit(tokenTypeGlobal)
			} else {
				l.emit(tokenTypeDollarSign)
			}
		case r == '@':
			l.emit(tokenTypeAtSign)
		case r == '.':
			if l.accept(".") {
				l.emit(tokenTypeRange)
				return lexAnything
			}

			l.emit(tokenTypeDot)

			if l.accept(validMethodNameRunes) {
				l.acceptRun(validMethodNameRunes)
				l.emit(tokenTypeReference)
			}
		case r == '|':
			if l.accept("|") {
				if l.accept("=") {
					l.emit(tokenTypeOrEquals)
				} else {
					l.emit(tokenTypeOperator)
				}
			} else {
				l.emit(tokenTypePipe)
			}
		case r == '/':
			return lexSlash
		case r == '&':
			if l.accept("&") {
				l.emit(tokenTypeOperator)
			} else {
				l.emit(tokenTypeAmpersand)
			}
		case r == '%':
			if l.accept("r") {
				if l.accept("(") {
					l.ignore()

					var r, prev rune
					for {
						prev = r
						switch r = l.next(); {
						case r == ')' && prev != '\\':
							l.backup()
							l.emit(tokenTypeRegex)
							l.next()
							l.ignore() // ignore closing paren
							return lexAnything
						case r == eof:
							l.emit(tokenTypeError)
							return lexAnything
						}
					}
				} else {
					l.emit(tokenTypeOperator)
				}
			} else {
				l.emit(tokenTypeOperator)
			}
		case r == '^':
			l.emit(tokenTypeOperator)
		case r == '`':
			return lexBacktics
		case r == eof:
			break
		default:
			var min, max int
			runesToShowNearByte := 100
			if l.start-runesToShowNearByte < 0 {
				min = 0
			} else {
				min = l.start - runesToShowNearByte
			}

			if l.pos+runesToShowNearByte >= len(l.input) {
				max = len(l.input) - 1
			} else {
				max = l.pos + runesToShowNearByte
			}

			msg := fmt.Sprintf("unknown rune encountered at byte %d: '%s' (aka '%d') (current parse is '%s')\nsurrounding context:\n\n==========\n%s\n==========", l.pos, string(r), r, l.input[l.start:l.pos], l.input[min:max])
			panic(msg)
		}

		if l.peek() == eof {
			break
		}
	}

	if l.start != len(l.input) {
		l.emit(tokenTypeError)
	}

	l.emit(tokenTypeEOF)
	return nil
}

func (l *StatefulRubyLexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *StatefulRubyLexer) backup() {
	l.pos -= l.width
}

// ignore skips over the pending input before this point.
func (l *StatefulRubyLexer) ignore() {
	l.start = l.pos
}

// peek returns but does not consume
// the next rune in the input.
func (l *StatefulRubyLexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

// accepts a single rune from valid
func (l *StatefulRubyLexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *StatefulRubyLexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *StatefulRubyLexer) emit(t tokenType) {
	newToken := token{
		typ:   t,
		value: l.input[l.start:l.pos],
	}
	l.tokens <- newToken
	l.lastToken = newToken

	l.start = l.pos
}

func (lexer *StatefulRubyLexer) Lex(lval *RubySymType) int {
	debug("Called Lex()")
	defer func() { debug("") }()

	for token := range lexer.tokens {
		switch token.typ {
		case tokenTypeInteger:
			debug("integer: %s", token.value)
			intVal, err := strconv.Atoi(token.value)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantInt{Value: intVal}
			return NODE
		case tokenTypeFloat:
			debug("float: %s", token.value)
			floatval, err := strconv.ParseFloat(token.value, 64)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantFloat{Value: floatval}
			return NODE
		case tokenTypeString:
			debug("string: '%s'", token.value)
			lval.genericValue = ast.SimpleString{Value: token.value}
			return NODE
		case tokenTypeDoubleQuoteString:
			debug("string: '%s'", token.value)
			lval.genericValue = ast.InterpolatedString{Value: token.value}
			return NODE
		case tokenTypeCharacter:
			debug("char: '%s'", token.value)
			lval.genericValue = ast.CharacterLiteral{Value: token.value}
			return NODE
		case tokenTypeSymbol:
			debug("symbol: %s", token.value)
			lval.genericValue = ast.Symbol{Name: token.value}
			return NODE
		case tokenTypeReference:
			debug("REF: %s", token.value)
			lval.genericValue = ast.BareReference{Name: token.value}
			return REF
		case tokenTypeCapitalizedReference:
			debug("CAPITAL REF: %s", token.value)
			lval.genericValue = ast.BareReference{Name: token.value}
			return CAPITAL_REF
		case tokenTypeGlobal:
			debug("REF: '%s'", token.value)
			lval.genericValue = ast.GlobalVariable{Name: token.value}
			return REF
		case tokenTypeLParen:
			debug("(")
			return LPAREN
		case tokenTypeRParen:
			debug(")")
			return RPAREN
		case tokenTypeComma:
			debug("COMMA")
			return COMMA
		case tokenTypeNewline:
			debug("NEWLINE")
			return NEWLINE
		case tokenTypeEOF:
			debug("EOF")
			return EOF
		case tokenTypeDEF:
			debug("DEF")
			return DEF
		case tokenTypeDO:
			debug("DO")
			return DO
		case tokenTypeEND:
			debug("END")
			return END
		case tokenTypeIF:
			debug("IF")
			return IF
		case tokenTypeELSE:
			debug("ELSE")
			return ELSE
		case tokenTypeELSIF:
			debug("ELSIF")
			return ELSIF
		case tokenTypeUNLESS:
			debug("UNLESS")
			return UNLESS
		case tokenTypeCLASS:
			debug("CLASS")
			return CLASS
		case tokenTypeMODULE:
			debug("MODULE")
			return MODULE
		case tokenTypeTRUE:
			debug("TRUE")
			return TRUE
		case tokenTypeFALSE:
			debug("FALSE")
			return FALSE
		case tokenTypeLessThan:
			debug("<")
			return LESSTHAN
		case tokenTypeGreaterThan:
			debug(">")
			return GREATERTHAN
		case tokenTypeColon:
			debug(":")
			return COLON
		case tokenTypeDoubleColon:
			debug("::")
			return DOUBLECOLON
		case tokenTypeSemicolon:
			debug(";")
			return SEMICOLON
		case tokenTypeEqual:
			debug("=")
			return EQUALTO
		case tokenTypeBang:
			debug("!")
			return BANG
		case tokenTypeTilde:
			debug("~")
			return COMPLEMENT
		case tokenTypePlus:
			debug("+")
			return POSITIVE
		case tokenTypeMinus:
			debug("-")
			return NEGATIVE
		case tokenTypeStar:
			debug("*")
			return STAR
		case tokenTypeLBracket:
			debug("[")
			return LBRACKET
		case tokenTypeRBracket:
			debug("]")
			return RBRACKET
		case tokenTypeLBrace:
			debug("{")
			return LBRACE
		case tokenTypeRBrace:
			debug("}")
			return RBRACE
		case tokenTypeDollarSign:
			debug("$")
			return DOLLARSIGN
		case tokenTypeAtSign:
			debug("@")
			return ATSIGN
		case tokenType__FILE__:
			debug("__FILE__")
			lval.genericValue = ast.FileNameConstReference{}
			return REF
		case tokenTypeDot:
			debug(".")
			return DOT
		case tokenTypePipe:
			debug("|")
			return PIPE
		case tokenTypeForwardSlash:
			debug("/")
			return SLASH
		case tokenTypeAmpersand:
			debug("&")
			return AMPERSAND
		case tokenTypeSubshell:
			debug("subshell : '%s'", token.value)
			lval.genericValue = ast.Subshell{Command: token.value}
			return NODE
		case tokenTypeOperator:
			debug("Operator: %s", token.value)
			lval.operator = token.value
			return OPERATOR
		case tokenTypeBEGIN:
			debug("BEGIN")
			return BEGIN
		case tokenTypeRESCUE:
			debug("RESCUE")
			return RESCUE
		case tokenTypeENSURE:
			debug("ENSURE")
			return ENSURE
		case tokenTypeBREAK:
			debug("BREAK")
			return BREAK
		case tokenTypeNEXT:
			debug("NEXT")
			return NEXT
		case tokenTypeREDO:
			debug("REDO")
			return REDO
		case tokenTypeRETRY:
			debug("RETRY")
			return RETRY
		case tokenTypeRETURN:
			debug("RETURN")
			return RETURN
		case tokenTypeYIELD:
			debug("YIELD")
			return YIELD
		case tokenTypeQuestionMark:
			debug("?")
			return QUESTIONMARK
		case tokenTypeMethodName:
			debug("Method: '%s'", token.value)
			lval.genericValue = ast.BareReference{Name: token.value}
			return SPECIAL_CHAR_REF
		case tokenTypeWHILE:
			debug("WHILE")
			return WHILE
		case tokenTypeAND:
			debug("AND")
			return AND
		case tokenTypeOR:
			debug("OR")
			return OR
		case tokenTypeLAMBDA:
			debug("LAMBDA")
			return LAMBDA
		case tokenTypeCASE:
			debug("CASE")
			return CASE
		case tokenTypeWHEN:
			debug("WHEN")
			return WHEN
		case tokenTypeOrEquals:
			debug("||=")
			return OR_EQUALS
		case tokenTypeRange:
			debug(".. (range)")
			return RANGE
		case tokenTypeRegex:
			debug("regex: '%s'", token.value)
			lval.genericValue = ast.Regex{Value: token.value}
			return NODE
		case tokenTypeError:
			panic(fmt.Sprintf("error, unknown token: '%s'", token.value))
		default:
			panic(fmt.Sprintf("unknown token: '%s'", token))
		}
	}

	return 0
}

func (lexer *StatefulRubyLexer) Error(error string) {
	lexer.LastError = errors.New(fmt.Sprintf("syntax error: %s\n", error))
}

func debug(formatString string, args ...interface{}) {
	var msg string
	if len(args) > 0 {
		msg = fmt.Sprintf(formatString, args...)
	} else {
		msg = formatString
	}

	DebugStatements = append(DebugStatements, msg)
}
