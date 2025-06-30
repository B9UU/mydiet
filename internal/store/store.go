package store

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	FoodStore  *FoodStore
	MealsStore *MealsStore
}

func NewStore(db *sqlx.DB) Store {
	return Store{
		FoodStore:  &FoodStore{db},
		MealsStore: &MealsStore{},
	}
}
