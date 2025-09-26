package services

import (
	"errors"
	"mydiet/internal/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFoodStore struct {
	mock.Mock
}

func (m *MockFoodStore) InsertLog(fl store.LoggingFood) error {
	args := m.Called(fl)
	return args.Error(0)
}

func (m *MockFoodStore) GetLogs(meal store.MealType, date time.Time) (store.Foods, error) {
	args := m.Called(meal, date)
	return args.Get(0).(store.Foods), args.Error(1)
}

func (m *MockFoodStore) Search(name string) (store.Foods, error) {
	args := m.Called(name)
	return args.Get(0).(store.Foods), args.Error(1)
}

func (m *MockFoodStore) GetAll(name string) (store.Foods, error) {
	args := m.Called(name)
	return args.Get(0).(store.Foods), args.Error(1)
}

func (m *MockFoodStore) GetUnits(id int) ([]store.FoodUnits, error) {
	args := m.Called(id)
	return args.Get(0).([]store.FoodUnits), args.Error(1)
}

func (m *MockFoodStore) GetByID(id int) (*store.Food, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*store.Food), args.Error(1)
}

var _ store.FoodStoreInterface = &MockFoodStore{}

func TestLogFood(t *testing.T) {
	// New test to specifically verify date logging
	t.Run("verify date logging", func(t *testing.T) {
		mockStore := store.Store{
			FoodStore: &MockFoodStore{},
		}
		specificDate := time.Date(2025, 9, 26, 0, 0, 0, 0, time.Local)

		// Setup mock to capture the logged date
		var capturedDate time.Time
		mockStore.FoodStore.(*MockFoodStore).On("InsertLog", mock.MatchedBy(func(logEntry store.LoggingFood) bool {
			capturedDate = logEntry.Timestamp
			return true
		})).Return(nil)

		service := NewNutritionService(mockStore)
		req := LogFoodRequest{
			UserID:     1,
			FoodID:     1,
			FoodUnitID: 1,
			Quantity:   2.0,
			Meal:       store.Breakfast,
		}

		err := service.LogFood(req, specificDate)
		assert.NoError(t, err)

		// Verify the date was logged correctly
		assert.Equal(t, specificDate.Year(), capturedDate.Year(), "Year should match")
		assert.Equal(t, specificDate.Month(), capturedDate.Month(), "Month should match")
		assert.Equal(t, specificDate.Day(), capturedDate.Day(), "Day should match")
	})
	tests := []struct {
		name        string
		request     LogFoodRequest
		setupMock   func(*MockFoodStore)
		expectError bool
		errorMsg    string
	}{
		{
			name: "successful food logging",
			request: LogFoodRequest{
				UserID:     1,
				FoodID:     1,
				FoodUnitID: 1,
				Quantity:   2.0,
				Meal:       store.Breakfast,
			},
			setupMock: func(mfs *MockFoodStore) {
				mfs.On("InsertLog", mock.AnythingOfType("store.LoggingFood")).Return(nil)
			},
			expectError: false,
		},
		{
			name: "invalid user ID",
			request: LogFoodRequest{
				UserID:     0,
				FoodID:     1,
				FoodUnitID: 1,
				Quantity:   2.0,
				Meal:       store.Breakfast,
			},
			setupMock:   func(*MockFoodStore) {},
			expectError: true,
			errorMsg:    "invalid user ID",
		},
		{
			name: "invalid food ID",
			request: LogFoodRequest{
				UserID:     1,
				FoodID:     0,
				FoodUnitID: 1,
				Quantity:   2.0,
				Meal:       store.Breakfast,
			},
			setupMock:   func(*MockFoodStore) {},
			expectError: true,
			errorMsg:    "invalid food ID",
		},
		{
			name: "negative quantity",
			request: LogFoodRequest{
				UserID:     1,
				FoodID:     1,
				FoodUnitID: 1,
				Quantity:   -1.0,
				Meal:       store.Breakfast,
			},
			setupMock:   func(*MockFoodStore) {},
			expectError: true,
			errorMsg:    "quantity must be positive",
		},
		{
			name: "empty meal type",
			request: LogFoodRequest{
				UserID:     1,
				FoodID:     1,
				FoodUnitID: 1,
				Quantity:   2.0,
				Meal:       "",
			},
			setupMock:   func(*MockFoodStore) {},
			expectError: true,
			errorMsg:    "meal type is required",
		},
		{
			name: "database error",
			request: LogFoodRequest{
				UserID:     1,
				FoodID:     1,
				FoodUnitID: 1,
				Quantity:   2.0,
				Meal:       store.Breakfast,
			},
			setupMock: func(mfs *MockFoodStore) {
				mfs.On("InsertLog", mock.AnythingOfType("store.LoggingFood")).Return(errors.New("db error"))
			},
			expectError: true,
			errorMsg:    "db error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := store.Store{
				FoodStore: &MockFoodStore{},
			}
			tt.setupMock(mockStore.FoodStore.(*MockFoodStore))

			service := NewNutritionService(mockStore)
			err := service.LogFood(tt.request, time.Now())

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}

			mockStore.FoodStore.(*MockFoodStore).AssertExpectations(t)
		})
	}
}

