package button

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	onPress tea.Cmd
	label   string
	Focused bool
}

func NewModel(label string) Model {
	return Model{
		// onPress: onPress,
		label:   label,
		Focused: false,
	}
}

func (m Model) Focus() {
	m.Focused = true
}

func (m Model) Blur() {
	m.Focused = false
}
