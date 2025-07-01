package store

import (
	"mydiet/internal/logger"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
)

// TODO: maybe food log rather than meals

func (s FoodStore) GetLogs(meal MealType) (Foods, error) {

	stmt := ` SELECT 
    fl.id as log_id,
    f.name,
    fl.meal,
    fl.quantity,
    fu.unit,
    (fl.quantity * fu.size_in_grams) AS grams,
    f.calories * (fl.quantity * fu.size_in_grams / 100.0) AS calories,
    f.protein * (fl.quantity * fu.size_in_grams / 100.0) AS protein,
    f.fat * (fl.quantity * fu.size_in_grams / 100.0) AS fat,
    f.carbs * (fl.quantity * fu.size_in_grams / 100.0) AS carbs,
    f.fiber * (fl.quantity * fu.size_in_grams / 100.0) AS fiber,
    f.sugar * (fl.quantity * fu.size_in_grams / 100.0) AS sugar,
    f.sodium * (fl.quantity * fu.size_in_grams / 100.0) AS sodium
FROM 
    food_logs fl
JOIN 
    foods f ON fl.food_id = f.id
JOIN 
    food_units fu ON fl.food_unit_id = fu.id
WHERE 
    fl.user_id = ? AND fl.meal = ? COLLATE NOCASE`

	userID := 1

	args := []any{userID, strings.ToLower(string(meal))}
	foods := Foods{}
	logger.Log.Info(len(foods), "from store", args)
	err := s.DB.Select(&foods, stmt, args...)
	if err != nil {
		logger.Log.Error(err)
		return Foods{}, err
	}

	return foods, nil
}

func (s FoodStore) Delete(meal MealType, row table.Row) Foods {
	return Foods{}
	// id, err := strconv.Atoi(row[0])
	// if err != nil {
	// 	logger.Log.Fatal("Unable to parse id Err: %v", err)
	// }
	//
	// meals := allMeals[meal]
	// newMeals := make(Foods, 0, len(meals))
	//
	// for _, m := range meals {
	// 	if m.ID != id {
	// 		newMeals = append(newMeals, m)
	// 	}
	// }
	//
	// allMeals[meal] = newMeals
	// return newMeals
}

func (f FoodStore) InsertLog(fl LoggingFood) error {
	fl.UserId = 1
	fl.Timestamp = time.Now()
	stmt := `INSERT INTO food_logs (
    user_id, food_id,
    food_unit_id, quantity,
	meal, timestamp) VALUES (
		:user_id, :food_id,
		:food_unit_id, :quantity,
		:meal, :timestamp
	)`
	_, err := f.DB.NamedExec(stmt, fl)
	if err != nil {
		return err
	}
	return nil
}
