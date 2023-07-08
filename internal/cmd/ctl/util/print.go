package util

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yukitaka/longlong/internal/interface/output"
	"gopkg.in/yaml.v3"
	"os"
)

type Options struct {
	data    interface{}
	columns []table.Column
	rows    []table.Row
}

func NewPrinter(data interface{}, columns []table.Column, rows []table.Row) *Options {
	return &Options{
		data:    data,
		columns: columns,
		rows:    rows,
	}
}

func (o *Options) Print() {
	if o.columns == nil {
		if organizationsYaml, err := yaml.Marshal(&o.data); err == nil {
			fmt.Println(string(organizationsYaml))
		}
	} else {
		t := table.New(
			table.WithColumns(o.columns),
			table.WithRows(o.rows),
			table.WithFocused(true),
			table.WithHeight(len(o.rows)),
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
