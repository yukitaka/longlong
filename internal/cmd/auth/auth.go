package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/util"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewAuthOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func NewCmdAuth(parent string, streams cli.IOStream) *cobra.Command {
	o := NewAuthOptions(parent, streams)

	cmd := &cobra.Command{
		Use:     "auth",
		Aliases: []string{"a"},
		Short:   "Manage authentication",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(args))
		},
	}

	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Authorize access to Longlong",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Login(cmd, args))
		},
	}
	loginCmd.PersistentFlags().StringP("output", "o", "table", "output format")
	cmd.AddCommand(loginCmd)

	return cmd
}

func (o *Options) Run(args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Login(cmd *cobra.Command, args []string) error {
	return nil
}
