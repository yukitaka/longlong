package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"github.com/yukitaka/longlong/internal/util"
	"golang.org/x/term"
	"syscall"
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
		Use:   "login [ORGANIZATION] [ACCOUNT]",
		Short: "Authorize access to Longlong",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Login(args))
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

func (o *Options) Login(args []string) error {
	authRep := repository.NewAuthenticationsRepository()
	defer authRep.Close()
	organizationRep := repository.NewOrganizationsRepository()
	defer organizationRep.Close()
	belongingRep := repository.NewOrganizationBelongingsRepository()
	defer belongingRep.Close()
	itr := usecase.NewAuthentication(authRep, organizationRep, belongingRep)

	fmt.Print("Password: ")
	pw, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return err
	}

	id, err := itr.Auth(args[0], args[1], string(pw))
	if err != nil {
		return fmt.Errorf("\nAuthentication failure (%s)", err)
	}
	fmt.Println()
	fmt.Printf("Login %s %s %d.\n", args[0], args[1], id)

	return nil
}
