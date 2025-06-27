package searchbox

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DODGEBLUE = lipgloss.Color("33")
	DARKCAYAN = lipgloss.Color("36")
	GREY      = lipgloss.Color("240")
)

func (m Model) inputBoxStyle() string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(getColor(m.input.Focused(), DODGEBLUE)).
		Width(50).Render(m.input.View())
}

func (m Model) mainBox() string {
	inputBox := m.inputBoxStyle()
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground().
		Width(55).
		Align(lipgloss.Center).
		Render(lipgloss.JoinVertical(lipgloss.Center, m.table.View(), inputBox))
}

func getColor(b bool, c lipgloss.Color) lipgloss.Color {
	if b {
		return c
	}
	return GREY
}
