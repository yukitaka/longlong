package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/cmd/create"
	"github.com/yukitaka/longlong/internal/cmd/get"
)

type LlctlOptions struct {
	CmdHandler Handler
	Arguments  []string
	cli.IOStream
}

func NewLlctlCommand() *cobra.Command {
	return NewLlctlCommandWithArgs(LlctlOptions{
		CmdHandler: NewDefaultHandler([]string{"llctl"}),
		Arguments:  os.Args,
		IOStream: cli.IOStream{
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
	cmds.AddCommand(get.NewCmdGet("llctl", o.IOStream))
	cmds.AddCommand(create.NewCmdCreate("llctl", o.IOStream))

	if len(o.Arguments) > 1 {
		cmdArgs := o.Arguments[1:]
		if _, _, err := cmds.Find(cmdArgs); err != nil {
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
				if err := HandleCommand(o.CmdHandler, cmdArgs); err != nil {
					fmt.Fprintf(o.IOStream.ErrOut, "Error: %v\n", cmdName)
					os.Exit(1)
				}
			}
		}
	}

	return cmds
}

type Handler interface {
	Lookup(filename string) (string, bool)
	Execute(executablePath string, cmdArgs, environment []string) error
}

type DefaultHandler struct {
	ValidPrefixes []string
}

func NewDefaultHandler(validPrefixes []string) *DefaultHandler {
	return &DefaultHandler{
		ValidPrefixes: validPrefixes,
	}
}

func (h *DefaultHandler) Lookup(filename string) (string, bool) {
	for _, prefix := range h.ValidPrefixes {
		path, err := exec.LookPath(fmt.Sprintf("%s-%s", prefix, filename))
		fmt.Println(err)
		if len(path) == 0 {
			continue
		}
		return path, true
	}
	return "", false
}

func (h *DefaultHandler) Execute(executablePath string, cmdArgs, environment []string) error {
	//TODO implement me
	panic("implement me")
}

func HandleCommand(cmdHandler Handler, cmdArgs []string) error {
	var remainingArgs []string
	for _, arg := range cmdArgs {
		if strings.HasPrefix(arg, "-") {
			break
		}
		remainingArgs = append(remainingArgs, strings.Replace(arg, "-", "_", -1))
	}

	foundBinaryPath := ""
	for len(remainingArgs) > 0 {
		path, found := cmdHandler.Lookup(strings.Join(remainingArgs, "-"))
		if !found {
			remainingArgs = remainingArgs[:len(remainingArgs)-1]
			continue
		}

		foundBinaryPath = path
		break
	}

	if len(foundBinaryPath) == 0 {
		return nil
	}

	if err := cmdHandler.Execute(foundBinaryPath, cmdArgs[len(remainingArgs):], os.Environ()); err != nil {
		return err
	}
	return nil
}
