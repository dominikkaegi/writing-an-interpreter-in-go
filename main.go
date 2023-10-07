package main

import (
	"fmt"
	"os"
	"os/user"

	"kdo.com/go/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free t type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
