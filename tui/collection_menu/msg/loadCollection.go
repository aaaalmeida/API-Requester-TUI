package collection_menu

import (
	"api-requester/collection"
)

type LoadCollectionMsg struct {
	Collections []collection.Collection
	Err         error
}
