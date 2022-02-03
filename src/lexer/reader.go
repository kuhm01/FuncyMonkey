package lexer

import (
	"funcymonkey/src/token"
)

/*peek the Character*/
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
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

/*Read Identifier*/
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*Read digit*/
func (l *Lexer) readNumber() (token.TokenType, string) {
	position := l.position

	flag := true
	ilflag := false

	for isDigit(l.ch) {
		if isDot(l.ch) {
			flag = false
			ilflag = true
		} else {
			ilflag = false
		}
		l.readChar()
	}

	if ilflag {
		return token.ILLEGAL, l.input[position:l.position]
	}

	if flag {
		return token.INT, l.input[position:l.position]
	} else {
		return token.FLOAT, l.input[position:l.position]
	}
}
