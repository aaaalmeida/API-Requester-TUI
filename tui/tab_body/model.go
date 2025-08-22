package tab_body

import (
	"api-requester/request"
	"api-requester/tui/input"
	"api-requester/tui/select_menu"
)

type Model struct {
	request      *request.Request
	selectMethod select_menu.Model
	input        input.Model
}

func NewModel(req *request.Request) Model {
	return Model{
		request:      req,
		selectMethod: select_menu.NewModel(),
		input:        input.NewModel(&req.Url),
	}
}
