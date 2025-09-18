package commands

import (
	"api-requester/shared/selectable"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

// func UserPressEnterInSelect(label string, value any) tea.Cmd {
func UserPressEnterInSelect(obj selectable.Selectable) tea.Cmd {
	return func() tea.Msg {
		return msg.SendSelectValue{
			Label: obj.Label(),
			Value: obj.Value(),
		}
	}
}
