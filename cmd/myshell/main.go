package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/builtins"
)

func main() {

	// infinite repl
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		prompt, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}
		args := strings.Fields(prompt)
		cmd := args[0]

		if f, ok := builtins.Match(cmd); ok {
			f(args)
		} else {
			fmt.Printf("%v: command not found\n", cmd)
		}
	}

}
