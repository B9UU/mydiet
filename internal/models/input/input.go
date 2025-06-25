package input

import (
	"mydiet/internal/models/textinput"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textinput.Model
	keys keyMap
	help help.Model
	Used bool
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func New() Model {
	return Model{
		Model: textinput.New(),
		keys:  keys,
		help:  help.New(),
		Used:  true,
	}
}

func (m Model) GetHelp() string {
	return m.help.View(m.keys)

}
