package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

type GetOptions struct {
	CmdParent string
}

func NewGetOptions(parent string) *GetOptions {
	return &GetOptions{
		CmdParent: parent,
	}
}

func NewCmdGet(parent string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	return cmd
}
