package request_header_box

import (
	cmds "api-requester/tui/commands"
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.cursor > 0 {
				m.subcomponents[m.cursor].Blur()
				m.cursor--
				m.subcomponents[m.cursor].Focus()
			}
			return m, nil
		case "right":
			if m.cursor < len(m.subcomponents)-1 {
				m.subcomponents[m.cursor].Blur()
				m.cursor++
				m.subcomponents[m.cursor].Focus()
			}
			return m, nil
		case "enter", " ":
			m.selectedComponentIndex = m.cursor

			_, cmd := m.subcomponents[m.selectedComponentIndex].Update(msg)
			return m, cmd
		}

	// USER PRESSED BUTTON TO SEND REQUEST
	// CallRequestCmd from Button
	case messages.ButtonPressedMsg:
		if msg.Action == "Send Request" && m.request != nil {
			return m, cmds.CallRequestCmd(m.request)
		}

	// USER PRESSED ENTER IN HEADER_COMPONENT
	// UserPressEnterInRequestCmd from Header Component
	case messages.LoadRequestMsg:
		m.request = msg.Request
		_, cmd := m.subcomponents[INPUT_URL_INDEX].Update(
			messages.SendStringMsg{Value: msg.Request.Url})
		return m, cmd

	// USER PRESSED ENTER IN COLLECTION_MENU
	// SendRequestToTabCmd from Collection_menu
	case messages.SendRequestMsg:
		m.request = msg.Request
		_, cmd := m.subcomponents[INPUT_URL_INDEX].Update(
			messages.SendStringMsg{Value: msg.Request.Url})
		return m, cmd

	// INITIAL CMD
	// FETCHES METHODS FROM DB AND SEND TO SELECT COMPONENT
	case messages.LoadMethodsMsg:
		m.subcomponents[SELECT_MENU_INDEX].Update(msg)

	case messages.SendStringMsg:
		m.context.Logger.Println("ISSO Q VEIO DO INPUT", msg)
	}

	return m, nil
}
