package parser

import (
	"funcymonkey/src/ast"
	"funcymonkey/src/lexer"
	"funcymonkey/src/token"
)

func Hello() string {
	return "Hello! This is parser.go! in parser package"
}

const (
	_int = iota
	LOWEST
	ASSIGNMENT
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	INDEX
)

type Parser struct {
	l *lexer.Lexer

	/*현재 읽고있는 토큰*/
	curToken token.Token

	/*다음 차례의 토큰*/
	peekToken token.Token
	errors    []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

/*연산자 우선순위 정의*/
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
	token.ASSIGN:   ASSIGNMENT,
	token.GOASSIGN: ASSIGNMENT,
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.nextToken()
	p.nextToken()

	return p
}
