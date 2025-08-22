package tab_header

import (
	cmd "api-requester/tui/commands"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if !m.isSelected {
				m.isSelected = true
				return m, cmd.SendRequestCmd(m.request)
			}
		}
	}

	return m, nil
}
