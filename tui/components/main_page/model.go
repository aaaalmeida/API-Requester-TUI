package main_page

import (
	"api-requester/context"
	"api-requester/tui/components/collection_menu"
	"api-requester/tui/components/header"
	"api-requester/tui/components/request_header_box"
	"api-requester/tui/components/request_response_box"
	"api-requester/tui/components/search_collection"

	tea "github.com/charmbracelet/bubbletea"
)

// SearchCollectionIndex components index
const SearchCollectionIndex int = 0
const CollectionMenuIndex int = 1
const HeaderIndex int = 2
const RequestHeadersIndex int = 3
const RequestResponseIndex int = 4

// Model Main TUI component.
// encapsulate everything
type Model struct {
	activeComponentIndex int
	subcomponents        []tea.Model
	context              *context.AppContext
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		context:              ctx,
		activeComponentIndex: CollectionMenuIndex,
		subcomponents: []tea.Model{
			search_collection.NewModel(ctx),
			collection_menu.NewModel(ctx),
			header.NewModel(ctx),
			request_header_box.NewModel(ctx),
			request_response_box.NewModel(),
		},
	}
}
