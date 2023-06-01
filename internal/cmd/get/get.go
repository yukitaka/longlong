package get

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/datastore"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	cmdutil "github.com/yukitaka/longlong/internal/cmd/util"
	"github.com/yukitaka/longlong/internal/util"
)

type Options struct {
	CmdParent string
	Operator  *entity.OrganizationMember
	cli.IOStream
}

func NewGetOptions(parent string, streams cli.IOStream, operator *entity.OrganizationMember) *Options {
	return &Options{
		CmdParent: parent,
		Operator:  operator,
		IOStream:  streams,
	}
}

func NewCmdGet(parent string, streams cli.IOStream, operator *entity.OrganizationMember) *cobra.Command {
	o := NewGetOptions(parent, streams, operator)

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Display one or many resources",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(args))
		},
	}

	organizationCmd := &cobra.Command{
		Use:     "organization",
		Aliases: []string{"organ"},
		Short:   "Display one or many organizations",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Organization(cmd, args))
		},
	}
	organizationCmd.PersistentFlags().StringP("output", "o", "table", "output format")

	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Display one or many users",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.User(cmd, args))
		},
	}
	userCmd.PersistentFlags().StringP("output", "o", "table", "output format")
	cmd.AddCommand(organizationCmd)
	cmd.AddCommand(userCmd)

	return cmd
}

func (o *Options) Run(args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Organization(cmd *cobra.Command, args []string) error {
	con, _ := datastore.NewSqliteOpen()
	organizationRep := repository.NewOrganizationsRepository(con)
	itr := usecase.NewOrganizationFinder(organizationRep)

	var err error
	if outputFlag, err := cmd.PersistentFlags().GetString("output"); err == nil {
		var columns []table.Column
		var rows []table.Row
		if outputFlag == "table" {
			columns = []table.Column{
				{Title: "PID", Width: 4},
				{Title: "ID", Width: 4},
				{Title: "Name", Width: 16},
			}
		}
		if len(args) > 0 {
			if id, err := strconv.Atoi(args[0]); err == nil {
				if organization, err := itr.Find(id); err == nil {
					if outputFlag == "table" {
						rows = append(rows, table.Row{strconv.Itoa(organization.ParentId), strconv.Itoa(organization.Id), organization.Name})
					}
					printer := cmdutil.NewPrinter(organization, columns, rows)
					printer.Print()
				}
			}
		} else {
			if organizations, err := itr.List(); err == nil {
				if outputFlag == "table" {
					for _, o := range *organizations {
						rows = append(rows, table.Row{strconv.Itoa(o.ParentId), strconv.Itoa(o.Id), o.Name})
					}
				}
				printer := cmdutil.NewPrinter(organizations, columns, rows)
				printer.Print()
			}
		}
	}

	return err
}

func (o *Options) User(cmd *cobra.Command, args []string) error {
	con, _ := datastore.NewSqliteOpen()
	individualRep := repository.NewIndividualsRepository()
	organizationRep := repository.NewOrganizationsRepository(con)
	memberRep := repository.NewOrganizationMembersRepository()
	rep := usecase.NewUserAssignedRepository(individualRep, organizationRep, memberRep)
	defer rep.Close()

	itr := usecase.NewUserAssigned(rep)
	organizations, err := itr.OrganizationList(o.Operator)
	if err != nil {
		return err
	}

	members := map[string][]entity.OrganizationMember{}
	for _, organization := range *organizations {
		managerRep := usecase.NewOrganizationManagerRepository(organizationRep, memberRep, individualRep)
		manager := usecase.NewOrganizationManager(organization.Organization, managerRep)
		m, err := manager.Members()
		if err != nil {
			continue
		}
		if _, ok := members[organization.Organization.Name]; !ok {
			members[organization.Organization.Name] = []entity.OrganizationMember{}
		}
		members[organization.Organization.Name] = append(members[organization.Organization.Name], *m...)
	}

	if outputFlag, err := cmd.PersistentFlags().GetString("output"); err == nil {
		var columns []table.Column
		var rows []table.Row
		if outputFlag == "table" {
			columns = []table.Column{
				{Title: "Organization", Width: 16},
				{Title: "ID", Width: 4},
				{Title: "Name", Width: 16},
				{Title: "Role", Width: 16},
			}
			for n, ms := range members {
				for _, m := range ms {
					rows = append(rows, table.Row{
						n,
						strconv.Itoa(int(m.Individual.Id)),
						m.Individual.Name,
						m.Role.String(),
					})
				}
			}
		}
		printer := cmdutil.NewPrinter(organizations, columns, rows)
		printer.Print()
	}

	return nil
}
