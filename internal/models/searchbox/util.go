package searchbox

import (
	"mydiet/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) GetSuggestions(query string) tea.Cmd {
	return func() tea.Msg {

		foods, err := m.Store.FoodStore.Search(query)
		if err != nil {
			return types.ErrMsg(err)
		}
		return types.SuccessRequest(foods)
	}
}
