package tab

import (
	"slices"

	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		// PASS COMMAND TO TAB HEADER AND BODY
	case messages.SendRequestToTabMsg:
		var headerCmd, bodyCmd tea.Cmd
		m.tab_header, headerCmd = m.tab_header.Update(msg)
		m.tab_body, bodyCmd = m.tab_body.Update(msg)

		return m, tea.Batch(headerCmd, bodyCmd)

		// NEW REQUEST LOADED
	case messages.LoadRequestMsg:
		if !slices.Contains(m.requests, msg.Request) {
			m.requests = append(m.requests, msg.Request)
		}
	}

	return m, nil
}
