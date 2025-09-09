package request_header_box

import (
	"api-requester/context"
	"api-requester/domain/request"
	"api-requester/shared/focusable"
	"api-requester/tui/components/button"
	"api-requester/tui/components/input"
	"api-requester/tui/components/select_menu"
)

const SELECT_MENU_INDEX int = 0
const INPUT_URL_INDEX int = 1
const SEND_REQUEST_BUTTON_INDEX int = 2

type Model struct {
	context                *context.AppContext
	request                *request.Request
	selectedComponentIndex int
	cursor                 int
	subcomponents          []focusable.Focusable
}

func NewModel(ctx *context.AppContext) Model {
	placeholder := "Inform url"
	return Model{
		context:                ctx,
		request:                nil,
		selectedComponentIndex: SELECT_MENU_INDEX,
		cursor:                 SELECT_MENU_INDEX,
		subcomponents: []focusable.Focusable{
			select_menu.NewModel(ctx),
			input.NewModel(40, &placeholder, ctx),
			button.NewModel(">>"),
		},
	}
}
