package request_header_box

import (
	"api-requester/context"
	"api-requester/domain/request"
	"api-requester/tui/components/button"
	"api-requester/tui/components/input"
	"api-requester/tui/components/select_menu"
)

type Model struct {
	context           *context.AppContext
	request           *request.Request
	selectMethod      select_menu.Model
	inputMethod       input.Model
	sendRequestButton button.Model
}

func NewModel(ctx *context.AppContext) Model {
	pl := "bolotinha"
	return Model{
		context:           ctx,
		request:           nil,
		selectMethod:      select_menu.NewModel(),
		inputMethod:       input.NewModel(40, &pl),
		sendRequestButton: button.NewModel(nil, &pl),
	}
}
