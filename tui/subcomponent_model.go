package tui

import tea "github.com/charmbracelet/bubbletea"

type Subcomponent interface {
	Init() tea.Cmd
	Update(tea.Msg) (Subcomponent, tea.Cmd)
	View() string
}
