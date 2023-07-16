package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/admin/internal/cmd/initialize"
	"github.com/yukitaka/longlong/server/admin/internal/cmd/server"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/cmd"
	"github.com/yukitaka/longlong/server/core/pkg/interface/config"
	"os"
)

type AdminOptions struct {
	CmdHandler cmd.Handler
	Arguments  []string
	*config.Config
	cli.IOStream
}

func NewAdminCommand() *cobra.Command {
	conf, err := config.LoadFromFile("config", "yaml", "$HOME/.config/llctl")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return NewAdminCommandWithArgs(AdminOptions{
		CmdHandler: cmd.NewDefaultHandler([]string{"lladmin"}),
		Arguments:  os.Args,
		Config:     &conf,
	})
}

func NewAdminCommandWithArgs(o AdminOptions) *cobra.Command {
	cmdGroup := &cobra.Command{
		Use:   "lladmin",
		Short: "lladmin manages the LongLong server",
		Long: `
lladmin manages the LongLong server.

Find more information at:
https://github.com/yukitaka/longlong/`,
	}
	cmdGroup.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of LongLong",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("LongLong version %s\n", "0.0.1")
		},
	})
	cmdGroup.AddCommand(initialize.NewCmdInit("lladmin", o.Config, o.IOStream))
	cmdGroup.AddCommand(server.NewCmdServer("lladmin", o.IOStream))

	if len(o.Arguments) > 1 {
		cmdArgs := o.Arguments[1:]
		if err := o.CmdHandler.Run(cmdGroup, cmdArgs, o.IOStream); err != nil {
			_, err := fmt.Fprintf(o.IOStream.ErrOut, "Error: %#v\n", err)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return nil
			}
			os.Exit(1)
		}
	}

	return cmdGroup
}
