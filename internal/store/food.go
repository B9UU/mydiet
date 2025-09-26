package store

import (
	"mydiet/internal/logger"

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
	logger.Log.Info("Search term", "name", name)
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
	logger.Log.Info("Retrieving all foods")
	stmt := "SELECT * FROM foods;"
	args := []any{name}
	logger.Log.Info("Searching with name parameter", "name", name)
	food := Foods{}
	err := f.DB.Select(&food, stmt, args...)
	if err != nil {
		return nil, err
	}
	logger.Log.Info("Food count", "count", len(food))

	return food, nil
}

func (s Foods) GetId(id int) *Food {
	for _, v := range s {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

// GetByID retrieves a single food by ID
func (f FoodStore) GetByID(id int) (*Food, error) {
	stmt := "SELECT * FROM foods WHERE id = ?"
	food := &Food{}
	err := f.DB.Get(food, stmt, id)
	if err != nil {
		return nil, err
	}
	return food, nil
}
