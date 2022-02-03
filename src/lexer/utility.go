package lexer

/*Skip WhiteSpace*/
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*the Function judge letter*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/*the Function judge digit*/
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

/*the Function judge dot for RealNumber*/
func isDot(ch byte) bool {
	return ch == '.'
}
