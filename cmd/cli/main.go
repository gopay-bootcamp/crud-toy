package main

import (
	"crud-toy/internal/app/cli/command"
	"crud-toy/internal/app/cli/daemon"
	"crud-toy/internal/app/cli/utility/io"
)

func main() {
	printer := io.GetPrinter()
	daemon.StartClient(printer)

	command.Execute()
}
