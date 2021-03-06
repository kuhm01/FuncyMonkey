package lexer

import (
	"funcymonkey/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	func Main() {}
	func determine(a, b) {}
	for a := range(5) {}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BFUNCTION, "func"},
		{token.MAINF, "Main"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.BFUNCTION, "func"},
		{token.IDENT, "determine"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.FOR, "for"},
		{token.IDENT, "a"},
		{token.GOASSIGN, ":="},
		{token.IDENT, "range"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype worng. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
