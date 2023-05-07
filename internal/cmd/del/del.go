package del

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/util"
)

type Options struct {
	CmdParent string
	Operator  *entity.OrganizationMember
	cli.IOStream
}

func NewDeleteOptions(parent string, streams cli.IOStream, operator *entity.OrganizationMember) *Options {
	return &Options{
		CmdParent: parent,
		Operator:  operator,
		IOStream:  streams,
	}
}

func NewCmdDelete(parent string, streams cli.IOStream, operator *entity.OrganizationMember) *cobra.Command {
	o := NewDeleteOptions(parent, streams, operator)

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete one or many resources",
	}

	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Delete one or many users",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.User(cmd, args))
		},
	}
	userCmd.PersistentFlags().StringP("output", "o", "table", "output format")
	cmd.AddCommand(userCmd)

	return cmd
}

func (o *Options) User(cmd *cobra.Command, args []string) error {
	fmt.Printf("%v\n", args)

	return nil
}
