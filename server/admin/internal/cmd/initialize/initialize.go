package initialize

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/util"
)

type Options struct {
	CmdParent string
	*sqlx.DB
	cli.IOStream
}

func NewCmdInit(parent string, db *sqlx.DB, streams cli.IOStream) *cobra.Command {
	o := newInitOptions(parent, db, streams)

	cmd := &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   "Initialize Longlong",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(args))
		},
	}

	return cmd
}

func newInitOptions(parent string, db *sqlx.DB, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		DB:        db,
		IOStream:  streams,
	}
}

func (o *Options) Run(args []string) error {
	return nil
}
