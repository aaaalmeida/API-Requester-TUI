package button

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	onPress tea.Cmd
	text    *string
}

func NewModel(onPress tea.Cmd, text *string) Model {
	return Model{
		onPress: onPress,
		text:    text,
	}
}
