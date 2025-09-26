package viewmodels

import (
	"fmt"
	"mydiet/internal/store"
)

// FoodLogViewModel handles display formatting for food log entries
type FoodLogViewModel struct {
	LogID    int
	Name     string
	Quantity string
	Grams    string
	Calories string
	Protein  string
	Fat      string
	Carbs    string
	Fiber    string
	Sugar    string
	Sodium   string
}

// NewFoodLogViewModel creates a view model from a domain Food object
func NewFoodLogViewModel(food store.Food) FoodLogViewModel {
	return FoodLogViewModel{
		LogID:    food.LogID,
		Name:     food.Name,
		Quantity: fmt.Sprintf("%.1f %s", food.QTY, food.Unit),
		Grams:    formatNumber(food.Grams, 1),
		Calories: formatNumber(food.Calories, 1),
		Protein:  formatNumber(food.Protein, 1),
		Fat:      formatNumber(food.Fat, 1),
		Carbs:    formatNumber(food.Carbs, 1),
		Fiber:    formatNumber(food.Fiber, 1),
		Sugar:    formatNumber(food.Sugar, 1),
		Sodium:   formatNumber(food.Sodium, 1),
	}
}

// NewFoodLogViewModels creates multiple view models from Foods slice
func NewFoodLogViewModels(foods store.Foods) []FoodLogViewModel {
	viewModels := make([]FoodLogViewModel, len(foods))
	for i, food := range foods {
		viewModels[i] = NewFoodLogViewModel(food)
	}
	return viewModels
}

// ToStringSlice converts the view model to a slice of strings for table display
func (vm FoodLogViewModel) ToStringSlice() []string {
	return []string{
		fmt.Sprintf("%d", vm.LogID),
		vm.Name,
		vm.Quantity,
		vm.Grams,
		vm.Calories,
		vm.Protein,
		vm.Fat,
		vm.Carbs,
		vm.Fiber,
		vm.Sugar,
		vm.Sodium,
	}
}

// formatNumber formats a float with specified decimal places
func formatNumber(value float64, decimals int) string {
	format := fmt.Sprintf("%%.%df", decimals)
	return fmt.Sprintf(format, value)
}