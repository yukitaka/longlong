package initialize

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/util"
)

type Options struct {
	CmdParent string
	*sqlx.DB
	cli.IOStream
}

func NewCmdInit(parent string, streams cli.IOStream) *cobra.Command {
	o := newInitOptions(parent, streams)

	cmd := &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   "Initialize Longlong",
		RunE: func(cmd *cobra.Command, args []string) error {
			driver, err := cmd.Flags().GetString("driver")
			if err != nil {
				return err
			}
			source, err := cmd.Flags().GetString("source")
			if err != nil {
				return err
			}
			util.CheckErr(o.Run(driver, source))
			return nil
		},
	}
	cmd.Flags().StringP("driver", "d", "postgres", "database driver")
	cmd.Flags().StringP("source", "s", "user=postgres password=postgres dbname=postgres sslmode=disable", "database source")

	return cmd
}

func newInitOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func (o *Options) Run(driver, source string) error {
	open, err := datastore.NewConnectionOpen(driver, source)
	if err != nil {
		return err
	}
	o.DB = open

	fmt.Printf("%#v\n", o.DB)

	return nil
}
