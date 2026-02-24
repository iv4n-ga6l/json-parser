package parser

import (
	"errors"
	"iv4n-ga6l/json-parser/internal/lexer"
)

// Parse validates a sequence of tokens as a valid JSON object.
func Parse(tokens []lexer.Token) error {
	if len(tokens) < 3 {
		return errors.New("incomplete JSON object")
	}

	if tokens[0].Type != lexer.TokenLeftBrace {
		return errors.New("JSON must start with '{'")
	}

	if tokens[1].Type != lexer.TokenRightBrace {
		return errors.New("expected '}' after '{'")
	}

	if tokens[2].Type != lexer.TokenEOF {
		return errors.New("unexpected tokens after '}'")
	}

	return nil
}