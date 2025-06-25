package views

import "github.com/charmbracelet/lipgloss"

var (
	FOCUSED   = lipgloss.Color("33")
	UNFOCUSED = lipgloss.Color("240")
)

func (m Search) getColor() lipgloss.Color {
	if m.textInput.Focused() {
		return FOCUSED
	}
	return UNFOCUSED
}
