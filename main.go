package main

import (
	"fmt"
	"github.com/GreatGodApollo/monk/repl"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) == 2 {
		repl.RunProgram(os.Stdin, os.Stdout, os.Args[1])
	} else {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! Welcome to the SMPLR programming language\n",
			usr.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
