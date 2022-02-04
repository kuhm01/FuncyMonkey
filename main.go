package main

import (
	"fmt"
	"funcymonkey/src/evaluator"
	"funcymonkey/src/route"
	"os"
)

func main() {
	fmt.Printf("Hello! This is FuncyMonkey!\n")
	fmt.Printf("%s\n", evaluator.Hello())

	args := os.Args
	route.Main(args)
}
