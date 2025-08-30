package input

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	textInput textinput.Model
}

func (m *Model) Focus() {
	m.textInput.Focus()
}

func (m *Model) Blur() {
	m.textInput.Blur()
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
