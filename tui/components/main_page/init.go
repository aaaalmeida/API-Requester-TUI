package main_page

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return m.subcomponents[m.activeComponentIndex].Init()
}
