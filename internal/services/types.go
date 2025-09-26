package services

import (
	"mydiet/internal/store"
	"time"
)

// LogFoodRequest represents a request to log food for a meal
type LogFoodRequest struct {
	UserID     int              `json:"user_id"`
	FoodID     int              `json:"food_id"`
	FoodUnitID int              `json:"food_unit_id"`
	Quantity   float64          `json:"quantity"`
	Meal       store.MealType   `json:"meal"`
}

// MealSummary contains nutritional totals for a single meal
type MealSummary struct {
	Calories float64 `json:"calories"`
	Protein  float64 `json:"protein"`
	Fat      float64 `json:"fat"`
	Carbs    float64 `json:"carbs"`
	Fiber    float64 `json:"fiber"`
	Sugar    float64 `json:"sugar"`
	Sodium   float64 `json:"sodium"`
}

// DailySummary contains nutritional totals for an entire day
type DailySummary struct {
	Date           time.Time                     `json:"date"`
	UserID         int                           `json:"user_id"`
	TotalCalories  float64                       `json:"total_calories"`
	TotalProtein   float64                       `json:"total_protein"`
	TotalFat       float64                       `json:"total_fat"`
	TotalCarbs     float64                       `json:"total_carbs"`
	MealSummaries  map[store.MealType]MealSummary `json:"meal_summaries"`
}

// SearchFoodRequest represents a food search query
type SearchFoodRequest struct {
	Query  string `json:"query"`
	UserID int    `json:"user_id"` // For future personalization
}

// GetMealLogsRequest represents a request for meal logs
type GetMealLogsRequest struct {
	UserID int              `json:"user_id"`
	Meal   store.MealType   `json:"meal"`
	Date   time.Time        `json:"date"` // Date to filter logs
}