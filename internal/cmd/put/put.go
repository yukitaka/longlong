package put

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"github.com/yukitaka/longlong/internal/util"
	"strconv"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewPutOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func NewCmdPut(parent string, streams cli.IOStream) *cobra.Command {
	o := NewPutOptions(parent, streams)

	cmd := &cobra.Command{
		Use:     "put",
		Aliases: []string{"p"},
		Short:   "Put one on a resource",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(args))
		},
	}

	organizationCmd := &cobra.Command{
		Use:     "organization",
		Aliases: []string{"organ"},
		Short:   "Put one on a organization",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Organization(cmd, args))
		},
	}
	organizationCmd.PersistentFlags().StringP("output", "o", "table", "output format")
	cmd.AddCommand(organizationCmd)

	return cmd
}

func (o *Options) Run(args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Organization(cmd *cobra.Command, args []string) error {
	var err error
	if id, err := strconv.ParseInt(args[0], 10, 64); err == nil {
		rep := repository.NewOrganizationsRepository()
		itr := usecase.NewOrganizationManager(id, rep, repository.NewOrganizationBelongingsRepository(id, rep))
		if avatarId, err := strconv.ParseInt(args[1], 10, 64); err == nil {
			if err := itr.Entry(avatarId); err != nil {
				return err
			}
		}
	}

	return err
}
