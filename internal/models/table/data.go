package tablelisting

import (
	"github.com/charmbracelet/bubbles/table"
)

var columns = []table.Column{
	{Title: "log_id", Width: 0},
	{Title: "Name", Width: 15},
	{Title: "Quantity", Width: 15},
	{Title: "Grams", Width: 5},

	{Title: "Calories", Width: 5},
	{Title: "Protein", Width: 5},
	{Title: "Fat", Width: 5},
	{Title: "Carbs", Width: 5},
	{Title: "Fiber", Width: 5},
	{Title: "Sugar", Width: 5},
	{Title: "Sodium mg", Width: 10},
}
