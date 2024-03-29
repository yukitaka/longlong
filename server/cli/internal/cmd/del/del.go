package del

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
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

func NewDeleteOptions(parent string, streams cli.IOStream, operator *entity.OrganizationMember, db *sqlx.DB) *Options {
	return &Options{
		CmdParent: parent,
		Operator:  operator,
		DB:        db,
		IOStream:  streams,
	}
}

func NewCmdDelete(parent string, streams cli.IOStream, operator *entity.OrganizationMember, db *sqlx.DB) *cobra.Command {
	o := NewDeleteOptions(parent, streams, operator, db)

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
	if len(args) != 1 {
		fmt.Println("Error: must also specify a name")
		return nil
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	organizationRep := repository.NewOrganizationsRepository(o.DB)
	memberRep := repository.NewOrganizationMembersRepository(o.DB)
	individualRep := repository.NewIndividualsRepository(o.DB)
	rep := usecase.NewOrganizationManagerRepository(organizationRep, memberRep, individualRep)
	defer rep.Close()

	organization, err := organizationRep.Find(o.Operator.Organization.Id)
	if err != nil {
		return err
	}

	itr := usecase.NewOrganizationManager(organization, rep)
	err = itr.QuitIndividual(o.Operator, id, "Delete by "+o.Operator.Individual.Name)
	if err != nil {
		return err
	}

	return nil
}
