package services

import (
	"mydiet/internal/apperrors"
	"mydiet/internal/logger"
	"mydiet/internal/store"
	"time"
)

// NutritionService handles all business logic related to nutrition tracking
type NutritionService struct {
	store store.Store
}

// NewNutritionService creates a new nutrition service
func NewNutritionService(store store.Store) *NutritionService {
	if logger.Log == nil {
		logger.Log = logger.NewLogger()
	}
	return &NutritionService{store: store}
}

// LogFood logs a food item for a specific meal
func (s *NutritionService) LogFood(req LogFoodRequest, dt time.Time) error {
	logger.Log.Printf("Logging food: userId=%d, foodId=%d, foodUnitId=%d, meal=%s, date=%v",
		req.UserID, req.FoodID, req.FoodUnitID, req.Meal, dt)

	if err := s.validateLogFoodRequest(req); err != nil {
		logger.Log.Error("Food logging validation failed: %v", err)
		return err
	}

	logEntry := store.LoggingFood{
		UserId:     req.UserID,
		FoodId:     req.FoodID,
		FoodUnitId: req.FoodUnitID,
		QTY:        req.Quantity,
		Meal:       req.Meal,
		Timestamp:  dt,
	}

	if err := s.store.FoodStore.InsertLog(logEntry); err != nil {
		logger.Log.Error("Failed to insert food log: %v", err)
		return apperrors.Wrap(err, "food log insertion failed")
	}

	return nil
}

// GetMealLogs retrieves food logs for a specific meal on a specific date
func (s *NutritionService) GetMealLogs(userID int, meal store.MealType, date time.Time) (store.Foods, error) {
	logger.Log.Printf("Retrieving meal logs: userId=%d, meal=%s, date=%s", userID, meal, date)

	if userID <= 0 {
		logger.Log.Error("Invalid user ID for meal logs")
		return nil, apperrors.New(apperrors.ErrValidation, "invalid user ID")
	}

	foods, err := s.store.FoodStore.GetLogs(meal, date)
	if err != nil {
		logger.Log.Error("Failed to retrieve meal logs: %v", err)
		return nil, apperrors.Wrap(err, "failed to retrieve meal logs")
	}

	return foods, nil
}

// SearchFoods searches for foods by name
func (s *NutritionService) SearchFoods(query string) (store.Foods, error) {
	logger.Log.Printf("Searching foods: query=%s", query)

	var foods store.Foods
	var err error

	if query == "" {
		foods, err = s.store.FoodStore.GetAll("")
	} else {
		foods, err = s.store.FoodStore.Search(query)
	}

	if err != nil {
		logger.Log.Error("Failed to search foods: %v", err)
		return nil, apperrors.Wrap(err, "food search failed")
	}

	return foods, nil
}

// GetFoodWithUnits retrieves a food item with its available units
func (s *NutritionService) GetFoodWithUnits(foodID int) (*store.Food, error) {
	logger.Log.Printf("Retrieving food with units: foodId=%d", foodID)

	if foodID <= 0 {
		logger.Log.Error("Invalid food ID")
		return nil, apperrors.New(apperrors.ErrValidation, "invalid food ID")
	}

	food, err := s.store.FoodStore.GetByID(foodID)
	if err != nil {
		logger.Log.Error("Failed to retrieve food: %v", err)
		return nil, apperrors.Wrap(err, "failed to retrieve food")
	}

	units, err := s.store.FoodStore.GetUnits(foodID)
	if err != nil {
		logger.Log.Error("Failed to retrieve food units: %v", err)
		return nil, apperrors.Wrap(err, "failed to retrieve food units")
	}

	food.Units = units
	return food, nil
}

// DeleteMealEntry removes a food entry from a meal
func (s *NutritionService) DeleteMealEntry(userID int, logID int) error {
	if userID <= 0 {
		return apperrors.New(apperrors.ErrValidation, "invalid user ID")
	}
	if logID <= 0 {
		return apperrors.New(apperrors.ErrValidation, "invalid log ID")
	}

	// Implementation would go here once the store method is implemented
	return apperrors.New(apperrors.ErrInternal, "delete functionality not yet implemented")
}

// GetDailySummary calculates nutritional totals for a day
func (s *NutritionService) GetDailySummary(userID int, date time.Time) (*DailySummary, error) {
	logger.Log.Printf("Retrieving daily summary: userId=%d, date=%s", userID, date)

	if userID <= 0 {
		logger.Log.Error("Invalid user ID for daily summary")
		return nil, apperrors.New(apperrors.ErrValidation, "invalid user ID")
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
			logger.Log.Error("Failed to retrieve meal logs for summary: %v", err)
			return nil, apperrors.Wrap(err, "failed to retrieve meal logs for summary")
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
		return apperrors.New(apperrors.ErrValidation, "invalid user ID")
	}
	if req.FoodID <= 0 {
		return apperrors.New(apperrors.ErrValidation, "invalid food ID")
	}
	if req.FoodUnitID <= 0 {
		return apperrors.New(apperrors.ErrValidation, "invalid food unit ID")
	}
	if req.Quantity <= 0 {
		return apperrors.New(apperrors.ErrValidation, "quantity must be positive")
	}
	if req.Meal == "" {
		return apperrors.New(apperrors.ErrValidation, "meal type is required")
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

