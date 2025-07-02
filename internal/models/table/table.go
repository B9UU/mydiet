package tablelisting

import (
	"mydiet/internal/store"
	"mydiet/internal/types"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	store    store.Store
	mealName store.MealType
	mealData store.Foods
	style    lipgloss.Style
	Table    table.Model
	Keys     KeyMap
}

func (m *Model) SyncRows() {
	meals, _ := m.store.FoodStore.GetLogs(m.mealName)
	m.Table.SetRows(meals.TableRowsFor())
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Delete):
			// m.mealData = m.store.MealsStore.Delete(
			// 	m.mealName, m.Table.SelectedRow())
			// m.Table.SetRows(m.mealData.TableRowsFor())
			return m, cmd
		case key.Matches(msg, m.Keys.Add):
			return m, func() tea.Msg {
				return types.ViewMessage{
					Msg:     m.mealName,
					NewView: types.SEARCHBOX,
				}
			}
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
			tableName.Render(string(m.mealName)),
			m.Table.View()) + "\n",
	)
}

// New creates a new model with default settings.
func New(mealName store.MealType, st store.Store) Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithHeight(7),
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
		store:    st,
		Keys:     Keys,
		mealName: mealName,
	}
	m.SyncRows()
	// m.mealData = m.store.MealsStore.Get(m.mealName)
	// m.Table.SetRows(m.mealData.TableRowsFor())
	return m
}
