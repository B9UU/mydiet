package views

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Details struct {
	table   table.Model
	data    []map[string]string
	columns []table.Column
}

func (m Details) Init() tea.Cmd { return nil }

func (m Details) Update(msg tea.Msg) (Details, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Details) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}
func (m Details) getRows() []table.Row {
	var rows []table.Row
	for i, row := range m.data {
		for k, v := range row {
			rows = append(rows, table.Row{string(i), k, v})
		}
	}
	return rows
}
func (m Details) newTable() table.Model {
	t := table.New(
		table.WithColumns(m.columns),
		table.WithRows(m.getRows()),
		table.WithFocused(true),
		table.WithHeight(7),
	)

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
	t.SetStyles(s)
	return t
}
func NewDetailView() Details {
	m := Details{
		data: []map[string]string{{
			"name":  "My Diet",
			"label": "grams",
		}},
		columns: []table.Column{
			{
				Title: "id",
				Width: 4,
			},
			{
				Title: "name",
				Width: 10,
			},
			{Title: "label",
				Width: 10,
			},
		},
	}
	m.table = m.newTable()
	return m
}
