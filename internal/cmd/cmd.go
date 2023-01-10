package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

type LlctlOptions struct {
	Arguments []string
	IOStream
}

type IOStream struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}

func NewLlctlCommand() *cobra.Command {
	return NewLlctlCommandWithArgs(LlctlOptions{
		Arguments: os.Args,
		IOStream: IOStream{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
	})
}

func NewLlctlCommandWithArgs(o LlctlOptions) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "llctl",
		Short: "llctl controls the LongLong manager",
		Long: `
llctl controls the LongLong manager.

Find more information at:
https://github.com/yukitaka/longlong/`,
	}

	if len(o.Arguments) > 1 {
		cmdArgs := o.Arguments[1:]
		var cmdName string
		for _, arg := range cmdArgs {
			if !strings.HasPrefix(arg, "-") {
				cmdName = arg
				break
			}
		}
		switch cmdName {
		case "help", cobra.ShellCompRequestCmd, cobra.ShellCompNoDescRequestCmd:
		default:
			fmt.Fprintf(o.IOStream.ErrOut, "Error: %v\n", cmdName)
			os.Exit(1)
		}
	}

	return cmds
}
