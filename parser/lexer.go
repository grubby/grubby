package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/grubby/grubby/ast"
)

var (
	integerRegexp = regexp.MustCompile("^[0-9]+$")
	floatRegexp   = regexp.MustCompile("^[0-9]+\\.[0-9]+$")
	stringRegexp  = regexp.MustCompile("^'[^']*'$")
)

type rubyLex struct {
	input        string
	position     int
	currentToken string

	lexIndex int
	tokens   []string
}

func NewLexer(input string) *rubyLex {
	lexer := &rubyLex{input: input}
	lexer.tokenize()

	return lexer
}

func (lexer *rubyLex) tokenize() {
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
}

func (lexer *rubyLex) appendNonEmptyTokens(tokens ...string) {
	defer func() {
		lexer.currentToken = ""
	}()

	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			lexer.tokens = append(lexer.tokens, token)
		}
	}
}

func (lexer *rubyLex) tokenizeString() {
	token := string(lexer.input[lexer.position])
	for i := lexer.position + 1; i < len(lexer.input); i += 1 {
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

func (l *rubyLex) Lex(lval *RubySymType) int {
	for l.lexIndex < len(l.tokens) {
		token := l.tokens[l.lexIndex]
		defer func() {
			l.lexIndex++
		}()

		switch {
		case integerRegexp.MatchString(token):
			intVal, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantInt{Value: intVal}
			return NODE
		case floatRegexp.MatchString(token):
			floatval, err := strconv.ParseFloat(token, 64)
			if err != nil {
				panic(err)
			}

			lval.genericValue = ast.ConstantFloat{Value: floatval}
			return NODE
		case stringRegexp.MatchString(token):
			lval.genericValue = ast.SimpleString{Value: token}
			return NODE
		case len(token) == 1: // ;  " " or another separator
			return int(rune(token[0]))
		default:
			panic("unknown token: " + token)
		}
	}

	return 0
}

func (l *rubyLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || r == ';'

}
