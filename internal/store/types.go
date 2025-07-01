package store

import (
	"time"
)

type MealType string

const (
	Breakfast MealType = "Breakfast"
	Lunch     MealType = "Lunch"
	Dinner    MealType = "Dinner"
	Snack     MealType = "Snack"
)

type Food struct {
	LogID    int      `db:"log_id"`
	ID       int      `db:"id"`
	Name     string   `db:"name"`
	Meal     MealType `db:"meal"`
	QTY      float64  `db:"quantity"`
	Unit     string   `db:"unit"`
	Grams    float64  `db:"grams"`
	Calories float64  `db:"calories"`
	Protein  float64  `db:"protein"`
	Fat      float64  `db:"fat"`
	Carbs    float64  `db:"carbs"`
	Fiber    float64  `db:"fiber"`
	Sugar    float64  `db:"sugar"`
	Sodium   float64  `db:"sodium"`
	Units    []FoodUnits
}
type FoodUnits struct {
	ID          int     `db:"id"`
	FoodID      int     `db:"food_id"`
	Unit        string  `db:"unit"`
	SizeInGrams float64 `db:"size_in_grams"`
}

type Foods []Food
type LoggingFood struct {
	UserId     int       `db:"user_id"`
	FoodId     int       `db:"food_id"`
	FoodUnitId int       `db:"food_unit_id"`
	QTY        float64   `db:"quantity"`
	Meal       MealType  `db:"meal"`
	Timestamp  time.Time `db:"timestamp"`
}
