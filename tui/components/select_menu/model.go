package select_menu

import (
	"api-requester/shared/selectable"
)

type Model struct {
	Options      []selectable.Selectable
	cursor       int
	selectedItem int
	isOpened     bool
	width        int
}

func NewModel() Model {
	return Model{
		cursor:       0,
		selectedItem: 0,
		Options:      nil,
		isOpened:     false,
		width:        8,
	}
}
