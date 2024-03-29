package main

import (
	"fmt"
	"github.com/ralaruri/interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Username)
	fmt.Printf("Feel ree to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
