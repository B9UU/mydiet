package app

import (
	"mydiet/internal/models/details"
	"mydiet/internal/models/form"
	"mydiet/internal/models/searchbox"
	"mydiet/internal/services"
	"mydiet/internal/store"
	"mydiet/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

// ViewCoordinator manages view transitions and coordinates between different views
type ViewCoordinator struct {
	activeView  types.View
	views       *Views
	service     *services.NutritionService
	currentUser int // Currently hardcoded to 1, will be replaced with proper user management
}

// Views holds all the application views
type Views struct {
	Details   details.Model
	Searchbox searchbox.Model
	Form      form.Model
}

// NewViewCoordinator creates a new view coordinator
func NewViewCoordinator(service *services.NutritionService, store store.Store) *ViewCoordinator {
	return &ViewCoordinator{
		activeView:  types.DETAILSVIEW,
		service:     service,
		currentUser: 1, // Will be replaced with proper user management
		views: &Views{
			Details:   details.New(store), // Keep existing for now, will refactor later
			Searchbox: searchbox.Model{Store: store},
		},
	}
}

// HandleViewMessage processes view transition messages
func (vc *ViewCoordinator) HandleViewMessage(msg types.ViewMessage) (*ViewCoordinator, tea.Cmd) {
	vc.activeView = msg.NewView

	switch msg.NewView {
	case types.DETAILSVIEW:
		return vc.handleDetailsView(msg)
	case types.FORMVIEW:
		return vc.handleFormView(msg)
	case types.SEARCHBOX:
		return vc.handleSearchboxView(msg, vc.views.Searchbox.Store)
	default:
		return vc, nil
	}
}

// GetActiveView returns the currently active view
func (vc *ViewCoordinator) GetActiveView() types.View {
	return vc.activeView
}

// GetViews returns all views
func (vc *ViewCoordinator) GetViews() *Views {
	return vc.views
}

// handleDetailsView handles transitions to the details view
func (vc *ViewCoordinator) handleDetailsView(msg types.ViewMessage) (*ViewCoordinator, tea.Cmd) {
	if msg.Msg == "updated" {
		// Business logic: log the food entry
		if vc.views.Form.FoodLog.FoodId > 0 {
			req := services.LogFoodRequest{
				UserID:     vc.currentUser,
				FoodID:     vc.views.Form.FoodLog.FoodId,
				FoodUnitID: vc.views.Form.FoodLog.FoodUnitId,
				Quantity:   vc.views.Form.FoodLog.QTY,
				Meal:       vc.views.Form.FoodLog.Meal,
			}

			if err := vc.service.LogFood(req); err != nil {
				return vc, func() tea.Msg {
					return types.ErrMsg(NewDatabaseError("Failed to log food", err))
				}
			}

			// Refresh the details view
			vc.views.Details.SyncRowsFor()
		}
	}
	return vc, nil
}

// handleFormView handles transitions to the form view
func (vc *ViewCoordinator) handleFormView(msg types.ViewMessage) (*ViewCoordinator, tea.Cmd) {
	food, ok := msg.Msg.(*store.Food)
	if !ok {
		return vc, func() tea.Msg {
			return types.ErrMsg(NewValidationError("Invalid food data"))
		}
	}

	// Use service to get food with units
	foodWithUnits, err := vc.service.GetFoodWithUnits(food.ID)
	if err != nil {
		return vc, func() tea.Msg {
			return types.ErrMsg(NewDatabaseError("Failed to load food details", err))
		}
	}

	foodWithUnits.Meal = food.Meal // Preserve the meal type
	vc.views.Form = form.New(foodWithUnits)

	return vc, nil
}

// handleSearchboxView handles transitions to the searchbox view
func (vc *ViewCoordinator) handleSearchboxView(msg types.ViewMessage, s store.Store) (*ViewCoordinator, tea.Cmd) {
	meal, ok := msg.Msg.(store.MealType)
	if !ok {
		return vc, func() tea.Msg {
			return types.ErrMsg(NewValidationError("Invalid meal type"))
		}
	}

	// Keep the existing searchbox creation, will be refactored to use service layer
	vc.views.Searchbox = searchbox.New(meal, s)

	return vc, nil
}
