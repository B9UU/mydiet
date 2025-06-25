package main

import (
	"fmt"
	"mydiet/internal/logger"
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
}

func (m *model) Init() tea.Cmd {
	switch m.activeView {
	case SPINNERVIEW:
		return m.Views.Spinner.Init()
	case CARTVIEW:
		return m.Views.Cart.Init()
	default:
		return nil
	}
}

// what the application shows
func (m *model) View() string {
	// The header
	switch m.activeView {
	case SPINNERVIEW:
		return m.Views.Spinner.View()
	default:
		return m.Views.Cart.View()
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case views.UpdateViewMessage:
		um := int(msg)
		m.activeView = um
		return m, m.Init()
	}
	switch m.activeView {
	case SPINNERVIEW:
		m.Views.Spinner, cmd = m.Views.Spinner.Update(msg)
		return m, cmd
	case CARTVIEW:
		m.Views.Cart, cmd = m.Views.Cart.Update(msg)

	}
	return m, cmd

}
func initialModel() *model {
	m := &model{
		activeView: 1,
		Views: allViews{
			Cart:    views.NewCartView(),
			Spinner: views.NewSpinnerView(),
		},
	}
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
