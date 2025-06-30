package types

import "mydiet/internal/store"

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

type FailedRequest error

type SuccessRequest store.Foods
