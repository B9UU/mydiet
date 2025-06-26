package views

import (
	"context"
	"encoding/json"
	"fmt"
	"mydiet/internal/logger"
	"mydiet/internal/models/input"
	"mydiet/internal/models/list"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Search struct {
	selected     map[int]struct{}
	spinnerIndex int
	textInput    input.Model
	listView     list.Model
	lastRequest  time.Time
}

func (m Search) View() string {
	inputBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.getColor()).
		Width(50).
		Render(m.textInput.View())

	mainBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36")).
		Width(55).
		Align(lipgloss.Center).
		Render(lipgloss.JoinVertical(lipgloss.Center, m.listView.View(), inputBox))

	mainBox += "\n" + m.textInput.GetHelp() + "\n"
	return mainBox
}

func (m Search) Init() tea.Cmd {
	return nil
}

func (m Search) Update(msg tea.Msg) (Search, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == tea.KeyCtrlC.String() {
			return m, tea.Quit
		}
		// m.textInput.Used = true
		switch msg.String() {

		// case tea.KeyEnter.String():
		// 	m.textInput.SetValue(m.textInput.CurrentSuggestion())
		// 	m.textInput.CursorEnd()
		// 	m.listView.Choices = m.textInput.MatchedSuggestions()
		// 	m.textInput.Blur()
		// 	m.listView.Focus()

		case tea.KeyEsc.String():
			// m.textInput.Blur()
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
		m.textInput.PromptStyle.GetBorder()

	case SuccessRequest:

		logger.Log.Infof("got the succes message: %d", len(msg.suggetions))
		newSug := append(
			m.textInput.AvailableSuggestions(),
			msg.suggetions...)
		m.textInput.SetSuggestions(newSug)

		m.listView.Choices = m.textInput.MatchedSuggestions()
		m.listView.Cursor = m.textInput.CurrentSuggestionIndex()
		return m, cmd
	case FailedRequest:
		logger.Log.Info("got the error message")
		logger.Log.Error(msg.err)
	case UpdateViewMessage:
		logger.Log.Info("got the message")
		return m, cmd
	}

	m.listView.Choices = m.textInput.MatchedSuggestions()
	m.listView.Cursor = m.textInput.CurrentSuggestionIndex()
	return m, cmd
}

func (m Search) GetSuggestions(query string) tea.Cmd {
	return func() tea.Msg {
		logger.Log.Info("in")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		url := "https://dummyjson.com/products/search?q="
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return FailedRequest{err}
		}
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return FailedRequest{err}
		}

		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return FailedRequest{
				fmt.Errorf("Failed Request with status: %d",
					response.StatusCode),
			}
		}
		var data DummyJson
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return FailedRequest{err}
		}
		logger.Log.Info("extracting titles")
		return SuccessRequest{extractTitles(data)}

	}
}

func extractTitles(data DummyJson) []string {
	titles := make([]string, 0, len(data.Products))
	for _, p := range data.Products {
		titles = append(titles, p.Title)
	}
	return titles
}

func NewCartView() Search {
	ti := input.New()
	ti.Placeholder = "search"
	ti.Focus()
	ti.ShowSuggestions = true
	ti.CharLimit = 64
	ti.Width = 20
	ti.SetSuggestions([]string{"Buy carrots", "Buy celery", "Buy kohlrabi", "4", "4"})
	return Search{
		textInput: ti,
		listView:  list.NewListView(ti.AvailableSuggestions()),
	}
}
