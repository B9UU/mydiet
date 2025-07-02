package date

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	datepicker "github.com/ethanefung/bubble-datepicker"
)

type Model struct {
	Date  datepicker.Model
	CTime time.Time
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	// Update logic here...
	var cmd tea.Cmd
	m.Date, cmd = m.Date.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var timeStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		// Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1).
		Border(lipgloss.NormalBorder(), true).
		BorderForeground(lipgloss.Color("#D0BCFF"))
	timeStr := m.Date.Time.Format("Mon Jan 2 15:04:05 2006") // nicer formatted time
	if !m.Date.Selected {
		return timeStyle.Render(timeStr)
	}
	return m.Date.View()
}

func New() Model {
	d := datepicker.New(time.Now())
	d.KeyMap = keys

	return Model{
		Date: d,
	}
}

var keys = datepicker.KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Left: key.NewBinding(
		key.WithKeys("←", "h"),
		key.WithHelp("←/h", "move left"),
	),
	FocusPrev: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "Focus Prev"),
	),
	FocusNext: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "Focus Next"),
	),
}
