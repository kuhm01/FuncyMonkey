package ast

import "bytes"

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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
