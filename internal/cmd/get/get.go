package get

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/util"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewGetOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func NewCmdGet(parent string, streams cli.IOStream) *cobra.Command {
	o := NewGetOptions(parent, streams)

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Display one or many resources",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(cmd, args))
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:     "organization",
		Aliases: []string{"organ"},
		Short:   "Display one or many organizations",
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
	itr := usecase.NewOrganizationFinder(rep)
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	itr.Find(id)

	return nil
}
