package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	// test input to the lexer
	input := `=+(){},;`

	// TestTokenType struct
	type TestTokenType struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	// slice of TestTokenType struct - dynamically sized
	tests := []TestTokenType{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
