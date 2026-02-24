package parser

import (
	"testing"

	"iv4n-ga6l/json-parser/internal/lexer"
)

func TestParse(t *testing.T) {
	tests := []struct {
		tokens []lexer.Token
		err   bool
	}{
		{[]lexer.Token{{Type: lexer.TokenLeftBrace, Value: "{"}, {Type: lexer.TokenRightBrace, Value: "}"}, {Type: lexer.TokenEOF, Value: ""}}, false},
		{[]lexer.Token{{Type: lexer.TokenLeftBrace, Value: "{"}, {Type: lexer.TokenEOF, Value: ""}}, true},
		{[]lexer.Token{{Type: lexer.TokenRightBrace, Value: "}"}, {Type: lexer.TokenEOF, Value: ""}}, true},
		{[]lexer.Token{{Type: lexer.TokenLeftBrace, Value: "{"}, {Type: lexer.TokenRightBrace, Value: "}"}}, true},
	}

	for _, test := range tests {
		err := Parse(test.tokens)
		if (err != nil) != test.err {
			t.Errorf("unexpected error status for tokens '%v': got %v, want error=%v", test.tokens, err, test.err)
		}
	}
}