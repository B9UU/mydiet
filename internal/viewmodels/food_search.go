package viewmodels

import (
	"fmt"
	"mydiet/internal/store"
)

// FoodSearchViewModel handles display formatting for food search results
type FoodSearchViewModel struct {
	ID       int
	Name     string
	Calories string
}

// NewFoodSearchViewModel creates a view model from a domain Food object
func NewFoodSearchViewModel(food store.Food) FoodSearchViewModel {
	return FoodSearchViewModel{
		ID:       food.ID,
		Name:     food.Name,
		Calories: formatNumber(food.Calories, 2),
	}
}

// NewFoodSearchViewModels creates multiple view models from Foods slice
func NewFoodSearchViewModels(foods store.Foods) []FoodSearchViewModel {
	viewModels := make([]FoodSearchViewModel, len(foods))
	for i, food := range foods {
		viewModels[i] = NewFoodSearchViewModel(food)
	}
	return viewModels
}

// ToStringSlice converts the view model to a slice of strings for table display
func (vm FoodSearchViewModel) ToStringSlice() []string {
	return []string{
		fmt.Sprintf("%d", vm.ID),
		vm.Name,
		vm.Calories,
	}
}