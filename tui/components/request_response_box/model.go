package request_response_box

import (
	"api-requester/context"
)

type Model struct {
	context *context.AppContext
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		context: ctx,
	}
}
