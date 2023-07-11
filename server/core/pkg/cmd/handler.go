package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"os"
	"os/exec"
	"strings"
)

type Handler interface {
	Lookup(filename string) (string, bool)
	Execute(executablePath string, cmdArgs, environment []string) error
	Run(cmdGroup *cobra.Command, cmdArgs []string, io cli.IOStream) error
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

func (h *DefaultHandler) Run(cmdGroup *cobra.Command, cmdArgs []string, io cli.IOStream) error {
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
			if err := HandleCommand(h, cmdArgs); err != nil {
				_, err := fmt.Fprintf(io.ErrOut, "Error: %v %v\n", cmdName, err)
				if err != nil {
					return err
				}
				os.Exit(1)
			}
		}
	}
	return nil
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
