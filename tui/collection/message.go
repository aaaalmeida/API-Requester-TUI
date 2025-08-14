package collection

import "api-requester/request"

type requestLoadedMsg struct {
	collection_id int
	requests      []request.Request
	err           error
}
