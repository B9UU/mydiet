package searchbox

import (
	"mydiet/internal/logger"
	"mydiet/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) GetSuggestions(query string) tea.Cmd {
	return func() tea.Msg {
		logger.Log.Info("in")

		foods, err := m.store.FoodStore.Search(query)
		if err != nil {
			return types.FailedRequest(err)
		}
		return types.SuccessRequest(foods)
	}
}
