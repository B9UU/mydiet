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
