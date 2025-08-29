package request_header_box

import (
	"api-requester/context"
	"api-requester/domain/request"
	"api-requester/shared/focusable"
	"api-requester/tui/components/button"
)

type Model struct {
	context       *context.AppContext
	request       *request.Request
	subcomponents []focusable.Focusable
	focusIndex    int
	// selectMethod      select_menu.Model
	// inputMethod       input.Model
	// sendRequestButton button.Model
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		context:    ctx,
		request:    nil,
		focusIndex: 0,
		subcomponents: []focusable.Focusable{
			// select_menu.NewModel(),
			// input.NewModel(40, &pl),
			button.NewModel("qweqwe"),
			button.NewModel("asdasdsa"),
		},
	}
}
