package tablelisting

import (
	"mydiet/internal/store"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	mealName string
	mealData store.MealsData
	style    lipgloss.Style
	Table    table.Model
	keys     KeyMap
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Esc):
			m.Table.Blur()
		case key.Matches(msg, m.keys.Enter):
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[2]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	style := m.style
	tableName := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true)
	if m.Table.Focused() {
		style = style.BorderForeground(lipgloss.Color("99"))
		tableName = tableName.Foreground(lipgloss.Color("205")) // pinkish
	} else {
		style = style.BorderForeground(lipgloss.Color("240"))   // grey border
		tableName = tableName.Foreground(lipgloss.Color("238")) // dark grey

		// Change selected row style when not focused
		s := table.DefaultStyles()
		s.Selected = s.Selected.Foreground(lipgloss.Color("240")).Background(lipgloss.Color("236"))
		m.Table.SetStyles(s)
	}

	return style.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			tableName.Render(m.mealName),
			m.Table.View()) + "\n",
	)
}

// New creates a new model with default settings.
func New(mealName string) Model {
	meals := store.Meals[store.MealType(mealName)]
	t := table.New(
		table.WithColumns(columns),
		table.WithHeight(7),
		table.WithRows(meals.TableRowsFor()),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)

	m := Model{
		style:    baseStyle,
		Table:    t,
		keys:     Keys,
		mealName: mealName,
		mealData: meals,
	}
	return m
}
