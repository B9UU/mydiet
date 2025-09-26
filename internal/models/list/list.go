package list

import (
	"mydiet/internal/logger"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Choices []string
	Cursor  int
	focus   bool

	keys   keyMap
	help   help.Model
	height int

	Width  int
	Border lipgloss.Border
	Color  lipgloss.Color

	showRows int
}

func (m Model) View() string {
	var tmp []string

	// Calculate start so that cursor is always visible, moving one row at a time
	start := m.Cursor - (m.showRows - 1)
	if start < 0 {
		start = 0
	}
	// Make sure window doesn't go beyond the last element
	if start+m.showRows > len(m.Choices) {
		start = len(m.Choices) - m.showRows
		if start < 0 {
			start = 0
		}
	}
	end := start + m.showRows
	if end > len(m.Choices) {
		end = len(m.Choices)
	}

	if start > 0 {
		tmp = append(tmp, arrowStyle().Render("↑"))
	} else {
		tmp = append(tmp, arrowStyle().Render(" "))
	}

	var suggestions []string
	for i := start; i < end; i++ {
		choice := m.Choices[i]
		if m.Cursor == i {
			styled := m.selectedRow().
				Render(" " + choice)
			suggestions = append(suggestions, styled)
		} else {
			suggestions = append(suggestions, " "+choice)
		}
	}
	tmp = append(tmp, suggestions...)

	if end < len(m.Choices) {
		tmp = append(tmp, arrowStyle().Render("↓"))
	} else {
		tmp = append(tmp, arrowStyle().Render(" "))
	}

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.getColor()).
		Width(m.Width).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left, tmp...,
			))
	return box
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if winMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.height = winMsg.Height
	}
	if !m.Focused() {
		return m, nil
	}
	var cmd tea.Cmd

	logger.Log.Println("Focused list")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// next

		case tea.KeyDown.String():
			logger.Log.Println("Cursor moved down")
			m.Cursor++
			if m.Cursor >= len(m.Choices) {
				m.Cursor = 0
			}
		// prev
		case tea.KeyUp.String():
			logger.Log.Println("Cursor moved up")
			m.Cursor--
			if m.Cursor < 0 {
				m.Cursor = len(m.Choices) - 1
			}
		}
	}

	logger.Log.Println("Current cursor position")
	return m, cmd
}

func (m Model) Init() tea.Cmd {

	return nil
}

func (m Model) Focused() bool {
	return m.focus
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}
func NewListView(v []string) Model {
	return Model{
		Choices:  v,
		Cursor:   0,
		Width:    50,
		showRows: 3,
		Border:   lipgloss.RoundedBorder(),
		Color:    lipgloss.Color("36"),
	}
}

func (m Model) GetHelp() string {
	return m.help.View(m.keys)

}
