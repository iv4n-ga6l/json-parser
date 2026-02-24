package lexer

import (
	"errors"
	"strings"
)

// TokenType represents the type of a token.
type TokenType string

const (
	TokenLeftBrace  TokenType = "LEFT_BRACE"
	TokenRightBrace TokenType = "RIGHT_BRACE"
	TokenEOF        TokenType = "EOF"
)

// Token represents a single token in the input.
type Token struct {
	Type  TokenType
	Value string
}

// Tokenize takes a JSON string and returns a slice of tokens or an error.
func Tokenize(input string) ([]Token, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return nil, errors.New("empty input")
	}

	tokens := []Token{}

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '{':
			tokens = append(tokens, Token{Type: TokenLeftBrace, Value: "{"})
		case '}':
			tokens = append(tokens, Token{Type: TokenRightBrace, Value: "}"})
		default:
			return nil, errors.New("unexpected character: " + string(input[i]))
		}
	}

	tokens = append(tokens, Token{Type: TokenEOF, Value: ""})
	return tokens, nil
}