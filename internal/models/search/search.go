package search

import (
	"context"
	"encoding/json"
	"fmt"
	"mydiet/internal/logger"
	"mydiet/internal/models/input"
	"mydiet/internal/models/list"
	"mydiet/internal/types"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	selected     map[int]struct{}
	spinnerIndex int
	textInput    input.Model
	listView     list.Model
	lastRequest  time.Time
}

func (m Model) View() string {

	mainBox := m.mainBox()
	mainBox += "\n" + m.textInput.GetHelp() + "\n"
	return mainBox
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == tea.KeyCtrlC.String() {
			return m, tea.Quit
		}
		switch msg.String() {
		case tea.KeyEsc.String():
			m.listView.Focus()
			return m, cmd
		}

		m.textInput, cmd = m.textInput.Update(msg)
		m.listView, cmd = m.listView.Update(msg)
		logger.Log.Info(len(m.textInput.Value()))
		if len(m.textInput.Value()) >= 3 {
			if time.Since(m.lastRequest) > 30*time.Second {
				cmd = tea.Batch(cmd,
					m.GetSuggestions(
						m.textInput.Value(),
					))
				m.lastRequest = time.Now()
			}
		}

	case types.SuccessRequest:

		newSug := append(
			m.textInput.AvailableSuggestions(),
			msg.Suggetions...)
		m.textInput.SetSuggestions(newSug)

		m.listView.Choices = m.textInput.MatchedSuggestions()
		m.listView.Cursor = m.textInput.CurrentSuggestionIndex()
		return m, cmd
	case types.FailedRequest:
		logger.Log.Info("got the error message")
		logger.Log.Error(msg.Err)
		return m, cmd
	case types.ViewMessage:
		logger.Log.Info("got the message")
		return m, cmd
	}

	m.listView.Choices = m.textInput.MatchedSuggestions()
	m.listView.Cursor = m.textInput.CurrentSuggestionIndex()
	return m, cmd
}

func (m Model) GetSuggestions(query string) tea.Cmd {
	return func() tea.Msg {
		logger.Log.Info("in")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		url := "https://dummyjson.com/products/search?q="
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return types.FailedRequest{Err: err}
		}
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return types.FailedRequest{Err: err}
		}

		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return types.FailedRequest{
				Err: fmt.Errorf("Failed Request with status: %d",
					response.StatusCode),
			}
		}
		var data DummyJson
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return types.FailedRequest{Err: err}
		}
		logger.Log.Info("extracting titles")
		return types.SuccessRequest{Suggetions: extractTitles(data)}
	}
}

func extractTitles(data DummyJson) []string {
	titles := make([]string, 0, len(data.Products))
	for _, p := range data.Products {
		titles = append(titles, p.Title)
	}
	return titles
}

func New() Model {
	ti := input.New()
	ti.Placeholder = "search"
	ti.Focus()
	ti.ShowSuggestions = true
	ti.CharLimit = 64
	ti.Width = 20
	ti.SetSuggestions([]string{"Buy carrots", "Buy celery", "Buy kohlrabi", "4", "4"})
	return Model{
		textInput: ti,
		listView:  list.NewListView(ti.AvailableSuggestions()),
	}
}
