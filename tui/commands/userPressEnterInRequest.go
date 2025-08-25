package commands

import (
	"api-requester/request"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func UserPressEnterInRequestCmd(request *request.Request) tea.Cmd {
	return func() tea.Msg {
		return msg.SendRequestToTabMsg{
			Request: request,
			Err:     nil,
		}
	}
}
