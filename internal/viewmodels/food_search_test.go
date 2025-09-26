package viewmodels

import (
	"mydiet/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFoodSearchViewModel(t *testing.T) {
	tests := []struct {
		name     string
		food     store.Food
		expected FoodSearchViewModel
	}{
		{
			name: "basic food search result",
			food: store.Food{
				ID:       1,
				Name:     "Apple",
				Calories: 52.123,
			},
			expected: FoodSearchViewModel{
				ID:       1,
				Name:     "Apple",
				Calories: "52.12", // Formatted to 2 decimals
			},
		},
		{
			name: "zero calories",
			food: store.Food{
				ID:       2,
				Name:     "Water",
				Calories: 0.0,
			},
			expected: FoodSearchViewModel{
				ID:       2,
				Name:     "Water",
				Calories: "0.00",
			},
		},
		{
			name: "high calorie food",
			food: store.Food{
				ID:       3,
				Name:     "Peanut Butter",
				Calories: 588.7890,
			},
			expected: FoodSearchViewModel{
				ID:       3,
				Name:     "Peanut Butter",
				Calories: "588.79",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFoodSearchViewModel(tt.food)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewFoodSearchViewModels(t *testing.T) {
	foods := store.Foods{
		{ID: 1, Name: "Apple", Calories: 52.0},
		{ID: 2, Name: "Banana", Calories: 89.0},
		{ID: 3, Name: "Orange", Calories: 47.0},
	}

	result := NewFoodSearchViewModels(foods)

	assert.Len(t, result, 3)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Apple", result[0].Name)
	assert.Equal(t, "52.00", result[0].Calories)

	assert.Equal(t, 2, result[1].ID)
	assert.Equal(t, "Banana", result[1].Name)
	assert.Equal(t, "89.00", result[1].Calories)

	assert.Equal(t, 3, result[2].ID)
	assert.Equal(t, "Orange", result[2].Name)
	assert.Equal(t, "47.00", result[2].Calories)
}

func TestFoodSearchViewModel_ToStringSlice(t *testing.T) {
	vm := FoodSearchViewModel{
		ID:       1,
		Name:     "Apple",
		Calories: "52.00",
	}

	expected := []string{
		"1",     // ID
		"Apple", // Name
		"52.00", // Calories
	}

	result := vm.ToStringSlice()
	assert.Equal(t, expected, result)
}

func TestEmptyFoodSearchViewModels(t *testing.T) {
	foods := store.Foods{}
	result := NewFoodSearchViewModels(foods)

	assert.Empty(t, result)
	assert.Len(t, result, 0)
}