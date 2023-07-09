package put

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/cli/internal/cli"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/usecase"
	"github.com/yukitaka/longlong/server/core/pkg/interface/repository"
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"strconv"
)

type Options struct {
	CmdParent string
	Operator  *entity.OrganizationMember
	*sqlx.DB
	cli.IOStream
}

func NewPutOptions(parent string, streams cli.IOStream, operator *entity.OrganizationMember, db *sqlx.DB) *Options {
	return &Options{
		CmdParent: parent,
		Operator:  operator,
		DB:        db,
		IOStream:  streams,
	}
}

func NewCmdPut(parent string, streams cli.IOStream, operator *entity.OrganizationMember, db *sqlx.DB) *cobra.Command {
	o := NewPutOptions(parent, streams, operator, db)

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
	if id, err := strconv.Atoi(args[0]); err == nil {
		organizationRep := repository.NewOrganizationsRepository(o.DB)
		memberRep := repository.NewOrganizationMembersRepository(o.DB)
		individualRep := repository.NewIndividualsRepository(o.DB)
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
