package header

import (
	messages "api-requester/tui/messages"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.selectedTabIndex = m.cursor
		case "left":
			if m.cursor == len(m.requests)-1 {
				m.cursor = 0
			} else {
				m.cursor++
			}
		case "right":
			if m.cursor == 0 {
				m.cursor = len(m.requests) - 1
			} else {
				m.cursor--
			}
		}

		// ADD REQUEST ONLY IF NOT ALREADY EXISTS
	case messages.SendRequestToTabMsg:
		if !slices.Contains(m.requests, msg.Request) {
			m.requests = append(m.requests, msg.Request)
		}
	}

	return m, nil
}
