package debug

import (
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewCmdDebug(parent string, streams cli.IOStream) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug",
		Short: "Debug Longlong",
	}

	return cmd
}
