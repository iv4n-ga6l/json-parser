package lexer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
		err      bool
	}{
		{"{}", []Token{{Type: TokenLeftBrace, Value: "{"}, {Type: TokenRightBrace, Value: "}"}, {Type: TokenEOF, Value: ""}}, false},
		{"{", nil, true},
		{"}", nil, true},
		{"", nil, true},
		{"{abc}", nil, true},
	}

	for _, test := range tests {
		tokens, err := Tokenize(test.input)
		if (err != nil) != test.err {
			t.Errorf("unexpected error status for input '%s': got %v, want error=%v", test.input, err, test.err)
		}

		if !reflect.DeepEqual(tokens, test.expected) {
			t.Errorf("unexpected tokens for input '%s': got %v, want %v", test.input, tokens, test.expected)
		}
	}
}