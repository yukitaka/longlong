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
	o := NewGetOptions(parent, streams)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Display one or many resources",
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(o.Complete(cmd, args))
		},
	}

	return cmd
}

func checkErr(err error) {
	return
}

func (o *GetOptions) Complete(cmd *cobra.Command, args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}
