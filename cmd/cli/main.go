package main

import (
	"crud-toy/internal/cli/command"
	"crud-toy/internal/cli/daemon"
	"crud-toy/internal/cli/printer"
)

func main() {
	io.InitPrinter()
	daemon.StartClient()

	command.Execute()
}
