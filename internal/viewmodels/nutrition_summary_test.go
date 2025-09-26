package viewmodels

import (
	"mydiet/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNutritionSummaryViewModel(t *testing.T) {
	tests := []struct {
		name     string
		foods    store.Foods
		expected NutritionSummaryViewModel
	}{
		{
			name: "single food item",
			foods: store.Foods{
				{
					Calories: 52.0,
					Protein:  0.3,
					Fat:      0.2,
					Carbs:    14.0,
					Fiber:    2.4,
					Sugar:    10.0,
					Sodium:   1.0,
				},
			},
			expected: NutritionSummaryViewModel{
				TotalCalories: "52.0",
				TotalProtein:  "0.3",
				TotalFat:      "0.2",
				TotalCarbs:    "14.0",
				TotalFiber:    "2.4",
				TotalSugar:    "10.0",
				TotalSodium:   "1.0",
			},
		},
		{
			name: "multiple food items",
			foods: store.Foods{
				{
					Calories: 52.0,
					Protein:  0.3,
					Fat:      0.2,
					Carbs:    14.0,
					Fiber:    2.4,
					Sugar:    10.0,
					Sodium:   1.0,
				},
				{
					Calories: 89.0,
					Protein:  1.1,
					Fat:      0.3,
					Carbs:    23.0,
					Fiber:    2.6,
					Sugar:    12.0,
					Sodium:   1.0,
				},
			},
			expected: NutritionSummaryViewModel{
				TotalCalories: "141.0", // 52 + 89
				TotalProtein:  "1.4",   // 0.3 + 1.1
				TotalFat:      "0.5",   // 0.2 + 0.3
				TotalCarbs:    "37.0",  // 14 + 23
				TotalFiber:    "5.0",   // 2.4 + 2.6
				TotalSugar:    "22.0",  // 10 + 12
				TotalSodium:   "2.0",   // 1 + 1
			},
		},
		{
			name:  "empty food list",
			foods: store.Foods{},
			expected: NutritionSummaryViewModel{
				TotalCalories: "0.0",
				TotalProtein:  "0.0",
				TotalFat:      "0.0",
				TotalCarbs:    "0.0",
				TotalFiber:    "0.0",
				TotalSugar:    "0.0",
				TotalSodium:   "0.0",
			},
		},
		{
			name: "zero nutrition values",
			foods: store.Foods{
				{
					Calories: 0.0,
					Protein:  0.0,
					Fat:      0.0,
					Carbs:    0.0,
					Fiber:    0.0,
					Sugar:    0.0,
					Sodium:   0.0,
				},
				{
					Calories: 0.0,
					Protein:  0.0,
					Fat:      0.0,
					Carbs:    0.0,
					Fiber:    0.0,
					Sugar:    0.0,
					Sodium:   0.0,
				},
			},
			expected: NutritionSummaryViewModel{
				TotalCalories: "0.0",
				TotalProtein:  "0.0",
				TotalFat:      "0.0",
				TotalCarbs:    "0.0",
				TotalFiber:    "0.0",
				TotalSugar:    "0.0",
				TotalSodium:   "0.0",
			},
		},
		{
			name: "decimal precision",
			foods: store.Foods{
				{
					Calories: 52.123,
					Protein:  0.345,
					Fat:      0.234,
					Carbs:    14.567,
					Fiber:    2.456,
					Sugar:    10.789,
					Sodium:   1.123,
				},
			},
			expected: NutritionSummaryViewModel{
				TotalCalories: "52.1", // Rounded to 1 decimal
				TotalProtein:  "0.3", // Rounded to 1 decimal
				TotalFat:      "0.2", // Rounded to 1 decimal
				TotalCarbs:    "14.6", // Rounded to 1 decimal
				TotalFiber:    "2.5", // Rounded to 1 decimal
				TotalSugar:    "10.8", // Rounded to 1 decimal
				TotalSodium:   "1.1", // Rounded to 1 decimal
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNutritionSummaryViewModel(tt.foods)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNutritionSummaryViewModel_ToStringSlice(t *testing.T) {
	vm := NutritionSummaryViewModel{
		TotalCalories: "141.0",
		TotalProtein:  "1.4",
		TotalFat:      "0.5",
		TotalCarbs:    "37.0",
		TotalFiber:    "5.0",
		TotalSugar:    "22.0",
		TotalSodium:   "2.0",
	}

	expected := []string{
		"",       // Empty LogID column
		"TOTAL",  // Name column
		"",       // Empty Quantity column
		"",       // Empty Grams column
		"141.0",  // Total Calories
		"1.4",    // Total Protein
		"0.5",    // Total Fat
		"37.0",   // Total Carbs
		"5.0",    // Total Fiber
		"22.0",   // Total Sugar
		"2.0",    // Total Sodium
	}

	result := vm.ToStringSlice()
	assert.Equal(t, expected, result)
}

func TestNutritionSummaryViewModel_LargeNumbers(t *testing.T) {
	foods := store.Foods{
		{Calories: 999.99, Protein: 99.99, Fat: 99.99, Carbs: 999.99, Fiber: 99.99, Sugar: 999.99, Sodium: 9999.99},
		{Calories: 1000.01, Protein: 100.01, Fat: 100.01, Carbs: 1000.01, Fiber: 100.01, Sugar: 1000.01, Sodium: 10000.01},
	}

	result := NewNutritionSummaryViewModel(foods)

	// Check that large numbers are handled correctly
	assert.Equal(t, "2000.0", result.TotalCalories)
	assert.Equal(t, "200.0", result.TotalProtein)
	assert.Equal(t, "200.0", result.TotalFat)
	assert.Equal(t, "2000.0", result.TotalCarbs)
	assert.Equal(t, "200.0", result.TotalFiber)
	assert.Equal(t, "2000.0", result.TotalSugar)
	assert.Equal(t, "20000.0", result.TotalSodium)
}