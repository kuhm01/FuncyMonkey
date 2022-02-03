package route

import (
	"fmt"
	"funcymonkey/src/route/repl"
	"os"
	"os/user"
)

func Main(args []string) {
	switch len(args) {
	case 1:
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is FuncyMonkey!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)

	default:
		fmt.Printf("Critical Error!\n")
	}
}
