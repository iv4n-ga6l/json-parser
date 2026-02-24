package main

import (
	"fmt"
	"os"

	"iv4n-ga6l/json-parser/internal/lexer"
	"iv4n-ga6l/json-parser/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jsonparser <json-string>")
		os.Exit(1)
	}

	input := os.Args[1]

	// Tokenize the input
	tokens, err := lexer.Tokenize(input)
	if err != nil {
		fmt.Println("Invalid JSON:", err)
		os.Exit(1)
	}

	// Parse the tokens
	if err := parser.Parse(tokens); err != nil {
		fmt.Println("Invalid JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Valid JSON")
	os.Exit(0)
}