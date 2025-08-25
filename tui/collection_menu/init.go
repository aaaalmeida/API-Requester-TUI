package collection_menu

import (
	cmd "api-requester/tui/commands"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return cmd.FetchCollectionsFromDBCmd(m.context)
}
