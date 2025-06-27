package store

import (
	"mydiet/internal/logger"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type Store struct{}

func (s *Store) Get(meal MealType) MealsData {
	return allMeals[meal]
}

func (s *Store) Delete(meal MealType, row table.Row) MealsData {
	id, err := strconv.Atoi(row[0])
	if err != nil {
		logger.Log.Fatal("Unable to parse id Err: %v", err)
	}

	meals := allMeals[meal]
	newMeals := make(MealsData, 0, len(meals))

	for _, m := range meals {
		if m.Id != id {
			newMeals = append(newMeals, m)
		}
	}

	allMeals[meal] = newMeals
	return newMeals
}

func (s *Store) Add(meal MealType, row MealData) MealsData {
	allMeals[meal] = append(allMeals[meal], row)
	return allMeals[meal]
}

var allMeals = map[MealType]MealsData{
	Breakfast: {
		{Id: 1, Name: "Oatmeal", Calories: 150, Carbs: 27, Protein: 5},
		{Id: 2, Name: "Scrambled Eggs", Calories: 200, Carbs: 2, Protein: 12},
		{Id: 3, Name: "Pancakes", Calories: 350, Carbs: 50, Protein: 6},
		{Id: 4, Name: "Avocado Toast", Calories: 300, Carbs: 30, Protein: 7},
		{Id: 5, Name: "Breakfast Burrito", Calories: 450, Carbs: 35, Protein: 20},
		{Id: 6, Name: "Fruit Smoothie", Calories: 180, Carbs: 35, Protein: 3},
	},
	Lunch: {
		{Id: 7, Name: "Grilled Chicken Salad", Calories: 350, Carbs: 15, Protein: 30},
		{Id: 8, Name: "Turkey Sandwich", Calories: 450, Carbs: 40, Protein: 25},
		{Id: 9, Name: "Veggie Wrap", Calories: 300, Carbs: 35, Protein: 10},
		{Id: 10, Name: "Tuna Salad", Calories: 320, Carbs: 10, Protein: 28},
		{Id: 11, Name: "Buddha Bowl", Calories: 500, Carbs: 55, Protein: 20},
		{Id: 12, Name: "Quinoa & Veggie Mix", Calories: 400, Carbs: 45, Protein: 18},
	},
	Dinner: {
		{Id: 13, Name: "Spaghetti Bolognese", Calories: 600, Carbs: 70, Protein: 35},
		{Id: 14, Name: "Grilled Salmon", Calories: 500, Carbs: 0, Protein: 40},
		{Id: 15, Name: "Beef Stir Fry", Calories: 550, Carbs: 30, Protein: 38},
		{Id: 16, Name: "Roasted Chicken & Veggies", Calories: 480, Carbs: 25, Protein: 36},
		{Id: 17, Name: "Lentil Curry", Calories: 430, Carbs: 50, Protein: 22},
		{Id: 18, Name: "Shrimp Tacos", Calories: 470, Carbs: 35, Protein: 30},
	},
	Snack: {
		{Id: 19, Name: "Protein Bar", Calories: 200, Carbs: 20, Protein: 15},
		{Id: 20, Name: "Greek Yogurt", Calories: 120, Carbs: 8, Protein: 10},
		{Id: 21, Name: "Apple with Peanut Butter", Calories: 180, Carbs: 22, Protein: 4},
		{Id: 22, Name: "Trail Mix", Calories: 250, Carbs: 16, Protein: 6},
		{Id: 23, Name: "Boiled Eggs", Calories: 140, Carbs: 1, Protein: 13},
		{Id: 24, Name: "Cottage Cheese", Calories: 110, Carbs: 4, Protein: 12},
		{Id: 24, Name: "Cottage Cheese", Calories: 110, Carbs: 4, Protein: 12},
		{Id: 24, Name: "Cottage Cheese", Calories: 110, Carbs: 4, Protein: 12},
		{Id: 24, Name: "Cottage Cheese", Calories: 110, Carbs: 4, Protein: 12},
	},
}
