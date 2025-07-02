package store

import (
	"fmt"
	"mydiet/internal/logger"

	"github.com/charmbracelet/bubbles/table"
	"github.com/jmoiron/sqlx"
)

type FoodStore struct {
	DB *sqlx.DB
}

func (f FoodStore) Search(name string) (Foods, error) {
	stmt := "SELECT * FROM foods WHERE name LIKE ?"
	args := []any{name + "%"}
	food := Foods{}
	err := f.DB.Select(&food, stmt, args...)
	if err != nil {
		return nil, err
	}
	logger.Log.Info(name)
	return food, nil
}
func (f FoodStore) GetUnits(id int) ([]FoodUnits, error) {
	stmt := "SELECT * FROM food_units WHERE food_id = ?"
	args := []any{id}
	foodUnits := []FoodUnits{}
	err := f.DB.Select(&foodUnits, stmt, args...)
	if err != nil {
		return nil, err
	}

	return foodUnits, nil
}
func (f FoodStore) GetAll(name string) (Foods, error) {
	stmt := "SELECT * FROM foods;"
	args := []any{name}
	food := Foods{}
	err := f.DB.Select(&food, stmt, args...)
	if err != nil {
		return nil, err
	}

	return food, nil
}

func (s Foods) TableRowsFor() []table.Row {
	var rows []table.Row
	for _, meal := range s {
		row := table.Row{
			fmt.Sprintf("%d", meal.LogID),
			meal.Name,
			fmt.Sprintf("%.1f %s", meal.QTY, meal.Unit),
			fmt.Sprintf("%.1f", meal.Grams),

			fmt.Sprintf("%.1f", meal.Calories),
			fmt.Sprintf("%.1f", meal.Protein),
			fmt.Sprintf("%.1f", meal.Fat),
			fmt.Sprintf("%.1f", meal.Carbs),
			fmt.Sprintf("%.1f", meal.Fiber),
			fmt.Sprintf("%.1f", meal.Sugar),
			fmt.Sprintf("%.1f", meal.Sodium),
		}

		rows = append(rows, row)
	}
	return rows
}

func (s Foods) SearchRows() []table.Row {
	var rows []table.Row
	for _, meal := range s {
		row := table.Row{
			fmt.Sprintf("%d", meal.ID),
			meal.Name,
			fmt.Sprintf("%.2f", meal.Calories),
		}
		rows = append(rows, row)
	}
	return rows
}
func (s Foods) GetId(id int) *Food {
	for _, v := range s {
		if v.ID == id {
			return &v
		}
	}
	return nil
}
