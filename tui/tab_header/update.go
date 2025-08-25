package tab_header

import (
	messages "api-requester/tui/messages"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// TODO
		case "enter":
		case "left":
		case "right":
		}

	case messages.SendRequestToTabMsg:
		if !slices.Contains(m.requests, msg.Request) {
			m.requests = append(m.requests, msg.Request)
		}
	}

	return m, nil
}
