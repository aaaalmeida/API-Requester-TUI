package tab

import (
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// go to next tab
		case "right":
			m.activeTab = min(m.activeTab+1, len(m.requests)-1)
			return m, nil

		// go to previos tab
		case "left":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
		}

	case messages.LoadRequestMsg:
		m.context.Logger.Println(msg.Request.ID)
		m.context.Logger.Println(msg.Request.Name)
		m.context.Logger.Println(msg.Request.Url)
		m.requests = append(m.requests, msg.Request)
	}

	return m, nil
}
