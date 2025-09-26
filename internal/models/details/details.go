package details

import (
	"mydiet/internal/logger"
	"mydiet/internal/models/date"
	tablelisting "mydiet/internal/models/table"
	"mydiet/internal/store"
	"mydiet/internal/types"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var AllMeals = []store.MealType{
	store.Breakfast, store.Lunch, store.Dinner, store.Snack}

type Model struct {
	style lipgloss.Style

	tables      map[store.MealType]tablelisting.Model
	active      store.MealType
	currentDate time.Time // Track current date for change detection

	date date.Model
	help help.Model
	keys keyMap
}

func (m Model) GetDate() time.Time {
	return m.date.Date.Time
}

// GetCurrentDate returns the current date tracked by the model
func (m Model) GetCurrentDate() time.Time {
	return m.currentDate
}

// GetDateModelTime returns the time from the date model
func (m Model) GetDateModelTime() time.Time {
	return m.date.Date.Time
}

func (m *Model) SyncRowsFor() {
	selectedDate := m.date.Date.Time
	t := m.tables[m.active]
	t.SyncRows(selectedDate)
	m.tables[m.active] = t
}

// SyncAllTables updates all tables with data for the selected date
func (m *Model) SyncAllTables() {
	selectedDate := m.date.Date.Time
	for _, meal := range AllMeals {
		t := m.tables[meal]
		t.SyncRows(selectedDate)
		m.tables[meal] = t
	}
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

	helpView := m.help.View(help.KeyMap(m))
	mainBox := m.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.date.View(),
			upper,
			lower,
			helpView,
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
		case key.Matches(msg, m.keys.Toggle):
			if !m.date.Date.Selected {
				m = m.BlurAll()
				m.date.Date.SelectDate()
				return m, cmd
			} else {
				m = m.SetActive(m.active)
				m.date.Date.UnselectDate()
				return m, cmd
			}
		}

		if m.date.Date.Selected {
			oldDate := m.currentDate
			m.date, cmd = m.date.Update(msg)
			newDate := m.date.Date.Time

			// If date changed, sync all tables with new date
			if !oldDate.Equal(newDate.Truncate(24 * time.Hour)) {
				logger.Log.Printf("Date changed: oldDate=%v (truncated=%v), newDate=%v (truncated=%v)",
					oldDate, oldDate.Truncate(24 * time.Hour), newDate, newDate.Truncate(24 * time.Hour))
				m.currentDate = newDate.Truncate(24 * time.Hour)
				m.SyncAllTables()
			}

			return m, cmd
		}
		switch {

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
func New(s store.Store) Model {
	var tables = make(map[store.MealType]tablelisting.Model)
	for _, k := range AllMeals {
		tables[k] = tablelisting.New(k, s)
	}
	now := time.Now().Truncate(24 * time.Hour)
	m := Model{
		style: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")).
			Align(lipgloss.Center),
		tables:      tables,
		help:        help.New(),
		keys:        keys,
		currentDate: now,
		date:        date.New(),
	}
	return m.SetActive(store.Breakfast)

}

func (m Model) BlurAll() Model {
	for _, i := range AllMeals {
		c := m.tables[i]
		c.Table.Blur()
		m.tables[i] = c
	}
	return m
}
func (m Model) SetActive(meal store.MealType) Model {
	m = m.BlurAll()
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

	Toggle key.Binding
	First  key.Binding
	Second key.Binding
	Third  key.Binding
	Fourth key.Binding
}

func (m Model) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
func (m Model) ShortHelp() []key.Binding {
	dd := []key.Binding{}
	if m.date.Date.Selected {
		keys.Toggle.SetHelp("t", "select table")
		dd = append(dd,
			m.date.Date.KeyMap.Up,
			m.date.Date.KeyMap.Down,
			m.date.Date.KeyMap.Right,
			m.date.Date.KeyMap.Left,
		)
	} else {
		dd = append(dd,
			keys.Add,
			keys.Delete,
			keys.PageUp,
			keys.PageDown,
		)
	}
	return append(
		dd,
		keys.Toggle,
		keys.Help,
		keys.Quit,
	)
}

var keys = keyMap{
	KeyMap: tablelisting.Keys,

	First: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "Breakfast"),
	),
	Second: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "Launch"),
	),
	Third: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "Dinner"),
	),
	Fourth: key.NewBinding(
		key.WithKeys("4"),
		key.WithHelp("4", "Snack"),
	),

	Toggle: key.NewBinding(
		key.WithKeys("t", "T"),
		key.WithHelp("t", "select Time"),
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
