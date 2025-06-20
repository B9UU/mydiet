package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type CartView struct {
	selected     map[int]struct{}
	choices      []string
	cursor       int
	spinnerIndex int
	textInput    textinput.Model
}

func (m CartView) View() string {
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\n" + m.textInput.View() + "\n"
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func (m CartView) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."

	return nil
}

// update the appliatoin state
func (m CartView) Update(msg tea.Msg) (CartView, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		case "q":

			return m, func() tea.Msg {
				return UpdateViewMessage(1)
			}

		case "up":
			if m.cursor > 0 {
				m.cursor--
				return m, cmd
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
				return m, cmd
			}

		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

			return m, cmd
		case "enter":
			m.choices = append(m.choices, string(m.textInput.Value()))
		}
		if len(m.textInput.Value()) >= 3 {
			m.textInput.SetSuggestions(m.choices)
		}
		m.textInput, cmd = m.textInput.Update(msg)
		m.choices = append(m.choices, string(m.textInput.Value()))
	}

	return m, cmd
}
func NewCartView() CartView {
	ti := textinput.New()
	ti.Placeholder = "search"
	ti.Focus()
	ti.ShowSuggestions = true
	ti.CharLimit = 64
	ti.Width = 20
	return CartView{

		textInput: ti,
		selected:  make(map[int]struct{}),
		choices:   []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
	}

}
