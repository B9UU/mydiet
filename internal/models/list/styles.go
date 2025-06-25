package list

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	FOCUSED   = lipgloss.Color("33")
	UNFOCUSED = lipgloss.Color("240")
)

func (m Model) getColor() lipgloss.Color {
	if m.Focused() {
		return m.Color
	}
	return UNFOCUSED
}
func (m Model) selectedRow() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("236")).
		Width(m.Width).
		Bold(true)
}
func arrowStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Width(10).
		Align(lipgloss.Center)
}
