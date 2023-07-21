package server

import (
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/admin/internal/interface/server"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/util"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewCmdServer(parent string, streams cli.IOStream) *cobra.Command {
	o := newServerOptions(parent, streams)

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

func newServerOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func (o *Options) Run(port int) error {
	s := server.NewServer()
	s.Run(port)
	return nil
}
