package tab

import (
	"slices"

	messages "api-requester/tui/messages"
	"api-requester/tui/tab_body"
	"api-requester/tui/tab_header"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// go to next tab
		case "right":
			m.cursor = min(m.cursor+1, len(m.tabsHeader)-1)
			return m, nil

		// go to previos tab
		case "left":
			m.cursor = max(m.cursor-1, 0)
			return m, nil

		case "enter":
			m.activeTabIndex = m.cursor
		}

		// NEW REQUEST LOADED
	case messages.LoadRequestMsg:
		if !slices.Contains(m.requests, msg.Request) {
			m.context.Logger.Println(msg.Request)
			m.context.Logger.Println(m.requests)
			m.requests = append(m.requests, msg.Request)

			th := tab_header.NewModel(msg.Request)
			m.tabsHeader = append(m.tabsHeader, &th)

			tb := tab_body.NewModel(msg.Request)
			m.body = &tb
		}
	}

	return m, nil
}
