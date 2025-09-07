package input

import (
	cmd "api-requester/tui/commands"
	messages "api-requester/tui/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, cmd.SendStringToInputCmd(m.textInput.Value())
		}

	// USER PRESS ENTER IN HEADER_COMPONENT
	// USER PRESS ENTER IN COLLECTION_MENU
	case messages.SendStringMsg:
		m.textInput.SetValue(msg.Value)
	}

	_, cmd := m.textInput.Update(msg)
	return m, cmd
}