func TestGetMealLogs(t *testing.T) {
	tests := []struct {
		name        string
		userID      int
		meal        store.MealType
		date        time.Time
		setupMock   func(*MockFoodStore)
		expected    store.Foods
		expectError bool
	}{
		{
			name:   "successful retrieval",
			userID: 1,
			meal:   store.Breakfast,
			date:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			setupMock: func(mfs *MockFoodStore) {
				expectedFoods := store.Foods{
					{ID: 1, Name: "Apple", Calories: 52},
				}
				mfs.On("GetLogs", store.Breakfast, mock.AnythingOfType("time.Time")).Return(expectedFoods, nil)
			},
			expected: store.Foods{
				{ID: 1, Name: "Apple", Calories: 52},
			},
			expectError: false,
		},
		{
			name:        "invalid user ID",
			userID:      0,
			meal:        store.Breakfast,
			date:        time.Now(),
			setupMock:   func(*MockFoodStore) {},
			expectError: true,
		},
		{
			name:   "database error",
			userID: 1,
			meal:   store.Breakfast,
			date:   time.Now(),
			setupMock: func(mfs *MockFoodStore) {
				mfs.On("GetLogs", store.Breakfast, mock.AnythingOfType("time.Time")).Return(store.Foods{}, errors.New("db error"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := store.Store{
				FoodStore: &MockFoodStore{},
			}
			tt.setupMock(mockStore.FoodStore.(*MockFoodStore))

			service := NewNutritionService(mockStore)
			result, err := service.GetMealLogs(tt.userID, tt.meal, tt.date)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			mockStore.FoodStore.(*MockFoodStore).AssertExpectations(t)
		})
	}
}

func TestSearchFoods(t *testing.T) {
	tests := []struct {
		name      string
		query     string
		setupMock func(*MockFoodStore)
		expected  store.Foods
	}{
		{
			name:  "search with query",
			query: "apple",
			setupMock: func(mfs *MockFoodStore) {
				expectedFoods := store.Foods{
					{ID: 1, Name: "Apple", Calories: 52},
				}
				mfs.On("Search", "apple").Return(expectedFoods, nil)
			},
			expected: store.Foods{
				{ID: 1, Name: "Apple", Calories: 52},
			},
		},
		{
			name:  "empty query returns all",
			query: "",
			setupMock: func(mfs *MockFoodStore) {
				expectedFoods := store.Foods{
					{ID: 1, Name: "Apple", Calories: 52},
					{ID: 2, Name: "Banana", Calories: 89},
				}
				mfs.On("GetAll", "").Return(expectedFoods, nil)
			},
			expected: store.Foods{
				{ID: 1, Name: "Apple", Calories: 52},
				{ID: 2, Name: "Banana", Calories: 89},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := store.Store{
				FoodStore: &MockFoodStore{},
			}
			tt.setupMock(mockStore.FoodStore.(*MockFoodStore))

			service := NewNutritionService(mockStore)
			result, err := service.SearchFoods(tt.query)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)

			mockStore.FoodStore.(*MockFoodStore).AssertExpectations(t)
		})
	}
}

func TestGetFoodWithUnits(t *testing.T) {
	tests := []struct {
		name        string
		foodID      int
		setupMock   func(*MockFoodStore)
		expected    *store.Food
		expectError bool
	}{
		{
			name:   "successful retrieval",
			foodID: 1,
			setupMock: func(mfs *MockFoodStore) {
				food := &store.Food{ID: 1, Name: "Apple", Calories: 52}
				units := []store.FoodUnits{
					{ID: 1, Unit: "piece", SizeInGrams: 182},
				}
				mfs.On("GetByID", 1).Return(food, nil)
				mfs.On("GetUnits", 1).Return(units, nil)
			},
			expected: &store.Food{
				ID: 1, Name: "Apple", Calories: 52,
				Units: []store.FoodUnits{{ID: 1, Unit: "piece", SizeInGrams: 182}},
			},
			expectError: false,
		},
		{
			name:        "invalid food ID",
			foodID:      0,
			setupMock:   func(*MockFoodStore) {},
			expectError: true,
		},
		{
			name:   "food not found",
			foodID: 999,
			setupMock: func(mfs *MockFoodStore) {
				mfs.On("GetByID", 999).Return(nil, errors.New("food not found"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := store.Store{
				FoodStore: &MockFoodStore{},
			}
			tt.setupMock(mockStore.FoodStore.(*MockFoodStore))

			service := NewNutritionService(mockStore)
			result, err := service.GetFoodWithUnits(tt.foodID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			mockStore.FoodStore.(*MockFoodStore).AssertExpectations(t)
		})
	}
}

