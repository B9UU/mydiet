package store

import (
	"mydiet/internal/logger"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
)

// Food log operations for meal tracking

func (s FoodStore) GetLogs(meal MealType, date time.Time) (Foods, error) {
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
    fl.user_id = ?
    AND fl.meal = ? COLLATE NOCASE
    AND DATE(fl.timestamp) = DATE(?)`

	userID := 1

	// Format date to ensure consistent comparison
	dateStr := date.Format("2006-01-02")

	args := []any{userID, strings.ToLower(string(meal)), dateStr}
	foods := Foods{}
	logger.Log.Info("Getting logs for", "meal", meal, "date", dateStr, "user", userID)
	err := s.DB.Select(&foods, stmt, args...)
	if err != nil {
		logger.Log.Error("Error getting logs", "error", err)
		return Foods{}, err
	}

	logger.Log.Info("Found logs", "count", len(foods))
	return foods, nil
}

// Delete removes a food log entry (implementation pending)
func (s FoodStore) Delete(meal MealType, row table.Row) Foods {
	// Implementation to be added when delete functionality is required
	return Foods{}
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
