package get

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/usecase"
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
			util.CheckErr(o.Run(cmd, args))
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

func (o *Options) Run(cmd *cobra.Command, args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Organization(cmd *cobra.Command, args []string) error {
	rep := repository.NewOrganizationsRepository()
	itr := usecase.NewOrganizationFinder(rep)

	var err error
	if output, err := cmd.PersistentFlags().GetString("output"); err == nil {
		if len(args) > 0 {
			if id, err := strconv.ParseInt(args[0], 10, 64); err == nil {
				if organization, err := itr.Find(id); err == nil {
					o.print(output, organization)
				}
			}
		} else {
			if organizations, err := itr.List(); err == nil {
				o.print(output, organizations)
			}
		}
	}

	return err
}

var baseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func (o *Options) print(output string, data interface{}) {
	if output == "yaml" {
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
			rows = append(rows, table.Row{strconv.FormatInt(o.ParentID, 10), strconv.FormatInt(o.ID, 10), o.Name})
		} else if orgs, ok := data.(*[]entity.Organization); ok {
			for _, o := range *orgs {
				rows = append(rows, table.Row{strconv.FormatInt(o.ParentID, 10), strconv.FormatInt(o.ID, 10), o.Name})
			}
		}
		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithHeight(len(rows)),
		)

		m := model{t}
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program: ", err)
			os.Exit(1)
		}
	}
}
