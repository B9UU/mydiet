package main

import (
	"context"
	"errors"
	"fmt"
	"mydiet/internal/logger"
	"mydiet/internal/models/details"
	"mydiet/internal/models/form"
	searchbox "mydiet/internal/models/searchbox"
	"mydiet/internal/store"
	"mydiet/internal/types"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type model struct {
	activeView types.View
	Views      allViews
	store      store.Store
}

type allViews struct {
	Detail details.Model
	Search searchbox.Model
	Form   form.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

// what the application shows
func (m model) View() string {
	switch m.activeView {
	case types.SEARCHBOX:
		return m.Views.Search.View()
	case types.FORMVIEW:
		return m.Views.Form.View()
	default:
		return m.Views.Detail.View()
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case types.ViewMessage:
		m.activeView = msg.NewView
		switch msg.NewView {
		case types.DETAILSVIEW:
			if msg.Msg == "updated" {
				logger.Log.Info("Updated")

				err := m.store.FoodStore.InsertLog(m.Views.Form.FoodLog)
				if err != nil {
					return m, func() tea.Msg {
						return types.ErrMsg(err)
					}
				}
				m.Views.Detail.SyncRowsFor()
				// m.Views.Detail = details.New(m.store)
			}
		case types.FORMVIEW:
			food, ok := msg.Msg.(*store.Food)
			if !ok {
				return m, func() tea.Msg {
					return types.ErrMsg(errors.New("Invalid food "))
				}
			}
			var err error
			food.Units, err = m.store.FoodStore.GetUnits(food.ID)
			if err != nil {
				return m, func() tea.Msg {
					return types.ErrMsg(errors.New("Invalid food "))
				}
			}
			m.Views.Form = form.New(food)
			m.activeView = types.FORMVIEW

		case types.SEARCHBOX:
			m.Views.Search = searchbox.New(msg.Msg.(store.MealType), m.store)
		}
	}
	switch m.activeView {

	case types.SEARCHBOX:
		m.Views.Search, cmd = m.Views.Search.Update(msg)

	case types.FORMVIEW:
		m.Views.Form, cmd = m.Views.Form.Update(msg)
	default:
		m.Views.Detail, cmd = m.Views.Detail.Update(msg)
	}
	return m, cmd
}
func initialModel() *model {
	db, err := openDB()
	if err != nil {
		logger.Log.Fatal(err)
	}
	s := store.NewStore(db)
	m := &model{
		activeView: 1,
		store:      s,
		Views: allViews{
			Detail: details.New(s),
			// Form:   form.New(),
		},
	}
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
func openDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./nutrition.db")
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
