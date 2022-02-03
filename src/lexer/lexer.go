package lexer

import (
	"funcymonkey/src/token"
)

func Hello() string {
	return "Hello! This is lexer.go! in lexer package"
}

//Hello! Package

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

/*Lexer initial and input string setting*/
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//Lexer implement

/*Read the next token and return*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.ThatisIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type, tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

/*return newToken*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
