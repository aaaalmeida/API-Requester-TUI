package input

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	textInput textinput.Model
}

func NewModel(placeholder *string) Model {
	ti := textinput.New()
	ti.Placeholder = *placeholder

	return Model{
		textInput: ti,
	}
}
