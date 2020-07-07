package main

import (
	"fmt"
	"os"
	"setec/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
