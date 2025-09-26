# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Build and Run
- **Build**: `go build -o mydiet ./cmd/main.go`
- **Run**: `./mydiet` (after building)
- **Quick run**: `go run ./cmd/main.go`

### Development
- **Test**: No test suite currently exists in the codebase
- **Database**: SQLite database file `nutrition.db` is created automatically on first run

## Architecture

### High-level Structure
MyDiet is a terminal-based nutrition tracking application built with Go using clean architecture principles and the Bubble Tea TUI framework. The application is organized into distinct layers with clear separation of concerns.

### Core Architecture Layers

**Application Layer** (`internal/app/`):
- `config.go`: Centralized configuration management with environment variable support
- `coordinator.go`: ViewCoordinator manages view transitions and coordinates between UI and business logic
- `errors.go`: Structured error types (ValidationError, DatabaseError, NotFoundError, InternalError)

**Service Layer** (`internal/services/`):
- `nutrition.go`: Core business logic for nutrition tracking (LogFood, SearchFoods, GetMealLogs, etc.)
- `types.go`: Service contracts and request/response types
- Handles all business rules, validation, and domain logic
- Clean APIs independent of UI framework

**Data Layer** (`internal/store/`):
- `Store` struct provides centralized data access with `FoodStore` for database operations
- Uses jmoiron/sqlx for database operations with SQLite backend
- Database schema managed through migration files in `migrations/`

**UI Layer** (`internal/models/`):
- `details/`: Main daily food log view (4-panel meal display)
- `searchbox/`: Food search interface with autocomplete
- `form/`: Food entry form using Huh form library
- Other directories contain reusable UI components (date, input, list, table, textinput)

**Main Application** (`cmd/main.go`):
- Clean application entry point with proper dependency injection
- Application struct manages lifecycle and resource cleanup
- AppModel handles Bubble Tea integration with minimal logic
- ViewCoordinator handles all view transitions

### Key Dependencies
- **Bubble Tea**: Core TUI framework for terminal interfaces
- **Huh**: Form library for interactive input forms
- **Lipgloss**: Styling and layout for terminal UI
- **SQLx**: Database operations with SQLite
- **Clipboard**: Cross-platform clipboard access

### Data Flow
1. User interactions trigger Bubble Tea messages
2. AppModel routes messages to ViewCoordinator
3. ViewCoordinator delegates business logic to NutritionService
4. NutritionService validates and processes requests via Store
5. Results flow back through layers to update UI

### Configuration
The application supports environment-based configuration:
- `DB_PATH`: Database file path (default: "./nutrition.db")
- `LOG_LEVEL`: Logging level (default: "info")
- `LOG_FILE`: Log file path (default: "debug.log")
- `LOG_CONSOLE`: Console logging (default: true)

### Database Schema
The application uses SQLite with migrations in `migrations/`:
- `users` table: User account information
- `foods` table: Food items with nutritional data
- `food_units` table: Different measurement units for foods
- `food_logs` table: User meal entries

### Error Handling
Structured error handling with specific error types:
- **ValidationError**: Input validation failures
- **DatabaseError**: Database operation failures
- **NotFoundError**: Resource not found
- **InternalError**: Unexpected application errors

### Development Notes
- **Architecture**: Clean architecture with dependency inversion
- **Testing**: No test suite exists yet (high priority improvement needed)
- **Logging**: Uses custom logger, consider migrating to structured logging
- **User Management**: Currently hardcoded to user ID 1
- **Original Code**: Preserved in `cmd/main_original.go` for reference