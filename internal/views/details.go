package views

import (
	tablelisting "mydiet/internal/models/table"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Meal int

const (
	Breakfast Meal = iota
	Lunch
	Dinner
	Snacks
	MealCount
)

type Details struct {
	style  lipgloss.Style
	tables [MealCount]tablelisting.Model
	active Meal
	help   help.Model
	keys   keyMap
}

func (m Details) View() string {

	upper := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.tables[0].View(),
		m.tables[1].View())
	lower := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.tables[2].View(),
		m.tables[3].View())

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

func (m Details) Update(msg tea.Msg) (Details, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			// return m, tea.Println("help")
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

	}
	m.tables[m.active], cmd = m.tables[m.active].Update(msg)
	return m, cmd
}
func (m Details) Init() tea.Cmd {
	return nil
}

// New creates a new model with default settings.
func New() Details {
	var tables [MealCount]tablelisting.Model

	for i := Meal(0); i < MealCount; i++ {
		meal := Meal(i)
		tables[i] = tablelisting.New(meal.String())
	}
	m := Details{
		style: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")).
			Align(lipgloss.Center),
		tables: tables,
		help:   help.New(),
		keys:   keys,
	}
	return m.SetActive(0)

}
func (m Details) SetActive(meal Meal) Details {
	for i := Meal(0); i < MealCount; i++ {
		m.tables[i].Table.Blur()
	}
	m.tables[meal].Table.Focus()
	m.active = meal
	return m
}

type keyMap struct {
	tablelisting.KeyMap
	Help key.Binding
	Quit key.Binding
}

func (k keyMap) FullHelp() [][]key.Binding {
	dd := k.KeyMap.FullHelp()
	dd[2] = append(
		dd[2],
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

	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys(tea.KeyCtrlC.String(), "q"),
		key.WithHelp("Ctrl-c/q", "Quit"),
	),
}

func (m Meal) String() string {
	switch m {
	case Breakfast:
		return "Breakfast"
	case Lunch:
		return "Lunch"
	case Dinner:
		return "Dinner"
	case Snacks:
		return "Snacks"
	default:
		return "Unknown"
	}
}
