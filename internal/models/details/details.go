package details

import (
	tablelisting "mydiet/internal/models/table"
	"mydiet/internal/store"
	"mydiet/internal/types"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var AllMeals = []store.MealType{
	store.Breakfast, store.Lunch, store.Dinner, store.Snack}

type Model struct {
	style  lipgloss.Style
	tables map[store.MealType]tablelisting.Model
	active store.MealType
	help   help.Model
	keys   keyMap
}

func (m Model) View() string {

	upper := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.tables[store.Breakfast].View(),
		m.tables[store.Lunch].View())
	lower := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.tables[store.Dinner].View(),
		m.tables[store.Snack].View())

	mainBox := m.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			upper,
			lower,
			m.help.View(m.keys),
		),
	)
	return mainBox
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case types.ViewMessage:
		return m, cmd

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			// return m, tea.Println("help")
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.First):
			m = m.SetActive(store.Breakfast)
		case key.Matches(msg, m.keys.Second):
			m = m.SetActive(store.Lunch)
		case key.Matches(msg, m.keys.Third):
			m = m.SetActive(store.Dinner)
		case key.Matches(msg, m.keys.Fourth):
			m = m.SetActive(store.Snack)
		}
	}
	m.tables[m.active], cmd = m.tables[m.active].Update(msg)
	return m, cmd
}
func (m Model) Init() tea.Cmd {
	return nil
}

// New creates a new model with default settings.
func New(s *store.Store) Model {
	var tables = make(map[store.MealType]tablelisting.Model)
	for _, k := range AllMeals {
		tables[k] = tablelisting.New(k, s)
	}
	m := Model{
		style: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")).
			Align(lipgloss.Center),
		tables: tables,
		help:   help.New(),
		keys:   keys,
	}
	return m.SetActive(store.Breakfast)

}

func (m Model) SetActive(meal store.MealType) Model {
	for _, i := range AllMeals {
		c := m.tables[i]
		c.Table.Blur()
		m.tables[i] = c
	}
	t := m.tables[meal]
	t.Table.Focus()
	m.tables[meal] = t

	m.active = meal
	return m
}

type keyMap struct {
	tablelisting.KeyMap
	Help key.Binding
	Quit key.Binding

	First  key.Binding
	Second key.Binding
	Third  key.Binding
	Fourth key.Binding
}

func (k keyMap) FullHelp() [][]key.Binding {
	dd := k.KeyMap.FullHelp()
	dd[2] = append(
		dd[2],
		k.First,
		k.Second,
		k.Third,
		k.Fourth,
		k.Help,
		k.Quit,
	)
	return dd
}
func (k keyMap) ShortHelp() []key.Binding {
	dd := k.KeyMap.ShortHelp()
	return append(
		dd,
		k.Help,
		// k.Quit,
	)
}

var keys = keyMap{
	KeyMap: tablelisting.Keys,

	First: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("?", "Help"),
	),
	Second: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("?", "Help"),
	),
	Third: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("?", "Help"),
	),
	Fourth: key.NewBinding(
		key.WithKeys("4"),
		key.WithHelp("?", "Help"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys(tea.KeyCtrlC.String(), "q"),
		key.WithHelp("Ctrl-c/q", "Quit"),
	),
}
