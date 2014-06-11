package token

import (
	"strings"
	"unicode"
)

func Tokenize(input string) []string {
	var (
		tokens        = []string{}
		currentToken  = ""
		parsingString = false
	)

	for _, char := range input {
		switch {
		case char == '\'':
			if parsingString {
				parsingString = false
				currentToken += string(char)
			} else {
				parsingString = true
				currentToken += string(char)
			}
		case !parsingString && isSeparator(char) && strings.TrimSpace(currentToken) != "":
			tokens = append(tokens, strings.TrimSpace(currentToken))
			currentToken = ""
		default:
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || r == ';'
}
