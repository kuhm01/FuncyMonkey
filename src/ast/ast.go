package ast

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
