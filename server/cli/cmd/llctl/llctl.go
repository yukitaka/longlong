package main

import (
	"github.com/yukitaka/longlong/server/cli/internal/cli"
	"github.com/yukitaka/longlong/server/cli/internal/cmd/ctl"
)

func main() {
	command := ctl.NewLlctlCommand()
	cli.Run(command)
}
