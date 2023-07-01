package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/cmd/config"
	"github.com/yukitaka/longlong/internal/interface/authentication"
	"github.com/yukitaka/longlong/internal/util"
	"log"
)

type Options struct {
	CmdParent string
	*config.Config
	*sqlx.DB
	cli.IOStream
}

func NewAuthOptions(parent string, config *config.Config, db *sqlx.DB, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		Config:    config,
		DB:        db,
		IOStream:  streams,
	}
}

func NewCmdAuth(parent string, config *config.Config, db *sqlx.DB, streams cli.IOStream) *cobra.Command {
	o := NewAuthOptions(parent, config, db, streams)

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
	oauth := authentication.NewOAuth(o.Config.Authorize.AccessToken, o.Config.Authorize.RefreshToken, o.Config.Authorize.Expiry)
	err := oauth.Run(o.DB)
	o.Config.Store(oauth.AccessToken, oauth.RefreshToken, oauth.Expiry)
	if err != nil {
		return err
	}

	return nil
}
