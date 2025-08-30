package header

import (
	"api-requester/tui/commands"
	"api-requester/tui/messages"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":

			if len(m.requests) == 0 || m.cursor < 0 || m.cursor >= len(m.requests) {
				return m, nil
			}
			m.selectedTabIndex = m.cursor
			return m, commands.UserPressEnterInRequestCmd(m.requests[m.selectedTabIndex])

		case "right":
			if len(m.requests) == 0 {
				return m, nil
			}
			if m.cursor == len(m.requests)-1 {
				m.cursor = 0
			} else {
				m.cursor++
			}

		case "left":
			if len(m.requests) == 0 {
				return m, nil
			}
			if m.cursor == 0 {
				m.cursor = len(m.requests) - 1
			} else {
				m.cursor--
			}

		}

		// USER PRESS ENTER IN COLLECTION MENU
		// ADD REQUEST ONLY IF NOT ALREADY EXISTS
	case messages.SendRequestMsg:
		if !slices.Contains(m.requests, msg.Request) {
			m.requests = append(m.requests, msg.Request)
			if len(m.requests) == 1 {
				m.cursor = 0
				m.selectedTabIndex = 0
			}
		}
	}

	return m, nil
}
