package searchbox

import (
	"mydiet/internal/logger"
	"mydiet/internal/store"
	"mydiet/internal/types"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	input    textinput.Model
	table    table.Model
	data     store.Foods
	help     help.Model
	keys     searchBoxKeys
	mealType store.MealType
	store    store.Store
}

func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center, m.mainBox(),
		m.help.View(m.keys),
	)
}

func (m Model) Init() tea.Cmd {
	return nil
}
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Cancel):
			return m, func() tea.Msg {
				return types.ViewMessage{
					NewView: types.DETAILSVIEW,
				}
			}
		case key.Matches(msg, m.keys.Search):
			logger.Log.Info("enter input")
			cmd = m.GetSuggestions(m.input.Value())
			return m, cmd

		case key.Matches(msg, m.keys.Up),
			key.Matches(msg, m.keys.Down):
			m.table, cmd = m.table.Update(msg)
			return m, cmd

		case key.Matches(msg, m.keys.Select):
			r := m.table.SelectedRow()
			id, _ := strconv.Atoi(r[0])
			m.store.MealsStore.Add(m.mealType, store.Food{
				ID:   id,
				Name: r[1],
			})
			return m, func() tea.Msg {
				return types.ViewMessage{
					Msg:     "updated",
					NewView: types.DETAILSVIEW,
				}
			}
		}
	case types.SuccessRequest:
		if len(msg) == 0 {
			return m, nil
		}
		m.data = store.Foods(msg)
		logger.Log.Info("Got messages: ", len(m.data), msg)
		m.table.SetRows(m.data.SearchRows())
		return m, nil

	case types.FailedRequest:

		logger.Log.Info("Failed")
	}
	m.input, cmd = m.input.Update(msg)
	// m.table, cmd = m.table.Update(msg)
	return m, cmd
}
func New(mType store.MealType, s store.Store) Model {
	t := table.New(
		table.WithColumns([]table.Column{
			{Title: "Id", Width: 0},
			{Title: "Name", Width: 10},
			{Title: "Calories", Width: 10},
		}),
		table.WithHeight(7),
	)

	t.Focus()
	t.SetStyles(table.DefaultStyles())

	ti := textinput.New()
	ti.Placeholder = string(mType)
	ti.ShowSuggestions = true
	ti.CharLimit = 64
	ti.Width = 20
	ti.Focus()
	m := Model{
		table:    t,
		input:    ti,
		help:     help.New(),
		keys:     keys,
		mealType: mType,
		store:    s,
	}
	f, err := m.store.FoodStore.GetAll("")
	if err != nil {
		logger.Log.Error(err)
	}
	m.data = f
	m.table.SetRows(m.data.SearchRows())

	return m
}
