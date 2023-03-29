package get

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/output"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/util"
	"gopkg.in/yaml.v3"
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
	cmd.AddCommand(organizationCmd)

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
		if len(args) > 0 {
			if id, err := strconv.ParseInt(args[0], 10, 64); err == nil {
				if organization, err := itr.Find(id); err == nil {
					o.print(outputFlag, organization)
				}
			}
		} else {
			if organizations, err := itr.List(); err == nil {
				o.print(outputFlag, organizations)
			}
		}
	}

	return err
}

func (o *Options) print(format string, data interface{}) {
	if format == "yaml" {
		if organizationsYaml, err := yaml.Marshal(&data); err == nil {
			fmt.Println(string(organizationsYaml))
		}
	} else {
		columns := []table.Column{
			{Title: "PID", Width: 4},
			{Title: "ID", Width: 4},
			{Title: "Name", Width: 16},
		}
		var rows []table.Row
		if o, ok := data.(*entity.Organization); ok {
			rows = append(rows, table.Row{strconv.FormatInt(o.ParentId, 10), strconv.FormatInt(o.Id, 10), o.Name})
		} else if organizations, ok := data.(*[]entity.Organization); ok {
			for _, o := range *organizations {
				rows = append(rows, table.Row{strconv.FormatInt(o.ParentId, 10), strconv.FormatInt(o.Id, 10), o.Name})
			}
		}
		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(len(rows)),
		)

		m := output.NewModel(func(id string) tea.Cmd {
			return tea.Printf("unfold %s", id)
		}, t)
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program: ", err)
			os.Exit(1)
		}
	}
}
