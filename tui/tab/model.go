package tab

import (
	"api-requester/context"
	"api-requester/request"
	"api-requester/tui/tab_body"
	"api-requester/tui/tab_header"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	activeTabIndex int
	cursor         int
	requests       []*request.Request
	tab_header     tea.Model
	tab_body       tea.Model
	context        *context.AppContext
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:        ctx,
		activeTabIndex: 0,
		cursor:         0,
		requests:       nil,
		tab_header:     tab_header.NewModel(ctx),
		tab_body:       tab_body.NewModel(ctx),
	}
}
