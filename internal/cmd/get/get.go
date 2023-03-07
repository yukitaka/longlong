package get

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
)

type GetOptions struct {
	CmdParent string
	cli.IOStream
}

func NewGetOptions(parent string, streams cli.IOStream) *GetOptions {
	return &GetOptions{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func NewCmdGet(parent string, streams cli.IOStream) *cobra.Command {
	_ = NewGetOptions(parent, streams)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Display one or many resources",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("in run")
			fmt.Println(args)
		},
	}

	return cmd
}

func checkErr(err error) {
	return
}
