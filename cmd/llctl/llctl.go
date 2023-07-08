package main

import (
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/cmd/ctl"
)

func main() {
	command := ctl.NewLlctlCommand()
	cli.Run(command)
}
