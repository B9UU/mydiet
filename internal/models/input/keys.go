package input

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Esc   key.Binding
	Quit  key.Binding
}

func (k keyMap) FullHelp() [][]key.Binding {
	return nil
}
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Enter,
		k.Up,
		k.Down,
		// k.Esc,
		k.Quit,
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys(
			tea.KeyEnter.String(),
			tea.KeyTab.String()),
		key.WithHelp("Enter", "Select"),
	),

	Esc: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp("Esc", "Unfocus"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
