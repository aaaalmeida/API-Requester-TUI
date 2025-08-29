package request_response_box

import (
	"api-requester/domain/request"
)

type Model struct {
	request *request.Request
}

func NewModel() Model {
	return Model{
		// request: request,
	}
}
