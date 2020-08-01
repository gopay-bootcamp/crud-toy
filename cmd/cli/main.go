package main

import (
	"crud-toy/internal/app/cli/command"
	"crud-toy/internal/app/cli/daemon"
	"crud-toy/internal/app/cli/utility/io"
)

func main() {
	io.InitPrinter()
	daemon.StartClient()

	command.Execute()
}
