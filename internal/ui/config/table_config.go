package config

import "github.com/charmbracelet/bubbles/table"

// TableConfig defines column layouts and formatting rules for different tables
type TableConfig struct{}

// NewTableConfig creates a new table configuration
func NewTableConfig() *TableConfig {
	return &TableConfig{}
}

// FoodLogColumns returns the column configuration for food log tables
func (tc *TableConfig) FoodLogColumns() []table.Column {
	return []table.Column{
		{Title: "ID", Width: 0},       // Hidden ID column for selection
		{Title: "Food", Width: 20},
		{Title: "Quantity", Width: 12},
		{Title: "Grams", Width: 8},
		{Title: "Calories", Width: 8},
		{Title: "Protein", Width: 8},
		{Title: "Fat", Width: 6},
		{Title: "Carbs", Width: 8},
		{Title: "Fiber", Width: 6},
		{Title: "Sugar", Width: 6},
		{Title: "Sodium", Width: 8},
	}
}

// FoodSearchColumns returns the column configuration for food search tables
func (tc *TableConfig) FoodSearchColumns() []table.Column {
	return []table.Column{
		{Title: "Id", Width: 0},       // Hidden for selection
		{Title: "Name", Width: 30},
		{Title: "Calories", Width: 10},
	}
}

// CompactFoodLogColumns returns a more compact column configuration for smaller screens
func (tc *TableConfig) CompactFoodLogColumns() []table.Column {
	return []table.Column{
		{Title: "ID", Width: 0},       // Hidden ID column
		{Title: "Food", Width: 25},
		{Title: "Qty", Width: 10},
		{Title: "Cal", Width: 6},
		{Title: "Pro", Width: 6},
		{Title: "Fat", Width: 6},
		{Title: "Carb", Width: 6},
	}
}