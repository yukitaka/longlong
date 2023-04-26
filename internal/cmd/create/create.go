package create

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Options struct {
	CmdParent      string
	UserId         int64
	OrganizationId int64
	cli.IOStream
}

func NewCreateOptions(parent string, streams cli.IOStream, userId, organizationId int64) *Options {
	return &Options{
		CmdParent:      parent,
		UserId:         userId,
		OrganizationId: organizationId,
		IOStream:       streams,
	}
}

func NewCmdCreate(parent string, streams cli.IOStream, userId, organizationId int64) *cobra.Command {
	o := NewCreateOptions(parent, streams, userId, organizationId)

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create one resource",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Error: must also specify a resource like an organization")
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:     "organization",
		Aliases: []string{"organ"},
		Short:   "Create one organization",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Organization(args))
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "user",
		Short: "Create one user",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.User(args))
		},
	})

	return cmd
}

func (o *Options) Organization(args []string) error {
	if len(args) != 1 {
		fmt.Println("Error: must also specify a name")
		return nil
	}
	organizationRep := repository.NewOrganizationsRepository()
	defer organizationRep.Close()
	belongingRep := repository.NewOrganizationBelongingsRepository()
	defer belongingRep.Close()
	itr := usecase.NewOrganizationCreator(organizationRep, belongingRep)

	individual := entity.Individual{UserId: o.UserId}

	name := args[0]
	id, err := itr.Create(name, individual)
	if err != nil {
		return err
	}
	fmt.Printf("create a organization %s which id is %d\n", name, id)

	return nil
}

func (o *Options) User(args []string) error {
	if len(args) != 1 {
		fmt.Println("Error: must also specify a name")
		return nil
	}
	userRep := repository.NewUsersRepository()
	defer userRep.Close()
	individualRep := repository.NewIndividualsRepository()
	defer individualRep.Close()
	belongingRep := repository.NewOrganizationBelongingsRepository()
	defer belongingRep.Close()

	itr := usecase.NewUserCreator(userRep, individualRep, belongingRep)
	name := args[0]
	id, err := itr.New(o.OrganizationId, name)
	if err != nil {
		return err
	}
	fmt.Printf("create a user %s which id is %d\n", name, id)

	return nil
}
