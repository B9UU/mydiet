package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoods_GetId(t *testing.T) {
	foods := Foods{
		{ID: 1, Name: "Apple"},
		{ID: 2, Name: "Banana"},
		{ID: 3, Name: "Orange"},
	}

	tests := []struct {
		name     string
		id       int
		expected *Food
	}{
		{
			name:     "existing food",
			id:       2,
			expected: &Food{ID: 2, Name: "Banana"},
		},
		{
			name:     "first food",
			id:       1,
			expected: &Food{ID: 1, Name: "Apple"},
		},
		{
			name:     "last food",
			id:       3,
			expected: &Food{ID: 3, Name: "Orange"},
		},
		{
			name:     "non-existing food",
			id:       999,
			expected: nil,
		},
		{
			name:     "zero id",
			id:       0,
			expected: nil,
		},
		{
			name:     "negative id",
			id:       -1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := foods.GetId(tt.id)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFoods_GetId_EmptySlice(t *testing.T) {
	foods := Foods{}
	result := foods.GetId(1)
	assert.Nil(t, result)
}

func TestFoods_GetId_SingleItem(t *testing.T) {
	foods := Foods{
		{ID: 42, Name: "Single Item"},
	}

	// Found
	result := foods.GetId(42)
	expected := &Food{ID: 42, Name: "Single Item"}
	assert.Equal(t, expected, result)

	// Not found
	result = foods.GetId(43)
	assert.Nil(t, result)
}

func TestMealType_Constants(t *testing.T) {
	// Test that meal type constants have expected values
	assert.Equal(t, MealType("Breakfast"), Breakfast)
	assert.Equal(t, MealType("Lunch"), Lunch)
	assert.Equal(t, MealType("Dinner"), Dinner)
	assert.Equal(t, MealType("Snack"), Snack)
}

func TestFoodUnits_Structure(t *testing.T) {
	// Test that FoodUnits struct can be created with expected fields
	unit := FoodUnits{
		ID:          1,
		FoodID:      2,
		Unit:        "piece",
		SizeInGrams: 182.0,
	}

	assert.Equal(t, 1, unit.ID)
	assert.Equal(t, 2, unit.FoodID)
	assert.Equal(t, "piece", unit.Unit)
	assert.Equal(t, 182.0, unit.SizeInGrams)
}

func TestLoggingFood_Structure(t *testing.T) {
	// Test that LoggingFood struct can be created with expected fields
	log := LoggingFood{
		UserId:     1,
		FoodId:     2,
		FoodUnitId: 3,
		QTY:        2.5,
		Meal:       Breakfast,
	}

	assert.Equal(t, 1, log.UserId)
	assert.Equal(t, 2, log.FoodId)
	assert.Equal(t, 3, log.FoodUnitId)
	assert.Equal(t, 2.5, log.QTY)
	assert.Equal(t, Breakfast, log.Meal)
}

func TestFood_Structure(t *testing.T) {
	// Test that Food struct can be created with all expected fields
	food := Food{
		LogID:    1,
		ID:       2,
		Name:     "Apple",
		Meal:     Breakfast,
		QTY:      1.0,
		Unit:     "piece",
		Grams:    182.0,
		Calories: 52.0,
		Protein:  0.3,
		Fat:      0.2,
		Carbs:    14.0,
		Fiber:    2.4,
		Sugar:    10.0,
		Sodium:   1.0,
		Units: []FoodUnits{
			{ID: 1, Unit: "piece", SizeInGrams: 182.0},
			{ID: 2, Unit: "gram", SizeInGrams: 1.0},
		},
	}

	assert.Equal(t, 1, food.LogID)
	assert.Equal(t, 2, food.ID)
	assert.Equal(t, "Apple", food.Name)
	assert.Equal(t, Breakfast, food.Meal)
	assert.Equal(t, 1.0, food.QTY)
	assert.Equal(t, "piece", food.Unit)
	assert.Equal(t, 182.0, food.Grams)
	assert.Equal(t, 52.0, food.Calories)
	assert.Equal(t, 0.3, food.Protein)
	assert.Equal(t, 0.2, food.Fat)
	assert.Equal(t, 14.0, food.Carbs)
	assert.Equal(t, 2.4, food.Fiber)
	assert.Equal(t, 10.0, food.Sugar)
	assert.Equal(t, 1.0, food.Sodium)
	assert.Len(t, food.Units, 2)
	assert.Equal(t, "piece", food.Units[0].Unit)
	assert.Equal(t, "gram", food.Units[1].Unit)
}

// Benchmark test for GetId with large dataset
func BenchmarkFoods_GetId(b *testing.B) {
	// Create a large dataset
	foods := make(Foods, 1000)
	for i := 0; i < 1000; i++ {
		foods[i] = Food{ID: i + 1, Name: "Food " + string(rune(i))}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Search for item in the middle
		foods.GetId(500)
	}
}

func BenchmarkFoods_GetId_NotFound(b *testing.B) {
	// Create a dataset
	foods := make(Foods, 100)
	for i := 0; i < 100; i++ {
		foods[i] = Food{ID: i + 1, Name: "Food " + string(rune(i))}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Search for non-existing item (worst case)
		foods.GetId(9999)
	}
}