package main

import (
	"fmt"
	"funcymonkey/src/ast"
	"funcymonkey/src/evaluator"
	"funcymonkey/src/lexer"
	"funcymonkey/src/parser"
	"funcymonkey/src/token"
)

func main() {
	fmt.Printf("Hello! This is FuncyMonkey!\n")
	fmt.Printf("%s\n", ast.Hello())
	fmt.Printf("%s\n", evaluator.Hello())
	fmt.Printf("%s\n", lexer.Hello())
	fmt.Printf("%s\n", parser.Hello())
	fmt.Printf("%s\n", token.Hello())
}