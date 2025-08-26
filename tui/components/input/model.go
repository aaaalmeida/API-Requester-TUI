package input

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	textInput textinput.Model
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "qweqw"

	return Model{
		textInput: ti,
	}
}
