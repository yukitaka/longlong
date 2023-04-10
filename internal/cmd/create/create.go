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
	CmdParent string
	UserId    int64
	cli.IOStream
}

func NewCreateOptions(parent string, streams cli.IOStream, userId int64) *Options {
	return &Options{
		CmdParent: parent,
		UserId:    userId,
		IOStream:  streams,
	}
}

func NewCmdCreate(parent string, streams cli.IOStream, userId int64) *cobra.Command {
	o := NewCreateOptions(parent, streams, userId)

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
	rep := repository.NewOrganizationsRepository()
	defer rep.Close()
	belongingsRep := repository.NewOrganizationBelongingsRepository(rep, -1)
	defer belongingsRep.Close()
	itr := usecase.NewOrganizationCreator(rep, belongingsRep)

	avatar := entity.Avatar{UserId: o.UserId}

	name := args[0]
	id, err := itr.Create(name, avatar)
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
	rep := repository.NewUsersRepository()
	defer rep.Close()
	avatarRep := repository.NewAvatarsRepository()
	defer avatarRep.Close()

	itr := usecase.NewUserCreator(rep, avatarRep)
	name := args[0]
	id, err := itr.New(name)
	if err != nil {
		return err
	}
	fmt.Printf("create a user %s which id is %d\n", name, id)

	return nil
}
