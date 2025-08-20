package tab

import (
	"api-requester/context"
	"api-requester/request"
)

type model struct {
	context   *context.AppContext
	activeTab int
	requests  []request.Request
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:   ctx,
		activeTab: 0,
		requests:  nil,
	}
}
