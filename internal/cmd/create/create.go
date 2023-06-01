package create

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/datastore"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Options struct {
	CmdParent string
	Operator  *entity.OrganizationMember
	cli.IOStream
}

func NewCreateOptions(parent string, streams cli.IOStream, member *entity.OrganizationMember) *Options {
	return &Options{
		CmdParent: parent,
		Operator:  member,
		IOStream:  streams,
	}
}

func NewCmdCreate(parent string, streams cli.IOStream, member *entity.OrganizationMember) *cobra.Command {
	o := NewCreateOptions(parent, streams, member)

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

	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Create one user",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.User(cmd, args))
		},
	}
	userCmd.PersistentFlags().StringP("role", "r", "member", "user role")
	cmd.AddCommand(userCmd)

	profileCmd := &cobra.Command{
		Use:   "profile",
		Short: "Create one profile",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Profile(cmd, args))
		},
	}
	cmd.AddCommand(profileCmd)

	return cmd
}

func (o *Options) Organization(args []string) error {
	if len(args) != 1 {
		fmt.Println("Error: must also specify a name")
		return nil
	}
	organizationRep := repository.NewOrganizationsRepository()
	memberRep := repository.NewOrganizationMembersRepository()
	rep := usecase.NewOrganizationCreatorRepository(organizationRep, memberRep)
	defer rep.Close()
	itr := usecase.NewOrganizationCreator(rep)

	name := args[0]
	id, err := itr.New(name, *o.Operator.Individual)
	if err != nil {
		return err
	}
	fmt.Printf("create a organization %s which id is %d\n", name, id)

	return nil
}

func (o *Options) User(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		fmt.Println("Error: must also specify a name")
		return nil
	}

	if role, err := cmd.PersistentFlags().GetString("role"); err == nil {
		con, _ := datastore.NewSqliteOpen()
		userRep := repository.NewUsersRepository(con)
		defer userRep.Close()
		individualRep := repository.NewIndividualsRepository()
		defer individualRep.Close()
		memberRep := repository.NewOrganizationMembersRepository()
		defer memberRep.Close()

		rep := usecase.NewUserCreatorRepository(userRep, individualRep, memberRep)
		itr := usecase.NewUserCreator(rep)
		name := args[0]
		id, err := itr.New(o.Operator, name, role)
		if err != nil {
			return err
		}
		fmt.Printf("create a user %s which id is %d\n", name, id)
	}

	return nil
}

func (o *Options) Profile(cmd *cobra.Command, args []string) error {
	if len(args) != 3 {
		fmt.Println("Error: must also specify a nickname and full name and bio")
		return nil
	}
	con, _ := datastore.NewSqliteOpen()
	itr := usecase.NewProfileCreator(repository.NewProfilesRepository(con))
	_, err := itr.New(o.Operator, args[0], args[1], args[2])
	if err != nil {
		return err
	}

	return nil
}
