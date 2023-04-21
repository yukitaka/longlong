package get

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/usecase"
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
	UserId    int64
	cli.IOStream
}

func NewGetOptions(parent string, streams cli.IOStream, userId int64) *Options {
	return &Options{
		CmdParent: parent,
		UserId:    userId,
		IOStream:  streams,
	}
}

func NewCmdGet(parent string, streams cli.IOStream, userId int64) *cobra.Command {
	o := NewGetOptions(parent, streams, userId)

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
	rep := repository.NewOrganizationsRepository()
	itr := usecase.NewOrganizationFinder(rep)

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
			if id, err := strconv.ParseInt(args[0], 10, 64); err == nil {
				if organization, err := itr.Find(id); err == nil {
					if outputFlag == "table" {
						rows = append(rows, table.Row{strconv.FormatInt(organization.ParentId, 10), strconv.FormatInt(organization.Id, 10), organization.Name})
					}
					printer := cmdutil.NewPrinter(organization, columns, rows)
					printer.Print()
				}
			}
		} else {
			if organizations, err := itr.List(); err == nil {
				if outputFlag == "table" {
					for _, o := range *organizations {
						rows = append(rows, table.Row{strconv.FormatInt(o.ParentId, 10), strconv.FormatInt(o.Id, 10), o.Name})
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
	rep := repository.NewIndividualsRepository()

	itr := usecase.NewUserAssigned(o.UserId, rep, repository.NewOrganizationsRepository(), repository.NewOrganizationBelongingsRepository())
	organizations, err := itr.OrganizationList()
	if err != nil {
		return err
	}

	if outputFlag, err := cmd.PersistentFlags().GetString("output"); err == nil {
		var columns []table.Column
		var rows []table.Row
		if outputFlag == "table" {
			columns = []table.Column{
				{Title: "ID", Width: 4},
				{Title: "Organization", Width: 16},
				{Title: "Name", Width: 16},
				{Title: "Role", Width: 16},
			}
			for _, o := range *organizations {
				rows = append(rows, table.Row{
					strconv.FormatInt(o.Organization.Id, 10),
					o.Organization.Name,
					o.Individual.Name,
					o.Role.String(),
				})
			}
		}
		printer := cmdutil.NewPrinter(organizations, columns, rows)
		printer.Print()
	}

	return nil
}
