package tablelisting

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	table.KeyMap

	Delete key.Binding
	Enter  key.Binding
	Esc    key.Binding
	Add    key.Binding
}

func (k KeyMap) FullHelp() [][]key.Binding {
	dd := k.KeyMap.FullHelp()
	dd = append(
		dd,
		[]key.Binding{
			k.Enter,
			k.Esc,
		},
	)
	return dd
}
func (k KeyMap) ShortHelp() []key.Binding {
	dd := k.KeyMap.ShortHelp()
	// dd = append(
	// 	dd,
	// 	k.Enter,
	// 	k.Esc,
	// )
	return dd
}

var Keys = KeyMap{
	KeyMap: table.DefaultKeyMap(),
	Add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add row"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete row"),
	),
	Enter: key.NewBinding(
		key.WithKeys(tea.KeyEnter.String()),
		key.WithHelp("Enter", "Select"),
	),
	Esc: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp("Esc", "Unfocus"),
	),
}
