package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd := NewLlctlCommand()
	cmd.Execute()
}

func NewLlctlCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "llctl",
		Short: "llctl controls the LongLong manager",
		Long: `
llctl controls the LongLong manager.

Find more information at:
https://github.com/yukitaka/longlong/`,
	}
}
