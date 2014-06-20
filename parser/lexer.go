package parser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type rubyLexer struct {
	index        int
	currentToken string

	lexIndex int
	tokens   []string
}

type Lexer interface {
	Tokenize(input string) []string
	Lex(lval *RubySymType) int
}

func NewLexer() Lexer {
	return &rubyLexer{}
}

func (lexer *rubyLexer) Tokenize(input string) []string {
	lexer.tokens = []string{}
	lexer.currentToken = ""
	lexer.index = 0

	for lexer.index < len(input) {
		char := rune(input[lexer.index])

		switch {
		case char == '(' || char == ')':
			lexer.appendNonEmptyTokens(lexer.currentToken, string(char))
		case char == '\'':
			lexer.tokenizeString(input)
		case char == '[' || char == ']':
			lexer.appendNonEmptyTokens(lexer.currentToken, string(char))
		case isSeparator(char):
			lexer.appendNonEmptyTokens(lexer.currentToken)
		default:
			lexer.currentToken += string(char)
		}

		lexer.index++
	}

	lexer.appendNonEmptyTokens(lexer.currentToken)
	return lexer.tokens
}

func (lexer *rubyLexer) Lex(lval *RubySymType) int {
	if lexer.lexIndex == len(lexer.tokens) {
		return 0
	}

	defer func() {
		lexer.lexIndex++
	}()

	token := lexer.tokens[lexer.lexIndex]
	switch {
	case integerRegexp.MatchString(token):
		lval.Val, _ = strconv.Atoi(token)
		return DIGIT
	}

	return 0
}

func (lexer *rubyLexer) Error(s string) {
	panic(fmt.Sprintf("syntax error: %s\n", s))
}

func (lexer *rubyLexer) appendNonEmptyTokens(tokens ...string) {
	defer func() {
		lexer.currentToken = ""
	}()

	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			lexer.tokens = append(lexer.tokens, token)
		}
	}
}

func (lexer *rubyLexer) tokenizeString(input string) {
	token := string(input[lexer.index])
	for i := lexer.index + 1; i < len(input); i += 1 {
		char := input[i]
		token += string(char)

		if char == '\'' {
			lexer.index = i
			break
		}
	}

	lexer.currentToken = token
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || r == ';'
}
