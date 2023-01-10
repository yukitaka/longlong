package main

import (
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/cmd"
)

func main() {
	command := cmd.NewLlctlCommand()
	cli.Run(command)
}
