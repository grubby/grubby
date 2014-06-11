package token

import (
	"strings"
	"unicode"
)

func Tokenize(input string) []string {
	tokens := []string{}
	currentToken := ""

	for _, char := range input {
		if unicode.IsSpace(char) && strings.TrimSpace(currentToken) != "" {
			tokens = append(tokens, strings.TrimSpace(currentToken))
			currentToken = ""
		} else {
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}
