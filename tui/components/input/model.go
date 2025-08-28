package input

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	textInput textinput.Model
}

func NewModel(width int, placeholder *string) Model {
	ti := textinput.New()
	ti.Prompt = "#"
	ti.Width = width
	if placeholder != nil {
		ti.Placeholder = *placeholder
	}

	return Model{
		textInput: ti,
	}
}
