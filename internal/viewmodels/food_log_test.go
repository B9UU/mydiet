package viewmodels

import (
	"mydiet/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFoodLogViewModel(t *testing.T) {
	tests := []struct {
		name     string
		food     store.Food
		expected FoodLogViewModel
	}{
		{
			name: "complete food entry",
			food: store.Food{
				LogID:    1,
				Name:     "Apple",
				QTY:      2.0,
				Unit:     "piece",
				Grams:    364.0,
				Calories: 104.0,
				Protein:  0.6,
				Fat:      0.4,
				Carbs:    28.0,
				Fiber:    4.8,
				Sugar:    20.0,
				Sodium:   2.0,
			},
			expected: FoodLogViewModel{
				LogID:    1,
				Name:     "Apple",
				Quantity: "2.0 piece",
				Grams:    "364.0",
				Calories: "104.0",
				Protein:  "0.6",
				Fat:      "0.4",
				Carbs:    "28.0",
				Fiber:    "4.8",
				Sugar:    "20.0",
				Sodium:   "2.0",
			},
		},
		{
			name: "zero values",
			food: store.Food{
				LogID:    0,
				Name:     "Water",
				QTY:      0.0,
				Unit:     "ml",
				Grams:    0.0,
				Calories: 0.0,
				Protein:  0.0,
				Fat:      0.0,
				Carbs:    0.0,
				Fiber:    0.0,
				Sugar:    0.0,
				Sodium:   0.0,
			},
			expected: FoodLogViewModel{
				LogID:    0,
				Name:     "Water",
				Quantity: "0.0 ml",
				Grams:    "0.0",
				Calories: "0.0",
				Protein:  "0.0",
				Fat:      "0.0",
				Carbs:    "0.0",
				Fiber:    "0.0",
				Sugar:    "0.0",
				Sodium:   "0.0",
			},
		},
		{
			name: "decimal precision",
			food: store.Food{
				LogID:    2,
				Name:     "Banana",
				QTY:      1.5,
				Unit:     "piece",
				Grams:    177.0,
				Calories: 133.5,
				Protein:  1.65,
				Fat:      0.45,
				Carbs:    34.5,
				Fiber:    3.9,
				Sugar:    18.0,
				Sodium:   1.5,
			},
			expected: FoodLogViewModel{
				LogID:    2,
				Name:     "Banana",
				Quantity: "1.5 piece",
				Grams:    "177.0",
				Calories: "133.5",
				Protein:  "1.6", // Rounded to 1 decimal
				Fat:      "0.5", // Rounded to 1 decimal
				Carbs:    "34.5",
				Fiber:    "3.9",
				Sugar:    "18.0",
				Sodium:   "1.5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFoodLogViewModel(tt.food)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewFoodLogViewModels(t *testing.T) {
	foods := store.Foods{
		{LogID: 1, Name: "Apple", QTY: 1.0, Unit: "piece", Calories: 52.0},
		{LogID: 2, Name: "Banana", QTY: 1.0, Unit: "piece", Calories: 89.0},
	}

	result := NewFoodLogViewModels(foods)

	assert.Len(t, result, 2)
	assert.Equal(t, "Apple", result[0].Name)
	assert.Equal(t, "Banana", result[1].Name)
	assert.Equal(t, "1.0 piece", result[0].Quantity)
	assert.Equal(t, "1.0 piece", result[1].Quantity)
}

func TestFoodLogViewModel_ToStringSlice(t *testing.T) {
	vm := FoodLogViewModel{
		LogID:    1,
		Name:     "Apple",
		Quantity: "1.0 piece",
		Grams:    "182.0",
		Calories: "52.0",
		Protein:  "0.3",
		Fat:      "0.2",
		Carbs:    "14.0",
		Fiber:    "2.4",
		Sugar:    "10.0",
		Sodium:   "1.0",
	}

	expected := []string{
		"1",      // LogID
		"Apple",  // Name
		"1.0 piece", // Quantity
		"182.0",  // Grams
		"52.0",   // Calories
		"0.3",    // Protein
		"0.2",    // Fat
		"14.0",   // Carbs
		"2.4",    // Fiber
		"10.0",   // Sugar
		"1.0",    // Sodium
	}

	result := vm.ToStringSlice()
	assert.Equal(t, expected, result)
}

func TestFormatNumber(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		decimals int
		expected string
	}{
		{"zero with 1 decimal", 0.0, 1, "0.0"},
		{"integer with 1 decimal", 5.0, 1, "5.0"},
		{"decimal with 1 decimal", 5.25, 1, "5.2"}, // Rounded down
		{"decimal with 2 decimals", 5.256, 2, "5.26"}, // Rounded down
		{"small number", 0.1234, 2, "0.12"},
		{"large number", 1234.5678, 1, "1234.6"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatNumber(tt.value, tt.decimals)
			assert.Equal(t, tt.expected, result)
		})
	}
}