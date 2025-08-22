package tab

import (
	"api-requester/context"
	"api-requester/request"
	"api-requester/tui/tab_body"
	"api-requester/tui/tab_header"
)

type model struct {
	context        *context.AppContext
	activeTabIndex int
	cursor         int
	requests       []*request.Request
	tabsHeader     []*tab_header.Model
	body           *tab_body.Model
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:        ctx,
		activeTabIndex: 0,
		cursor:         0,
		tabsHeader:     nil,
		requests:       nil,
		body:           nil,
	}
}
