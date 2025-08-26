package select_menu

import (
	"api-requester/shared/selectable"
)

type Model struct {
	options      []selectable.SelectOption
	cursor       int
	selectedItem int
	isOpened     bool
	width        int
}

func NewModel() Model {
	return Model{
		cursor:       0,
		selectedItem: 0,
		options:      nil,
		isOpened:     false,
		width:        20,
	}
}
