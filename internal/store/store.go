package store

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type FoodStoreInterface interface {
	InsertLog(fl LoggingFood) error
	GetLogs(meal MealType, date time.Time) (Foods, error)
	Search(name string) (Foods, error)
	GetAll(name string) (Foods, error)
	GetUnits(id int) ([]FoodUnits, error)
	GetByID(id int) (*Food, error)
}

type Store struct {
	FoodStore FoodStoreInterface
}

func NewStore(db *sqlx.DB) Store {
	return Store{
		FoodStore: &FoodStore{db},
	}
}
