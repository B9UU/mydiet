package services

import (
	"errors"
	"mydiet/internal/store"
	"time"
)

// NutritionService handles all business logic related to nutrition tracking
type NutritionService struct {
	store store.Store
}

// NewNutritionService creates a new nutrition service
func NewNutritionService(store store.Store) *NutritionService {
	return &NutritionService{store: store}
}

// LogFood logs a food item for a specific meal
func (s *NutritionService) LogFood(req LogFoodRequest) error {
	if err := s.validateLogFoodRequest(req); err != nil {
		return err
	}

	logEntry := store.LoggingFood{
		UserId:     req.UserID,
		FoodId:     req.FoodID,
		FoodUnitId: req.FoodUnitID,
		QTY:        req.Quantity,
		Meal:       req.Meal,
		Timestamp:  time.Now(),
	}

	return s.store.FoodStore.InsertLog(logEntry)
}

// GetMealLogs retrieves food logs for a specific meal on a specific date
func (s *NutritionService) GetMealLogs(userID int, meal store.MealType, date time.Time) (store.Foods, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.store.FoodStore.GetLogs(meal, date)
}

// SearchFoods searches for foods by name
func (s *NutritionService) SearchFoods(query string) (store.Foods, error) {
	if query == "" {
		return s.store.FoodStore.GetAll("")
	}
	return s.store.FoodStore.Search(query)
}

// GetFoodWithUnits retrieves a food item with its available units
func (s *NutritionService) GetFoodWithUnits(foodID int) (*store.Food, error) {
	if foodID <= 0 {
		return nil, errors.New("invalid food ID")
	}

	food, err := s.store.FoodStore.GetByID(foodID)
	if err != nil {
		return nil, err
	}

	units, err := s.store.FoodStore.GetUnits(foodID)
	if err != nil {
		return nil, err
	}

	food.Units = units
	return food, nil
}

// DeleteMealEntry removes a food entry from a meal
func (s *NutritionService) DeleteMealEntry(userID int, logID int) error {
	if userID <= 0 {
		return errors.New("invalid user ID")
	}
	if logID <= 0 {
		return errors.New("invalid log ID")
	}

	// Implementation would go here once the store method is implemented
	return errors.New("delete functionality not yet implemented")
}

// GetDailySummary calculates nutritional totals for a day
func (s *NutritionService) GetDailySummary(userID int, date time.Time) (*DailySummary, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	summary := &DailySummary{
		Date:          date,
		UserID:        userID,
		MealSummaries: make(map[store.MealType]MealSummary),
	}

	// Get logs for all meals
	meals := []store.MealType{store.Breakfast, store.Lunch, store.Dinner, store.Snack}

	for _, meal := range meals {
		foods, err := s.GetMealLogs(userID, meal, date)
		if err != nil {
			return nil, err
		}

		summary.MealSummaries[meal] = s.calculateMealSummary(foods)
		summary.TotalCalories += summary.MealSummaries[meal].Calories
		summary.TotalProtein += summary.MealSummaries[meal].Protein
		summary.TotalFat += summary.MealSummaries[meal].Fat
		summary.TotalCarbs += summary.MealSummaries[meal].Carbs
	}

	return summary, nil
}

func (s *NutritionService) validateLogFoodRequest(req LogFoodRequest) error {
	if req.UserID <= 0 {
		return errors.New("invalid user ID")
	}
	if req.FoodID <= 0 {
		return errors.New("invalid food ID")
	}
	if req.FoodUnitID <= 0 {
		return errors.New("invalid food unit ID")
	}
	if req.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	if req.Meal == "" {
		return errors.New("meal type is required")
	}
	return nil
}

func (s *NutritionService) calculateMealSummary(foods store.Foods) MealSummary {
	var summary MealSummary

	for _, food := range foods {
		summary.Calories += food.Calories
		summary.Protein += food.Protein
		summary.Fat += food.Fat
		summary.Carbs += food.Carbs
		summary.Fiber += food.Fiber
		summary.Sugar += food.Sugar
		summary.Sodium += food.Sodium
	}

	return summary
}