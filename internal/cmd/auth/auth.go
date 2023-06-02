package auth

import (
	"database/sql"
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
	*sql.DB
	cli.IOStream
}

func NewAuthOptions(parent string, streams cli.IOStream, db *sql.DB) *Options {
	return &Options{
		CmdParent: parent,
		DB:        db,
		IOStream:  streams,
	}
}

func NewCmdAuth(parent string, streams cli.IOStream, db *sql.DB) *cobra.Command {
	o := NewAuthOptions(parent, streams, db)

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
	authRep := repository.NewAuthenticationsRepository(o.DB)
	organizationRep := repository.NewOrganizationsRepository(o.DB)
	memberRep := repository.NewOrganizationMembersRepository(o.DB)
	rep := usecase.NewAuthenticationRepository(authRep, organizationRep, memberRep)
	defer rep.Close()

	itr := usecase.NewAuthentication(rep)

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
