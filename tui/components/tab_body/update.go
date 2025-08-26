package tab_body

import (
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case messages.SendRequestToTabMsg:
		m.request = msg.Request
	}

	return m, nil
}
