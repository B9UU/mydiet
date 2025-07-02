package form

import (
	"errors"
	"mydiet/internal/logger"
	"mydiet/internal/store"
	"mydiet/internal/types"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form    *huh.Form
	food    *store.Food
	FoodLog store.LoggingFood
}

func New(f *store.Food) Model {
	options := make([]huh.Option[int], len(f.Units))
	for i, u := range f.Units {
		options[i] = huh.NewOption(u.Unit, u.ID)

	}
	return Model{
		food: f,
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Select Unit").
					Key("ID").
					Options(options...),
				huh.NewInput().
					Title("How much?").
					Key("Size").
					Validate(func(str string) error {
						_, err := strconv.Atoi(str)
						if err != nil {
							return errors.New("invalid number")
						}
						return nil
					}),
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return nil

}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	// ...

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Interrupt
		case "q":
			return m, tea.Quit
		case "esc":
			return m, func() tea.Msg {
				return types.ViewMessage{
					Msg:     "back",
					NewView: types.SEARCHBOX,
				}
			}
		}
	}

	if m.form.State == huh.StateCompleted {
		logger.Log.Info("completed")
		unitId := m.form.Get("ID").(int)
		unitSize := m.form.Get("Size").(string)
		uSize, err := strconv.ParseFloat(string(unitSize), 64)
		if err != nil {
			return m, func() tea.Msg {
				return types.ErrMsg(errors.New("Invalid food "))
			}
		}

		m.FoodLog = store.LoggingFood{
			FoodId:     m.food.ID,
			FoodUnitId: unitId,
			QTY:        uSize,
			Meal:       m.food.Meal,
		}
		return m, func() tea.Msg {
			return types.ViewMessage{
				NewView: types.DETAILSVIEW,
				Msg:     "updated",
			}

		}
	}
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m Model) View() string {
	// m.form.Init()
	return m.form.View()
}
