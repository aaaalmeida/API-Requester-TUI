package commands

import (
	"api-requester/request"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func SendRequestCmd(request *request.Request) tea.Cmd {
	return func() tea.Msg {
		return msg.LoadRequestMsg{
			Request: request,
			Err:     nil,
		}
	}
}
