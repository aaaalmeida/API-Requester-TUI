package focusable

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Focusable interface {
	Focus()
	Blur()
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}
