package tablelisting

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
)

type KeyMap struct {
	table.KeyMap

	Delete key.Binding
	Edit   key.Binding
	Add    key.Binding
}

var Keys = KeyMap{
	KeyMap: table.DefaultKeyMap(),
	Add: key.NewBinding(
		key.WithKeys("a", "A"),
		key.WithHelp("a", "add row"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d", "D"),
		key.WithHelp("d", "delete row"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e", "E"),
		key.WithHelp("Enter", "Select"),
	),
}
