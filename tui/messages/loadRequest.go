package messages

import "api-requester/request"

type LoadRequestMsg struct {
	Request request.Request
	Err     error
}
