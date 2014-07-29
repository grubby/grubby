package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/grubby/grubby/ast"
)

var (
	integerRegexp          = regexp.MustCompile("^[0-9]+$")
	floatRegexp            = regexp.MustCompile("^[0-9]+\\.[0-9]+$")
	stringRegexp           = regexp.MustCompile("^'[^']*'$")
	referenceRegexp        = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")
	capitalReferenceRegexp = regexp.MustCompile("^[A-Z][a-zA-Z0-9_]*$")
	DebugStatements        = []string{}
)

type RubyLex struct {
	input        string
	position     int
	currentToken string

	lexIndex int
	tokens   []string

	LastError error
}

func NewLexer(input string) RubyLexer {
	lexer := &RubyLex{input: input}
	lexer.tokenize()

	return lexer
}

func (lexer *RubyLex) tokenize() {
	for lexer.position < len(lexer.input) {
		char := rune(lexer.input[lexer.position])

		switch {
		case isSeparator(char):
			lexer.appendNonEmptyTokens(lexer.currentToken)
			lexer.tokens = append(lexer.tokens, string(char))
		case char == '\'':
			lexer.tokenizeString()
		default:
			lexer.currentToken += string(char)
		}

		lexer.position++
	}

	lexer.appendNonEmptyTokens(lexer.currentToken)
	debug("Lexed %d tokens", len(lexer.tokens))
}

func (lexer *RubyLex) appendNonEmptyTokens(tokens ...string) {
	defer func() {
		lexer.currentToken = ""
	}()

	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			lexer.tokens = append(lexer.tokens, token)
		}
	}
}

func (lexer *RubyLex) tokenizeString() {
	token := string(lexer.input[lexer.position])
	for i := lexer.position + 1; i < len(lexer.input); i++ {
		char := lexer.input[i]
		token += string(char)

		if char == '\'' {
			lexer.position = i
			break
		}
	}

	// FIXME: needs to blow up if the string was never closed
	lexer.currentToken = token
}

func (l *RubyLex) Lex(lval *RubySymType) int {
	debug("Called Lex()")
	defer func() {
		debug("")
	}()

	for l.lexIndex < len(l.tokens) {
		token := l.tokens[l.lexIndex]
		defer func() {
			l.lexIndex++
		}()

		switch {
		case token == "def":
			debug("def")
			return DEF
		case token == "end":
			debug("end")
			return END
		case token == "class":
			debug("class")
			return CLASS
		case token == "module":
			debug("module")
			return MODULE
		case token == "true":
			debug("TRUE")
			return TRUE
		case token == "false":
			debug("FALSE")
			return FALSE
		case integerRegexp.MatchString(token):
			debug("integer: %s", token)
			intVal, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantInt{Value: intVal}
			return NODE
		case floatRegexp.MatchString(token):
			debug("float: %s", token)
			floatval, err := strconv.ParseFloat(token, 64)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantFloat{Value: floatval}
			return NODE
		case stringRegexp.MatchString(token):
			debug("string: %s", token)
			lval.genericValue = ast.SimpleString{Value: token}
			return NODE
		case capitalReferenceRegexp.MatchString(token):
			debug("capitalized ref: %s", token)
			lval.genericValue = ast.BareReference{Name: token}
			return CAPITAL_REF
		case referenceRegexp.MatchString(token):
			debug("ref: %s", token)
			lval.genericValue = ast.BareReference{Name: token}
			return REF
		case token == "<":
			debug("LESS THAN")
			return LESSTHAN
		case token == ">":
			debug("GREATER THAN")
			return GREATERTHAN
		case token == "=":
			debug("EQUAL")
			return EQUALTO
		case token == "(":
			debug("LPAREN")
			return LPAREN
		case token == ")":
			debug("RPAREN")
			return RPAREN
		case token == ",":
			debug("COMMA")
			return COMMA
		case token == ":":
			debug("COLON")
			return COLON
		case token == "!":
			debug("BANG")
			return BANG
		case token == "~":
			debug("COMPLEMENT")
			return COMPLEMENT
		case token == "+":
			debug("POSITIVE")
			return POSITIVE
		case token == "-":
			debug("NEGATIVE")
			return NEGATIVE
		case len(token) == 1: // ;  " " or another separator
			debug("separator: '%s'", strings.Replace(token, string('\n'), "\\n", -1))
			return int(rune(token[0]))
		default:
			panic(fmt.Sprintf("unknown token: '%s'", token))
		}
	}

	return 0
}

func (l *RubyLex) Error(error string) {
	l.LastError = errors.New(fmt.Sprintf("syntax error: %s\n", error))
}

// FIXME: is there a concept of a character set that would make this faster?
//        barring that, maybe just a map lookup?
func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || r == ';' || r == ')' || r == '(' || r == ',' || r == ':' || r == '!' || r == '~' || r == '+' || r == '-'
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
