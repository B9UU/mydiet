package search

import "github.com/charmbracelet/lipgloss"

var (
	FOCUSED   = lipgloss.Color("33")
	UNFOCUSED = lipgloss.Color("240")
)

func (m Model) inputBoxStyle() string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.getColor()).
		Width(50).Render(m.textInput.View())
}

func (m Model) mainBox() string {
	inputBox := m.inputBoxStyle()
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36")).
		Width(55).
		Align(lipgloss.Center).
		Render(lipgloss.JoinVertical(lipgloss.Center, m.listView.View(), inputBox))
}

func (m Model) getColor() lipgloss.Color {
	if m.textInput.Focused() {
		return FOCUSED
	}
	return UNFOCUSED
}
