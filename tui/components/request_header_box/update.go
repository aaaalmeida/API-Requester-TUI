package request_header_box

import (
	"api-requester/shared/focusable"
	cmds "api-requester/tui/commands"
	"api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.selectedComponentIndex > 0 {
				m.subcomponents[m.selectedComponentIndex].Blur()
				m.selectedComponentIndex--
				m.subcomponents[m.selectedComponentIndex].Focus()
			}
			return m, nil

		case "right":
			if m.selectedComponentIndex < len(m.subcomponents)-1 {
				m.subcomponents[m.selectedComponentIndex].Blur()
				m.selectedComponentIndex++
				m.subcomponents[m.selectedComponentIndex].Focus()
			}
			return m, nil

		default:
			aux, cmd := m.subcomponents[m.selectedComponentIndex].Update(msg)
			m.subcomponents[m.selectedComponentIndex] = aux.(focusable.Focusable)
			return m, cmd
		}

	// USER PRESSED BUTTON TO SEND REQUEST
	// CallRequestCmd from Button
	case messages.ButtonPressedMsg:
		if msg.Action == ">>" && m.request != nil {
			return m, cmds.CallRequestCmd(m.request)
		}

	// USER PRESSED ENTER IN HEADER_COMPONENT
	// UserPressEnterInRequestCmd from Header Component
	case messages.LoadRequestMsg:
		m.request = msg.Request
		aux, cmd := m.subcomponents[INPUT_URL_INDEX].Update(
			messages.SendStringMsg{Value: msg.Request.Url})
		m.subcomponents[INPUT_URL_INDEX] = aux.(focusable.Focusable)
		return m, cmd

	// USER PRESSED ENTER IN COLLECTION_MENU
	// SendRequestToTabCmd from Collection_menu
	case messages.SendRequestMsg:
		m.request = msg.Request
		aux, cmd := m.subcomponents[INPUT_URL_INDEX].Update(
			messages.SendStringMsg{Value: msg.Request.Url})
		m.subcomponents[INPUT_URL_INDEX] = aux.(focusable.Focusable)
		return m, cmd

	// INITIAL CMD
	// FETCHES METHODS FROM DB AND SEND TO SELECT COMPONENT
	case messages.LoadMethodsMsg:
		// do not return cmd back or program will get in infinite loop
		aux, _ := m.subcomponents[SELECT_MENU_INDEX].Update(msg)
		m.subcomponents[SELECT_MENU_INDEX] = aux.(focusable.Focusable)

	case messages.InputChangedMsg:
		m.request.Url = msg.Value
	}

	return m, nil
}
