package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/darragh-downey/twerp/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the TWERP programming language!\n", user.Username)
	fmt.Printf("Feel free to type in some commands!!\n")
	repl.Start(os.Stdin, os.Stdout)
}
