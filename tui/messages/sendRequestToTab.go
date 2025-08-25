package messages

import "api-requester/request"

type SendRequestToTabMsg struct {
	Request *request.Request
	Err     error
}
