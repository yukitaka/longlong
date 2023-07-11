package cmd

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/cli/internal/cmd/auth"
	"github.com/yukitaka/longlong/server/cli/internal/cmd/config"
	"github.com/yukitaka/longlong/server/cli/internal/cmd/create"
	"github.com/yukitaka/longlong/server/cli/internal/cmd/del"
	"github.com/yukitaka/longlong/server/cli/internal/cmd/get"
	"github.com/yukitaka/longlong/server/cli/internal/cmd/put"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/interface/repository"
	"os"
	"os/exec"
	"strings"
)

type LlctlOptions struct {
	CmdHandler Handler
	Arguments  []string
	Operator   entity.OrganizationMember
	*config.Config
	*sqlx.DB
	cli.IOStream
}

func NewLlctlCommand() *cobra.Command {
	conf, err := config.LoadFromFile("config", "yaml", "$HOME/.config/llctl")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	con, _ := datastore.NewConnectionOpen(conf.Datastore.Driver, conf.Datastore.Source)
	itr := usecase.NewOrganizationMemberFinder(repository.NewOrganizationMembersRepository(con))
	member, err := itr.FindById(conf.Authorize.OrganizationId, conf.Authorize.UserId)
	if err != nil {
		panic("pkg/cmd/cmd.go:51 " + err.Error())
	}
	if member == nil {
		panic("Not found the operator.")
	}
	operator := *member

	return NewLlctlCommandWithArgs(LlctlOptions{
		CmdHandler: NewDefaultHandler([]string{"llctl"}),
		Arguments:  os.Args,
		Operator:   operator,
		Config:     &conf,
		DB:         con,
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
	cmdGroup.AddCommand(auth.NewCmdAuth("llctl", o.Config, o.DB, o.IOStream))
	cmdGroup.AddCommand(get.NewCmdGet("llctl", o.IOStream, &o.Operator, o.DB))
	cmdGroup.AddCommand(put.NewCmdPut("llctl", o.IOStream, &o.Operator, o.DB))
	cmdGroup.AddCommand(del.NewCmdDelete("llctl", o.IOStream, &o.Operator, o.DB))
	cmdGroup.AddCommand(create.NewCmdCreate("llctl", o.IOStream, &o.Operator, o.DB))

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
