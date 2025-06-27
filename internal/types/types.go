package types

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type View int

type ViewMessage struct {
	Msg     any
	NewView View
}

const (
	_ View = iota
	DETAILSVIEW
	SEARCHVIEW
	SEARCHBOX
)

type FailedRequest struct {
	Err error
}
type SuccessRequest struct {
	Suggetions []string
}

type DummyJson struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type SuccessDummy struct {
	Data DummyJson
}

func (s DummyJson) TableRowsFor() []table.Row {
	var rows []table.Row

	for _, meal := range s.Products {
		rows = append(rows, table.Row{
			strconv.Itoa(meal.Id),
			meal.Title,
			meal.Description,
		})
	}
	return rows
}
