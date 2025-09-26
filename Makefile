.PHONY: build test test-unit test-integration test-coverage clean help run lint fmt

# Build variables
BINARY_NAME=mydiet
BUILD_DIR=./build
CMD_DIR=./cmd

# Go variables
GO_FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)/main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

run: build ## Build and run the application
	@echo "Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

test: test-unit test-integration ## Run all tests

test-unit: ## Run unit tests only
	@echo "Running unit tests..."
	@go test -v -race ./internal/...

test-integration: ## Run integration tests only
	@echo "Running integration tests..."
	@go test -v -race ./test/integration/...

test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-benchmark: ## Run benchmark tests
	@echo "Running benchmark tests..."
	@go test -bench=. -benchmem ./internal/...

lint: ## Run linter (requires golangci-lint)
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	@golangci-lint run

fmt: ## Format Go code
	@echo "Formatting code..."
	@gofmt -s -w $(GO_FILES)
	@go mod tidy

clean: ## Clean build artifacts and test files
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@rm -f test_*.db
	@rm -f debug.log
	@go clean

deps: ## Download and tidy dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

install: build ## Install the application to GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	@go install $(CMD_DIR)/main.go

dev: ## Development mode - build and run with hot reload (requires air)
	@echo "Starting development mode..."
	@which air > /dev/null || (echo "air not installed. Run: go install github.com/cosmtrek/air@latest" && exit 1)
	@air

# CI/CD targets
ci-test: deps test lint ## Run CI tests (dependencies, tests, linting)
	@echo "CI tests completed successfully"

release: clean fmt test build ## Prepare a release build
	@echo "Release build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# Database targets
db-reset: ## Reset the database (removes nutrition.db)
	@echo "Resetting database..."
	@rm -f nutrition.db
	@echo "Database reset complete"

# Development tools
tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/cosmtrek/air@latest
	@echo "Development tools installed"