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

type Food struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	ServingSize float64 `db:"serving_size"`
	ServingUnit string  `db:"serving_unit"`
	Calories    float64 `db:"calories"`
	Protein     float64 `db:"protein"`
	Fat         float64 `db:"fat"`
	Carbs       float64 `db:"carbs"`
	Fiber       float64 `db:"fiber"`
	Sugar       float64 `db:"sugar"`
	Sodium      float64 `db:"sodium"`
}
type Foods []Food

func (f FoodStore) Search(name string) (Foods, error) {
	stmt := "SELECT * FROM foods WHERE name LIKE ?"
	args := []any{name + "%"}
	food := Foods{}
	logger.Log.Info(name, food)
	err := f.DB.Select(&food, stmt, args...)
	if err != nil {
		return nil, err
	}

	logger.Log.Info(name, food)
	return food, nil
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
			fmt.Sprintf("%d", meal.ID),
			meal.Name,
			fmt.Sprintf("%.2f", meal.Calories),
			fmt.Sprintf("%.2f", meal.Carbs),
			fmt.Sprintf("%.2f", meal.Protein),
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
