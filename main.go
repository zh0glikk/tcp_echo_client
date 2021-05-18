package main

import (
	"os"

	"tcp_echo_client/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
