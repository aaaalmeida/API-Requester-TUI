package collection_menu

import "api-requester/request"

type LoadCollectionMsg struct {
	Collection_id int
	Requests      []request.Request
	Err           error
}
