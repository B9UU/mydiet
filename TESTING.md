# Testing Guide

This document describes the testing strategy and setup for the MyDiet application.

## Overview

The project uses a comprehensive testing approach with:
- **Unit Tests**: Test individual components in isolation
- **Integration Tests**: Test component interactions with real database
- **Benchmarks**: Performance testing for critical paths

## Test Structure

```
test/
├── integration/       # Integration tests
├── unit/             # Additional unit test helpers (if needed)
└── testdata/         # Test fixtures and data files

internal/
├── services/
│   └── nutrition_test.go    # Service layer unit tests
├── store/
│   └── food_test.go         # Store layer unit tests
└── viewmodels/
    ├── food_log_test.go     # View model tests
    ├── food_search_test.go
    └── nutrition_summary_test.go
```

## Running Tests

### Using Make (Recommended)

```bash
# Run all tests
make test

# Run only unit tests
make test-unit

# Run only integration tests
make test-integration

# Run tests with coverage report
make test-coverage

# Run benchmark tests
make test-benchmark
```

### Using Go directly

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...

# Run specific test package
go test ./internal/services/

# Run specific test
go test -run TestNutritionService_LogFood ./internal/services/

# Run benchmarks
go test -bench=. ./internal/store/
```

## Test Categories

### Unit Tests

**Services (`internal/services/nutrition_test.go`)**
- Tests business logic in isolation using mocked dependencies
- Validates input validation, error handling, and business rules
- Covers success and failure scenarios

**Store (`internal/store/food_test.go`)**
- Tests data structures and utility functions
- No database dependencies
- Focuses on logic like `Foods.GetId()`

**View Models (`internal/viewmodels/`)**
- Tests data transformation and formatting
- Validates string formatting and number precision
- Tests edge cases like empty data and large numbers

### Integration Tests

**Database Integration (`test/integration/database_test.go`)**
- Tests the complete data flow from database to application
- Uses temporary SQLite databases for each test
- Validates SQL queries and database operations
- Tests date filtering and meal-specific queries

## Testing Best Practices

### 1. Test Structure
```go
func TestFeatureName(t *testing.T) {
    tests := []struct {
        name        string
        input       InputType
        setupMock   func(*MockType)
        expected    ExpectedType
        expectError bool
    }{
        // Test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### 2. Assertions
- Use `github.com/stretchr/testify/assert` for readable assertions
- Use `require` for assertions that should stop the test if failed
- Prefer specific assertions over generic ones

### 3. Mocking
- Use `github.com/stretchr/testify/mock` for interface mocking
- Mock external dependencies (database, file system, network)
- Keep mocks simple and focused

### 4. Test Data
- Use table-driven tests for multiple similar test cases
- Create helper functions for common test setup
- Use meaningful test names that describe the scenario

## Test Coverage

Target coverage levels:
- **Services**: >90% (critical business logic)
- **Store**: >80% (data access layer)
- **View Models**: >95% (simple data transformation)
- **Overall**: >85%

Generate coverage reports:
```bash
make test-coverage
```

Open the generated `coverage.html` file to see detailed coverage information.

## Continuous Integration

The test suite is designed to run in CI environments:

```bash
# Run all CI checks
make ci-test
```

This includes:
1. Dependency installation
2. All test suites
3. Code linting
4. Coverage reporting

## Writing New Tests

### For Services
1. Create mocks for dependencies
2. Test both success and failure scenarios
3. Validate input validation
4. Test edge cases

### For Store Layer
1. Focus on logic that doesn't require database
2. Test data structures and utility functions
3. Use simple, predictable test data

### For Integration Tests
1. Use temporary databases
2. Test complete workflows
3. Validate SQL query results
4. Clean up resources in defer statements

## Debugging Tests

### Failed Tests
```bash
# Run specific test with verbose output
go test -v -run TestSpecificFunction ./package/path/

# Run with race detector
go test -race ./...

# Run with additional logging
go test -v ./... 2>&1 | tee test.log
```

### Performance Issues
```bash
# Profile tests
go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=.

# Analyze profiles
go tool pprof cpu.prof
go tool pprof mem.prof
```

## Common Issues

### Database Tests
- Ensure temporary databases are cleaned up
- Use unique database names for parallel tests
- Don't rely on existing data

### Race Conditions
- Always run tests with `-race` flag
- Be careful with shared state
- Use proper synchronization

### Flaky Tests
- Avoid time-based assertions
- Use deterministic test data
- Clean up global state between tests

## Example Test

```go
func TestNutritionService_LogFood(t *testing.T) {
    // Arrange
    mockStore := &MockStore{FoodStore: &MockFoodStore{}}
    mockStore.FoodStore.On("InsertLog", mock.AnythingOfType("store.LoggingFood")).Return(nil)

    service := NewNutritionService(store.Store{FoodStore: mockStore.FoodStore})
    request := LogFoodRequest{
        UserID: 1,
        FoodID: 1,
        // ... other fields
    }

    // Act
    err := service.LogFood(request)

    // Assert
    assert.NoError(t, err)
    mockStore.FoodStore.AssertExpectations(t)
}
```