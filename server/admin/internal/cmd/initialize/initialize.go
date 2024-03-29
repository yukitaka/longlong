package initialize

import (
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/interface/config"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/util"
)

type Options struct {
	CmdParent string
	*config.Config
	*datastore.Connection
	cli.IOStream
}

func NewCmdInit(parent string, config *config.Config, streams cli.IOStream) *cobra.Command {
	o := newInitOptions(parent, config, streams)

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

func newInitOptions(parent string, config *config.Config, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		Config:    config,
		IOStream:  streams,
	}
}

func (o *Options) Run(driver, source string) error {
	con, err := datastore.NewConnectionOpen(driver, source)
	if err != nil {
		return err
	}
	defer con.Close()

	o.Connection = con

	o.Config.SetDatastore(driver, source)
	if err := o.Config.Store(); err != nil {
		return err
	}
	if err := NewDatabase(con).Init(); err != nil {
		return err
	}

	return nil
}
