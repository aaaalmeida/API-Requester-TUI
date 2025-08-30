package request_header_box

import (
	"api-requester/context"
	"api-requester/domain/request"
	"api-requester/shared/focusable"
	"api-requester/tui/components/button"
	"api-requester/tui/components/input"
)

type Model struct {
	context       *context.AppContext
	request       *request.Request
	subcomponents []focusable.Focusable
	focusIndex    int
}

func NewModel(ctx *context.AppContext) Model {
	pl := "teste"

	inputModel := input.NewModel(40, &pl)
	return Model{
		context:    ctx,
		request:    nil,
		focusIndex: 0,
		subcomponents: []focusable.Focusable{
			&inputModel,
			button.NewModel("qweqwe"),
			button.NewModel("asdasdsa"),
		},
	}
}
