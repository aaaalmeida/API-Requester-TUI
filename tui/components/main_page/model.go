package main_page

import (
	"api-requester/context"
	"api-requester/tui/components/collection_menu"
	"api-requester/tui/components/search_collection"
	"api-requester/tui/components/tab"

	tea "github.com/charmbracelet/bubbletea"
)

// Main TUI component.
// encapsulate everything
type model struct {
	active_component_index int
	subcomponents          []tea.Model
	context                *context.AppContext
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:                ctx,
		active_component_index: 1,
		subcomponents: []tea.Model{
			search_collection.NewModel(ctx),
			collection_menu.NewModel(ctx),
			tab.NewModel(ctx),
		},
	}
}
