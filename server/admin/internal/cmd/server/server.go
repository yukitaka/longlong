package server

import (
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/admin/internal/interface/server"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/interface/config"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"github.com/yukitaka/longlong/server/core/pkg/util"
)

type Options struct {
	CmdParent string
	*config.Config
	cli.IOStream
}

func NewCmdServer(parent string, config *config.Config, streams cli.IOStream) *cobra.Command {
	o := newServerOptions(parent, config, streams)

	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"s", "serv"},
		Short:   "Serv Longlong",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				return err
			}
			util.CheckErr(o.Run(port))
			return nil
		},
	}
	cmd.Flags().IntP("port", "p", 8080, "port number")

	return cmd
}

func newServerOptions(parent string, config *config.Config, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		Config:    config,
		IOStream:  streams,
	}
}

func (o *Options) Run(port int) error {
	con, err := datastore.NewConnectionOpen(o.Config.Datastore.Driver, o.Config.Datastore.Source)
	if err != nil {
		return err
	}
	defer con.Close()

	s := server.NewServer()
	s.Run(port)
	return nil
}
