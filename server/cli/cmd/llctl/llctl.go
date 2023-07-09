package main

import (
	"github.com/yukitaka/longlong/server/cli/internal/cli"
	"github.com/yukitaka/longlong/server/cli/internal/cmd"
)

func main() {
	command := cmd.NewLlctlCommand()
	cli.Run(command)
}
