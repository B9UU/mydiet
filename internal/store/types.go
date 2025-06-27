package store

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type MealData struct {
	Id       int
	Name     string
	Calories int
	Carbs    int
	Protein  int
}

type MealsData []MealData

type MealType string

const (
	Breakfast MealType = "Breakfast"
	Lunch     MealType = "Lunch"
	Dinner    MealType = "Dinner"
	Snack     MealType = "Snack"
)

func (s MealsData) TableRowsFor() []table.Row {
	var rows []table.Row

	for _, meal := range s {
		rows = append(rows, table.Row{
			strconv.Itoa(meal.Id),
			meal.Name,
			strconv.Itoa(meal.Calories),
			strconv.Itoa(meal.Carbs),
			strconv.Itoa(meal.Protein),
		})
	}
	return rows
}
