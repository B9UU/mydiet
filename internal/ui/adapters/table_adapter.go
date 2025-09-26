package adapters

import (
	"mydiet/internal/viewmodels"

	"github.com/charmbracelet/bubbles/table"
)

// TableAdapter handles converting view models to table rows
type TableAdapter struct{}

// NewTableAdapter creates a new table adapter
func NewTableAdapter() *TableAdapter {
	return &TableAdapter{}
}

// FoodLogToRows converts food log view models to table rows
func (ta *TableAdapter) FoodLogToRows(viewModels []viewmodels.FoodLogViewModel) []table.Row {
	rows := make([]table.Row, len(viewModels))
	for i, vm := range viewModels {
		rows[i] = table.Row(vm.ToStringSlice())
	}
	return rows
}

// FoodLogToRowsWithSummary converts food log view models to table rows and adds a summary row
func (ta *TableAdapter) FoodLogToRowsWithSummary(viewModels []viewmodels.FoodLogViewModel, summary viewmodels.NutritionSummaryViewModel) []table.Row {
	rows := ta.FoodLogToRows(viewModels)

	// Add summary row if there are entries
	if len(viewModels) > 0 {
		summaryRow := table.Row(summary.ToStringSlice())
		rows = append(rows, summaryRow)
	}

	return rows
}

// FoodSearchToRows converts food search view models to table rows
func (ta *TableAdapter) FoodSearchToRows(viewModels []viewmodels.FoodSearchViewModel) []table.Row {
	rows := make([]table.Row, len(viewModels))
	for i, vm := range viewModels {
		rows[i] = table.Row(vm.ToStringSlice())
	}
	return rows
}