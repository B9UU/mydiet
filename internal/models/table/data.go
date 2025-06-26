package tablelisting

import (
	"github.com/charmbracelet/bubbles/table"
)

var columns = []table.Column{
	{Title: "Index", Width: 4},
	{Title: "Name", Width: 10},
	{Title: "Calories", Width: 10},
	{Title: "Carbs", Width: 10},
	{Title: "Protein", Width: 10},
}
var rows = []table.Row{
	{"1", "Eggs", "2", "3", "10"},
	{"1", "Eggs", "2", "3", "10"},
	{"1", "Eggs", "2", "3", "10"},
	{"1", "Eggs", "2", "3", "10"},
	{"1", "Eggs", "2", "3", "10"},
}
