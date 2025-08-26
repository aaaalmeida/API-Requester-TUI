package main_page

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return m.subcomponents[COLLECTION_MENU_INDEX].Init()
}
