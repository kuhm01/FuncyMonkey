package ast

import (
	"funcymonkey/src/token"
)

func Hello() string {
	return "Hello! This is ast.go! in ast package"
}

//Hello! Package

type Node interface {
	TokenLiteral() string
	String() string
}

/*Statement 명령문*/
type Statement interface {
	Node
	statementNode()
}

/*Expression 표현식*/
type Expression interface {
	Node
	expressionNode()
}

/*Program Node ast의 첫 노드. 명령문을 내포*/
type Program struct {
	/*Set of Statement*/
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}
