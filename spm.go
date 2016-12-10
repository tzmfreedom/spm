package main

import (
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func main() {
	cli := &CLI{}
	err := cli.Run(os.Args)
	statusCode := ExitCodeOK
	if err != nil {
		statusCode = ExitCodeError
	}
	os.Exit(statusCode)
}
