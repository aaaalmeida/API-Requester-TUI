package commands

import (
	"api-requester/context"
	"api-requester/domain/request"
	msg "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func UpdateRequestCmd(ctx *context.AppContext, request_id int, req *request.Request) tea.Cmd {
	err := request.UpdateRequest(ctx, request_id, req)
	return func() tea.Msg {
		return msg.UpdateRequestMsg{
			Err: err,
		}
	}
}
