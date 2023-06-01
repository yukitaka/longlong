package put

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/datastore"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"github.com/yukitaka/longlong/internal/util"
	"strconv"
)

type Options struct {
	CmdParent string
	Operator  *entity.OrganizationMember
	cli.IOStream
}

func NewPutOptions(parent string, streams cli.IOStream, operator *entity.OrganizationMember) *Options {
	return &Options{
		CmdParent: parent,
		Operator:  operator,
		IOStream:  streams,
	}
}

func NewCmdPut(parent string, streams cli.IOStream, operator *entity.OrganizationMember) *cobra.Command {
	o := NewPutOptions(parent, streams, operator)

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
	con, _ := datastore.NewSqliteOpen()
	if id, err := strconv.Atoi(args[0]); err == nil {
		organizationRep := repository.NewOrganizationsRepository(con)
		memberRep := repository.NewOrganizationMembersRepository(con)
		individualRep := repository.NewIndividualsRepository(con)
		rep := usecase.NewOrganizationManagerRepository(organizationRep, memberRep, individualRep)
		defer rep.Close()

		organization, err := organizationRep.Find(id)
		if err != nil {
			return err
		}
		itr := usecase.NewOrganizationManager(organization, rep)
		if individualId, err := strconv.Atoi(args[1]); err == nil {
			if err := itr.AssignIndividual(individualId); err != nil {
				return err
			}
		}
	}

	return err
}
