package token

import (
	"strings"
	"unicode"
)

func Tokenize(input string) []string {
	var (
		tokens       = []string{}
		currentToken = ""
	)

	index := 0
	for index < len(input) {
		char := rune(input[index])

		switch {
		case char == '(' || char == ')':
			tokens = append(
				tokens,
				strings.TrimSpace(currentToken),
				string(char),
			)
			currentToken = ""
		case char == '\'':
			currentToken = tokenizeString(input, &index)
		case isSeparator(char) && strings.TrimSpace(currentToken) != "":
			tokens = append(tokens, strings.TrimSpace(currentToken))
			currentToken = ""
		default:
			currentToken += string(char)
		}

		index += 1
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func tokenizeString(input string, index *int) string {
	token := string(input[*index])
	for i := *index + 1; i < len(input); i += 1 {
		char := input[i]
		token += string(char)

		if char == '\'' {
			(*index) = i
			break
		}
	}

	return token
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || r == ';'
}
