package tab_body

import (
	"api-requester/context"
	"api-requester/domain/request"
	"api-requester/tui/components/input"
	"api-requester/tui/components/select_menu"
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
