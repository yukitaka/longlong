package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/interface/authentication"
	"github.com/yukitaka/longlong/internal/util"
	"log"
)

type Options struct {
	CmdParent string
	*sqlx.DB
	cli.IOStream
}

func NewAuthOptions(parent string, streams cli.IOStream, db *sqlx.DB) *Options {
	return &Options{
		CmdParent: parent,
		DB:        db,
		IOStream:  streams,
	}
}

func NewCmdAuth(parent string, streams cli.IOStream, db *sqlx.DB) *cobra.Command {
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
		Use:   "login [ORGANIZATION]",
		Short: "Authorize access to Longlong",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
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
	log.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Login(args []string) error {
	log.Println("Start login.")
	oauth := authentication.NewOAuth()
	err := oauth.Run(o.DB)
	if err != nil {
		return err
	}

	return nil
}
