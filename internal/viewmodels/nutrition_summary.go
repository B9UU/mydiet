package viewmodels

import (
	"mydiet/internal/store"
)

// NutritionSummaryViewModel handles display formatting for nutrition totals
type NutritionSummaryViewModel struct {
	TotalCalories string
	TotalProtein  string
	TotalFat      string
	TotalCarbs    string
	TotalFiber    string
	TotalSugar    string
	TotalSodium   string
}

// NewNutritionSummaryViewModel creates a view model from Foods slice
func NewNutritionSummaryViewModel(foods store.Foods) NutritionSummaryViewModel {
	var totalCalories, totalProtein, totalFat, totalCarbs float64
	var totalFiber, totalSugar, totalSodium float64

	for _, food := range foods {
		totalCalories += food.Calories
		totalProtein += food.Protein
		totalFat += food.Fat
		totalCarbs += food.Carbs
		totalFiber += food.Fiber
		totalSugar += food.Sugar
		totalSodium += food.Sodium
	}

	return NutritionSummaryViewModel{
		TotalCalories: formatNumber(totalCalories, 1),
		TotalProtein:  formatNumber(totalProtein, 1),
		TotalFat:      formatNumber(totalFat, 1),
		TotalCarbs:    formatNumber(totalCarbs, 1),
		TotalFiber:    formatNumber(totalFiber, 1),
		TotalSugar:    formatNumber(totalSugar, 1),
		TotalSodium:   formatNumber(totalSodium, 1),
	}
}

// ToStringSlice converts the summary to a slice of strings for table display
func (vm NutritionSummaryViewModel) ToStringSlice() []string {
	return []string{
		"", // Empty LogID column
		"TOTAL",
		"",
		"",
		vm.TotalCalories,
		vm.TotalProtein,
		vm.TotalFat,
		vm.TotalCarbs,
		vm.TotalFiber,
		vm.TotalSugar,
		vm.TotalSodium,
	}
}