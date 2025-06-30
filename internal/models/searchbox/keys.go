package searchbox

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type searchBoxKeys struct {
	Up     key.Binding
	Down   key.Binding
	Search key.Binding
	Select key.Binding
	Cancel key.Binding
	Quit   key.Binding
}

func (k searchBoxKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Search, k.Select, k.Quit}
}

func (m searchBoxKeys) FullHelp() [][]key.Binding {
	return nil
}

var keys = searchBoxKeys{

	Up: key.NewBinding(
		key.WithKeys(tea.KeyUp.String()),
		key.WithHelp("↑", "Up"),
	),
	Down: key.NewBinding(
		key.WithKeys(tea.KeyDown.String()),
		key.WithHelp("↓", "Down"),
	),
	Search: key.NewBinding(
		key.WithKeys(tea.KeyEnter.String()),
		key.WithHelp("Enter", "Search.."),
	),

	Select: key.NewBinding(
		key.WithKeys(tea.KeyTab.String()),
		key.WithHelp("Tab", "Select.."),
	),
	Cancel: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp("Esc", "Cancell"),
	),
	Quit: key.NewBinding(
		key.WithKeys(tea.KeyCtrlC.String()),
		key.WithHelp("ctrl+c", "quit"),
	),
}
