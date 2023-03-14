package create

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewCreateOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func NewCmdCreate(parent string, streams cli.IOStream) *cobra.Command {
	o := NewCreateOptions(parent, streams)

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create one resource",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(cmd, args))
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:     "organization",
		Aliases: []string{"organ"},
		Short:   "Create one organization",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Organization(cmd, args))
		},
	})

	return cmd
}

func (o *Options) Run(cmd *cobra.Command, args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Organization(cmd *cobra.Command, args []string) error {
	rep := repository.NewOrganizationsRepository()
	defer rep.Close()
	itr := usecase.NewOrganizationCreator(rep)
	itr.Create(args[0])

	return nil
}
