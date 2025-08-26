package main_page

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return m.subcomponents[1].Init()
}
