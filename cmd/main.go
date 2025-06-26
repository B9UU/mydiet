package main

import (
	"fmt"
	"mydiet/internal/logger"
	"mydiet/internal/models/textinput"
	"mydiet/internal/views"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	_ = iota
	SPINNERVIEW
	CARTVIEW
	DETAILSVIEW
)

type model struct {
	activeView int
	Views      allViews
}

type allViews struct {
	Cart    views.Search
	Spinner views.SpinnerView
	Detail  views.Details
}

func (m model) Init() tea.Cmd {
	return nil
}

// what the application shows
func (m model) View() string {
	// The header

	return m.Views.Detail.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Views.Detail, cmd = m.Views.Detail.Update(msg)
	return m, cmd
}
func initialModel() *model {
	m := &model{
		activeView: 1,
		Views: allViews{
			Cart:    views.NewCartView(),
			Spinner: views.NewSpinnerView(),
			Detail:  views.New(),
		},
	}
	textinput.New()
	m.Views.Spinner.ResetSpinner()
	return m
}
func main() {
	logger.Log = logger.NewLogger()
	defer logger.LogFile.Close()
	f, err := tea.LogToFile("debug.log", "help")
	if err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	}
	defer f.Close() // nolint:errcheck
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
