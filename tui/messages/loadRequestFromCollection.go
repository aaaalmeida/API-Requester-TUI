package messages

import "api-requester/request"

type LoadRequestFromCollectionMsg struct {
	Collection_id int
	Requests      []request.Request
	Err           error
}
