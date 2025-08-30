package header

import (
	"api-requester/context"
	"api-requester/domain/request"
)

type Model struct {
	context          *context.AppContext
	cursor           int
	selectedTabIndex int
	requests         []*request.Request
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		cursor:           -1,
		selectedTabIndex: -1,
		requests:         make([]*request.Request, 0),
		context:          ctx,
	}
}
