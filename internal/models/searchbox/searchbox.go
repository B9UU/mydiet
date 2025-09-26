package searchbox

import (
	"mydiet/internal/logger"
	"mydiet/internal/store"
	"mydiet/internal/types"
	"mydiet/internal/ui/adapters"
	"mydiet/internal/ui/config"
	"mydiet/internal/viewmodels"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	input        textinput.Model
	table        table.Model
	data         store.Foods
	help         help.Model
	keys         searchBoxKeys
	mealType     store.MealType
	Store        store.Store
	tableAdapter *adapters.TableAdapter
	tableConfig  *config.TableConfig
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
			logger.Log.Println("Entering search input")
			cmd = m.GetSuggestions(m.input.Value())
			return m, cmd

		case key.Matches(msg, m.keys.Up),
			key.Matches(msg, m.keys.Down):
			m.table, cmd = m.table.Update(msg)
			return m, cmd

		case key.Matches(msg, m.keys.Select):
			r := m.table.SelectedRow()
			id, err := strconv.Atoi(r[0])
			if err != nil {
				return m, func() tea.Msg {
					return types.ErrMsg(err)
				}
			}
			// m.store.MealsStore.Add(m.mealType, store.Food{
			// 	ID:   id,
			// 	Name: r[1],
			// })
			food := m.data.GetId(id)
			food.Meal = m.mealType
			return m, func() tea.Msg {
				return types.ViewMessage{
					Msg:     food,
					NewView: types.FORMVIEW,
				}
			}
		}
	case types.SuccessRequest:
		if len(msg) == 0 {
			return m, nil
		}
		m.data = store.Foods(msg)
		logger.Log.Printf("Search results retrieved: %d", len(m.data))

		// Convert to view models and then to table rows
		searchViewModels := viewmodels.NewFoodSearchViewModels(m.data)
		rows := m.tableAdapter.FoodSearchToRows(searchViewModels)
		m.table.SetRows(rows)

		return m, nil

	case types.ErrMsg:
		logger.Log.Println("Search request failed")
	}
	m.input, cmd = m.input.Update(msg)
	// m.table, cmd = m.table.Update(msg)
	return m, cmd
}
func New(mType store.MealType, s store.Store) Model {
	// Initialize adapters and config
	tableAdapter := adapters.NewTableAdapter()
	tableConfig := config.NewTableConfig()

	// Create table with columns from config
	t := table.New(
		table.WithColumns(tableConfig.FoodSearchColumns()),
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
		table:        t,
		input:        ti,
		help:         help.New(),
		keys:         keys,
		Store:        s,
		mealType:     mType,
		tableAdapter: tableAdapter,
		tableConfig:  tableConfig,
	}

	// Load initial data and convert to view models
	f, err := m.Store.FoodStore.GetAll("")
	if err != nil {
		logger.Log.Printf("Error retrieving initial data: %v", err)
		panic(err)
	}

	m.data = f
	searchViewModels := viewmodels.NewFoodSearchViewModels(f)
	rows := m.tableAdapter.FoodSearchToRows(searchViewModels)
	m.table.SetRows(rows)

	return m
}
