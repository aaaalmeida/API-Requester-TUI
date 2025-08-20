package messages

import "api-requester/collection"

type LoadCollectionsMsg struct {
	Collections []collection.Collection
	Err         error
}
