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

/*read next Character*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

//Lexer implement

func (l *Lexer) NewToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '{':
		return newToken(token.LBRACE, l.ch)
	}

	return tok
}

/*return newToken*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	var tok token.Token
	return tok
}
