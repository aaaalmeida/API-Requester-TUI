package tab_body

import (
	"api-requester/context"
	"api-requester/request"
	"api-requester/tui/input"
	"api-requester/tui/select_menu"
)

type Model struct {
	request      *request.Request
	selectMethod select_menu.Model
	inputMethod  input.Model
	context      *context.AppContext
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		context:      ctx,
		request:      nil,
		selectMethod: select_menu.NewModel(),
		inputMethod:  input.NewModel(),
	}
}
