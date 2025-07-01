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
	FORMVIEW
	SEARCHBOX
)

type ErrMsg error

type SuccessRequest store.Foods
type FormSubbmited bool
