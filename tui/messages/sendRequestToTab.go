package messages

import "api-requester/domain/request"

type SendRequestToTabMsg struct {
	Request *request.Request
	Err     error
}
