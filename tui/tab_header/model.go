package tab_header

import "api-requester/request"

type Model struct {
	isSelected bool
	request    *request.Request
}

func NewModel(req *request.Request) Model {
	return Model{
		isSelected: false,
		request:    req,
	}
}
