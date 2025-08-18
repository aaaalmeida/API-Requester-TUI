package collection_menu

import (
	collection_menu "api-requester/tui/collection_menu/command"

	tea "github.com/charmbracelet/bubbletea"
)

// no init I/O command
func (m model) Init() tea.Cmd {

	return collection_menu.FetchCollectionsCmd(m.context)
}
