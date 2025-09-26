package main

import (
	"context"
	"fmt"
	"mydiet/internal/app"
	"mydiet/internal/logger"
	"mydiet/internal/services"
	"mydiet/internal/store"
	"mydiet/internal/types"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Application is the main application struct
type Application struct {
	config      *app.Config
	coordinator *app.ViewCoordinator
	service     *services.NutritionService
}

// NewApplication creates a new application instance
func NewApplication() (*Application, error) {
	// Load configuration
	config := app.LoadConfig()

	// Initialize logger
	logger.Log = logger.NewLogger()

	// Initialize database
	db, err := initializeDatabase(config.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize store and service
	store := store.NewStore(db)
	service := services.NewNutritionService(store)

	// Initialize view coordinator
	coordinator := app.NewViewCoordinator(service, store)

	return &Application{
		config:      config,
		coordinator: coordinator,
		service:     service,
	}, nil
}

// Run starts the application
func (a *Application) Run() error {
	// Initialize Bubble Tea logging
	f, err := tea.LogToFile(a.config.Logging.FilePath, "mydiet")
	if err != nil {
		return fmt.Errorf("couldn't open log file: %w", err)
	}
	defer f.Close()

	// Create and run the Bubble Tea program
	model := &AppModel{
		coordinator: a.coordinator,
		config:      a.config,
	}

	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		return fmt.Errorf("application error: %w", err)
	}

	return nil
}

// Close cleans up application resources
func (a *Application) Close() error {
	if logger.LogFile != nil {
		return logger.LogFile.Close()
	}
	return nil
}

// AppModel is the main Bubble Tea model
type AppModel struct {
	coordinator *app.ViewCoordinator
	config      *app.Config
	err         error
}

func (m *AppModel) Init() tea.Cmd {
	return nil
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Global key handlers
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case types.ViewMessage:
		// Handle view transitions through coordinator
		m.coordinator, cmd = m.coordinator.HandleViewMessage(msg)
		return m, cmd

	case types.ErrMsg:
		// Handle application errors
		m.err = error(msg)
		logger.Log.Error("Application error", "error", m.err)
		// Could show error in UI here
		return m, nil
	}

	// Delegate to the active view
	views := m.coordinator.GetViews()
	switch m.coordinator.GetActiveView() {
	case types.SEARCHBOX:
		views.Searchbox, cmd = views.Searchbox.Update(msg)
	case types.FORMVIEW:
		views.Form, cmd = views.Form.Update(msg)
	default:
		views.Details, cmd = views.Details.Update(msg)
	}

	return m, cmd
}

func (m *AppModel) View() string {
	// Show error if there's one
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress Ctrl+C to quit", m.err)
	}

	// Delegate to active view
	views := m.coordinator.GetViews()
	switch m.coordinator.GetActiveView() {
	case types.SEARCHBOX:
		return views.Searchbox.View()
	case types.FORMVIEW:
		return views.Form.View()
	default:
		return views.Details.View()
	}
}

// initializeDatabase sets up and connects to the database
func initializeDatabase(config app.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", config.Path)
	if err != nil {
		return nil, err
	}

	// Test connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectionTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	app, err := NewApplication()
	if err != nil {
		fmt.Printf("Failed to initialize application: %v\n", err)
		os.Exit(1)
	}
	defer app.Close()

	if err := app.Run(); err != nil {
		fmt.Printf("Application failed: %v\n", err)
		os.Exit(1)
	}
}
