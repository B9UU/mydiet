package views

import (
	"context"
	"encoding/json"
	"fmt"
	"mydiet/internal/logger"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CartView struct {
	selected     map[int]struct{}
	cursor       int
	spinnerIndex int
	textInput    textinput.Model
	lastRequest  time.Time
}

func (m CartView) View() string {
	s := "What should we buy at the market?\n\n"

	var suggetions []string
	for i, choice := range m.textInput.MatchedSuggestions() {
		cursor := " "
		if m.textInput.CurrentSuggestionIndex() == i {
			cursor = ">"
		}
		suggetions = append(suggetions, fmt.Sprintf("%s %s", cursor, choice))
	}
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(lipgloss.Color("63"))).
		Width(50).
		Render(lipgloss.JoinVertical(lipgloss.Left, suggetions...))

	s += box
	// s += strings.Join(suggetions, "")
	s += "\n\n" + m.textInput.View() + "\n"
	s += "\nPress q to quit.\n"
	return s
}

func (m CartView) Init() tea.Cmd {

	return nil
}

func (m CartView) Update(msg tea.Msg) (CartView, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			return m, func() tea.Msg {
				return UpdateViewMessage(1)
			}
		case "enter":
			m.textInput.SetValue(m.textInput.CurrentSuggestion())
			m.textInput.CursorEnd()
			m.textInput.SetSuggestions(nil)

		}
		// update the input field
		m.textInput, cmd = m.textInput.Update(msg)

		logger.Log.Info(m.lastRequest)
		if len(m.textInput.Value()) >= 3 {
			if time.Since(m.lastRequest) > 30*time.Second {
				logger.Log.Info("time has past")
				cmd = tea.Batch(cmd,
					m.GetSuggestions(
						m.textInput.Value(),
					))
				m.lastRequest = time.Now()
			} else {
				logger.Log.Info("wait 30 sec")
			}
		}
	case SuccessRequest:
		logger.Log.Info("got the succes message")
		m.textInput.SetSuggestions(msg.suggetions)
	case FailedRequest:
		logger.Log.Info("got the error message")
		logger.Log.Error(msg.err)
	}

	return m, cmd
}

func (m CartView) GetSuggestions(query string) tea.Cmd {
	return func() tea.Msg {

		logger.Log.Info("preparing the requests")
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

		logger.Log.Info("request sent")
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

func NewCartView() CartView {
	ti := textinput.New()
	ti.Placeholder = "search"
	ti.Focus()
	ti.ShowSuggestions = true
	ti.CharLimit = 64
	ti.Width = 20
	ti.SetSuggestions([]string{"Buy carrots", "Buy celery", "Buy kohlrabi"})
	return CartView{
		textInput: ti,
		selected:  make(map[int]struct{}),
	}

}
