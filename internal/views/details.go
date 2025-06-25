package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Details struct {
	data map[string]any
}

func (m Details) View() string {
	mainBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36")).
		Width(55).
		Align(lipgloss.Center).
		Render(lipgloss.JoinVertical(lipgloss.Center, m.View()))
	return mainBox
}

func (m Details) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
func (m Details) Init() tea.Cmd {
	return nil
}
