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
	isFocused    bool
}

func NewModel() *Model {
	return &Model{
		isFocused:    false,
		cursor:       0,
		selectedItem: 0,
		Options:      nil,
		isOpened:     false,
		width:        8,
	}
}

func (M *Model) Focus() {
	M.isFocused = true
}

func (M *Model) Blur() {
	M.isFocused = false
}
