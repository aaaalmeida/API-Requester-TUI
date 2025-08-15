package main_page

import (
	"api-requester/context"
	"api-requester/tui/collection_menu"
	"api-requester/tui/tab"

	tea "github.com/charmbracelet/bubbletea"
)

// Main TUI component.
// encapsulate everything else
type model struct {
	active_component_index int
	subcomponents          []tea.Model
}

func NewModel(ctx *context.AppContext) model {
	return model{
		active_component_index: 0,
		subcomponents: []tea.Model{
			// FIXME: adicionar os subcomponentes de verdade
			collection_menu.NewModel(ctx),
			tab.NewModel(ctx),
		},
	}
}
