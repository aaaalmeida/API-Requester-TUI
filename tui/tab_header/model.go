package tab_header

import (
	"api-requester/context"
	"api-requester/request"
)

type Model struct {
	context     *context.AppContext
	cursor      int
	selectedTab int
	requests    []*request.Request
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		cursor:      0,
		selectedTab: 0,
		requests:    nil,
		context:     ctx,
	}
}
