package utils

import "os"

func init() {

	// Init Custom Flag [flags.go]
	Flags()

	// Init Command [variables.go | Cmd string]
	Cmd = os.Args[1]
}
