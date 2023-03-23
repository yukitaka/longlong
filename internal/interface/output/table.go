package output

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	help      = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type Model struct {
	enter func(string) tea.Cmd
	table table.Model
}

func NewModel(enter func(string) tea.Cmd, tbl table.Model) Model {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	tbl.SetStyles(s)

	return Model{
		enter: enter,
		table: tbl,
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				m.enter(m.table.SelectedRow()[1]),
			)
		}
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return baseStyle.Render(m.table.View()) + "\n" + help.Render(" Quit[q] Up[↑/k] Down[↓/j]") + "\n"
}
