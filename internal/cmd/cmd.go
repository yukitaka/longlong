package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/cmd/auth"
	"github.com/yukitaka/longlong/internal/cmd/create"
	"github.com/yukitaka/longlong/internal/cmd/get"
	initialize "github.com/yukitaka/longlong/internal/cmd/init"
	"github.com/yukitaka/longlong/internal/cmd/put"
)

type LlctlOptions struct {
	CmdHandler Handler
	Arguments  []string
	UserId     int64
	cli.IOStream
}

func NewLlctlCommand() *cobra.Command {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/llctl")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return NewLlctlCommandWithArgs(LlctlOptions{
		CmdHandler: NewDefaultHandler([]string{"llctl"}),
		Arguments:  os.Args,
		UserId:     1,
		IOStream: cli.IOStream{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
	})
}

func NewLlctlCommandWithArgs(o LlctlOptions) *cobra.Command {
	cmdGroup := &cobra.Command{
		Use:   "llctl",
		Short: "llctl controls the LongLong manager",
		Long: `
llctl controls the LongLong manager.

Find more information at:
https://github.com/yukitaka/longlong/`,
	}
	cmdGroup.AddCommand(initialize.NewCmdInit("llctl", o.IOStream))
	cmdGroup.AddCommand(auth.NewCmdAuth("llctl", o.IOStream))
	cmdGroup.AddCommand(get.NewCmdGet("llctl", o.IOStream, o.UserId))
	cmdGroup.AddCommand(put.NewCmdPut("llctl", o.IOStream, o.UserId))
	cmdGroup.AddCommand(create.NewCmdCreate("llctl", o.IOStream, o.UserId))

	if len(o.Arguments) > 1 {
		cmdArgs := o.Arguments[1:]
		if _, _, err := cmdGroup.Find(cmdArgs); err != nil {
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
					_, err := fmt.Fprintf(o.IOStream.ErrOut, "Error: %v %v\n", cmdName, err)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
						return nil
					}
					os.Exit(1)
				}
			}
		}
	}

	return cmdGroup
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
