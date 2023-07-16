package server

import (
	"github.com/spf13/cobra"
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
			util.CheckErr(o.Run())
			return nil
		},
	}

	return cmd
}

func newServerOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func (o *Options) Run() error {
	return nil
}
